package controllers

import (
	"backend-konstruksi/config"
	"backend-konstruksi/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllProyek(c *fiber.Ctx) error {
	proyekCollection := config.GetCollection("proyek")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var proyekList []models.Proyek

	cursor, err := proyekCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var proyek models.Proyek
		if err := cursor.Decode(&proyek); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal decode data"})
		}
		proyekList = append(proyekList, proyek)
	}

	return c.JSON(proyekList)
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
