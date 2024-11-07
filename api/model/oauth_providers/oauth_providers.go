package oauth_providers

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ProviderType 型の定義
type ProviderType string

const (
	GitHub   ProviderType = "GITHUB"
	LinkedIn ProviderType = "LINKEDIN"
)

// IsValid メソッド：ProviderTypeが有効かどうかを確認
func (p ProviderType) IsValid() bool {
	switch p {
	case GitHub, LinkedIn:
		return true
	}
	return false
}

// OauthProviders モデル
type OauthProviders struct {
	ID   uuid.UUID    `db:"id"`
	Name ProviderType `db:"name"`
}

// GetProviderID 名前からIDを取得する関数
func GetProviderID(db any, providerName ProviderType) (uuid.UUID, error) {
	if !providerName.IsValid() {
		return uuid.Nil, fmt.Errorf("invalid provider name: %s", providerName)
	}

	q := `SELECT id FROM oauth_providers WHERE name = $1`

	var id uuid.UUID
	switch conn := db.(type) {
	// DBを使った場合
	case *sqlx.DB:
		if err := conn.QueryRowx(q, providerName).Scan(&id); err != nil {
			return uuid.Nil, err
		}
		return id, nil

	// トランザクションを使った場合
	case *sqlx.Tx:
		if err := conn.QueryRowx(q, providerName).Scan(&id); err != nil {
			return uuid.Nil, err
		}
		return id, nil

	default:
		return uuid.Nil, fmt.Errorf("GetProviderID unsupported type: %T", db)
	}
}
