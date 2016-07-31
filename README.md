# sessions package for Go(go programing language) fasthttp
[![GoDoc](https://godoc.org/github.com/clevergo/sessions?status.svg)](https://godoc.org/github.com/clevergo/sessions) [![Build Status](https://travis-ci.org/clevergo/sessions.png?branch=master)](https://travis-ci.org/clevergo/sessions)

**This repo fork from [gorilla/sessions](https://github.com/gorilla/sessions)**

gorilla/sessions provides cookie and filesystem sessions and infrastructure for
custom session backends.

The key features are:

* Simple API: use it as an easy way to set signed (and optionally
  encrypted) cookies.
* Built-in backends to store sessions in cookies or the filesystem.
* Flash messages: session values that last until read.
* Convenient way to switch session persistency (aka "remember me") and set
  other attributes.
* Mechanism to rotate authentication and encryption keys.
* Multiple sessions per request, even using different backends.
* Interfaces and infrastructure for custom session backends: sessions from
  different stores can be retrieved and batch-saved using a common API.

Let's start with an example that shows the sessions API in a nutshell:

```go
package main

import (
	"fmt"
	"github.com/clevergo/sessions"
	"github.com/valyala/fasthttp"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

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
	fasthttp.ListenAndServe(":8080", MyHandler)
}
```

First we initialize a session store calling `NewCookieStore()` and passing a
secret key used to authenticate the session. Inside the handler, we call
`store.Get()` to retrieve an existing session or a new one. Then we set some
session values in session.Values, which is a `map[interface{}]interface{}`.
And finally we call `session.Save()` to save the session in the response.

See also [examples](examples)

Important Note:  you need to **context.Clear(ctx)** as or else you will leak memory!

## License

BSD licensed. See the LICENSE file for details.
