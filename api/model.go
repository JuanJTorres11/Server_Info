package api

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func UpdateDomain(servers []Server, simple_servers []ServerSSL, domain_name, grade, image, title string) (bool, string) {
	r_servers := []ServerSSL{}
	change := false
	ssl := grade
	time_now := time.Now()
	var time_sql time.Time

	db, err := sql.Open("postgres",
		"postgresql://user_servers:pwd@localhost:26257/domains?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.user_servers.key&sslcert=certs/client.user_servers.crt")
	if err != nil {
		log.Panicln("error connecting to the database: ", err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT time, sslgrade FROM domain WHERE domain='" + domain_name + "'")
	if err := row.Scan(&time_sql, &ssl); err != nil {
		if _, err := db.Exec(fmt.Sprintf(
			"INSERT INTO domain VALUES ('%s', '%s', '%s', '%s', '%s')",
			domain_name, grade, image, title, pq.FormatTimestamp(time_now))); err != nil {
			log.Panicln("Error inserting rows to the table domain ", err)
		}

		for i := 0; i < len(servers); i++ {
			if _, err := db.Exec(fmt.Sprintf(
				"INSERT INTO server VALUES ('%s', '%s', '%s', '%s', '%s')",
				domain_name, servers[i].Address, servers[i].SslGrade, servers[i].Country, servers[i].Owner)); err != nil {
				log.Panicln("Error inserting rows to the table server ", err)
			}
		}
		return change, ssl
	}

	if time_now.Sub(time_sql).Hours() >= 1 {
		rows_servers, err := db.Query("SELECT address, sslgrade FROM server WHERE domain='" + domain_name + "'")
		if err != nil {
			log.Panicln("Error retrieving the list of servers of the domain", domain_name, "Error message:", err)
		}
		defer rows_servers.Close()
		for rows_servers.Next() {
			var add, ssl string
			err := rows_servers.Scan(&add, &ssl)
			if err != nil {
				log.Panicln("Error reading rows of table server", err)
			}
			s := ServerSSL{add, ssl}
			r_servers = append(r_servers, s)
		}

		if !cmp.Equal(r_servers, servers) {
			change = true
			if _, err := db.Exec("DELETE FROM server WHERE domain='" + domain_name + "'"); err != nil {
				log.Panicln("Error reading rows of table server", err)
			}
			for i := 0; i < len(servers); i++ {
				if _, err := db.Exec(fmt.Sprintf(
					"INSERT INTO server VALUES ('%s', '%s', '%s', '%s', '%s')",
					domain_name, servers[i].Address, servers[i].SslGrade, servers[i].Country, servers[i].Owner)); err != nil {
					log.Panicln("Error inserting rows to the table server", err)
				}
			}
		}
	}

	return change, ssl
}

func GetDomains(list ServerList) ServerList {

	db, err := sql.Open("postgres",
		"postgresql://user_servers:pwd@localhost:26257/domains?ssl=true&sslmode=require&sslrootcert=certs/ca.crt&sslkey=certs/client.user_servers.key&sslcert=certs/client.user_servers.crt")
	if err != nil {
		log.Panicln("error connecting to the database: ", err)
	}
	defer db.Close()

	domains := []SimpleDomain{}

	rows_domains, err := db.Query("SELECT domain, sslgrade, logo, title, time FROM domain")
	if err != nil {
		log.Panicln("Error retrieving the list of domains ", err)
	}
	defer rows_domains.Close()
	for rows_domains.Next() {
		var dom, ssl, logo, title string
		var time time.Time
		err := rows_domains.Scan(&dom, &ssl, &logo, &title, &time)
		if err != nil {
			log.Panicln("Error reading rows of table domains", err)
		}
		servers := []Server{}
		rows_servers, err := db.Query("SELECT address, sslgrade, country, owner FROM server WHERE domain='" + dom + "'")
		if err != nil {
			log.Panicln("Error retrieving the list of servers of the domain ", dom, "Error message: ", err)
		}
		defer rows_servers.Close()
		for rows_servers.Next() {
			var add, ssl, country, owner string
			err := rows_servers.Scan(&add, &ssl, &country, &owner)
			if err != nil {
				log.Panicln("Error reading rows of table server", err)
			}
			s := Server{add, ssl, country, owner}
			servers = append(servers, s)
		}
		d := SimpleDomain{dom, servers, ssl, logo, title, time}
		domains = append(domains, d)
	}

	return ServerList{domains}
}
