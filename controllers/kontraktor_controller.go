package controllers

import (
	"context"
	"time"

	"backend-konstruksi/config"
	"backend-konstruksi/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
