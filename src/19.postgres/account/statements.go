package account

import (
	"database/sql"
	"fmt"
	"log"
)

func InsertAccount(account models.Account) string {
	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()
	// the inserted id will store in this id
	var id string
	// hardcoded
	stmt := `INSERT INTO accounts (owner, balance, currency) VALUES ($1, $2, $3) RETURNING id, owner, balance, currency, created_at`

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(stmt, account.Owner, account.Balance, account.Currency).Scan(&id)

	CheckError(err)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

func GetAccountId(id string) (models.account, error) {
	// hardcoded
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a accounts of models.Account type
	var user models.Account

	// create the select sql query
	stmt := `SELECT * FROM accounts WHERE userid=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

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
