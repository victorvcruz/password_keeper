CREATE TABLE users (
                       id BIGSERIAL UNIQUE NOT NULL,
                       name VARCHAR NOT NULL,
                       email VARCHAR NOT NULL,
                       master_password VARCHAR NOT NULL,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL,
                       deleted_at TIMESTAMP NULL,
                       PRIMARY KEY (id)
);
CREATE UNIQUE INDEX user_email_index ON users(email, deleted_at) WHERE deleted_at IS NULL;