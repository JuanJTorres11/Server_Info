CREATE USER IF NOT EXISTS user_servers;
CREATE DATABASE IF NOT EXISTS domains;
GRANT ALL ON DATABASE domains TO user_servers;
CREATE TABLE IF NOT EXISTS domains.domain (
    domain VARCHAR(50) PRIMARY KEY,
    sslgrade VARCHAR(2),
    logo VARCHAR(50),
    title VARCHAR(50),
    time TIMESTAMP
    );
CREATE TABLE IF NOT EXISTS domains.server (
    domain VARCHAR(50),
    address VARCHAR(39) NOT NULL REFERENCES domains.domain (domain) ON DELETE CASCADE,
    sslgrade VARCHAR(2),
    country VARCHAR(30),
    owner VARCHAR(50),
    PRIMARY KEY (address, domain)
);