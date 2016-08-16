// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sessions

import (
	"github.com/valyala/fasthttp"
)

// Store is an interface for custom session stores.
//
// See CookieStore and FilesystemStore for examples.
type Store interface {
	// Get should return a session.
	Get(ctx *fasthttp.RequestCtx, name string) (*Session, error)

	// Save should persist session to the underlying store implementation.
	Save(ctx *fasthttp.RequestCtx, s *Session) error
}
