package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Proyek struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nama           string             `bson:"nama" json:"nama"`
	Lokasi         string             `bson:"lokasi" json:"lokasi"`
	TanggalMulai   string             `bson:"tanggal_mulai" json:"tanggal_mulai"`
	TanggalSelesai string             `bson:"tanggal_selesai" json:"tanggal_selesai"`
	KontraktorID   string             `bson:"kontraktor_id" json:"kontraktor_id"`
	Status         string             `bson:"status" json:"status"`
}
