package jsonapi

import (
	"encoding/json"
	"net/http"
)

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

// Response is the default JSON structure that will be written.
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// respond writes a JSON-encoded body to http.ResponseWriter.
//
// he data argument is optional on all methods. If omitted, the response data field
// will be set to the HTTP status text. If provided, the response data field will be
// set to the first argument, and all other arguments will be ignored.
func respond(w http.ResponseWriter, statusCode int, data ...interface{}) {
	r := new(Response)
	r.Code = statusCode

	if len(data) == 0 {
		r.Data = http.StatusText(statusCode)
	} else {
		r.Data = data[0]
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(r); err != nil {
		panic(err)
	}
}

// Respond writes data with a custom status.
func Respond(w http.ResponseWriter, status int, data ...interface{}) {
	respond(w, status, data...)
}

// Continue writes data with status code 100.
func Continue(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusContinue, data...)
}

// SwitchingProtocols writes data with status code 101.
func SwitchingProtocols(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusSwitchingProtocols, data...)
}

// Processing writes data with status code 102.
func Processing(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusProcessing, data...)
}

// OK writes data with status code 200.
func OK(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusOK, data...)
}

// Created writes data with status code 201.
func Created(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusCreated, data...)
}

// Accepted writes data with status code 202.
func Accepted(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusAccepted, data...)
}

// NonAuthoritativeInfo writes data with status code 203.
func NonAuthoritativeInfo(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusNonAuthoritativeInfo, data...)
}

// NoContent writes data with status code 204.
func NoContent(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusNoContent, data...)
}

// ResetContent writes data with status code 205.
func ResetContent(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusResetContent, data...)
}

// PartialContent writes data with status code 206.
func PartialContent(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusPartialContent, data...)
}

// MultiStatus writes data with status code 207.
func MultiStatus(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusMultiStatus, data...)
}

// AlreadyReported writes data with status code 208.
func AlreadyReported(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusAlreadyReported, data...)
}

// IMUsed writes data with status code 226.
func IMUsed(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusIMUsed, data...)
}

// MultipleChoices writes data with status code 300.
func MultipleChoices(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusMultipleChoices, data...)
}

// MovedPermanently writes data with status code 301.
func MovedPermanently(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusMovedPermanently, data...)
}

// Found writes data with status code 302.
func Found(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusFound, data...)
}

// SeeOther writes data with status code 303.
func SeeOther(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusSeeOther, data...)
}

// NotModified writes data with status code 304.
func NotModified(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusNotModified, data...)
}

// UseProxy writes data with status code 305.
func UseProxy(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusUseProxy, data...)
}

// TemporaryRedirect writes data with status code 307.
func TemporaryRedirect(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusTemporaryRedirect, data...)
}

// PermanentRedirect writes data with status code 308.
func PermanentRedirect(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusPermanentRedirect, data...)
}

// BadRequest writes data with status code 400.
func BadRequest(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusBadRequest, data...)
}

// Unauthorized writes data with status code 401.
func Unauthorized(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusUnauthorized, data...)
}

// PaymentRequired writes data with status code 402.
func PaymentRequired(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusPaymentRequired, data...)
}

// Forbidden writes data with status code 403.
func Forbidden(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusForbidden, data...)
}

// NotFound writes data with status code 404.
func NotFound(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusNotFound, data...)
}

// MethodNotAllowed writes data with status code 405.
func MethodNotAllowed(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusMethodNotAllowed, data...)
}

// NotAcceptable writes data with status code 406.
func NotAcceptable(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusNotAcceptable, data...)
}

// ProxyAuthRequired writes data with status code 407.
func ProxyAuthRequired(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusProxyAuthRequired, data...)
}

// RequestTimeout writes data with status code 408.
func RequestTimeout(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusRequestTimeout, data...)
}

// Conflict writes data with status code 409.
func Conflict(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusConflict, data...)
}

// Gone writes data with status code 410.
func Gone(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusGone, data...)
}

// LengthRequired writes data with status code 411.
func LengthRequired(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusLengthRequired, data...)
}

// PreconditionFailed writes data with status code 412.
func PreconditionFailed(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusPreconditionFailed, data...)
}

// RequestEntityTooLarge writes data with status code 413.
func RequestEntityTooLarge(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusRequestEntityTooLarge, data...)
}

// RequestURITooLong writes data with status code 414.
func RequestURITooLong(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusRequestURITooLong, data...)
}

// UnsupportedMediaType writes data with status code 415.
func UnsupportedMediaType(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusUnsupportedMediaType, data...)
}

// RequestedRangeNotSatisfiable writes data with status code 416.
func RequestedRangeNotSatisfiable(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusRequestedRangeNotSatisfiable, data...)
}

// ExpectationFailed writes data with status code 417.
func ExpectationFailed(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusExpectationFailed, data...)
}

// Teapot writes data with status code 418.
func Teapot(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusTeapot, data...)
}

// UnprocessableEntity writes data with status code 422.
func UnprocessableEntity(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusUnprocessableEntity, data...)
}

// Locked writes data with status code 423.
func Locked(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusLocked, data...)
}

// FailedDependency writes data with status code 424.
func FailedDependency(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusFailedDependency, data...)
}

// UpgradeRequired writes data with status code 426.
func UpgradeRequired(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusUpgradeRequired, data...)
}

// PreconditionRequired writes data with status code 428.
func PreconditionRequired(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusPreconditionRequired, data...)
}

// TooManyRequests writes data with status code 429.
func TooManyRequests(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusTooManyRequests, data...)
}

// RequestHeaderFieldsTooLarge writes data with status code 431.
func RequestHeaderFieldsTooLarge(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusRequestHeaderFieldsTooLarge, data...)
}

// UnavailableForLegalReasons writes data with status code 451.
func UnavailableForLegalReasons(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusUnavailableForLegalReasons, data...)
}

// InternalServerError writes data with status code 500.
func InternalServerError(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusInternalServerError, data...)
}

// NotImplemented writes data with status code 501.
func NotImplemented(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusNotImplemented, data...)
}

// BadGateway writes data with status code 502.
func BadGateway(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusBadGateway, data...)
}

// ServiceUnavailable writes data with status code 503.
func ServiceUnavailable(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusServiceUnavailable, data...)
}

// GatewayTimeout writes data with status code 504.
func GatewayTimeout(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusGatewayTimeout, data...)
}

// HTTPVersionNotSupported writes data with status code 505.
func HTTPVersionNotSupported(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusHTTPVersionNotSupported, data...)
}

// VariantAlsoNegotiates writes data with status code 506.
func VariantAlsoNegotiates(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusVariantAlsoNegotiates, data...)
}

// InsufficientStorage writes data with status code 507.
func InsufficientStorage(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusInsufficientStorage, data...)
}

// LoopDetected writes data with status code 508.
func LoopDetected(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusLoopDetected, data...)
}

// NotExtended writes data with status code 510.
func NotExtended(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusNotExtended, data...)
}

// NetworkAuthenticationRequired writes data with status code 511.
func NetworkAuthenticationRequired(w http.ResponseWriter, data ...interface{}) {
	respond(w, http.StatusNetworkAuthenticationRequired, data...)
}
