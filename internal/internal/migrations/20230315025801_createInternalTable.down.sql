CREATE TABLE internal (
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