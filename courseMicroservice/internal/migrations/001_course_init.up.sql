CREATE TABLE IF NOT EXISTS courses
(
    id          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    description TEXT         NOT NULL,
    photo_url   TEXT         NOT NULL
);

CREATE TABLE IF NOT EXISTS video_lectures
(
    id        BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    course_id BIGINT REFERENCES courses (id) NOT NULL,
    url       TEXT                   NOT NULL
);

CREATE TABLE modules
(
    id        BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name      VARCHAR(255) NOT NULL,
    course_id BIGINT NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS lectures
(
    id        BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title     VARCHAR(100)                   NOT NULL,
    module_id BIGINT REFERENCES modules (id) NOT NULL,
    lecture   TEXT                           NOT NULL
);

CREATE TABLE IF NOT EXISTS tests (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    lecture_id BIGINT NOT NULL,
    FOREIGN KEY (lecture_id) REFERENCES lectures(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS questions (
    id         BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    text       TEXT    NOT NULL,
    is_correct BOOLEAN NOT NULL,
    test_id    BIGINT  NOT NULL,
    FOREIGN KEY (test_id) REFERENCES tests (id) ON DELETE CASCADE
);
