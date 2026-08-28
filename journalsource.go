// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
)

// JournalSourceService contains methods and other services that help with
// interacting with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJournalSourceService] method instead.
type JournalSourceService struct {
	Options []option.RequestOption
}

// NewJournalSourceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJournalSourceService(opts ...option.RequestOption) (r *JournalSourceService) {
	r = &JournalSourceService{}
	r.Options = opts
	return
}

// Retrieve a specific journal source
func (r *JournalSourceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("api/journal_sources/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

// Retrieve a list of journal sources
func (r *JournalSourceService) List(ctx context.Context, query JournalSourceListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "api/journal_sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return err
}

type JournalSourceListParams struct {
	// The ID of the journal entry
	JournalEntryID param.Field[string] `query:"journal_entry_id"`
	// The ID of the journal report
	JournalReportID param.Field[string] `query:"journal_report_id"`
	// Page number for pagination
	Page param.Field[int64] `query:"page"`
	// Number of items per page
	PerPage param.Field[int64] `query:"per_page"`
	// The ID of the source object
	SourceID param.Field[string] `query:"source_id"`
	// The type of the source object
	SourceType param.Field[string] `query:"source_type"`
}

// URLQuery serializes [JournalSourceListParams]'s query parameters as
// `url.Values`.
func (r JournalSourceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
