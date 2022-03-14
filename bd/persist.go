package bd

import (
	"context"
	"time"

	"github.com/altairmagampo/altair-go-test-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Registrar(u models.Usuario) (string, bool, error) {

	//El contexto el background es el contexto TODO, al cual con esta linea le estamos agregando este contexto que manjeará el timeout
	//Timeout de 15 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	//Se setea aquí pero se ejecuta al final de la función, es como un finally
	defer cancel() //Cancel, cancela el timeout, cancela el context.WithTimeout

	bd := MongoCN.Database("test-go-aa")
	collection := bd.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)
	result, err := collection.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
