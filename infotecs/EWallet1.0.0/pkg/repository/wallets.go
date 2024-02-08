package repository

import (
	"context"
	"infotecs/pkg/models"
	"time"

	"github.com/google/uuid"
)

func (repo *PGRepo) CreateWallet() (*models.Wallet, error) {
	id := uuid.New().String()
	_, err := repo.pool.Exec(context.Background(), `
		insert into Wallets (id, balance)
		values ($1,100);`,
		id,
	)
	if err != nil {
		return nil, err
	}
	ans, err := repo.GetWallet(id)
	if err != nil {
		return nil, err
	}
	return ans, nil
}
func (repo *PGRepo) GetWallet(id string) (*models.Wallet, error) {
	ans := models.Wallet{}
	err := repo.pool.QueryRow(context.Background(), `
		select ID, Balance
		from Wallets
		where ID = $1;`,
		id,
	).Scan(
		&ans.ID,
		&ans.Balance,
	)
	if err != nil {
		return nil, err
	}
	return &ans, nil
}
func (repo *PGRepo) SendMoney(id, to string, amount float64) error {
	_, err := repo.pool.Exec(context.Background(), `
		update Wallets
		set Balance = Balance - $1
		where ID = $2;`,
		amount,
		id,
	)
	if err != nil {
		return err
	}
	_, err = repo.pool.Exec(context.Background(), `
		update Wallets
		set Balance = Balance + $1
		where ID = $2;`,
		amount,
		to,
	)
	if err != nil {
		return err
	}
	return nil
}
func (repo *PGRepo) SaveTransfer(id, to string, amount float64) error {
	timeRFC3339 := time.Now().Format(time.RFC3339)
	_, err := repo.pool.Exec(context.Background(), `
		insert into History (Event_time, IDfrom, IDto, Amount)
		values ($1,$2,$3,$4);`,
		timeRFC3339,
		id,
		to,
		amount,
	)
	if err != nil {
		return err
	}
	return nil
}
func (repo *PGRepo) GetHistory(id string) ([]models.SavedTransfer, error) {
	rows, err := repo.pool.Query(context.Background(), `
		select Event_time, IDfrom, IDto, Amount
		from History
		where IDto = $1 or IDfrom = $1;`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := []models.SavedTransfer{}
	for rows.Next() {
		savedTransfer := models.SavedTransfer{}
		var transferTime time.Time
		err = rows.Scan(
			&transferTime,
			&savedTransfer.From,
			&savedTransfer.TO,
			&savedTransfer.Amount,
		)
		if err != nil {
			return nil, err
		}
		savedTransfer.Time = transferTime.Format(time.RFC3339)
		data = append(data, savedTransfer)
	}
	return data, nil
}
