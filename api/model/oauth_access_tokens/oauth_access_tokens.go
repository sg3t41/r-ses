package oauth_access_tokens

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type OauthAccessTokens struct {
	ID                  uuid.UUID `db:"id"`
	UserOauthProviderID uuid.UUID `db:"user_oauth_provider_id"`
	AccessToken         string    `db:"access_token"`
	ExpiresAt           time.Time `db:"expires_at"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
	DeletedAt           time.Time `db:"deleted_at"`
}

// Create - oauth_access_tokens を作成する
func Create(db any, userOauthProviderID uuid.UUID, accessToken string, expires *time.Time) (uuid.UUID, error) {
	q := `
		INSERT INTO oauth_access_tokens (
			user_oauth_provider_id, access_token, expires_at
		) 
		VALUES (
			$1, $2, $3
		) 
		RETURNING id
	`

	var expiresAt *time.Time

	if expires.IsZero() {
		expiresAt = nil
	} else {
		expiresAt = expires
	}

	var id uuid.UUID
	switch conn := db.(type) {
	// DB を使用する場合
	case *sqlx.DB:
		if err := conn.QueryRowx(q, userOauthProviderID, accessToken).Scan(&id); err != nil {
			return uuid.Nil, err
		}
		return id, nil

	// トランザクションを使用する場合
	case *sqlx.Tx:
		if err := conn.QueryRowx(q, userOauthProviderID, accessToken, &expiresAt).Scan(&id); err != nil {
			return uuid.Nil, err
		}
		return id, nil

	default:
		return uuid.Nil, fmt.Errorf("CreateOauthAccessToken unsupported type: %T", db)
	}
}
