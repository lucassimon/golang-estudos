package databases

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo estrutura para conexao com o mongo
type Mongo struct {
	db       *mongo.Client
	mongoCtx context.Context
}

// Connect metodo para conctar ao mysql
func (m *Mongo) Connect() error {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	m.mongoCtx = ctx
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	db, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(err)
	}

	m.Ping()

	// return the connection
	m.db = db
	return nil
}

// Ping testa se a conexao está ok
func (m *Mongo) Ping() {
	// check the connection
	err := m.db.Ping(m.mongoCtx, nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

}

// GetNow busca a data atual no servidor postgres
func (m *Mongo) GetNow() (*time.Time, error) {

	t := &time.Time{}

	return t, nil
}

// Close fecha a conexao do mysql
func (m *Mongo) Close() error {
	return m.db.Disconnect(m.mongoCtx)
}
