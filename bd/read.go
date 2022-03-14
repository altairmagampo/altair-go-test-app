package bd

import (
	"context"
	"time"

	"github.com/altairmagampo/altair-go-test-app/models"
	"go.mongodb.org/mongo-driver/bson"
)

func BuscarUsuarioPorEmail(mail string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoCN.Database("test-go-aa")
	collection := bd.Collection("usuarios")

	//La codición de mi find. M es un formato que devuelve un MapStream
	condicion := bson.M{"email": mail}
	var resultado models.Usuario
	err := collection.FindOne(ctx, condicion).Decode(&resultado)
	//Converir el OnjectID (_ID) a un hexadecimal para devolverlo
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
