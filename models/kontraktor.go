package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Kontraktor struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nama    string             `bson:"nama" json:"nama"`
	Alamat  string             `bson:"alamat" json:"alamat"`
	Telepon string             `bson:"telepon" json:"telepon"`
	Email   string             `bson:"email" json:"email"`
}
