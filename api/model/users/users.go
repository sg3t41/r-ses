package users

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sg3t41/api/model"
)

type Users struct {
	model.Model
}

func CreateUser(db sqlx.Ext) (uuid.UUID, error) {
	q := `INSERT INTO users DEFAULT VALUES RETURNING id`

	var id uuid.UUID
	switch conn := db.(type) {

	// トランザクションを未使用の場合
	case *sqlx.DB:
		if err := conn.QueryRowx(q).Scan(&id); err != nil {
			return uuid.Nil, err
		}
		return id, nil

	// トランザクションを使用する場合
	case *sqlx.Tx:
		if err := conn.QueryRowx(q).Scan(&id); err != nil {
			return uuid.Nil, err
		}
		return id, nil

	default:
		return uuid.Nil, fmt.Errorf("CreateUser unsupported type: %T", db)
	}
}
