package repository

import (
	"context"
	"workWithDB/pkg/models"
)

func (repo *PGRepo) GetBooks() ([]models.Book, error) {
	rows, err := repo.pool.Query(context.Background(), `
		select ID,Name,Author_id,Genre_id,Price
		from Books;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []models.Book
	for rows.Next() {
		var item models.Book
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Author_id,
			&item.Genre_id,
			&item.Price,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}
func (repo *PGRepo) CreateBook(item models.Book) error {
	_, err := repo.pool.Exec(context.Background(), `
		insert into Books (Name, Author_id, Genre_id, price)
		values ($1, $2, $3, $4)`,
		item.Name,
		item.Author_id,
		item.Genre_id,
		item.Price,
	)
	return err
}
func (repo *PGRepo) DeleteBookByID(pos int) error {
	_, err := repo.pool.Exec(context.Background(), `
		delete from Books where ID = ($1)`,
		pos,
	)
	return err
}
func (repo *PGRepo) GetBookByID(pos int) (*models.Book, error) {
	var b models.Book
	err := repo.pool.QueryRow(context.Background(), `
		select ID,Name,Author_id,Genre_id,Price
		from Books
		where ID = $1;`,
		pos,
	).Scan(
		&b.ID,
		&b.Name,
		&b.Author_id,
		&b.Genre_id,
		&b.Price,
	)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
func (repo *PGRepo) ChangeBookByID(pos int, b models.Book) error {
	_, err := repo.pool.Exec(context.Background(), `
		update Books set Name = $1, Author_id = $2,Genre_id = $3, Price = $4 where ID = $5`,
		b.Name,
		b.Author_id,
		b.Genre_id,
		b.Price,
		pos,
	)
	return err
}
