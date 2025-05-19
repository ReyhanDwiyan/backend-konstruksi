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

func GetMaterialByID(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("Fetching material with ID: %s", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	collection := config.GetCollection("material")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var material models.Material
	filter := bson.M{"_id": objID}
	log.Printf("Searching with filter: %v", filter)

	err = collection.FindOne(ctx, filter).Decode(&material)
	if err != nil {
		log.Printf("Error finding material: %v", err)
		return c.Status(404).JSON(fiber.Map{"error": "Material tidak ditemukan"})
	}

	log.Printf("Successfully found material: %+v", material)
	return c.JSON(material)
}

func GetAllMaterial(c *fiber.Ctx) error {
	collection := config.GetCollection("material")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("Fetching all materials")

	var result []models.Material
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error fetching materials: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item models.Material
		if err := cursor.Decode(&item); err != nil {
			log.Printf("Error decoding material: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": "Gagal decode data"})
		}
		result = append(result, item)
	}

	// Initialize empty array if no results
	if result == nil {
		result = []models.Material{}
	}

	log.Printf("Successfully fetched %d materials", len(result))
	return c.JSON(result)
}

func CreateMaterial(c *fiber.Ctx) error {
	collection := config.GetCollection("material")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data models.Material
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

func UpdateMaterial(c *fiber.Ctx) error {
	collection := config.GetCollection("material")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	log.Printf("Updating material with ID: %s", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	var updateData models.Material
	if err := c.BodyParser(&updateData); err != nil {
		log.Printf("Invalid request body: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Data tidak valid"})
	}

	// Pastikan ID tetap sama
	updateData.ID = objID

	update := bson.M{
		"$set": bson.M{
			"nama":           updateData.Nama,
			"stok":           updateData.Stok,
			"satuan":         updateData.Satuan,
			"harga_per_unit": updateData.HargaPerUnit,
			"proyek_id":      updateData.ProyekID,
			"image_url":      updateData.ImageURL,
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		log.Printf("Error updating material: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Gagal update data"})
	}

	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Material tidak ditemukan"})
	}

	log.Printf("Successfully updated material with ID: %s", id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Berhasil update data",
		"data":    updateData,
	})
}

func DeleteMaterial(c *fiber.Ctx) error {
	collection := config.GetCollection("material")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	log.Printf("Attempting to delete material with ID: %s", id)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Printf("Error deleting material: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus data"})
	}

	if result.DeletedCount == 0 {
		log.Printf("No material found with ID: %s", id)
		return c.Status(404).JSON(fiber.Map{"error": "Material tidak ditemukan"})
	}

	log.Printf("Successfully deleted material with ID: %s", id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Berhasil menghapus data",
		"id":      id,
	})
}
