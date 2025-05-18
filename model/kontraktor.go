package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Kontraktor struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Nama    string             `bson:"nama"`
	Alamat  string             `bson:"alamat"`
	Telepon string             `bson:"telepon"`
}
