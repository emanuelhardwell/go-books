package store

import (
	"database/sql"
	"go-books/internal/model"
)

type Store interface {
	GetAll() ([]*model.Book, error)
	GetById(id int) (*model.Book, error)
	Create(book *model.Book) (*model.Book, error)
	Update(id int, book *model.Book) (*model.Book, error)
	Delete(id int) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &store{db: db}
}

func (s *store) GetAll() ([]*model.Book, error) {
	query := "select * from Books"

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*model.Book
	for rows.Next() {
		book := model.Book{}
		if err := rows.Scan(&book.Id, &book.Title, &book.Author); err != nil {
			return nil, err
		}

		books = append(books, &book)
	}

	return books, nil
}

func (s *store) GetById(id int) (*model.Book, error) {
	query := "select * from Books where id = ?"

	book := model.Book{}
	err := s.db.QueryRow(query, id).Scan(&book.Id, &book.Title, &book.Author)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (s *store) Create(book *model.Book) (*model.Book, error) {
	query := "insert into table Books (title, author) values (?, ?)"

	resp, err := s.db.Exec(query, book.Title, book.Author)
	if err != nil {
		return nil, err
	}

	id, err := resp.LastInsertId()
	if err != nil {
		return nil, err
	}

	book.Id = int(id)
	return book, nil
}

func (s *store) Update(id int, book *model.Book) (*model.Book, error) {
	query := "update Books set title= ? author= ? where id= ?"

	_, err := s.db.Exec(query, book.Title, book.Author, id)
	if err != nil {
		return nil, err
	}

	book.Id = id
	return book, nil
}

func (s *store) Delete(id int) error {
	query := "delete from Books where id= ?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
