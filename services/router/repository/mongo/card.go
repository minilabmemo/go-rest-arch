package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/minilabmemo/go-rest-arch/services/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongo_driver "go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type mongoCardRepository struct {
	Database   string
	Collection string
	Conn       *mongo_driver.Client
}

// input is a connected connection
func NewMongoCardRepository(conn *mongo_driver.Client, database, collection string) models.CardRpository {
	return &mongoCardRepository{Conn: conn, Database: database, Collection: collection}
}

func (m *mongoCardRepository) fetch(ctx context.Context) (result []models.Card, err error) {
	res := []models.Card{}
	cursor, err := m.Conn.Database(m.Database).Collection(m.Collection).Find(ctx, bson.D{})
	if err != nil {
		return res, err
	}
	var results []models.Card

	for cursor.Next(ctx) {
		//	fmt.Println(cursor)   //&{{"_id": {"$oid":"63c39d754be73bb880adb763"},"title": "string","content": "string"}
		var card models.Card
		if err := cursor.Decode(&card); err != nil {
			fmt.Println(err)
			return res, err
		}

		results = append(results, card)
	}

	return results, nil
}

func (m *mongoCardRepository) collectionDocuments(ctx context.Context) (int64, error) {
	return m.Conn.Database(m.Database).Collection(m.Collection).CountDocuments(ctx, bson.D{})
}

func (m *mongoCardRepository) FetchAll(ctx context.Context) ([]models.Card, string, error) {

	res, err := m.fetch(ctx)
	if err != nil {
		return nil, "", err
	}
	total, err := m.collectionDocuments(ctx)
	if err != nil {
		return nil, "", err
	}
	return res, fmt.Sprintf("%d", total), nil
}

func (m *mongoCardRepository) Store(ctx context.Context, card *models.CardUpdate) (string, error) {
	if card == nil {
		return "", errors.New("card is empty")
	}
	collection := m.Conn.Database(m.Database).Collection(m.Collection)
	res, err := collection.InsertOne(ctx, card)
	if err != nil {
		return "", err
	}

	id := fmt.Sprintf("%v", res.InsertedID)
	return id, nil
}

func (m *mongoCardRepository) Update(ctx context.Context, id string, card *models.CardUpdate) (string, error) {
	if card == nil {
		return "", errors.New("card is empty")
	}
	objID, err := primitive.ObjectIDFromHex(id) //比較特別的是要轉成objectID
	if err != nil {
		return "", err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	update := bson.D{{"$set", card}}
	collection := m.Conn.Database(m.Database).Collection(m.Collection)
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", err
	}
	zap.S().Debugf("Update:%s res:%#v", id, res)
	return id, nil
}

func (m *mongoCardRepository) Delete(ctx context.Context, id string) (string, error) {
	objID, err := primitive.ObjectIDFromHex(id) //比較特別的是要轉成objectID
	if err != nil {
		return "", err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	collection := m.Conn.Database(m.Database).Collection(m.Collection)
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}
	zap.S().Debugf("DeleteOne:%s res:%#v", id, res)
	return id, nil
}
