package controllers

import (
	"context"
	"log"
	"time"

	"backend-konstruksi/config"
	"backend-konstruksi/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetKontraktorByID(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("Fetching kontraktor with ID: %s", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	collection := config.GetCollection("kontraktor")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var kontraktor models.Kontraktor
	filter := bson.M{"_id": objID}
	log.Printf("Searching with filter: %v", filter)

	err = collection.FindOne(ctx, filter).Decode(&kontraktor)
	if err != nil {
		log.Printf("Error finding kontraktor: %v", err)
		return c.Status(404).JSON(fiber.Map{"error": "Kontraktor tidak ditemukan"})
	}

	log.Printf("Successfully found kontraktor: %+v", kontraktor)
	return c.JSON(kontraktor)
}

func GetAllKontraktor(c *fiber.Ctx) error {
	collection := config.GetCollection("kontraktor")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result []models.Kontraktor
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item models.Kontraktor
		if err := cursor.Decode(&item); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal decode data"})
		}
		result = append(result, item)
	}

	return c.JSON(result)
}

func CreateKontraktor(c *fiber.Ctx) error {
	collection := config.GetCollection("kontraktor")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data models.Kontraktor
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	data.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal simpan data"})
	}

	return c.Status(201).JSON(data)
}

func UpdateKontraktor(c *fiber.Ctx) error {
	collection := config.GetCollection("kontraktor")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	log.Printf("Updating kontraktor with ID: %s", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	var updateData models.Kontraktor
	if err := c.BodyParser(&updateData); err != nil {
		log.Printf("Invalid request body: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	// Pastikan ID tetap sama
	updateData.ID = objID

	update := bson.M{
		"$set": bson.M{
			"nama":    updateData.Nama,
			"alamat":  updateData.Alamat,
			"telepon": updateData.Telepon,
			"email":   updateData.Email,
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		log.Printf("Error updating kontraktor: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Gagal update data"})
	}

	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Kontraktor tidak ditemukan"})
	}

	log.Printf("Successfully updated kontraktor with ID: %s", id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Berhasil update data",
		"data":    updateData,
	})
}

func DeleteKontraktor(c *fiber.Ctx) error {
	proyekCollection := config.GetCollection("kontraktor")
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
