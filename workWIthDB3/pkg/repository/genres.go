package repository

import (
	"context"
	"workWithDB/pkg/models"
)

func (repo *PGRepo) GetGenres() ([]models.Genre, error) {
	rows, err := repo.pool.Query(context.Background(), `
		select ID,Genre 
		from Genres;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []models.Genre
	for rows.Next() {
		var item models.Genre
		err = rows.Scan(
			&item.ID,
			&item.Genre,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	return data, nil
}
func (repo *PGRepo) CreateGenre(item models.Genre) error {
	_, err := repo.pool.Exec(context.Background(), `
		insert into Genres (Genre)
		values ($1)`,
		item.Genre,
	)
	return err
}
func (repo *PGRepo) DeleteGenreByID(pos int) error {
	_, err := repo.pool.Exec(context.Background(), `
		delete from Genres where ID = ($1)`,
		pos,
	)
	return err
}
func (repo *PGRepo) GetGenreByID(pos int) (*models.Genre, error) {
	var b models.Genre
	err := repo.pool.QueryRow(context.Background(), `
		select ID,Genre
		from Genres
		where ID = $1;`,
		pos,
	).Scan(
		&b.ID,
		&b.Genre,
	)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
func (repo *PGRepo) ChangeGenreByID(pos int, b models.Genre) error {
	_, err := repo.pool.Exec(context.Background(), `
		update Genres set Genre = $1 where ID = $2`,
		b.Genre,
		pos,
	)
	return err
}
