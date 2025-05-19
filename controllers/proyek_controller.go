package controllers

import (
	"backend-konstruksi/config"
	"backend-konstruksi/models"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllProyek(c *fiber.Ctx) error {
	collection := config.GetCollection("proyek")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result []models.Proyek
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error finding projects: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var proyek models.Proyek
		if err := cursor.Decode(&proyek); err != nil {
			log.Printf("Error decoding project: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "Gagal decode data"})
		}
		result = append(result, proyek)
	}

	if result == nil {
		result = []models.Proyek{}
	}

	return c.JSON(result)
}

// GetProyekByID
func GetProyekByID(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("Attempting to fetch project with ID: %s", id)

	// Validasi ID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error converting ID to ObjectID: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	collection := config.GetCollection("proyek")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var proyek models.Proyek
	filter := bson.M{"_id": objID}
	log.Printf("Executing MongoDB query with filter: %v", filter)

	err = collection.FindOne(ctx, filter).Decode(&proyek)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No project found with ID: %s", id)
			return c.Status(404).JSON(fiber.Map{"error": "Proyek tidak ditemukan"})
		}
		log.Printf("Database error: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data proyek"})
	}

	log.Printf("Successfully found project: %+v", proyek)
	return c.JSON(proyek)
}

// Tambah Proyek
func CreateProyek(c *fiber.Ctx) error {
	proyekCollection := config.GetCollection("proyek")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var proyek models.Proyek
	if err := c.BodyParser(&proyek); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	proyek.ID = primitive.NewObjectID()

	_, err := proyekCollection.InsertOne(ctx, proyek)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan data"})
	}

	return c.Status(201).JSON(proyek)
}

// Update Proyek
func UpdateProyek(c *fiber.Ctx) error {
	proyekCollection := config.GetCollection("proyek")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	var proyek models.Proyek
	if err := c.BodyParser(&proyek); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	update := bson.M{
		"$set": proyek,
	}

	_, err = proyekCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal update data"})
	}

	return c.JSON(fiber.Map{"message": "Berhasil update data"})
}

// Hapus Proyek
func DeleteProyek(c *fiber.Ctx) error {
	proyekCollection := config.GetCollection("proyek")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	_, err = proyekCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus data"})
	}

	return c.JSON(fiber.Map{"message": "Berhasil menghapus data"})
}

// Tambahkan fungsi ini di file yang sama
