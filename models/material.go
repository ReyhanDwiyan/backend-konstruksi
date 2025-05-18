package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Material struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nama         string             `bson:"nama" json:"nama"`
	Stok         int                `bson:"stok" json:"stok"`
	Satuan       string             `bson:"satuan" json:"satuan"`
	HargaPerUnit int                `bson:"harga_per_unit" json:"harga_per_unit"`
	ProyekID     string             `bson:"proyek_id" json:"proyek_id"`
}
