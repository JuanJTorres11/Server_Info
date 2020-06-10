CREATE USER IF NOT EXISTS user_servers;
CREATE DATABASe domains;
GRANT ALL ON DATABASE domains TO user_servers;
CREATE TABLE domains.servers (domain STRING, servers STRING, sslgrade STRING, time TIMESTAMP, PRIMARY KEY (domain, servers));