package api

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func UpdateDomain(servers []Server, domain_name, grade, image, title string) (change bool, prev_grade string) {
	db, err := sql.Open("postgres",
		"postgresql://root@localhost:26257/domains?sslcert=certs%2Fclient.root.crt&sslkey=certs%2Fclient.user_servers.key&sslmode=verify-full&sslrootcert=certs%2Fca.crt")
	if err != nil {
		log.Panicln("error connecting to the database: ", err)
	}
	defer db.Close()

	return false, "NA"
}
