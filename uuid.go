package events

import gouuid "github.com/satori/go.uuid"

func uuid() (string, error) {
	id, err := gouuid.NewV4()

	if err != nil {
		return "", err
	}

	return id.String(), nil
}
