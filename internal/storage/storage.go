package storage

import (
	"fmt"
	"httpservice/internal/model"
)

type WorkStorage struct {
	Id   int
	repo map[int]*model.User
}

func NewWorkStorage() *WorkStorage {
	return &WorkStorage{
		repo: map[int]*model.User{},
		Id:   1000,
	}
}

func (ws *WorkStorage) Add(u *model.User) error {
	ws.Id++
	u.ID = ws.Id
	if _, ok := ws.repo[u.ID]; !ok {
		ws.repo[u.ID] = u
		return nil
	}
	return fmt.Errorf("service already exist")
}

func (ws *WorkStorage) Read(id int) (*model.User, error) {
	if ws.contains(id) {
		return ws.repo[id], nil
	}
	return nil, fmt.Errorf("service not exist")
}

func (ws *WorkStorage) ReadAll() []*model.User {
	values := make([]*model.User, 0, len(ws.repo))
	for _, val := range ws.repo {
		values = append(values, val)
	}
	return values
}

func (ws *WorkStorage) Update(u *model.User) error {
	if !ws.contains(u.ID) {
		return fmt.Errorf("service not exist")
	}
	ws.repo[u.ID] = u
	return nil
}

func (ws *WorkStorage) Delete(id int) error {
	if ws.contains(id) {
		delete(ws.repo, id)
		return nil
	}
	return fmt.Errorf("service not exit")
}

func (ws *WorkStorage) contains(id int) bool {
	for userID := range ws.repo {
		if userID == id {
			return true
		}
	}
	return false
}
