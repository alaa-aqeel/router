package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Login(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "access_token!\n")
}

func Profile(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}

func UpdateProfile(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hi, %s, %s!\n", ctx.UserValue("name"), ctx.UserValue("word"))
}

func main() {
	r := router.New()
	r.POST("/login", Login)
	group := r.Group("/profile/")
	group.Middleware(func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			token := ctx.Request.Header.Peek("Authorization")
			if token != nil {
				next(ctx)
				return
			}
			ctx.Error("unauthorization", 401)
		}
	}) // call before hanlder
	group.GET("/", Profile)
	group.POST("/update", UpdateProfile)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
