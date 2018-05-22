package common

import (
	"fmt"
	"testing"
)

func TestAddMovieId(t *testing.T) {
	val, err := AddMovieId(45)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
}

func TestIsExistMovieId(t *testing.T) {
	ok, err := IsExistMovieId(453)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ok)
}
