package model

import (
	"fmt"
)

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Friends Friends `json:"friends"`
}

type Friends []*User

func NewUser() *User {
	return &User{}
}

func (u *User) FriendContains(id int) error {
	for _, fs := range u.Friends {
		if fs.ID == id {
			return fmt.Errorf("already in a friend list")
		}
	}
	return nil
}

func (u *User) AddFriend(t *User) error {
	if err := u.FriendContains(t.ID); err != nil {
		return err
	}
	u.Friends = append(u.Friends, t)
	return nil
}

func (u *User) getFriendIndex(id int) int {
	for i, user := range u.Friends {
		if user.ID == id {
			return i
		}
	}
	return -1
}

func (u *User) DeleteFriend(id int) error {
	i := u.getFriendIndex(id)
	if i < 0 {
		return fmt.Errorf("not in friend list")
	}
	u.Friends, u.Friends[len(u.Friends)-1] = append(u.Friends[:i], u.Friends[i+1:]...), nil
	return nil
}

func (u *User) String() string {
	return fmt.Sprintf(
		"ID: %s\nName: %s\nAge: %d\n",
		u.ID,
		u.Name,
		u.Age)
}
