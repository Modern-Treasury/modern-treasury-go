// File generated from our OpenAPI spec by Stainless.

package moderntreasury_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"

	moderntreasury "github.com/Modern-Treasury/modern-treasury-go"
	"github.com/Modern-Treasury/modern-treasury-go/internal/testutil"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

func TestDocumentNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	t.Skip("skipped: prism mock server is broken for file uploads")
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Documents.New(
		context.TODO(),
		moderntreasury.DocumentNewParamsDocumentableTypeCases,
		"string",
		moderntreasury.DocumentNewParams{
			File:           moderntreasury.F(io.Reader(bytes.NewBuffer([]byte("some file contents")))),
			DocumentType:   moderntreasury.F("string"),
			IdempotencyKey: moderntreasury.F("string"),
		},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Documents.Get(
		context.TODO(),
		moderntreasury.DocumentGetParamsDocumentableTypeCases,
		"string",
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDocumentListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := moderntreasury.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithOrganizationID("my-organization-ID"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := c.Documents.List(
		context.TODO(),
		moderntreasury.DocumentListParamsDocumentableTypeCases,
		"string",
		moderntreasury.DocumentListParams{
			AfterCursor: moderntreasury.F("string"),
			PerPage:     moderntreasury.F(int64(0)),
		},
	)
	if err != nil {
		var apierr *moderntreasury.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
