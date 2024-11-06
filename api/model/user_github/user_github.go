package user_github

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserGithub struct {
	UserProviderID   uuid.UUID `db:"user_provider_id"`   // ユーザープロバイダーID
	GithubID         int64     `db:"github_id"`          // GitHub ID
	Username         string    `db:"username"`           // ユーザー名
	Email            string    `db:"email"`              // メールアドレス
	AvatarURL        string    `db:"avatar_url"`         // アバターURL
	ProfileURL       string    `db:"profile_url"`        // プロフィールURL
	FullName         string    `db:"full_name"`          // フルネーム
	Bio              string    `db:"bio"`                // 自己紹介
	Location         string    `db:"location"`           // 所在地
	Company          string    `db:"company"`            // 会社名
	AccountCreatedAt time.Time `db:"account_created_at"` // アカウント作成日時
}

func ExistsByGithubId(db sqlx.Ext, githubId int64) (bool, error) {
	q := "SELECT EXISTS(SELECT 1 FROM user_github WHERE github_id=$1)"
	var exists bool

	switch conn := db.(type) {
	case *sqlx.DB:
		if err := conn.Get(&exists, q, githubId); err != nil {
			if err == sql.ErrNoRows {
				return false, nil
			}
			return false, err
		}

	case *sqlx.Tx:
		if err := conn.Get(&exists, q, githubId); err != nil {
			if err == sql.ErrNoRows {
				return false, nil
			}
			return false, err
		}

	default:
		return false, fmt.Errorf("ExistByGithubId unsupported type: %T", db)
	}

	return exists, nil
}
