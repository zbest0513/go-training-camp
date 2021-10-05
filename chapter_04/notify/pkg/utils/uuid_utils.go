package utils

import guid "github.com/google/uuid"

func UUID() string {
	return guid.New().String()
}
