package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/jonlaing/htmlmeta"
	"github.com/valyala/fasthttp"
	"github.com/xellio/whois"
)

type Server struct {
	Address  string `json:"address"`
	SslGrade string `json:"ssl_grade"`
	Country  string `json:"country"`
	Owner    string `json:"owner"`
}

type ServerSSL struct {
	Address  string `json:"ipAddress"`
	SslGrade string `json:"grade"`
}

type responseSSL struct {
	Endpoints []ServerSSL `json:"endpoints"`
}
type Domain struct {
	Servers          []Server `json:"endpoints"`
	ServerChanged    bool     `json:"servers_changed"`
	SSLGrade         string   `json:"ssl_grade"`
	PreviousSslGrade string   `json:"previous_ssl_grade"`
	Logo             string   `json:"logo"`
	Title            string   `json:"title"`
	Is_down          bool     `json:"is_down"`
}

type ServerList struct {
	Items []Domain `json:"items"`
}

func DomainInfo(ctx *fasthttp.RequestCtx) {
	domain_name := fmt.Sprint(ctx.UserValue("name"))
	servers := []Server{}
	is_down := false
	simple_servers := []ServerSSL{}

	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain_name)
	if err != nil {
		log.Panicln("There was an error trying to connect to ssl labs")
	}
	defer response.Body.Close()

	r, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panicln("There was an error trying to read the JSON from ssl labs")
	}

	ssl_json := responseSSL{}
	err = json.Unmarshal(r, &ssl_json)
	if err != nil {
		log.Panicln("There was an error trying to parse the JSON from ssl labs")
	}

	worst_grade := ssl_json.Endpoints[0].SslGrade
	for i := 0; i < len(ssl_json.Endpoints); i++ {
		address := ssl_json.Endpoints[i].Address
		grade := ssl_json.Endpoints[i].SslGrade
		s := ServerSSL{address, grade}
		simple_servers = append(simple_servers, s)
		ip := net.ParseIP(address)
		res, err := whois.QueryIP(ip)
		if err != nil {
			log.Panicln("There was an error trying to obtain the whois of the server " + address)
		}
		owner := strings.Join(res.Output["OrgName"], ",")
		country := strings.Join(res.Output["Country"], ",")
		server := Server{address, grade, country, owner}
		servers = append(servers, server)
		worst_grade = compareGrades(grade, worst_grade)
	}

	res, err := http.Get("http://" + domain_name)
	if err != nil {
		is_down = true
		log.Panicln("There was an error trying to connect to the domain " + domain_name)
	}
	defer res.Body.Close()
	meta := htmlmeta.Extract(res.Body)
	title := meta.OGTitle
	image := meta.OGImage

	server_changed, prev_grade := UpdateDomain(servers, simple_servers, domain_name, worst_grade, image, title)
	domain := Domain{servers, server_changed, worst_grade, prev_grade, image, title, is_down}
	json.NewEncoder(ctx).Encode(domain)

}

func ListServers(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "TODO!\n")
}

func compareGrades(grade1, grade2 string) string {
	if grade1 == "A+" {
		grade1 = "1"
	} else if grade1 == "A-" {
		grade1 = "2"
	}

	if grade2 == "A+" {
		grade2 = "1"
	} else if grade2 == "A-" {
		grade2 = "2"
	}

	var r string
	if grade1 < grade2 {
		r = grade2
	} else {
		r = grade1
	}

	if r == "1" {
		r = "A+"
	} else if r == "2" {
		r = "A-"
	}
	return r
}
