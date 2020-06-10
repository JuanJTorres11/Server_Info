package main

import (
	"fmt"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type Server struct {
	Address		string		`json:"address"`
	SslGrade	string		`json:"ssl_grade"`
	Country		string		`json:"country"`
	Owner		string		`json:"owner"`
}

type Domain struct {
	Servers 			[]Server	`json:"endpoints"`
	ServerChanged		bool		`json:"servers_changed"`
	SSLGrade			string		`json:"ssl_grade"`
	PreviousSslGrade	string		`json:"previous_ssl_grade"`
	Logo				string		`json:"logo"`
	Title				string		`json:"title"`
	Is_down				string		`json:"is_down"`
}

type ServerList struct {
	Items []Domain		`json:"items"`
}

func domainInfo(ctx *fasthttp.RequestCtx) {
	//domain := ctx.UserValue("name")
	server1 := Server{"addres1", "A", "Col", "Owner"}
	server2 := Server{"addres2", "A", "Col", "Owner"}
	servers := []Server{server1,server2}
	domain := Domain {servers, true, "A", "A", "TestLogo", "TestTitle", "Down"}
	json.NewEncoder(ctx).Encode(domain)
}

func listServers(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "TODO!\n")
}