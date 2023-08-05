package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func ComparePassword(hashedPassword []byte, entredPassword string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(entredPassword))
}
