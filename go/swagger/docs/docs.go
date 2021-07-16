// Package docs Documentation of our awesome API.
//
//     Title: テストアプリのAPI Document
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:8080
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package docs

import "github.com/pinkumohikan/gocha-gocha/go/swagger/api"

// swagger:route POST /greet PostGreet
// 挨拶をするすごいAPI
// responses:
//   200: greetResponse

// swagger:parameters PostGreet
type greetRequestWrapper struct {
	// in:body
	Body api.GreetRequest
}

// swagger:response greetResponse
type greetResponseWrapper struct {
	//in:body
	Body api.GreetResponse
}
