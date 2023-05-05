CREATE TABLE IF NOT EXISTS users (
    user_id BIGINT generated always as identity,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,

    CONSTRAINT pk_user_id PRIMARY KEY(user_id)
);