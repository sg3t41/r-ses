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
		Username:         "testuser222",
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

	data := getTestData()

	err := Create(model.DB, data)
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}
	t.Logf("ユーザー登録が成功しました。")
}
