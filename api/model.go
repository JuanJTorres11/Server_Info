package api

import (
	"log"
	"database/sql"
	
    _ "github.com/lib/pq"
)

func GetInfoServers() {
    // Connect to the "bank" database.
    db, err := sql.Open("postgres",
        "postgresql://root@localhost:26257/domains?sslcert=certs%2Fclient.root.crt&sslkey=certs%2Fclient.user_servers.key&sslmode=verify-full&sslrootcert=certs%2Fca.crt")
    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }
    defer db.Close()

}