package users

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/lib/pq"
	databases "github.com/lucassimon/golang-estudos/src/19.postgres/databases"
	utils "github.com/lucassimon/golang-estudos/src/19.postgres/utils"
)

func makeCreateStatement() string {
	return `INSERT INTO users (name, age, active) VALUES ($1, $2, $3) RETURNING id, name, age, active, created_at`
}

// Create user
func Create(user User) error {
	// create the postgres db connection
	db := databases.CreateConnection()
	// close the db connection
	defer db.Close()

	intNull := sql.NullInt64{}
	// boolNull := sql.NullBool{}

	// hardcoded
	query := makeCreateStatement()

	// execute the sql statement
	stmt, err := db.Prepare(query)
	defer stmt.Close()
	utils.CheckError(err)

	if user.Age == 0 {
		intNull.Valid = false
	} else {
		intNull.Int64 = int64(user.Age)
	}

	row, err := stmt.Exec(user.Name, user.Age, user.Active)
	utils.CheckError(err)

	affected, _ := row.RowsAffected()

	if affected != 1 {
		return errors.New("Error: Nothing was created")
	}

	return nil

}

func makeGetStatement() string {
	return `SELECT * FROM users WHERE id=$1`
}

// Get Fetch user by id
func Get(id string) (User, error) {

	// create the postgres db connection
	db := databases.CreateConnection()
	// close the db connection
	defer db.Close()

	// create a users of models.users type
	var user User

	// create the select sql query
	query := makeGetStatement()

	// execute the sql statement
	row := db.QueryRow(query, id)

	// unmarshal the row object to user
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Active, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return user, err
}

func makeAllStatement() string {
	return `SELECT * FROM users`
}

// All users
func All(limit, offset int64) (users []User, err error) {

	// create the postgres db connection
	db := databases.CreateConnection()
	// close the db connection
	defer db.Close()

	// create the select sql query
	query := makeAllStatement()

	updatedAtNull := pq.NullTime{}
	deletedAtNull := pq.NullTime{}
	ageNull := sql.NullInt64{}
	boolNull := sql.NullBool{}

	rows, err := db.Query(query)
	defer rows.Close()
	utils.CheckError(err)

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Name, &ageNull, &boolNull, &user.CreatedAt, &updatedAtNull, &deletedAtNull)
		utils.CheckError(err)

		user.UpdatedAt = updatedAtNull.Time
		user.DeletedAt = deletedAtNull.Time
		user.Age = int16(ageNull.Int64)
		user.Active = boolNull.Bool

		users = append(users, user)
	}

	return users, nil
}

func makeUpdateStatement() string {
	return `UPDATE users SET name = $1, age = $2, active = $3, updated_at = now() where id = $4`
}

// Update user
func Update(user User) error {
	// create the postgres db connection
	db := databases.CreateConnection()
	// close the db connection
	defer db.Close()

	// create the select sql query
	query := makeUpdateStatement()

	stmt, err := db.Prepare(query)
	defer stmt.Close()
	utils.CheckError(err)

	row, err := stmt.Exec(user.Name, user.Age, user.Active, user.ID)
	utils.CheckError(err)

	affected, _ := row.RowsAffected()

	if affected != 1 {
		return errors.New("Error: Nothing was created")
	}

	return nil
}

func makeDeleteStatement() string {
	return `UPDATE users SET active = false, deleted_at = now() where id = $1`
}

// Delete user
func Delete(user User) error {
	// create the postgres db connection
	db := databases.CreateConnection()
	// close the db connection
	defer db.Close()

	// create the select sql query
	query := makeDeleteStatement()

	stmt, err := db.Prepare(query)
	defer stmt.Close()
	utils.CheckError(err)

	row, err := stmt.Exec(user.ID)
	utils.CheckError(err)

	affected, _ := row.RowsAffected()

	if affected != 1 {
		return errors.New("Error othing was created")
	}

	return nil
}
