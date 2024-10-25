package util

import (
	"errors"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type HashHelperIntf interface {
	HashPassword(pwd string) (string, error)
	CheckPassword(pwd string, hash []byte) (bool, error)
}

type HashHelperImpl struct{}

func (h *HashHelperImpl) HashPassword(pwd string) (string, error) {
	costStr := os.Getenv("HASH_COST")

	cost, err := strconv.Atoi(costStr)
	if err != nil {
		return "", errors.New("error while reading env")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h *HashHelperImpl) CheckPassword(pwd string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pwd))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
