package repository

import (
	"context"
	"workWithDB/models"
)

func (repo *PGRepo) GetAuthors() ([]models.Author, error) {
	rows, err := repo.pool.Query(context.Background(), `
		select ID,Name 
		from Authors;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []models.Author
	for rows.Next() {
		var item models.Author
		err = rows.Scan(
			&item.ID,
			&item.Name,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}
func (repo *PGRepo) CreateAuthor(item models.Author) error {
	_, err := repo.pool.Exec(context.Background(), `
		insert into Authors (Name)
		values ($1)`,
		item.Name,
	)
	return err
}
func (repo *PGRepo) DeleteAuthorByID(pos int) error {
	_, err := repo.pool.Exec(context.Background(), `
		delete from Authors where ID = ($1)`,
		pos,
	)
	return err
}
func (repo *PGRepo) GetAuthorByID(pos int) (*models.Author, error) {
	var b models.Author
	err := repo.pool.QueryRow(context.Background(), `
		select ID,Name
		from Authors
		where ID = $1;`,
		pos,
	).Scan(
		&b.ID,
		&b.Name,
	)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
func (repo *PGRepo) ChangeAuthorByID(pos int, b models.Author) error {
	_, err := repo.pool.Exec(context.Background(), `
		update Authors set Name = $1 where ID = $2`,
		b.Name,
		pos,
	)
	return err
}
