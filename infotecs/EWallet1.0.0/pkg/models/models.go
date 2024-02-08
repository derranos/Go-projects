package models

type Wallet struct {
	ID      string
	Balance float64
}
type Transfer struct {
	TO     string
	Amount float64
}
type SavedTransfer struct {
	Time   string
	From   string
	TO     string
	Amount float64
}
