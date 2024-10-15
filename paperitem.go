// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
	"github.com/Modern-Treasury/modern-treasury-go/v2/packages/pagination"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

// PaperItemService contains methods and other services that help with interacting
// with the Modern Treasury API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaperItemService] method instead.
type PaperItemService struct {
	Options []option.RequestOption
}

// NewPaperItemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPaperItemService(opts ...option.RequestOption) (r *PaperItemService) {
	r = &PaperItemService{}
	r.Options = opts
	return
}

// Get details on a single paper item.
func (r *PaperItemService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PaperItem, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/paper_items/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of all paper items.
func (r *PaperItemService) List(ctx context.Context, query PaperItemListParams, opts ...option.RequestOption) (res *pagination.Page[PaperItem], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/paper_items"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of all paper items.
func (r *PaperItemService) ListAutoPaging(ctx context.Context, query PaperItemListParams, opts ...option.RequestOption) *pagination.PageAutoPager[PaperItem] {
	return pagination.NewPageAutoPager(r.List(ctx, query, opts...))
}

type PaperItem struct {
	ID string `json:"id,required" format:"uuid"`
	// The account number on the paper item.
	AccountNumber string `json:"account_number,required,nullable"`
	// The last 4 digits of the account_number.
	AccountNumberSafe string `json:"account_number_safe,required,nullable"`
	// The amount of the paper item.
	Amount int64 `json:"amount,required"`
	// The check number on the paper item.
	CheckNumber string    `json:"check_number,required,nullable"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	// The currency of the paper item.
	Currency shared.Currency `json:"currency,required"`
	// The date the paper item was deposited into your organization's bank account.
	DepositDate time.Time `json:"deposit_date,required" format:"date"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// The identifier for the lockbox assigned by the bank.
	LockboxNumber string `json:"lockbox_number,required"`
	// The memo field on the paper item.
	MemoField string `json:"memo_field,required,nullable"`
	Object    string `json:"object,required"`
	// The name of the remitter on the paper item.
	RemitterName string `json:"remitter_name,required,nullable"`
	// The routing number on the paper item.
	RoutingNumber string `json:"routing_number,required,nullable"`
	// The current status of the paper item. One of `pending`, `completed`, or
	// `returned`.
	Status PaperItemStatus `json:"status,required"`
	// The ID of the reconciled Transaction or `null`.
	TransactionID string `json:"transaction_id,required,nullable" format:"uuid"`
	// The ID of the reconciled Transaction Line Item or `null`.
	TransactionLineItemID string        `json:"transaction_line_item_id,required,nullable" format:"uuid"`
	UpdatedAt             time.Time     `json:"updated_at,required" format:"date-time"`
	JSON                  paperItemJSON `json:"-"`
}

// paperItemJSON contains the JSON metadata for the struct [PaperItem]
type paperItemJSON struct {
	ID                    apijson.Field
	AccountNumber         apijson.Field
	AccountNumberSafe     apijson.Field
	Amount                apijson.Field
	CheckNumber           apijson.Field
	CreatedAt             apijson.Field
	Currency              apijson.Field
	DepositDate           apijson.Field
	LiveMode              apijson.Field
	LockboxNumber         apijson.Field
	MemoField             apijson.Field
	Object                apijson.Field
	RemitterName          apijson.Field
	RoutingNumber         apijson.Field
	Status                apijson.Field
	TransactionID         apijson.Field
	TransactionLineItemID apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PaperItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paperItemJSON) RawJSON() string {
	return r.raw
}

// The current status of the paper item. One of `pending`, `completed`, or
// `returned`.
type PaperItemStatus string

const (
	PaperItemStatusCompleted PaperItemStatus = "completed"
	PaperItemStatusPending   PaperItemStatus = "pending"
	PaperItemStatusReturned  PaperItemStatus = "returned"
)

func (r PaperItemStatus) IsKnown() bool {
	switch r {
	case PaperItemStatusCompleted, PaperItemStatusPending, PaperItemStatusReturned:
		return true
	}
	return false
}

type PaperItemListParams struct {
	AfterCursor param.Field[string] `query:"after_cursor"`
	// Specify an inclusive end date (YYYY-MM-DD) when filtering by deposit_date
	DepositDateEnd param.Field[time.Time] `query:"deposit_date_end" format:"date"`
	// Specify an inclusive start date (YYYY-MM-DD) when filtering by deposit_date
	DepositDateStart param.Field[time.Time] `query:"deposit_date_start" format:"date"`
	// Specify `lockbox_number` if you wish to see paper items that are associated with
	// a specific lockbox number.
	LockboxNumber param.Field[string] `query:"lockbox_number"`
	PerPage       param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [PaperItemListParams]'s query parameters as `url.Values`.
func (r PaperItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
