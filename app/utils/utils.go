package utils

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

func ConvertSlice[E uint16](in []interface{}) (out []E) {

	out = make([]E, 0, len(in))
	for _, v := range in {
		out = append(out, E(v.(float64)))
	}
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}

func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}
