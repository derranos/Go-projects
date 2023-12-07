package repository

import (
	"fmt"
	"log"

	"workWithDB/models"
)

const connStr = "postgres://postgres:psql@localhost:5432/postgres"

func Testing() {
	db, err := New(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	book := models.Book{Author_id: 1, Genre_id: 1, Name: "Идиот", Price: 198}
	err = db.CreateBook(book)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.DeleteByID(8)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.ChangeByID(9, book)
	if err != nil {
		log.Fatal(err.Error())
	}
	row, err := db.GetByID(9)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print(*row)
	data, err := db.GetBooks()
	if err != nil {
		fmt.Print("2\n")
		log.Fatal(err.Error())
	}
	for _, item := range data {
		fmt.Printf("%d: %s, price: %d, author_id: %d, genre_id: %d\n", item.ID, item.Name, item.Price, item.Author_id, item.Genre_id)
	}
}
