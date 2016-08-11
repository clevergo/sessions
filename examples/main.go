package main

import (
	"fmt"
	"github.com/clevergo/context"
	"github.com/clevergo/sessions"
	"github.com/clevergo/sessions/stores/redisstore"
	"github.com/valyala/fasthttp"
)

var (
	store sessions.Store
)

func MyHandler(ctx *fasthttp.RequestCtx) {
	defer context.Clear(ctx)
	// Get session from store
	session, _ := store.Get(ctx, "GOSESSION")
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
	var err error
	store, err = redisstore.NewRediStore(10, "tcp", ":6379", "", []byte("secret-key"))
	if err != nil {
		panic(err)
	}
	// store = sessions.NewCookieStore([]byte("something-very-secret"))
	fasthttp.ListenAndServe(":8080", MyHandler)
}
