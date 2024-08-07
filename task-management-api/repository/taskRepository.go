package repository

import (
	"context"
	"task-management-api/domain/entities"
	"task-management-api/domain/model"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (tr *taskRepository) GetTasks(ctx context.Context, userID string) ([]*model.TaskInfo, error) {
    var tasks []*model.TaskInfo

	filter := bson.M{
		"user_id": userID,
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


func (tr *taskRepository) GetTaskByID(ctx context.Context, id string, userID string) (*entities.Task, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"_id": id},
			{"user_id": userID},
		},
	}

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

func (tr *taskRepository) UpdateTask (ctx context.Context, id string, updatedTask entities.Task, userID string) error{
	filter := bson.M{
		"$and": []bson.M{
			{"_id": id},
			{"user_id": userID},
		},
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

func (tr *taskRepository) DeleteTask(ctx context.Context, id string, userID string) error{
	filter := bson.M{
		"$and": []bson.M{
			{"_id": id},
			{"user_id": userID},
		},
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

