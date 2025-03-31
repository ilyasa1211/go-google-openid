package repositories

import (
	"database/sql"

	"github.com/ilyasa1211/go-google-openid/internal/application/dto"
	"github.com/ilyasa1211/go-google-openid/internal/core/domain/user"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindAll() []*user.User {
	rows, err := r.Db.Query("SELECT * FROM users")

	if err != nil {
		panic(err)
	}

	var users []*user.User

	for rows.Next() {
		var user user.User
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

		users = append(users, &user)
	}

	return users
}
func (r *UserRepository) FindById(id string) *user.User {
	rows, err := r.Db.Query("SELECT * FROM users WHERE id = $1", id)

	if err != nil {
		panic(err)
	}

	var user user.User

	if rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	}

	return &user
}

func (r *UserRepository) FindByEmail(email string) *user.User {
	rows, err := r.Db.Query("SELECT * FROM users WHERE email = $1", email)

	if err != nil {
		panic(err)
	}

	var user user.User

	if rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	}

	return &user
}

func (r *UserRepository) Create(u *dto.CreateUserRequest) error {
	if _, err := r.Db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", u.Name, u.Email, u.Password); err != nil {
		return err
	}

	return nil
}
func (r *UserRepository) UpdateById(id string, u *dto.UpdateUserRequest) error {
	if _, err := r.Db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, id); err != nil {
		return err
	}

	return nil
}
func (r *UserRepository) DeleteById(id string) error {
	if _, err := r.Db.Exec("DELETE FROM users WHERE id = $1", id); err != nil {
		return err
	}

	return nil
}
