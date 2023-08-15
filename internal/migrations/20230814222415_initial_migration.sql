-- +goose Up
-- +goose StatementBegin
CREATE TABLE users_info (
    id VARCHAR(36) PRIMARY KEY,
    uun VARCHAR(50),
    username VARCHAR(50) UNIQUE,
    hashed_password VARCHAR(255),
    created_at DATETIME,
    email VARCHAR(100),
    totp_secret VARCHAR(255),
    role VARCHAR(20),
    status VARCHAR(20)
);

CREATE TABLE sessions (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36),
    session_exp DATETIME,
    FOREIGN KEY (user_id) REFERENCES users_info(id)
);

CREATE TABLE tracks_info (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36),
    action VARCHAR(255),
    timestamp DATETIME,
    FOREIGN KEY (user_id) REFERENCES users_info(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users_info;
DROP TABLE sessions;
DROP TABLE tracks_info;
-- +goose StatementEnd
