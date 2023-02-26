package service

import "httpservice/internal/model"

type Storage interface {
	Add(e *model.User) error
	Read(id int) (*model.User, error)
	ReadAll() []*model.User
	Update(e *model.User) error
	Delete(id int) error
}

type Service interface {
	GetAll() []*model.User
	CreateUser(e *model.User) error
	GetById(id int) (*model.User, error)
	DeleteById(id int) error
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (r *service) GetAll() []*model.User {
	return r.storage.ReadAll()
}

func (r *service) CreateUser(u *model.User) error {
	return r.storage.Add(u)
}

func (r *service) GetById(id int) (*model.User, error) {
	return r.storage.Read(id)
}

func (r *service) DeleteById(id int) error {
	return r.storage.Delete(id)
}
