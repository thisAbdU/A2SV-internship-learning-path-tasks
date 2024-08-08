package repository

import (
	"context"
	"fmt"
	"task-management-api/domain/entities"
	"task-management-api/domain/model"
	"task-management-api/mongo"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(database mongo.Database, collection string) entities.TaskRepository{
	return &taskRepository{
		database:   database,
		collection: collection,
	}
}

func (tr *taskRepository) GetTasks(ctx context.Context, userID string) ([]*model.TaskInfo, error) {
    var tasks []*model.TaskInfo

	filter := bson.M{
		"userid": userID,
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

    return tasks, nil
}


func (tr *taskRepository) GetTaskByID(ctx context.Context, id string, userID string) (*entities.Task, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"$and": []bson.M{
			{"_id": objectID},
			{"userid": userID},
		},
	}

    var task entities.Task

    err := tr.database.Collection(tr.collection).FindOne(ctx, filter).Decode(&task)
    if err != nil {
		log.Println(err)
        if err == mongo.ErrNoDocuments {
            return nil, err
        }
        return nil, err
    }

    return &task, nil
}
func (tr *taskRepository) UpdateTask(ctx context.Context, id string, updatedTask entities.Task, userID string) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Fatal(err)
        return err
    }

    filter := bson.M{
        "_id":    objectID,
        "userid": userID,
    }

	update := bson.M{
		"$set": bson.M{
			"title": updatedTask.Title,
			"description": updatedTask.Description,
			"status": updatedTask.Status,
		},
	}

    result, err := tr.database.Collection(tr.collection).UpdateOne(ctx, filter, update)
	
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("no documents updated")
	}
	

    return nil
}

func (tr *taskRepository) DeleteTask(ctx context.Context, id string, userID string) error{
	objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Fatal(err)
        return err
    }

	filter := bson.M{
		"$and": []bson.M{
			{"_id": objectID},
			{"userid": userID},
		},
	}
	
	numDeleted, err := tr.database.Collection(tr.collection).DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if numDeleted == 0 {
		return fmt.Errorf("no documents deleted")
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

