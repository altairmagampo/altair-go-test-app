package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID              primitive.ObjectID `bson: "_id, omitempty" json: "id"` //Se respeta la convención de mongo en _id, se omite si no viene el ID, y se va a formato json como id
	Nombre          string             `bson: "nombre, omitempty" json: "nombre"`
	Apellidos       string             `bson: "apellidos, omitempty" json: "apellidos"`
	FechaNacimiento time.Time          `bson: "fechaNacimiento, omitempty" json: "fechaNacimiento"`
	Email           string             `bson: "email" json: "email"`
	Password        string             `bson: "password" json: "password,omitempty"`
	Avatar          string             `bson: "avatar" json: "avatar,omitempty"`
	Banner          string             `bson: "banner" json: "banner,omitempty"`
	Biografia       string             `bson: "biografia" json: "biografia,omitempty"`
	SitioWeb        string             `bson: "sitioWeb" json: "sitioWeb,omitempty"`
}
