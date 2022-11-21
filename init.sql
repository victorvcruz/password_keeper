CREATE TABLE users (
	id BIGINT UNIQUE NOT NULL,
	name VARCHAR NOT NULL,
	email VARCHAR NOT NULL,
	master_password VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);
CREATE UNIQUE INDEX user_email_index ON users(email, deleted_at) WHERE deleted_at IS NULL;
CREATE SEQUENCE serial START 1;

CREATE TABLE folders (
                         id BIGINT UNIQUE NOT NULL,
                         user_id BIGINT NOT NULL,
                         name VARCHAR NOT NULL,
                         created_at TIMESTAMP NOT NULL,
                         updated_at TIMESTAMP NOT NULL,
                         deleted_at TIMESTAMP NULL,
                         PRIMARY KEY (id),
                         FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE vaults (
   id BIGINT UNIQUE NOT NULL,
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
   PRIMARY KEY (id),
   FOREIGN KEY (user_id) REFERENCES users (id),
   FOREIGN KEY (folder_id) REFERENCES folders (id)
);

CREATE TABLE reports (
	id BIGINT UNIQUE NOT NULL,
    user_id BIGINT NULL,
	vault_id BIGINT NULL,
	action VARCHAR,
	description VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (vault_id) REFERENCES vaults (id)
);


