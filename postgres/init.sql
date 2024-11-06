CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE provider_type AS ENUM ('GITHUB', 'LINKEDIN');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE user_tokens (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    jwt_access_token VARCHAR(500) NOT NULL,
    refresh_token VARCHAR(500) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP,
    refresh_expires_at TIMESTAMP,
    is_revoked BOOLEAN DEFAULT FALSE
);

CREATE TABLE oauth_providers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name provider_type NOT NULL UNIQUE
);

CREATE TABLE oauth_tokens (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    provider_id UUID REFERENCES oauth_providers(id) ON DELETE CASCADE,
    access_token VARCHAR(500) NOT NULL,
    refresh_token VARCHAR(500),
    expires_at TIMESTAMP,
    refresh_expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_provider (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    provider_id UUID REFERENCES oauth_providers(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, provider_id)
);

-- GitHubのOAuthデータ
CREATE TABLE oauth_github_data (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100),
    avatar_url VARCHAR(255),
    profile_url VARCHAR(255),
    full_name VARCHAR(100),
    bio TEXT,
    location VARCHAR(100),
    company VARCHAR(100),
    account_created_at TIMESTAMP
);

CREATE TABLE user_github (
    user_provider_id UUID REFERENCES user_provider(id) ON DELETE CASCADE,
    github_id BIGINT NOT NULL UNIQUE,
    user_github_data UUID REFERENCES oauth_github_data(id) ON DELETE CASCADE,
    PRIMARY KEY (user_provider_id)
);

-- LinkedInのOAuthデータ
CREATE TABLE oauth_linkedin_data (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100),
    location VARCHAR(100),
    industry VARCHAR(100),
    profile_url VARCHAR(255),
    headline VARCHAR(255),
    summary TEXT,
    public_profile_url VARCHAR(255),
    account_created_at TIMESTAMP
);

CREATE TABLE user_linkedin (
    user_provider_id UUID REFERENCES user_provider(id) ON DELETE CASCADE,
    linkedin_id VARCHAR(50) NOT NULL UNIQUE,
    user_linkedin_data UUID REFERENCES oauth_linkedin_data(id) ON DELETE CASCADE,
    PRIMARY KEY (user_provider_id)
);

-- Add OAuth provider (Github & LinkedIn)
INSERT INTO oauth_providers (name) VALUES
    ('GITHUB'),
    ('LINKEDIN')
ON CONFLICT (name) DO NOTHING;

