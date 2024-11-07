package user_oauth_providers

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// UserOauthProviders モデル
type UserOauthProviders struct {
	ID         uuid.UUID `db:"id"`
	UserID     uuid.UUID `db:"user_id"`
	ProviderID uuid.UUID `db:"provider_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	DeletedAt  time.Time `db:"deleted_at"`
}

// Create - user_oauth_providers を作成する関数
func Create(db any, userID, providerID uuid.UUID) (uuid.UUID, error) {
	q := `
		INSERT INTO user_oauth_providers (
			user_id, provider_id
		) 
		VALUES (
			$1, $2
		) 
		RETURNING id
	`

	var id uuid.UUID
	switch conn := db.(type) {
	case *sqlx.DB:
		if err := conn.QueryRowx(q, userID, providerID).Scan(&id); err != nil {
			return uuid.Nil, err
		}
		return id, nil

	// トランザクションを使った場合
	case *sqlx.Tx:
		if err := conn.QueryRowx(q, userID, providerID).Scan(&id); err != nil {
			return uuid.Nil, err
		}
		return id, nil

	default:
		return uuid.Nil, fmt.Errorf("CreateUserOauthProvider unsupported type: %T", db)
	}
}

