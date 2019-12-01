package events

import gouuid "github.com/google/uuid"

func uuid() (string, error) {
	generated, err := gouuid.NewRandom()

	if err != nil {
		return "", err
	}

	return generated.String(), nil
}
