package repository

import (
	"backend-konstruksi/config"
	"backend-konstruksi/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var kontraktorCollection = config.GetCollection("kontraktor")
var proyekCollection = config.GetCollection("proyek")

func CreateKontraktor(k model.Kontraktor) error {
	_, err := kontraktorCollection.InsertOne(context.TODO(), k)
	if err != nil {
		fmt.Println("Insert error:", err)
		return err
	}
	return nil
}

func CreateProyek(p model.Proyek) error {
	_, err := proyekCollection.InsertOne(context.TODO(), p)
	return err
}

func GetAllProyek() ([]model.Proyek, error) {
	var proyeks []model.Proyek
	cursor, err := proyekCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &proyeks)
	return proyeks, err
}

func GetProyekByID(id string) (model.Proyek, error) {
	var proyek model.Proyek
	objID, _ := primitive.ObjectIDFromHex(id)
	err := proyekCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&proyek)
	return proyek, err
}

func UpdateProyek(id string, p model.Proyek) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := proyekCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": p},
	)
	return err
}

func DeleteProyek(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := proyekCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
