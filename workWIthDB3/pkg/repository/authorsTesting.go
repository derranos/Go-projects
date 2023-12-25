package repository

import (
	"fmt"
	"log"

	"workWithDB/pkg/models"
)

func AuthorsTesing() {
	db, err := New(ConnStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	author := models.Author{Name: "Бунин"}
	err = db.CreateAuthor(author)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.DeleteAuthorByID(3)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.ChangeAuthorByID(1, author)
	if err != nil {
		log.Fatal(err.Error())
	}
	row, err := db.GetAuthorByID(2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print(*row, "\n")
	data, err := db.GetAuthors()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, item := range data {
		fmt.Printf("%d: %s\n", item.ID, item.Name)
	}
}
