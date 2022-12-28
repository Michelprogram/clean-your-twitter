package database

import (
	"api-clean-twitter/entities"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddUser(user entities.User) error {

	filter := bson.M{
		"user.data.twitterid": user.Profile.Data.TwitterId,
	}

	update := bson.M{
		"$set": bson.M{
			"user": user,
		},
	}

	opts := options.Update().SetUpsert(true)

	res, err := collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}

func GetUserByTwitterId(twitter_id string) (*entities.UserMongo, error) {

	filter := bson.M{
		"user.profile.data.twitter_id": twitter_id,
	}
	var user entities.UserMongo

	err := collection.FindOne(ctx, filter, nil).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func GetImageUserById(twitter_id string) (*entities.UserMongo, error) {

	filter := bson.M{
		"user.profile.data.twitter_id": twitter_id,
	}

	opts := options.FindOne().SetProjection(bson.M{
		"user.profile.data.profile_image_url": 1,
		"_id":                                 0,
	})

	var user entities.UserMongo

	err := collection.FindOne(ctx, filter, opts).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil

}
