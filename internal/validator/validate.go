package validator

import (
	"errors"

	"github.com/murphy6867/server/utils"
)

const MaxChirpLength = 140

var profaneWords = map[string]bool{
	"kerfuffle": true,
	"sharbert":  true,
	"fornax":    true,
}

func ValidateChirp(body string) (string, error) {
	if len(body) > MaxChirpLength {
		return "", errors.New("chirp is too long")
	}

	cleaned := utils.FilterWord(profaneWords, body, "****")
	return cleaned, nil
}
