package bd

import (
	"golang.org/x/crypto/bcrypt"
)

func EncriptarPassword(pass string) (string, error) {
	//Entre mayo el costo, mayor cantidad de pasadas, pero es mas demandante de tiempo. Se recomienda para un user normal un costo de 6, para usuario más delicados un 8
	costo := 8
	//Algoritmo 2 elevado al costo, en este caso 256 veces 2 a la 8, una ecriptación dentro de otra y otra hasta llegar al costo (2^8)
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
