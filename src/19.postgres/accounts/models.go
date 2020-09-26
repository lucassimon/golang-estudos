package accounts

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	databases "github.com/lucassimon/golang-estudos/src/19.postgres/databases"
	utils "github.com/lucassimon/golang-estudos/src/19.postgres/utils"
)

// User schema of the user table
type Account struct {
	ID        string    `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

// Create account
func Create(a Account) string {
	// create the postgres db connection
	db := databases.CreateConnection()
	// close the db connection
	defer db.Close()
	// the inserted id will store in this id
	var id string
	// hardcoded
	stmt := `INSERT INTO accounts (owner, balance, currency) VALUES ($1, $2, $3) RETURNING id, owner, balance, currency, created_at`

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(stmt, a.Owner, a.Balance, a.Currency).Scan(&id)

	utils.CheckError(err)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

// Get Fetch account by id
func Get(id string) (Account, error) {
	// hardcoded
	// create the postgres db connection
	db := databases.CreateConnection()
	// close the db connection
	defer db.Close()

	// create a accounts of models.Account type
	var account Account

	// create the select sql query
	stmt := `SELECT * FROM accounts WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(stmt, id)

	// unmarshal the row object to user
	err := row.Scan(&account.ID, &account.Owner, &account.Balance, &account.Currency)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return account, nil
	case nil:
		return account, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return account, err
}
