package repository

import (
	"api-clean-twitter/database"
	"api-clean-twitter/entities"
	"api-clean-twitter/twitter"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func AddUser(user entities.User) (*mongo.UpdateResult, error) {
	var client = database.ClientDB

	filter := bson.M{
		"user.profile.data.twitter_id": user.Profile.Data.TwitterId,
	}

	update := bson.M{
		"$set": bson.M{
			"user": user,
		},
	}

	opts := options.Update().SetUpsert(true)

	res, err := client.Collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func DeleteUserByTwitterId(twitter_id string) (*mongo.DeleteResult, error) {
	var client = database.ClientDB

	filter := bson.M{
		"user.profile.data.twitter_id": twitter_id,
	}

	res, err := client.Collection.DeleteOne(ctx, filter, nil, nil)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetUserByTwitterId(twitter_id string) (*entities.UserMongo, error) {
	var client = database.ClientDB

	filter := bson.M{
		"user.profile.data.twitter_id": twitter_id,
	}
	var user entities.UserMongo

	err := client.Collection.FindOne(ctx, filter, nil).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func GetImageUserById(twitter_id string) (*entities.UserMongo, error) {
	var client = database.ClientDB

	filter := bson.M{
		"user.profile.data.twitter_id": twitter_id,
	}

	opts := options.FindOne().SetProjection(bson.M{
		"user.profile.data.profile_image_url": 1,
		"user.profile.data.username":          1,
		"user.profile.data.name":              1,
		"_id":                                 0,
	})

	var user entities.UserMongo

	err := client.Collection.FindOne(ctx, filter, opts).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil

}

func UpdateTokenUser(twitter_id string, new_token *twitter.Token) (*mongo.UpdateResult, error) {
	var client = database.ClientDB

	filter := bson.M{
		"user.profile.data.twitter_id": twitter_id,
	}

	update := bson.M{
		"$set": bson.M{
			"user.token": new_token,
		},
	}

	opts := options.Update()

	res, err := client.Collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		return nil, err
	}

	return res, nil
}
