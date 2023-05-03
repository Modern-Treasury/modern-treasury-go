# Modern Treasury Go API Library

<a href="https://pkg.go.dev/github.com/Modern-Treasury/modern-treasury-go"><img src="https://pkg.go.dev/badge/github.com/Modern-Treasury/modern-treasury-go.svg" alt="Go Reference"></a>

The Modern Treasury Go library provides convenient access to [the Modern Treasury REST
API](https://docs.moderntreasury.com) from applications written in Go.

## Installation

```go
import (
	"github.com/Modern-Treasury/modern-treasury-go" // imported as moderntreasury
)
```

Or to pin the version:

```sh
go get -u 'github.com/Modern-Treasury/modern-treasury-go@v0.0.1'
```

## Requirements

This library requires Go 1.18+.

## Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func main() {
	client := moderntreasury.NewClient(
		option.WithOrganizationID("my-organization-ID"), // defaults to os.LookupEnv("MODERN_TREASURY_ORGANIZATION_ID")
		option.WithAPIKey("my api key"),                 // defaults to os.LookupEnv("MODERN_TREASURY_API_KEY")
	)
	externalAccount, err := client.ExternalAccounts.New(context.TODO(), moderntreasury.ExternalAccountNewParams{
		CounterpartyID: moderntreasury.F("9eba513a-53fd-4d6d-ad52-ccce122ab92a"),
		Name:           moderntreasury.F("my bank"),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", externalAccount)
}

```

### Request Fields

All request parameters are wrapped in a generic `Field` type,
which we use to distinguish zero values from null or omitted fields.

This prevents accidentally sending a zero value if you forget a required parameter,
and enables explicitly sending `null`, `false`, `''`, or `0` on optional parameters.
Any field not specified is not sent.

To construct fields with values, use the helpers `String()`, `Int()`, `Float()`, or most commonly, the generic `F[T]()`.
To send a null, use `Null[T]()`, and to send a nonconforming value, use `Raw[T](any)`. For example:

```go
params := FooParams{
	Name: moderntreasury.F("hello"),

	// Explicitly send `"description": null`
	Description: moderntreasury.Null[string](),

	Point: moderntreasury.F(moderntreasury.Point{
		X: moderntreasury.Int(0),
		Y: moderntreasury.Int(1),

		// In cases where the API specifies a given type,
		// but you want to send something else, use `Raw`:
		Z: moderntreasury.Raw[int64](0.01), // sends a float
	}),
}
```

### Response Objects

All fields in response structs are value types (not pointers or wrappers).

If a given field is `null`, not present, or invalid, the corresponding field
will simply be its zero value.

All response structs also include a special `JSON` field, containing more detailed
information about each property, which you can use like so:

```go
if res.Name == "" {
	// true if `"name"` is either not present or explicitly null
	res.JSON.Name.IsNull()

	// true if the `"name"` key was not present in the repsonse JSON at all
	res.JSON.Name.IsMissing()

	// When the API returns data that cannot be coerced to the expected type:
	if res.JSON.Name.IsInvalid() {
		raw := res.JSON.Name.Raw()

		legacyName := struct{
			First string `json:"first"`
			Last  string `json:"last"`
		}{}
		json.Unmarshal([]byte(raw), &legacyName)
		name = legacyName.First + " " + legacyName.Last
	}
}
```

These `.JSON` structs also include an `Extras` map containing
any properties in the json response that were not specified
in the struct. This can be useful for API features not yet
present in the SDK.

```go
body := res.JSON.Extras["my_unexpected_field"].Raw()
```

### RequestOptions

This library uses the functional options pattern. Functions defined in the
`option` package return a `RequestOption`, which is a closure that mutates a
`RequestConfig`. These options can be supplied to the client or at individual
requests. For example:

```go
client := moderntreasury.NewClient(
	// Adds a header to every request made by the client
	option.WithHeader("X-Some-Header", "custom_header_info"),
)

client.ExternalAccounts.New(context.TODO(), ...,
	// Override the header
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	// Add an undocumented field to the request body, using sjson syntax
	option.WithJSONSet("some.json.path", map[string]string{"my": "object"}),
)
```

The full list of request options is [here](https://pkg.go.dev/github.com/Modern-Treasury/modern-treasury-go/option).

### Pagination

This library provides some conveniences for working with paginated list endpoints.

You can use `.ListAutoPaging()` methods to iterate through items across all pages:

```go
iter := client.ExternalAccounts.ListAutoPaging(context.TODO(), moderntreasury.ExternalAccountListParams{})
// Automatically fetches more pages as needed.
for iter.Next() {
	externalAccount := iter.Current()
	fmt.Printf("%+v\n", externalAccount)
}
if err := iter.Err(); err != nil {
	panic(err.Error())
}
```

Or you can use simple `.List()` methods to fetch a single page and receive a standard response object
with additional helper methods like `.GetNextPage()`, e.g.:

```go
page, err := client.ExternalAccounts.List(context.TODO(), moderntreasury.ExternalAccountListParams{})
for page != nil {
	for _, externalAccount := range page.Items {
		fmt.Printf("%+v\n", externalAccount)
	}
	page, err = page.GetNextPage()
}
if err != nil {
	panic(err.Error())
}
```

### Errors

When the API returns a non-success status code, we return an error with type
`*moderntreasury.Error`. This contains the `StatusCode`, `*http.Request`, and
`*http.Response` values of the request, as well as the JSON of the error body
(much like other response objects in the SDK).

To handle errors, we recommend that you use the `errors.As` pattern:

```go
_, err := client.ExternalAccounts.New(context.TODO(), moderntreasury.ExternalAccountNewParams{
	CounterpartyID: moderntreasury.F("missing"),
})
if err != nil {
	var apierr *moderntreasury.Error
	if errors.As(err, &apierr) {
		println(string(apierr.DumpRequest(true)))  // Prints the serialized HTTP request
		println(string(apierr.DumpResponse(true))) // Prints the serialized HTTP response

	}
	panic(err.Error()) // GET "/api/external_accounts": 400 Bad Request { ... }
}
```

When other errors occur, they are returned unwrapped; for example,
if HTTP transport fails, you might receive `*url.Error` wrapping `*net.OpError`.

### Timeouts

Requests do not time out by default; use context to configure a timeout for a request lifecycle.

Note that if a request is [retried](#retries), the context timeout does not start over.
To set a per-retry timeout, use `option.WithRequestTimeout()`.

```go
// This sets the timeout for the request, including all the retries.
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()
client.ExternalAccounts.List(
	ctx,
	moderntreasury.ExternalAccountListParams{
		PartyName: moderntreasury.F("my bank"),
	},
	// This sets the per-retry timeout
	option.WithRequestTimeout(20*time.Second),
)
```

## Retries

Certain errors will be automatically retried 2 times by default, with a short exponential backoff.
Connection errors (for example, due to a network connectivity problem), 409 Conflict, 429 Rate Limit,
and >=500 Internal errors will all be retried by default.

You can use the `WithMaxRetries` option to configure or disable this:

```go
// Configure the default for all requests:
client := moderntreasury.NewClient(
	option.WithOrganizationID("my-organization-ID"), // defaults to os.LookupEnv("MODERN_TREASURY_ORGANIZATION_ID")
	option.WithMaxRetries(0),                        // default is 2
)

// Override per-request:
client.ExternalAccounts.List(
	context.TODO(),
	moderntreasury.ExternalAccountListParams{},
	option.WithMaxRetries(5),
)
```

### Middleware

You may apply any middleware you wish by replacing the default `http.Client` with
`option.WithHTTPClient(client)`. An example of a basic logging middleware is given
below:

```go
type requestLogger struct {}

func (l *requestLogger) RoundTrip(req *http.Request) (*http.Response, error) {
	b, _ := httputil.DumpRequest(req, true)
	println(string(b))
	res, err := http.DefaultClient.Do(req)
	return res, err
}

client := moderntreasury.NewClient(
	option.WithHTTPClient(&http.Client{Transport: &requestLogger{}}),
)
```

## Status

This package is in beta. Its internals and interfaces are not stable and
subject to change without a major version bump; please reach out if you rely on
any undocumented behavior.

We are keen for your feedback; please email us at
[sdk-feedback@moderntreasury.com](mailto:sdk-feedback@moderntreasury.com) or open an issue with questions, bugs, or
suggestions.