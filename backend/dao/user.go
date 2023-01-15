package dao

import (
	"api-clean-twitter/database"
	"api-clean-twitter/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddUser(user models.User) (*mongo.UpdateResult, error) {
	var client = database.ClientDB

	filter := bson.M{
		"user.twitter_id": user.TwitterId,
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
		"user.twitter_id": twitter_id,
	}

	res, err := client.Collection.DeleteOne(ctx, filter, nil, nil)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetUserByTwitterId(twitter_id string) (*models.User, error) {
	var client = database.ClientDB

	filter := bson.M{
		"user.twitter_id": twitter_id,
	}
	var user models.UserMongo

	err := client.Collection.FindOne(ctx, filter, nil).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user.User, nil
}

func GetImageUserById(twitter_id string) (*models.User, error) {
	var client = database.ClientDB
	var user models.UserMongo

	filter := bson.M{
		"user.twitter_id": twitter_id,
	}

	opts := options.FindOne().SetProjection(bson.M{
		"user.profile_image_url": 1,
		"user.username":          1,
		"user.name":              1,
		"_id":                    0,
	})

	err := client.Collection.FindOne(ctx, filter, opts).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user.User, nil

}

func UpdateTokenUser(new_token models.Token, old_access_token string) (*mongo.UpdateResult, error) {
	var client = database.ClientDB

	filter := bson.M{
		"user.token.access_token": old_access_token,
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
