package randresources

import "github.com/google/uuid"

func createUUID() string {
	u, _ := uuid.NewV7()

	return u.String()
}
