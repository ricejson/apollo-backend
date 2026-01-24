package dao

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

type MongoToggleDAO struct {
	col *mongo.Collection
}

func NewMongoToggleDAO(col *mongo.Collection) *MongoToggleDAO {
	return &MongoToggleDAO{
		col: col,
	}
}

func (dao *MongoToggleDAO) List(ctx context.Context) ([]Toggle, error) {
	cur, err := dao.col.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	toggles := make([]Toggle, 0)
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var toggle Toggle
		if err = cur.Decode(&toggle); err != nil {
			log.Fatal(err)
			return nil, err
		}
		toggles = append(toggles, toggle)
	}
	return toggles, nil
}

func (dao *MongoToggleDAO) InsertOne(ctx context.Context, toggle Toggle) (bool, error) {
	result, err := dao.col.InsertOne(ctx, toggle)
	if err != nil {
		return false, err
	}
	if result.InsertedID == nil || !result.Acknowledged {
		return false, errors.New("inserted id is null")
	}
	return true, nil
}
