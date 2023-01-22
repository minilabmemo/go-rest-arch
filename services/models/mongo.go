package models

import (
	"context"

	mongo_driver "go.mongodb.org/mongo-driver/mongo"
)

// MongoRepository represent  the mongo client's repository contract
type MongoRepository interface {
	Close(ctx context.Context) (err error)
	Connection() (Conn *mongo_driver.Client)
}
