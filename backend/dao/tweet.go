package dao

import (
	"api-clean-twitter/database"
	"api-clean-twitter/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


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
			"tweets_ids": bson.M{
				"$each": tweets_ids,
			},
		},
	}

	opts := options.Update().SetUpsert(true)

	res, err := tweets_collection.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func FetchAllTweets() ([]models.Tweets, error) {
	var tweets_collection = database.ClientDB.Database.Collection("tweets")
	var results []models.Tweets

	filter := bson.M{}

	cursor, err := tweets_collection.Find(ctx, filter, nil)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

func RemoveTweets(tweets_ids []string, twitter_id string) error {

	var tweets_collection = database.ClientDB.Database.Collection("tweets")

	filter := bson.M{
		"twitter_id": twitter_id,
	}

	options := bson.M{
		"$pull": bson.M{
			"tweets_ids": bson.M{
				"$in": tweets_ids,
			},
		},
	}

	_, err := tweets_collection.UpdateOne(ctx, filter, options)

	if err != nil {
		return err
	}

	return nil

}

func CountTweets(twitter_id string) (int, error) {

	var tweets_collection = database.ClientDB.Database.Collection("tweets")

	type Count struct {
		Count int
	}

	var res []Count

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"twitter_id": "872417601846218752",
			},
		},
		{
			"$project": bson.M{
				"count": bson.M{
					"$size": "$tweets_ids",
				},
				"_id": 0,
			},
		},
	}

	cursor, err := tweets_collection.Aggregate(ctx, pipeline, nil)

	if err != nil {
		return 0, err
	}

	err = cursor.All(ctx, &res)

	if err != nil {
		return 0, err
	}

	return res[0].Count, nil
}

func RemoveUserFromTweets(twitter_id string) error {
	var tweets_collection = database.ClientDB.Database.Collection("tweets")

	filter := bson.M{
		"twitter_id": twitter_id,
	}

	_, err := tweets_collection.DeleteOne(ctx, filter, nil)

	if err != nil {
		return err
	}

	return nil
}

func UpdateTokenTweets(newToken models.Token, oldAccesToken string) error {
	var tweets_collection = database.ClientDB.Database.Collection("tweets")

	filter := bson.M{
		"token.access_token": oldAccesToken,
	}

	update := bson.M{
		"$set": bson.M{
			"token": newToken,
		},
	}

	_, err := tweets_collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil

}

// TODO : A rename
func UpdateTokenTweets2(user models.User) error {
	var tweets_collection = database.ClientDB.Database.Collection("tweets")

	filter := bson.M{
		"twitter_id": user.TwitterId,
	}

	update := bson.M{
		"$set": bson.M{
			"token": user.Token,
		},
	}

	_, err := tweets_collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}
