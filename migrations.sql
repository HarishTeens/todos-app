-- To set up DB run this


-- Create table
CREATE TABLE IF NOT EXISTS users (
    id BIGINT  AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(55) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS todos (
    id BIGINT  AUTO_INCREMENT NOT NULL PRIMARY KEY,
    todo VARCHAR(255) NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Insert data
INSERT INTO users (name) VALUES ('John Doe');
INSERT INTO users (name) VALUES ('Jane Doe');
INSERT INTO users (name) VALUES ('John Smith');

INSERT INTO todos (todo, user_id) VALUES ('Buy milk', 1);
INSERT INTO todos (todo, user_id) VALUES ('Buy bread', 1);
