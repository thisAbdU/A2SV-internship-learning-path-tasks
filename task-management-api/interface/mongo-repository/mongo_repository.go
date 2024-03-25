package task

import (
	"context"
	"errors"
	task "example/GO-PRACTICE-EXERCISE/GO-API-exercise/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "github.com/google/uuid"
)

type MongoTaskRepository struct {
    Collection *mongo.Collection
}

func NewMongoTaskRepository(db *mongo.Database) *MongoTaskRepository {
    return &MongoTaskRepository{
        Collection: db.Collection("task"),
    }
}

func (r *MongoTaskRepository) GetTasks() ([]task.Task, error) {

    var tasks []task.Task

    collection := r.Collection

    filter := bson.M{}

    ctx, cacnel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cacnel()
   
    cursor, err := collection.Find(ctx, filter)
    if err != nil {
        return nil, err 
    }
    defer cursor.Close(context.Background())

    // Iterate through the cursor and decode tasks
    for cursor.Next(context.Background()) {
        var task task.Task
        if err := cursor.Decode(&task); err != nil {
            return nil, err 
        }
        tasks = append(tasks, task) 
    }

    if err := cursor.Err(); err != nil {
        return nil, err 
    }

    return tasks, nil 
}

func (r * MongoTaskRepository) GetTaskByID (id string) (task.Task, error){
    var task task.Task
    filter := bson.M{"id": id}

    ctx, cacnel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cacnel()

    err := r.Collection.FindOne(ctx, filter).Decode(&task)
    if err != nil{
        return task, err
    }
    return task, nil
}

func (r *MongoTaskRepository) UpdateTask(id string) error {
    filter := bson.M{"id": id}

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Check if the task exists
    var existingTask task.Task
    err := r.Collection.FindOne(ctx, filter).Decode(&existingTask)
    if err == mongo.ErrNoDocuments {
        return errors.New("no task found with given ID")
    } else if err != nil {
        return err
    }

    // Update the task status to "done"
    update := bson.M{"$set": bson.M{"status": "done"}}

    // Perform the update operation
    res, err := r.Collection.UpdateOne(ctx, filter, update)
    if err != nil {
        return err
    }

    // Check if any task was modified
    if res.ModifiedCount < 1 {
        return errors.New("no task modified")
    }

    return nil
}

func (r *MongoTaskRepository) DeleteTask(id string) error {
    filter := bson.M{"id": id}

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var existingTask task.Task
    err := r.Collection.FindOne(ctx, filter).Decode(&existingTask)
    if err == mongo.ErrNoDocuments {
        return errors.New("no task found with given ID")
    } else if err != nil {
        return err
    }

    // Perform the delete operation
    _, err = r.Collection.DeleteOne(ctx, filter)
    if err != nil {
        return err
    }

    return nil
}

func (r *MongoTaskRepository) CreateTask(tsk task.Task) (string, error) {
    tsk.ID = uuid.New().String()
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := r.Collection.InsertOne(ctx, tsk)
    if err != nil {
        return "", err
    }
    return tsk.ID, nil
}