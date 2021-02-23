// Code generated by goa v3.2.2, DO NOT EDIT.
//
// user HTTP server types
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package server

import (
	user "github.com/tektoncd/hub/api/gen/user"
	goa "goa.design/goa/v3/pkg"
)

// RefreshAccessTokenResponseBody is the type of the "user" service
// "RefreshAccessToken" endpoint HTTP response body.
type RefreshAccessTokenResponseBody struct {
	// User Access JWT
	Data *AccessTokenResponseBody `form:"data" json:"data" xml:"data"`
}

// RefreshAccessTokenInternalErrorResponseBody is the type of the "user"
// service "RefreshAccessToken" endpoint HTTP response body for the
// "internal-error" error.
type RefreshAccessTokenInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RefreshAccessTokenInvalidTokenResponseBody is the type of the "user" service
// "RefreshAccessToken" endpoint HTTP response body for the "invalid-token"
// error.
type RefreshAccessTokenInvalidTokenResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RefreshAccessTokenInvalidScopesResponseBody is the type of the "user"
// service "RefreshAccessToken" endpoint HTTP response body for the
// "invalid-scopes" error.
type RefreshAccessTokenInvalidScopesResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// AccessTokenResponseBody is used to define fields on response body types.
type AccessTokenResponseBody struct {
	// Access Token for user
	Access *TokenResponseBody `form:"access,omitempty" json:"access,omitempty" xml:"access,omitempty"`
}

// TokenResponseBody is used to define fields on response body types.
type TokenResponseBody struct {
	// JWT
	Token string `form:"token" json:"token" xml:"token"`
	// Duration the token will Expire In
	RefreshInterval string `form:"refreshInterval" json:"refreshInterval" xml:"refreshInterval"`
	// Time the token will expires at
	ExpiresAt int64 `form:"expiresAt" json:"expiresAt" xml:"expiresAt"`
}

// NewRefreshAccessTokenResponseBody builds the HTTP response body from the
// result of the "RefreshAccessToken" endpoint of the "user" service.
func NewRefreshAccessTokenResponseBody(res *user.RefreshAccessTokenResult) *RefreshAccessTokenResponseBody {
	body := &RefreshAccessTokenResponseBody{}
	if res.Data != nil {
		body.Data = marshalUserAccessTokenToAccessTokenResponseBody(res.Data)
	}
	return body
}

// NewRefreshAccessTokenInternalErrorResponseBody builds the HTTP response body
// from the result of the "RefreshAccessToken" endpoint of the "user" service.
func NewRefreshAccessTokenInternalErrorResponseBody(res *goa.ServiceError) *RefreshAccessTokenInternalErrorResponseBody {
	body := &RefreshAccessTokenInternalErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRefreshAccessTokenInvalidTokenResponseBody builds the HTTP response body
// from the result of the "RefreshAccessToken" endpoint of the "user" service.
func NewRefreshAccessTokenInvalidTokenResponseBody(res *goa.ServiceError) *RefreshAccessTokenInvalidTokenResponseBody {
	body := &RefreshAccessTokenInvalidTokenResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRefreshAccessTokenInvalidScopesResponseBody builds the HTTP response body
// from the result of the "RefreshAccessToken" endpoint of the "user" service.
func NewRefreshAccessTokenInvalidScopesResponseBody(res *goa.ServiceError) *RefreshAccessTokenInvalidScopesResponseBody {
	body := &RefreshAccessTokenInvalidScopesResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRefreshAccessTokenPayload builds a user service RefreshAccessToken
// endpoint payload.
func NewRefreshAccessTokenPayload(refreshToken string) *user.RefreshAccessTokenPayload {
	v := &user.RefreshAccessTokenPayload{}
	v.RefreshToken = refreshToken

	return v
}