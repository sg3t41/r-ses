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

func TestCreateUserNoTx(t *testing.T) {
	t.Logf("ユーザー登録のテストを実施します。")
	id, err := Create(model.DB)
	if err != nil {
		t.Errorf(err.Error())
	}
	if id == uuid.Nil {
		t.Errorf("inesrt error")
	}
	t.Logf("ユーザー登録が成功しました。\n")
	t.Logf("[該当ユーザーID] %v\n", id)
}

func TestCreateUserWithTx(t *testing.T) {
	t.Logf("ユーザー登録のテスト（トランザクション使用）を実施します。")

	// トランザクションを開始
	tx, err := model.DB.Begin()
	if err != nil {
		t.Errorf("トランザクション開始エラー: %v", err)
		return
	}
	defer tx.Rollback() // テスト後にロールバック

	// トランザクション内でユーザー作成
	id, err := Create(tx)
	if err != nil {
		t.Errorf("ユーザー作成エラー: %v", err)
		return
	}
	if id == uuid.Nil {
		t.Errorf("ユーザー作成時に無効なIDが返されました")
	}

	t.Logf("ユーザー登録が成功しました（トランザクション）。\n")
	t.Logf("[該当ユーザーID] %v\n", id)

	// トランザクションをコミット（今回はテストなのでロールバックされるが、一般的にコミットする）
	err = tx.Commit()
	if err != nil {
		t.Errorf("トランザクションコミットエラー: %v", err)
	}
}
