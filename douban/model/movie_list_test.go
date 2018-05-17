package model

import (
	"fmt"
	"testing"
)

func TestGetAllMovie(t *testing.T) {
	m, err := GetAllMovie()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}
