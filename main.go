package main

import (
	"fmt"
	"log"

	"github.com/JuanJTorres11/Server_Info/api"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func main() {
	router := fasthttprouter.New()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("There was an error", r)
		}
	}()

	router.GET("/", Index)
	router.GET("/servers_info/:name", api.DomainInfo)
	router.GET("/servers", api.ListServers)

	log.Fatal(fasthttp.ListenAndServe(":4000", router.Handler))
}
