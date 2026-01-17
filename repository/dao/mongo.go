package dao

import (
	"context"
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
