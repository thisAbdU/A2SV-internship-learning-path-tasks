package repository

import (
	"context"
	"example/GO-PRACTICE-EXERCISE/GO-API-exercise/domain/entities"
	"example/GO-PRACTICE-EXERCISE/GO-API-exercise/domain/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	database   *mongo.Database
	collection string
}

func NewTaskRepository(database *mongo.Database, collection string) entities.TaskRepository{
	return &taskRepository{
		database:   database,
		collection: collection,
	}
}
func (tr *taskRepository) GetTasks(ctx context.Context, param string) ([]*model.TaskInfo, error) {
    var tasks []*model.TaskInfo

    filter := bson.M{
        "$or": []bson.M{
            {"title": primitive.Regex{Pattern: param, Options: "i"}},
            {"description": primitive.Regex{Pattern: param, Options: "i"}},
        },
    }

    cursor, err := tr.database.Collection(tr.collection).Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var task entities.Task
        if err := cursor.Decode(&task); err != nil {
            return nil, err
        }
        tasks = append(tasks, &model.TaskInfo{
            Title:       task.Title,
            Description: task.Description,
            DueDate:     (time.Now()).Add(3 * 24 * time.Hour).Format(time.RFC3339),
        })
    }
    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return tasks, nil
}


func (tr *taskRepository) GetTaskByID(ctx context.Context, id string) (*entities.Task, error) {
    filter := bson.M{"_id": id}

    var task entities.Task

    err := tr.database.Collection(tr.collection).FindOne(ctx, filter).Decode(&task)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, err
        }
        return nil, err
    }

    return &task, nil
}

func (tr *taskRepository) UpdateTask (ctx context.Context, id string, updatedTask entities.Task) error{
	filter := bson.M{
		"_id": id,
	}

	update := bson.M{
		"$set": updatedTask,
	}

	_, err := tr.database.Collection(tr.collection).UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
	
}

func (tr *taskRepository) DeleteTask(ctx context.Context, id string) error{
	filter := bson.M{
		"_id": id,
	}

	_, err := tr.database.Collection(tr.collection).DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
	
}

func (tr *taskRepository) CreateTask(ctx context.Context, newTask entities.Task) error {
	_, err := tr.database.Collection(tr.collection).InsertOne(ctx, newTask)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

