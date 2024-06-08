CREATE TABLE IF NOT EXISTS users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    firstname VARCHAR(100) NOT NULL,
    lastname VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NUll,
    login VARCHAR(100) NOT NULL,
    password BYTEA NOT NUll,
    haspremium BOOLEAN DEFAULT FALSE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS tokens (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT REFERENCES users(id) NOT NULL,
    token bytea NOT NULL,
    refresh bytea NOT NULL,
    expires_at timestamptz NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS roles (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS permissions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS users_roles (
    user_id BIGINT REFERENCES users(id) NOT NULL,
    role_id BIGINT REFERENCES roles(id) NOT NULL,
    PRIMARY KEY (user_id, role_id)
);

CREATE TABLE IF NOT EXISTS roles_permissions (
    role_id BIGINT REFERENCES roles(id) NOT NULL,
    permission_id BIGINT REFERENCES permissions(id) NOT NULL,
    PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE IF NOT EXISTS photo_url  (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT REFERENCES users(id) NOT NULL,
    url VARCHAR(300) NOT NULL
);

--permissions CRUD users
INSERT INTO permissions(name) VALUES('POST user');
INSERT INTO permissions(name) VALUES('GET user');
INSERT INTO permissions(name) VALUES('PUT user');
INSERT INTO permissions(name) VALUES('DELETE user');

--roles
INSERT INTO roles(name) VALUES('admin');
INSERT INTO roles(name) VALUES('user');

--admin permissions
INSERT INTO roles_permissions VALUES(1, 1);
INSERT INTO roles_permissions VALUES(1, 2);
INSERT INTO roles_permissions VALUES(1, 3);
INSERT INTO roles_permissions VALUES(1, 4);

--user permissions
INSERT INTO roles_permissions VALUES(2, 1);
INSERT INTO roles_permissions VALUES(2, 2);
INSERT INTO roles_permissions VALUES(2, 3);
INSERT INTO roles_permissions VALUES(2, 4);

--admin
INSERT INTO users(firstname, lastname, email, login, password) VALUES('admin','admin', 'admin@k.ru', 'admin', 'admin');
INSERT INTO users_roles VALUES(1, 1);
