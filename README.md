# Sessions package for Go(go programing language) fasthttp
[![GoDoc](https://godoc.org/github.com/clevergo/sessions?status.svg)](https://godoc.org/github.com/clevergo/sessions) [![Build Status](https://travis-ci.org/clevergo/sessions.png?branch=master)](https://travis-ci.org/clevergo/sessions).

**This repo fork from [gorilla/sessions](https://github.com/gorilla/sessions)**.

This package just support **fasthttp**, if you looking for a sessions package for net/http, the [gorilla/sessions](https://github.com/gorilla/sessions) is better.

clevergo/sessions provides cookie, filesystem and redis sessions and infrastructure for
custom session backends.

The key features are:

* Simple API: use it as an easy way to set signed (and optionally
  encrypted) cookies.
* Built-in backends to store sessions in cookies, the filesystem or redis.
* Flash messages: session values that last until read.
* Convenient way to switch session persistency (aka "remember me") and set
  other attributes.
* Mechanism to rotate authentication and encryption keys.
* Interfaces and infrastructure for custom session backends: sessions from
  different stores can be retrieved and batch-saved using a common API.

## Install
```
go get github.com/clevergo/sessions
```

## Examples
```
go run $GOPATH/src/github.com/clevergo/sessions/examples/main.go
```
See also [examples](examples)

## License

BSD licensed. See the LICENSE file for details.
