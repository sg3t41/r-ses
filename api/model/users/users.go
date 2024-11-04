package users

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sg3t41/api/model"
)

type Users struct {
	model.Model
}

func CreateById(db sqlx.Ext, users Users) (uuid.UUID, error) {
	q := `INSERT INTO users DEFAULT VALUES`
	id, err := model.Insert[Users](db, q, Users{})
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func Find() {

}

func IsExist() {

}

func Update() {

}

func Delete() {

}
