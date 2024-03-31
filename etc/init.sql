CREATE TABLE IF NOT EXISTS materials_map (
    material_id BIGINT REFERENCES materials(id) NOT NULL,
    map_id BIGINT REFERENCES map(id) NOT NULL,
    PRIMARY KEY (material_id, map_id)
);

CREATE TABLE IF NOT EXISTS users_calculate (
    calculate_id BIGINT REFERENCES calculate(id) NOT NULL,
    user_id BIGINT REFERENCES users(id) NOT NULL,
    PRIMARY KEY (calculate_id, user_id)
);

CREATE TABLE IF NOT EXISTS users_courses (
    user_id BIGINT REFERENCES users(id) NOT NULL,
    course_id BIGINT REFERENCES courses(id) NOT NULL,
    PRIMARY KEY (user_id, course_id),
    progress_course INTEGER CHECK (progress_course >= 0 AND progress_course <= 0) NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    start_date timestamptz NOT NULL DEFAULT now(),
    completed_at timestamptz   
);

CREATE TABLE IF NOT EXISTS users_materials (
    material_id BIGINT REFERENCES materials(id) NOT NULL,
    user_id BIGINT REFERENCES users(id) NOT NULL,
    PRIMARY KEY (material_id, user_id),
    favorite BOOLEAN NOT NULL DEFAULT FALSE
);
