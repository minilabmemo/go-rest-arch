package mongo

import (
	"context"

	"github.com/minilabmemo/go-rest-arch/services/config"
	"github.com/minilabmemo/go-rest-arch/services/models"
	"github.com/pkg/errors"
	mongo_driver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClient struct {
	Conn *mongo_driver.Client
}

func NewMongoClient(ctx context.Context, mongoInfo config.ClientInfo) (models.MongoRepository, error) {
	// MongoDB连接地址 ex:"mongodb://localhost:27017"
	uri := mongoInfo.Url("")
	// 连接MongoDB
	client, err := mongo_driver.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &mongoClient{Conn: client}, nil
}

func (m *mongoClient) Close(ctx context.Context) (err error) {
	if m.Conn == nil {
		return errors.Errorf("connection is not initialize.")
	}

	return m.Conn.Disconnect(ctx)
}

func (m *mongoClient) Connection() (Conn *mongo_driver.Client) {
	return m.Conn
}
