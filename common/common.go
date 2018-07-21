package common

import "crypto/rand"

type CallerID string

const runeOptions = "abcdefghijklmnopqrstuvwxyz0123456789"

func RandASCIIBytes(desiredLength int) []byte {
	output := make([]byte, desiredLength)
	randomness := make([]byte, desiredLength)

	_, err := rand.Read(randomness)
	if err != nil {
		panic(err)
	}

	for eachIndex := range output {
		random := uint8(randomness[eachIndex])
		randomPos := random % uint8(len(runeOptions))
		output[eachIndex] = runeOptions[randomPos]
	}

	return output
}
