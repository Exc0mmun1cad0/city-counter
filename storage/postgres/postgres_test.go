package postgres

import (
	"fmt"
	"testing"
)

func TestDB(t *testing.T) {
	storage, err := NewPostgresStorage("localhost", "5432", "postgres", "postgres", "postgres", "disable")
	fmt.Println(err == nil, err)

	val, err := storage.CountCitiesInCountry("BRA")
	fmt.Println(err == nil, err)
	fmt.Println(val)
}
