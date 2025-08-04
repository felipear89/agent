package user

import "fmt"

type Repository interface {
	FindAll() ([]User, error)
	FindByID(id int) (*User, error)
	Create(user User) (*User, error)
	Update(user User) (*User, error)
	Delete(id int) error
	FindByEmail(email string) (*User, error)
}

type InMemoryRepository struct {
	users []User
}

func (r *InMemoryRepository) FindByEmail(email string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user with email %s not found", email)
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users: []User{
			{ID: 1, Name: "John Doe", Email: "john@example.com", Username: "johndoe"},
			{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Username: "janesmith"},
			{ID: 3, Name: "Felipe Rodrigues", Email: "felipear89@gmail.com", Username: "felipe", Password: "123456"},
		},
	}
}

func (r *InMemoryRepository) FindAll() ([]User, error) {
	return r.users, nil
}

func (r *InMemoryRepository) FindByID(id int) (*User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found", id)
}

func (r *InMemoryRepository) Create(user User) (*User, error) {
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return &user, nil
}

func (r *InMemoryRepository) Update(user User) (*User, error) {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
			return &r.users[i], nil
		}
	}
	return nil, nil
}

func (r *InMemoryRepository) Delete(id int) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return nil
}
