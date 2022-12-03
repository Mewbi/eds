PRAGMA auto_vacuum = 1;

DROP TABLE IF EXISTS test_results;
CREATE TABLE test_results (
    id VARCHAR(36) NOT NULL PRIMARY KEY CHECK (LENGTH(id) > 0),
    name VARCHAR(250) NOT NULL CHECK (LENGTH(name) > 0), 
    email VARCHAR(250) NOT NULL CHECK (LENGTH(email) > 0), 
    responses VARCHAR(1000) NOT NULL CHECK(LENGTH(responses) > 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    confirmation BOOLEAN DEFAULT NULL
);

DROP TABLE IF EXISTS questions;
CREATE TABLE questions (
    id VARCHAR(36) NOT NULL PRIMARY KEY CHECK (LENGTH(id) > 0),
    content TEXT
);

DROP TABLE IF EXISTS questions_effectiveness;
CREATE TABLE questions_effectiveness (
    question_id VARCHAR(36) NOT NULL PRIMARY KEY CHECK (LENGTH(question_id) > 0),
    total INTEGER,
    effectiveness FLOAT
);

DROP TABLE IF EXISTS comments;
CREATE TABLE comments (
    id VARCHAR(36) NOT NULL PRIMARY KEY CHECK (LENGTH(id) > 0),
    name TEXT,
    comment TEXT
);
