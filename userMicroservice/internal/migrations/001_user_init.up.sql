CREATE TABLE IF NOT EXISTS users(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    login text NOT NULL,
    password bytea NOT NULL,
    name TEXT,
    surname TEXT,
    email TEXT,
    phone TEXT,
    place_work TEXT,
    position TEXT,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz
);

CREATE TABLE IF NOT EXISTS tokens(
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    user_id uuid REFERENCES users(id) NOT NULL,
    token TEXT NOT NULL,
    refresh TEXT NOT NULL,
    expires_at timestamptz NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz
);

CREATE TABLE IF NOT EXISTS roles (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS users_roles (
    user_id uuid REFERENCES users(id) NOT NULL,
    role_id BIGINT REFERENCES roles(id) NOT NULL,
    PRIMARY KEY (user_id, role_id)
);


CREATE TABLE IF NOT EXISTS permissions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    deleted_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS roles_permissions (
    role_id BIGINT REFERENCES roles(id) NOT NULL,
    permission_id BIGINT REFERENCES permissions(id) NOT NULL,
    PRIMARY KEY (role_id, permission_id)
);
