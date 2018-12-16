package events

import gouuid "github.com/satori/go.uuid"

func uuid() (string, error) {
	generated, err := gouuid.NewV4()

	if err != nil {
		return "", err
	}

	return generated.String(), nil
}
