package main

import (
	"fmt"
	"github.com/clevergo/context"
	"github.com/clevergo/sessions"
	"github.com/clevergo/sessions/stores/redisstore"
	"github.com/valyala/fasthttp"
)

var store sessions.Store

func redis(ctx *fasthttp.RequestCtx) {
	defer context.Clear(ctx)
	// Get session from store
	session, err := store.Get(ctx, "GOSESSION")
	if err != nil {
		ctx.SetBodyString(err.Error())
		return
	}
	// Save session.
	defer session.Save(ctx)

	if string(ctx.Path()) == "/set" {
		name := string(ctx.FormValue("name"))
		if len(name) > 0 {
			session.Values["name"] = name
			ctx.SetBodyString(fmt.Sprintf("name has been set as: %s\n", session.Values["name"]))
		} else {
			ctx.SetBodyString("No name specified.")
		}
		return
	}

	if name, ok := session.Values["name"]; ok {
		ctx.SetBodyString(fmt.Sprintf("name: %s\n", name))
		return
	}

	ctx.SetBodyString(`
	You should navigate to http://127.0.0.1:8080/set?name=yourname to set specified name.
	`)
}

func main() {
	redisStore, err := redisstore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	defer redisStore.Close()

	// Change session storage configuration for MaxAge = 10 days.
	redisStore.SetMaxAge(10 * 24 * 3600)

	store = redisStore

	fasthttp.ListenAndServe(":8080", redis)
}
