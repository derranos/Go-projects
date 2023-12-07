package repository

import (
	"context"
	"workWithDB/models"
)

func (repo *PGRepo) GetBooks() ([]models.Book, error) {
	rows, err := repo.pool.Query(context.Background(), `
		select id,name,author_id,genre_id,price
		from books;
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
		insert into books (name, author_id, genre_id, price)
		values ($1, $2, $3, $4)`,
		item.Name,
		item.Author_id,
		item.Genre_id,
		item.Price,
	)
	return err
}
func (repo *PGRepo) DeleteByID(pos int) error {
	_, err := repo.pool.Exec(context.Background(), `
		delete from books where id = ($1)`,
		pos,
	)
	return err
}
func (repo *PGRepo) GetByID(pos int) (*models.Book, error) {
	var b models.Book
	err := repo.pool.QueryRow(context.Background(), `
		select id,name,author_id,genre_id,price
		from books
		where id = $1;`,
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
func (repo *PGRepo) ChangeByID(pos int, b models.Book) error {
	_, err := repo.pool.Exec(context.Background(), `
		update books set name = $1, author_id = $2,genre_id = $3, price = $4 where id = $5`,
		b.Name,
		b.Author_id,
		b.Genre_id,
		b.Price,
		pos,
	)
	return err
}
