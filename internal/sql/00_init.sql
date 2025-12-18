CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid(),
    email VARCHAR(100) UNIQUE NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE table_primary (

);

CREATE TABLE table_secondary (

);
