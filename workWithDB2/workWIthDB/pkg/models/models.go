package models

type Book struct {
	ID        int
	Name      string
	Price     int
	Author_id int
	Genre_id  int
}
type Genre struct {
	ID    int
	Genre string
}
type Author struct {
	ID   int
	Name string
}
