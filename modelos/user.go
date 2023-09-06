package modelos

import "time"

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

func (usurious *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	usurious.Id = id
	usurious.Name = name
	usurious.CreatedAt = createdAt
	usurious.Status = status
}
