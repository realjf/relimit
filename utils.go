package relimit

import (
	"strings"

	"github.com/google/uuid"
)

func RandomName() string {
	name := uuid.New().String()
	name = strings.Replace(name, "-", "", -1)
	return name
}

func RandomSlice() string {
	slice := uuid.New().String()
	slice = strings.Replace(slice, "-", "", -1)
	return "/" + slice
}

func RandomGroup() string {
	group := uuid.New().String()
	group = strings.Replace(group, "-", "", -1)
	return group + ".slice"
}
