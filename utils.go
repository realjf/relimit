package relimit

import "github.com/google/uuid"

func RandomName() string {
	return uuid.New().String()
}

func RandomSlice() string {
	slice := uuid.New().String()
	return "/" + slice
}

func RandomGroup() string {
	group := uuid.New().String()
	return group + ".slice"
}
