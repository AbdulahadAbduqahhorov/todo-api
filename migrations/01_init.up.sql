CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    username varchar(255) NOT NULL UNIQUE,
    password_hash varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS todo_lists (
    id BIGSERIAL NOT NULL UNIQUE,
    title varchar(255) NOT NULL,
    description varchar(255) 
);

CREATE TABLE IF NOT EXISTS user_lists (
    id BIGSERIAL NOT NULL UNIQUE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    list_id INT REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL
     
);

CREATE TABLE IF NOT EXISTS todo_items (
    id BIGSERIAL NOT NULL UNIQUE,
    title varchar(255) NOT NULL,
    description varchar(255) ,
    done BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS list_items (
    id BIGSERIAL NOT NULL UNIQUE,
    item_id INT REFERENCES todo_items(id) ON DELETE CASCADE NOT NULL,
    list_id INT REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL
     
);
