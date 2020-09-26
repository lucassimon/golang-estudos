package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Postgres estrutura para conexao com o postgres
type Postgres struct {
	db *sql.DB
}

// Connect metodo para conctar ao mysql
func (p *Postgres) Connect() error {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	p.Ping()

	fmt.Println("Successfully connected!")
	// return the connection
	p.db = db
	return nil
}

// Ping checa a conexao com o banco
func (p *Postgres) Ping() {
	// check the connection
	err := p.db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

// GetNow busca a data atual no servidor postgres
func (p *Postgres) GetNow() (*time.Time, error) {

	t := &time.Time{}
	err := p.db.QueryRow("SELECT CURRENT_DATE").Scan(t)

	if err != nil {
		fmt.Printf("Erro ao buscar a data no servidor do mysql")
		return nil, err
	}

	return t, nil
}

// Close fecha a conexao do mysql
func (p *Postgres) Close() error {
	return p.db.Close()
}
