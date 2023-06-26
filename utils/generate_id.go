package utils

import (
	"github.com/kehiy/mushid"
)

func GenerateId() string {
	return MushId.New()
}