PRAGMA auto_vacuum = 1;

DROP TABLE IF EXISTS test_results;
CREATE TABLE test_results (
    id VARCHAR(36) NOT NULL PRIMARY KEY CHECK (LENGTH(id) > 0),
    name VARCHAR(250) NOT NULL CHECK (LENGTH(name) > 0), 
    email VARCHAR(250) NOT NULL CHECK (LENGTH(email) > 0), 
    responses VARCHAR(1000) NOT NULL CHECK(LENGTH(responses) > 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    confirmation BOOLEAN DEFAULT NULL,
);
