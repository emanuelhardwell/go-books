package service

import (
	"errors"
	"go-books/internal/model"
	"go-books/internal/store"
)

type Service struct {
	store store.Store
}

func New(s store.Store) *Service {
	return &Service{store: s}
}

func (s *Service) GetAllCtrl() ([]*model.Book, error) {
	return s.store.GetAll()
}

func (s *Service) GetByIdCtrl(id int) (*model.Book, error) {
	return s.store.GetById(id)
}

func (s *Service) CreateCtrl(book *model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("el campo title no puede estar vacio")
	}

	return s.store.Create(book)
}

func (s *Service) UpdateCtrl(id int, book *model.Book) (*model.Book, error) {
	if book.Title == "" {
		return nil, errors.New("el campo title no puede estar vacio")
	}

	return s.store.Update(id, book)
}

func (s *Service) DeleteCtrl(id int) error {
	if err := s.store.Delete(id); err != nil {
		return err
	}
	return nil
}
