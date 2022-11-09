package tools

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePasswords(hashedPassword string, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	} else if err != nil {
		panic(err)
	}
	return nil
}
