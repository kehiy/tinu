package utils

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateID() (string, error) {
	return gonanoid.New()
}
