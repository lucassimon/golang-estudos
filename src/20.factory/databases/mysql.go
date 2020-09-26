package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Mysql estrutura para conexao com o mysql
type Mysql struct {
	db *sql.DB
}

// Connect metodo para conctar ao mysql
func (m *Mysql) Connect() error {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))

	if err != nil {
		panic(err)
	}

	m.Ping()

	// return the connection
	m.db = db
	return nil
}

// Ping checa a conexao com o banco
func (m *Mysql) Ping() {
	// check the connection
	err := m.db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

// GetNow busca a data atual no servidor mysql
func (m *Mysql) GetNow() (*time.Time, error) {

	t := &time.Time{}
	err := m.db.QueryRow("SELECT CURDATE()").Scan(t)

	if err != nil {
		fmt.Printf("Erro ao buscar a data no servidor do mysql")
		return nil, err
	}

	return t, nil
}

// Close fecha a conexao do mysql
func (m *Mysql) Close() error {
	return m.db.Close()
}
