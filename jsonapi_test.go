package jsonapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomStatusCodeWithNoDataRespond(t *testing.T) {
	w := httptest.NewRecorder()

	Respond(w, http.StatusForbidden)

	statusCode := w.Result().StatusCode
	if statusCode != http.StatusForbidden {
		t.Errorf("expected status code %#v, got %#v", http.StatusForbidden, statusCode)
	}

	resp := &Response{}
	if err := json.NewDecoder(w.Body).Decode(resp); err != nil {
		t.Errorf("expected %#v, got %#v", nil, err)
	}
	if resp.Data != http.StatusText(statusCode) {
		t.Errorf("expected %#v, got %#v", http.StatusText(statusCode), resp.Data)
	}
}

func TestCustomDataRespond(t *testing.T) {
	w := httptest.NewRecorder()
	data := `{"foo": "bar"}`

	Respond(w, http.StatusOK, data)

	resp := &Response{}
	if err := json.NewDecoder(w.Body).Decode(resp); err != nil {
		t.Errorf("Expected to get %#v, got %#v", nil, err)
	}
	if resp.Data != data {
		t.Errorf("Expected to get %#v, got %#v", data, resp.Data)
	}
}

func TestExplicitNilDataRespond(t *testing.T) {
	w := httptest.NewRecorder()

	Respond(w, http.StatusOK, nil)

	resp := &Response{}
	if err := json.NewDecoder(w.Body).Decode(resp); err != nil {
		t.Errorf("Expected to get %#v, got %#v", nil, err)
	}
	if resp.Data != nil {
		t.Errorf("Expected to get %#v, got %#v", nil, resp.Data)
	}
}

func TestPanicRespond(t *testing.T) {
	w := httptest.NewRecorder()

	defer func() {
		err := recover()
		if _, ok := err.(*json.UnsupportedTypeError); !ok {
			t.Errorf("Expected to get error from recover json.UnsupportedTypeError, got %#v", err)
		}
	}()

	Respond(w, http.StatusNotFound, map[bool]string{
		true: "",
	})
}

func TestStandardHTTPResponds(t *testing.T) {
	for _, test := range []struct {
		f    func(w http.ResponseWriter, data ...interface{})
		code int
	}{
		{f: Continue, code: http.StatusContinue},
		{f: SwitchingProtocols, code: http.StatusSwitchingProtocols},
		{f: Processing, code: http.StatusProcessing},

		{f: OK, code: http.StatusOK},
		{f: Created, code: http.StatusCreated},
		{f: Accepted, code: http.StatusAccepted},
		{f: NonAuthoritativeInfo, code: http.StatusNonAuthoritativeInfo},
		{f: NoContent, code: http.StatusNoContent},
		{f: ResetContent, code: http.StatusResetContent},
		{f: PartialContent, code: http.StatusPartialContent},
		{f: MultiStatus, code: http.StatusMultiStatus},
		{f: AlreadyReported, code: http.StatusAlreadyReported},
		{f: IMUsed, code: http.StatusIMUsed},

		{f: MultipleChoices, code: http.StatusMultipleChoices},
		{f: MovedPermanently, code: http.StatusMovedPermanently},
		{f: Found, code: http.StatusFound},
		{f: SeeOther, code: http.StatusSeeOther},
		{f: NotModified, code: http.StatusNotModified},
		{f: UseProxy, code: http.StatusUseProxy},
		{f: TemporaryRedirect, code: http.StatusTemporaryRedirect},
		{f: PermanentRedirect, code: http.StatusPermanentRedirect},

		{f: BadRequest, code: http.StatusBadRequest},
		{f: Unauthorized, code: http.StatusUnauthorized},
		{f: PaymentRequired, code: http.StatusPaymentRequired},
		{f: Forbidden, code: http.StatusForbidden},
		{f: NotFound, code: http.StatusNotFound},
		{f: MethodNotAllowed, code: http.StatusMethodNotAllowed},
		{f: NotAcceptable, code: http.StatusNotAcceptable},
		{f: ProxyAuthRequired, code: http.StatusProxyAuthRequired},
		{f: RequestTimeout, code: http.StatusRequestTimeout},
		{f: Conflict, code: http.StatusConflict},
		{f: Gone, code: http.StatusGone},
		{f: LengthRequired, code: http.StatusLengthRequired},
		{f: PreconditionFailed, code: http.StatusPreconditionFailed},
		{f: RequestEntityTooLarge, code: http.StatusRequestEntityTooLarge},
		{f: RequestURITooLong, code: http.StatusRequestURITooLong},
		{f: UnsupportedMediaType, code: http.StatusUnsupportedMediaType},
		{f: RequestedRangeNotSatisfiable, code: http.StatusRequestedRangeNotSatisfiable},
		{f: ExpectationFailed, code: http.StatusExpectationFailed},
		{f: Teapot, code: http.StatusTeapot},
		{f: UnprocessableEntity, code: http.StatusUnprocessableEntity},
		{f: Locked, code: http.StatusLocked},
		{f: FailedDependency, code: http.StatusFailedDependency},
		{f: UpgradeRequired, code: http.StatusUpgradeRequired},
		{f: PreconditionRequired, code: http.StatusPreconditionRequired},
		{f: TooManyRequests, code: http.StatusTooManyRequests},
		{f: RequestHeaderFieldsTooLarge, code: http.StatusRequestHeaderFieldsTooLarge},
		{f: UnavailableForLegalReasons, code: http.StatusUnavailableForLegalReasons},

		{f: InternalServerError, code: http.StatusInternalServerError},
		{f: NotImplemented, code: http.StatusNotImplemented},
		{f: BadGateway, code: http.StatusBadGateway},
		{f: ServiceUnavailable, code: http.StatusServiceUnavailable},
		{f: GatewayTimeout, code: http.StatusGatewayTimeout},
		{f: HTTPVersionNotSupported, code: http.StatusHTTPVersionNotSupported},
		{f: VariantAlsoNegotiates, code: http.StatusVariantAlsoNegotiates},
		{f: InsufficientStorage, code: http.StatusInsufficientStorage},
		{f: LoopDetected, code: http.StatusLoopDetected},
		{f: NotExtended, code: http.StatusNotExtended},
		{f: NetworkAuthenticationRequired, code: http.StatusNetworkAuthenticationRequired},
	} {
		w := httptest.NewRecorder()
		test.f(w)
		if w.Result().StatusCode != test.code {
			t.Errorf("Expected to get %#v, got %#v", w.Result().StatusCode, test.code)
		}
	}
}
