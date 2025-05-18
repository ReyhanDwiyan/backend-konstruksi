package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	ConnectDB()
}

func ConnectDB() {
	mongoURI := os.Getenv("MONGOSTRING")
	if mongoURI == "" {
		fmt.Println("MONGOSTRING not found in environment variables")
		return
	}

	clientOpts := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		fmt.Println("MongoConnect error:", err)
		return
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		fmt.Println("Ping error:", err)
		return
	}

	fmt.Println("Connected to MongoDB successfully!")
	DB = client.Database("kontruksi")
}

func GetCollection(name string) *mongo.Collection {
	if DB == nil {
		fmt.Println("Database not initialized. Call ConnectDB() first.")
		return nil
	}
	return DB.Collection(name)
}
