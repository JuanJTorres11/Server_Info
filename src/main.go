package main

import (
	"fmt"
	"log"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/servers_info/:name", domainInfo)
	router.GET("/servers", listServers)

	log.Fatal(fasthttp.ListenAndServe(":8000", router.Handler))
}