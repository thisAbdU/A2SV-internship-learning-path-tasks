package repository

import (
	"context"
	"example/GO-PRACTICE-EXERCISE/GO-API-exercise/domain/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	database   *mongo.Database
	collection string
}

func NewUserRepository(database *mongo.Database, collection string) entities.UserRepository {
	return &userRepository{
		database:   database,
		collection: collection,
	}
}

func (ur *userRepository) GetUser(ctx context.Context, param string) ([]*entities.User, error) {
	var users []*entities.User

	filter := bson.M{
		"$or": []bson.M{
			{"username": primitive.Regex{Pattern: param, Options: "i"}},
			{"password": primitive.Regex{Pattern: param, Options: "i"}},
		},
	}

	cursor, err := ur.database.Collection(ur.collection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user entities.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *userRepository) GetUserByID(ctx context.Context, id string) (*entities.User, error) {
	filter := bson.M{
		"_id": id,
	}

	result := ur.database.Collection(ur.collection).FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var user entities.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, id string, updatedUser entities.User) error {
	filter := bson.M{
		"_id": id,
	}

	update := bson.M{
		"$set": updatedUser,
	}

	_, err := ur.database.Collection(ur.collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, id string) error {
	filter := bson.M{
		"_id": id,
	}

	_, err := ur.database.Collection(ur.collection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) CreateUser(ctx context.Context, newUser entities.User) error {
	_, err := ur.database.Collection(ur.collection).InsertOne(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}

