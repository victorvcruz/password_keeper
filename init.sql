CREATE SCHEMA user_service
CREATE TABLE IF NOT EXISTS user_service.users (
	id BIGSERIAL UNIQUE NOT NULL,
	name VARCHAR NOT NULL,
	email VARCHAR NOT NULL,
	master_password VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);
CREATE UNIQUE INDEX user_email_index ON user_service.users(email, deleted_at) WHERE deleted_at IS NULL;

CREATE SCHEMA folder_service
CREATE TABLE IF NOT EXISTS folder_service.folders (
                         id BIGSERIAL UNIQUE NOT NULL,
                         user_id BIGINT NOT NULL,
                         name VARCHAR NOT NULL,
                         created_at TIMESTAMP NOT NULL,
                         updated_at TIMESTAMP NOT NULL,
                         deleted_at TIMESTAMP NULL
);

CREATE SCHEMA vault_service
CREATE TABLE IF NOT EXISTS vault_service.vaults (
   id BIGSERIAL UNIQUE NOT NULL,
   user_id BIGINT NOT NULL,
   folder_id BIGINT NULL,
   username VARCHAR NOT NULL,
   name VARCHAR NOT NULL,
   password VARCHAR NOT NULL,
   url VARCHAR NOT NULL,
   notes VARCHAR NOT NULL,
   favorite VARCHAR NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL,
   deleted_at TIMESTAMP NULL,
   PRIMARY KEY (id)
);

CREATE SCHEMA report_service
CREATE TABLE IF NOT EXISTS report_service.reports (
	id BIGSERIAL UNIQUE NOT NULL,
    user_id BIGINT NULL,
	vault_id BIGINT NULL,
	action VARCHAR,
	description VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);

CREATE SCHEMA auth_service
CREATE TABLE IF NOT EXISTS auth_service.auth (
	id BIGSERIAL UNIQUE NOT NULL,
        user_id BIGINT NULL,
	token VARCHAR,
    created_at TIMESTAMP NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);

CREATE SCHEMA internal_service
CREATE TABLE internal_service.internal (
   id BIGSERIAL UNIQUE NOT NULL,
   service VARCHAR NOT NULL,
   token VARCHAR NOT NULL,
   password VARCHAR NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL,
   deleted_at TIMESTAMP NULL,
   PRIMARY KEY (id)
);
CREATE UNIQUE INDEX internal_service_index ON internal(service, deleted_at) WHERE deleted_at IS NULL;

