package scope

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Scope struct {
	Database *sqlx.DB
}

func NewScope() Scope {
	return Scope{
		Database: nil,
	}
}
func (s Scope) String() string {
	return fmt.Sprintf("Scope: {Database:%s }", s.Database)
}
