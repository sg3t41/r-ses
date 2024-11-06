package oauth_github_data

import (
	"github.com/google/uuid"
	"github.com/sg3t41/api/config"
	"github.com/sg3t41/api/model"
	"testing"
	"time"
)

func init() {
	// 設定とモデルの初期化
	config.Setup("../../config/app.ini")
	model.Setup()
}

// テスト用のデータ作成
func getTestData() OauthGithubData {
	return OauthGithubData{
		ID:               uuid.New(),
		Username:         "testuser",
		Email:            "test@example.com",
		AvatarURL:        "http://example.com/avatar.jpg",
		ProfileURL:       "http://example.com/profile",
		FullName:         "Test User",
		Bio:              "This is a bio",
		Location:         "Tokyo, Japan",
		Company:          "Example Corp",
		AccountCreatedAt: time.Now(),
	}
}

func TestCreate(t *testing.T) {
	t.Logf("Create OauthGithubDataのテストを実施します。")

	// テスト用データ
	data := getTestData()

	// データベース接続
	id, err := Create(model.DB, data)
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}
	if id == nil || *id == uuid.Nil {
		t.Errorf("Create failed: invalid ID returned")
	}
	t.Logf("ユーザー登録が成功しました。")
	t.Logf("[該当ユーザーID] %v\n", id)
}

func TestUpdate(t *testing.T) {
	t.Logf("Update OauthGithubDataのテストを実施します。")

	// まずCreateでデータを作成してから更新
	data := getTestData()

	// CreateでIDを取得
	id, err := Create(model.DB, data)
	if err != nil {
		t.Errorf("Create failed: %v", err)
		return
	}

	// 更新データの準備
	data.ID = *id
	data.Username = "updateduser"
	data.Email = "updated@example.com"

	// データベース接続
	updatedID, err := Update(model.DB, data)
	if err != nil {
		t.Errorf("Update failed: %v", err)
	}
	if updatedID == nil || *updatedID == 0 {
		t.Errorf("Update failed: invalid ID returned")
	}

	t.Logf("ユーザー更新が成功しました。")
	t.Logf("[更新されたユーザーID] %v\n", updatedID)
}

