package option

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/tidwall/sjson"
)

type RequestOption = func(*requestconfig.RequestConfig) error

func WithBaseURL(base string) RequestOption {
	u, err := url.Parse(base)
	if err != nil {
		log.Fatalf("failed to parse BaseURL: %s\n", err)
	}
	return func(r *requestconfig.RequestConfig) error {
		r.BaseURL = u
		return nil
	}
}

// WithHTTPClient changes the underlying [http.Client] used to make this
// request, which by default is [http.DefaultClient].
func WithHTTPClient(client *http.Client) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.HTTPClient = client
		return nil
	}
}

func WithMaxRetries(retries int) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.MaxRetries = retries
		return nil
	}
}

// WithHeader sets the header value to the associated key. It overwrites
// any value if there was one already present.
func WithHeader(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.Request.Header[key] = []string{value}
		return nil
	}
}

// WithHeaderAdd adds the header value to the associated key. It appends
// onto any existing values.
func WithHeaderAdd(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.Request.Header[key] = append(r.Request.Header[key], value)
		return nil
	}
}

// WithHeaderDel deletes the header value(s) associated with the given key
func WithHeaderDel(key string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		delete(r.Request.Header, key)
		return nil
	}
}

// WithQuery sets the query value to the associated key. It overwrites
// any value if there was one already present.
func WithQuery(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		query := r.Request.URL.Query()
		query.Set(key, value)
		r.Request.URL.RawQuery = query.Encode()
		return nil
	}
}

// WithQueryAdd adds the query value to the associated key. It appends
// onto any existing values.
func WithQueryAdd(key, value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		query := r.Request.URL.Query()
		query.Add(key, value)
		r.Request.URL.RawQuery = query.Encode()
		return nil
	}
}

// WithQueryDel deletes the query value(s) associated with the key
func WithQueryDel(key string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		query := r.Request.URL.Query()
		query.Del(key)
		r.Request.URL.RawQuery = query.Encode()
		return nil
	}
}

// WithJSONSet sets the body's JSON value associated with the key.
// The key accepts a string as defined by the [sjson format](https://github.com/tidwall/sjson)
func WithJSONSet(key string, value interface{}) RequestOption {
	return func(r *requestconfig.RequestConfig) (err error) {
		r.Buffer, err = sjson.SetBytes(r.Buffer, key, value)
		return err
	}
}

// WithJSONDel deletes the body's JSON value associated with the key.
// The key accepts a string as defined by the [sjson format](https://github.com/tidwall/sjson)
func WithJSONDel(key string) RequestOption {
	return func(r *requestconfig.RequestConfig) (err error) {
		r.Buffer, err = sjson.DeleteBytes(r.Buffer, key)
		return err
	}
}

// WithResponseBodyInto overwrites the deserialization target with
// the given destination. If provided, we don't deserialize into the default struct.
func WithResponseBodyInto(dst any) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.ResponseBodyInto = dst
		return nil
	}
}

// WithResponseInto copies the `*http.Response` into the given
// address.
func WithResponseInto(dst **http.Response) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.ResponseInto = dst
		return nil
	}
}

// WithRequestTimeout sets the timeout for each request attempt. This
// should be smaller than the timeout defined in the context, which spans all retries.
func WithRequestTimeout(dur time.Duration) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.RequestTimeout = dur
		return nil
	}
}

func WithAPIKey(key string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.APIKey = key
		return r.Apply(WithHeader("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(r.OrganizationID+":"+r.APIKey)))))
	}
}

func WithEnvironmentProduction() RequestOption {
	return WithBaseURL("https://app.moderntreasury.com/")
}

func WithOrganizationID(value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.OrganizationID = value
		err := r.Apply(WithHeader("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(r.OrganizationID+":"+r.APIKey)))))
		if err != nil {
			return err
		}
		return nil
	}
}

func WithWebhookKey(value string) RequestOption {
	return func(r *requestconfig.RequestConfig) error {
		r.WebhookKey = value
		return nil
	}
}
