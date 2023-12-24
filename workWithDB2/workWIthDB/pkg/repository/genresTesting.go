package repository

import (
	"fmt"
	"log"

	"workWithDB/pkg/models"
)

func GenresTesing() {
	db, err := New(ConnStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	genre := models.Genre{Genre: "Повесть"}
	err = db.CreateGenre(genre)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.DeleteGenreByID(3)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.ChangeGenreByID(1, genre)
	if err != nil {
		log.Fatal(err.Error())
	}
	row, err := db.GetGenreByID(2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print(*row, "\n")
	data, err := db.GetGenres()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, item := range data {
		fmt.Printf("%d: %s\n", item.ID, item.Genre)
	}
}
