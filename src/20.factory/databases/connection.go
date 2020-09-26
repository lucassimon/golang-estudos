package databases

import "time"

// DBConnection interface para criacao de conexoes
type DBConnection interface {
	Connect() error
	GetNow() (*time.Time, error)
	Close() error
}
