package moderntreasury

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/apiquery"
	"github.com/Modern-Treasury/modern-treasury-go/internal/field"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/internal/shared"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

// PaperItemService contains methods and other services that help with interacting
// with the Modern Treasury API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPaperItemService] method instead.
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
	path := fmt.Sprintf("api/paper_items/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of all paper items.
func (r *PaperItemService) List(ctx context.Context, query PaperItemListParams, opts ...option.RequestOption) (res *shared.Page[PaperItem], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
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
func (r *PaperItemService) ListAutoPaging(ctx context.Context, query PaperItemListParams, opts ...option.RequestOption) *shared.PageAutoPager[PaperItem] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

type PaperItem struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// The ID of the reconciled Transaction Line Item or `null`.
	TransactionLineItemID string `json:"transaction_line_item_id,required,nullable" format:"uuid"`
	// The ID of the reconciled Transaction or `null`.
	TransactionID string `json:"transaction_id,required,nullable" format:"uuid"`
	// The current status of the paper item. One of `pending`, `completed`, or
	// `returned`.
	Status PaperItemStatus `json:"status,required"`
	// The identifier for the lockbox assigned by the bank.
	LockboxNumber string `json:"lockbox_number,required"`
	// The date the paper item was deposited into your organization's bank account.
	DepositDate time.Time `json:"deposit_date,required" format:"date"`
	// The amount of the paper item.
	Amount int64 `json:"amount,required"`
	// The currency of the paper item.
	Currency shared.Currency `json:"currency,required,nullable"`
	// The account number on the paper item.
	AccountNumber string `json:"account_number,required,nullable"`
	// The last 4 digits of the account_number.
	AccountNumberSafe string `json:"account_number_safe,required,nullable"`
	// The routing number on the paper item.
	RoutingNumber string `json:"routing_number,required,nullable"`
	// The check number on the paper item.
	CheckNumber string `json:"check_number,required,nullable"`
	// The name of the remitter on the paper item.
	RemitterName string `json:"remitter_name,required,nullable"`
	// The memo field on the paper item.
	MemoField string `json:"memo_field,required,nullable"`
	JSON      paperItemJSON
}

// paperItemJSON contains the JSON metadata for the struct [PaperItem]
type paperItemJSON struct {
	ID                    apijson.Field
	Object                apijson.Field
	LiveMode              apijson.Field
	CreatedAt             apijson.Field
	UpdatedAt             apijson.Field
	TransactionLineItemID apijson.Field
	TransactionID         apijson.Field
	Status                apijson.Field
	LockboxNumber         apijson.Field
	DepositDate           apijson.Field
	Amount                apijson.Field
	Currency              apijson.Field
	AccountNumber         apijson.Field
	AccountNumberSafe     apijson.Field
	RoutingNumber         apijson.Field
	CheckNumber           apijson.Field
	RemitterName          apijson.Field
	MemoField             apijson.Field
	raw                   string
	Extras                map[string]apijson.Field
}

func (r *PaperItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PaperItemStatus string

const (
	PaperItemStatusCompleted PaperItemStatus = "completed"
	PaperItemStatusPending   PaperItemStatus = "pending"
	PaperItemStatusReturned  PaperItemStatus = "returned"
)

type PaperItemListParams struct {
	// Specify `lockbox_number` if you wish to see paper items that are associated with
	// a specific lockbox number.
	LockboxNumber field.Field[string] `query:"lockbox_number"`
	// Specify an inclusive start date (YYYY-MM-DD) when filtering by deposit_date
	DepositDateStart field.Field[time.Time] `query:"deposit_date_start" format:"date"`
	// Specify an inclusive end date (YYYY-MM-DD) when filtering by deposit_date
	DepositDateEnd field.Field[time.Time] `query:"deposit_date_end" format:"date"`
	AfterCursor    field.Field[string]    `query:"after_cursor,nullable"`
	PerPage        field.Field[int64]     `query:"per_page"`
}

// URLQuery serializes [PaperItemListParams]'s query parameters as `url.Values`.
func (r PaperItemListParams) URLQuery() (v url.Values) {
	return apiquery.Marshal(r)
}
