# jsonapi
A lightweight JSON response writer for Go. Simple, fully tested, expressive and uses only the standard library.

## Examples

The `data` argument is optional on all methods. If omitted, the response data field will be set to the HTTP status text. If provided, the response data field will be set to the first argument, and all other arguments will be ignored. This allows for optional arguments with default values without requiring any configuration or structs.

**Action:**

```go
func (w http.ResponseWriter, r *http.Request) {
    jsonapi.OK(w)
}
```

**Result:**

```go
200 OK
{"code": 200, "data": "OK"}
```

**Action:**

```go
func (w http.ResponseWriter, r *http.Request) {
    jsonapi.OK(w, map[string]string{"foo": "bar"})
}
```

**Result:**

```go
200 OK
{"code": 200, "data": {"foo": "bar"}}
```

**Action:**

```go
func (w http.ResponseWriter, r *http.Request) {
    type User struct {
        Name  string `json:"name"`
	Email string `json:"email"
    }
    jsonapi.Created(w, User{Name: "John Doe", email: "johndoe@domain.com")
}
```

**Result:**

```go
201 Created
{"code": 201, "data": {"name": "John Doe", "email", "johndoe@domain.com"}}
```

## Responder interface 

The interface contains methods for writing all standardized HTTP response codes provided by the standard library.

```go
// Responder is an interface for writing standardized JSON responses.
type Responder interface {
	Respond(w http.ResponseWriter, status int, data ...interface{})

	// 100 status codes.
	Continue(w http.ResponseWriter, data ...interface{})
	SwitchingProtocols(w http.ResponseWriter, data ...interface{})
	Processing(w http.ResponseWriter, data ...interface{})

	// 200 status codes.
	OK(w http.ResponseWriter, data ...interface{})
	Created(w http.ResponseWriter, data ...interface{})
	Accepted(w http.ResponseWriter, data ...interface{})
	NonAuthoritativeInfo(w http.ResponseWriter, data ...interface{})
	NoContent(w http.ResponseWriter, data ...interface{})
	ResetContent(w http.ResponseWriter, data ...interface{})
	PartialContent(w http.ResponseWriter, data ...interface{})
	MultiStatus(w http.ResponseWriter, data ...interface{})
	AlreadyReported(w http.ResponseWriter, data ...interface{})
	IMUsed(w http.ResponseWriter, data ...interface{})

	// 300 status codes.
	MultipleChoices(w http.ResponseWriter, data ...interface{})
	MovedPermanently(w http.ResponseWriter, data ...interface{})
	Found(w http.ResponseWriter, data ...interface{})
	SeeOther(w http.ResponseWriter, data ...interface{})
	NotModified(w http.ResponseWriter, data ...interface{})
	UseProxy(w http.ResponseWriter, data ...interface{})
	SwitchProxy(w http.ResponseWriter, data ...interface{})
	TemporaryRedirect(w http.ResponseWriter, data ...interface{})
	PermanentRedirect(w http.ResponseWriter, data ...interface{})

	// 400 status codes.
	BadRequest(w http.ResponseWriter, data ...interface{})
	Unauthorized(w http.ResponseWriter, data ...interface{})
	PaymentRequired(w http.ResponseWriter, data ...interface{})
	Forbidden(w http.ResponseWriter, data ...interface{})
	NotFound(w http.ResponseWriter, data ...interface{})
	MethodNotAllowed(w http.ResponseWriter, data ...interface{})
	NotAcceptable(w http.ResponseWriter, data ...interface{})
	ProxyAuthenticationRequired(w http.ResponseWriter, data ...interface{})
	RequestTimeout(w http.ResponseWriter, data ...interface{})
	Conflict(w http.ResponseWriter, data ...interface{})
	Gone(w http.ResponseWriter, data ...interface{})
	LengthRequired(w http.ResponseWriter, data ...interface{})
	PreconditionFailed(w http.ResponseWriter, data ...interface{})
	PayloadTooLarge(w http.ResponseWriter, data ...interface{})
	URITooLong(w http.ResponseWriter, data ...interface{})
	UnsupportedMediaType(w http.ResponseWriter, data ...interface{})
	RangeNotSatisfiable(w http.ResponseWriter, data ...interface{})
	ExpectationFailed(w http.ResponseWriter, data ...interface{})
	Teapot(w http.ResponseWriter, data ...interface{})
	MisdirectedRequest(w http.ResponseWriter, data ...interface{})
	UnprocessableEntity(w http.ResponseWriter, data ...interface{})
	Locked(w http.ResponseWriter, data ...interface{})
	FailedDependency(w http.ResponseWriter, data ...interface{})
	UpgradeRequired(w http.ResponseWriter, data ...interface{})
	PreconditionRequired(w http.ResponseWriter, data ...interface{})
	TooManyRequests(w http.ResponseWriter, data ...interface{})
	RequestHeaderFieldsTooLarge(w http.ResponseWriter, data ...interface{})
	UnavailableForLegalReasons(w http.ResponseWriter, data ...interface{})

	// 500 status codes.
	InternalServerError(w http.ResponseWriter, data ...interface{})
	NotImplemented(w http.ResponseWriter, data ...interface{})
	BadGateway(w http.ResponseWriter, data ...interface{})
	ServiceUnavailable(w http.ResponseWriter, data ...interface{})
	GatewayTimeout(w http.ResponseWriter, data ...interface{})
	HTTPVersionNotSupported(w http.ResponseWriter, data ...interface{})
	VariantAlsoNegotiates(w http.ResponseWriter, data ...interface{})
	InsufficientStorage(w http.ResponseWriter, data ...interface{})
	LoopDetected(w http.ResponseWriter, data ...interface{})
	NotExtended(w http.ResponseWriter, data ...interface{})
	NetworkAuthenticationRequired(w http.ResponseWriter, data ...interface{})
}
```

## License

Copyright (c) 2018-present [Lansana Camara](https://github.com/lansana)

Licensed under [MIT License](./LICENSE)
