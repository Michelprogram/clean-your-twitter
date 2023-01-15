package dao

import (
	"api-clean-twitter/database"
	"api-clean-twitter/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO : voir si l'upsert fonctionne bien
func AddTweets(tweets_ids []string, token models.Token, twitter_id string) (*mongo.UpdateResult, error) {
	var tweets_collection = database.ClientDB.Database.Collection("tweets")

	filter := bson.M{
		"twitter_id": twitter_id,
	}

	update := bson.M{
		"$set": bson.M{
			"token": token,
		},
		"$addToSet": bson.M{
			"tweets_ids": tweets_ids,
		},
	}

	opts := options.Update().SetUpsert(true)

	res, err := tweets_collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		return nil, err
	}

	return res, nil
}
