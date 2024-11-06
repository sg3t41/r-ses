package users

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sg3t41/api/config"
	"github.com/sg3t41/api/model"
)

func init() {
	config.Setup("../../config/app.ini")
	model.Setup()
}

func TestCreateUserWithOutTx(t *testing.T) {
	t.Logf("ユーザー登録のテストを実施します。")
	id, err := CreateUser(model.DB)
	if err != nil {
		t.Errorf(err.Error())
	}
	if id == uuid.Nil {
		t.Errorf("inesrt error")
	}
	t.Logf("ユーザー登録が成功しました。\n")
	t.Logf("[該当ユーザーID] %v\n", id)
}
