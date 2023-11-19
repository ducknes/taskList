package service

import (
	"fmt"
	"strconv"
)

func validateId(id string) (int, error) {
	result, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("невалидный id")
	}

	return result, nil
}
