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
    id SERIAL PRIMARY KEY,
    name provider_type NOT NULL UNIQUE
);

CREATE TABLE oauth_tokens (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    provider_id INTEGER REFERENCES oauth_providers(id) ON DELETE CASCADE,
    access_token VARCHAR(500) NOT NULL,
    refresh_token VARCHAR(500),
    expires_at TIMESTAMP,
    refresh_expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_provider (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    provider_id INTEGER REFERENCES oauth_providers(id) ON DELETE CASCADE
);

CREATE TABLE user_github (
    user_provider_id UUID REFERENCES user_provider(id) ON DELETE CASCADE,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100),
    avatar_url VARCHAR(255),
    profile_url VARCHAR(255),
    full_name VARCHAR(100),
    bio TEXT,
    location VARCHAR(100),
    company VARCHAR(100),
    account_created_at TIMESTAMP,
    PRIMARY KEY (user_provider_id)
);

CREATE TABLE user_linkedin (
    user_provider_id UUID REFERENCES user_provider(id) ON DELETE CASCADE,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100),
    location VARCHAR(100),
    industry VARCHAR(100),
    profile_url VARCHAR(255),
    PRIMARY KEY (user_provider_id)
);

-- Add OAuth provider (Github & LinkedIn)
INSERT INTO oauth_providers (name) VALUES
    ('GITHUB'),
    ('LINKEDIN')
ON CONFLICT (name) DO NOTHING;
