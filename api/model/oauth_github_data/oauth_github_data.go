package oauth_github_data

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type OauthGithubData struct {
	ID               uuid.UUID `db:"id"`                 // ユーザープロバイダーID
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

func Create(db sqlx.Ext, data OauthGithubData) error {
	// INSERTクエリを作成
	query := `
        INSERT INTO oauth_github_data (username, email, avatar_url, profile_url, full_name, bio, location, company, account_created_at)
        VALUES (:username, :email, :avatar_url, :profile_url, :full_name, :bio, :location, :company, :account_created_at)
    `

	// dbの型に応じて処理を分ける
	switch conn := db.(type) {
	case *sqlx.DB:
		// DB接続の場合
		result, err := conn.NamedExec(query, data)
		if err != nil {
			return fmt.Errorf("failed to create OauthGithubData: %w", err)
		}

		// 影響を受けた行数を確認
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("failed to retrieve affected rows: %w", err)
		}

		if rowsAffected == 0 {
			return fmt.Errorf("no rows were affected")
		}

	case *sqlx.Tx:
		// トランザクションの場合
		result, err := conn.NamedExec(query, data)
		if err != nil {
			return fmt.Errorf("failed to create OauthGithubData: %w", err)
		}

		// 影響を受けた行数を確認
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return fmt.Errorf("failed to retrieve affected rows: %w", err)
		}

		if rowsAffected == 0 {
			return fmt.Errorf("no rows were affected")
		}

	default:
		return fmt.Errorf("unsupported db type: %T", db)
	}

	// 成功
	return nil
}

func Update(db sqlx.Ext, data OauthGithubData) (*int64, error) {
	// UPDATEクエリを作成
	query := `
		UPDATE oauth_github_data 
		SET 
			username = :username,
			email = :email,
			avatar_url = :avatar_url,
			profile_url = :profile_url,
			full_name = :full_name,
			bio = :bio,
			location = :location,
			company = :company,
			account_created_at = :account_created_at
		WHERE id = :id
		RETURNING id
	`

	var id int64

	// dbの型に応じて処理を分ける
	switch conn := db.(type) {
	case *sqlx.DB:
		// DB接続の場合
		err := conn.Get(&id, query, data)
		if err != nil {
			return nil, fmt.Errorf("failed to update OauthGithubData: %w", err)
		}

	case *sqlx.Tx:
		// トランザクションの場合
		err := conn.Get(&id, query, data)
		if err != nil {
			return nil, fmt.Errorf("failed to update OauthGithubData: %w", err)
		}

	default:
		return nil, fmt.Errorf("unsupported db type: %T", db)
	}

	// 更新されたIDを返す
	return &id, nil
}
