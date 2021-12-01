package database

import (
	"context"
	"database/sql"
)

type Repository interface {
	Save(context context.Context, user User) (User, error)
	GetAllData(context context.Context) ([]User, error)
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) Save(ctx context.Context, user User) (User, error) {
	sql := "INSERT INTO user(name, email, picture) VALUES (?,?,?)"
	_, err := r.DB.ExecContext(ctx, sql, user.Name, user.Email, user.Picture)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetAllData(ctx context.Context) ([]User, error) {
	script := "SELECT * FROM user"
	rows, err := r.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []User
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Name, &user.Email, &user.Picture)
		users = append(users, user)
	}

	return users, nil
}
