package apierror

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
)

// TODO
type Error struct {
	JSON       errorJSON
	StatusCode int
	Request    *http.Request
	Response   *http.Response
}

// errorJSON contains the JSON metadata for the struct [Error]
type errorJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r *Error) Error() string {
	body, _ := io.ReadAll(r.Response.Body)
	return fmt.Sprintf("%s \"%s\": %d %s %s", r.Request.Method, r.Request.URL, r.Response.StatusCode, http.StatusText(r.Response.StatusCode), string(body))
}

func (r *Error) DumpRequest(body bool) []byte {
	out, _ := httputil.DumpRequestOut(r.Request, body)
	return out
}

func (r *Error) DumpResponse(body bool) []byte {
	out, _ := httputil.DumpResponse(r.Response, body)
	return out
}
