CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- UUIDでローカルIDを生成
    github_id BIGINT UNIQUE, -- GitHubユーザーの一意のID
    username VARCHAR(50) NOT NULL, -- GitHubのユーザー名
    email VARCHAR(100), -- GitHubのメールアドレス（公開されている場合のみ）
    avatar_url VARCHAR(255), -- プロフィール画像のURL
    profile_url VARCHAR(255), -- GitHubプロフィールページのURL
    full_name VARCHAR(100), -- フルネーム
    bio TEXT, -- 自己紹介
    location VARCHAR(100), -- 所在地
    company VARCHAR(100), -- 所属
    github_created_at TIMESTAMP, -- GitHubアカウント作成日
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE user_tokens (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    access_token VARCHAR(500) NOT NULL, -- JWT形式のアクセストークン
    refresh_token VARCHAR(500) NOT NULL, -- 非JWT形式のリフレッシュトークン
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP, -- アクセストークンの有効期限
    refresh_expires_at TIMESTAMP, -- リフレッシュトークンの有効期限
    is_revoked BOOLEAN DEFAULT FALSE -- トークンが無効化されているかどうか
);

