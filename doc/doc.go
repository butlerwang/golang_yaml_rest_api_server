// Copyright 2019 Butler Wang. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Hulu CDNWatch APIs
//
// This documentation describes Hulu CDNWatch APIs 
//
//     Schemes: http, https
//     BasePath: /v1
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Butler Wang <butler.wang@hulu.com>
//     Host: localhost
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta
package main

// swagger:parameters UpdateUserResponseWrapper
type UpdateUserRequest struct {
	// in: body
	Body string
}

// Update User Info
//
// swagger:response UpdateUserResponseWrapper
type UpdateUserResponseWrapper struct {
	// in: body
	Body string
}

// swagger:route POST /users users UpdateUserResponseWrapper
//
// Update User
//
// This will update user info
//
//     Responses:
//       200: UpdateUserResponseWrapper
