package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Laporan struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProyekID    string             `bson:"proyek_id" json:"proyek_id"`
	Judul       string             `bson:"judul" json:"judul"`
	Tanggal     string             `bson:"tanggal" json:"tanggal"`
	Isi         string             `bson:"isi" json:"isi"`
	DisusunOleh string             `bson:"disusun_oleh" json:"disusun_oleh"`
}
