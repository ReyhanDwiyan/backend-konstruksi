package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Proyek struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Nama    string             `bson:"nama"`
	Lokasi  string             `bson:"lokasi"`
	Tanggal time.Time          `bson:"tanggal"`
}
