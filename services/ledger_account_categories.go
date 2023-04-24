package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type LedgerAccountCategoryService struct {
	Options []option.RequestOption
}

func NewLedgerAccountCategoryService(opts ...option.RequestOption) (r *LedgerAccountCategoryService) {
	r = &LedgerAccountCategoryService{}
	r.Options = opts
	return
}

// Create a ledger account category.
func (r *LedgerAccountCategoryService) New(ctx context.Context, body *requests.LedgerAccountCategoryNewParams, opts ...option.RequestOption) (res *responses.LedgerAccountCategory, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ledger_account_categories"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// Get the details on a single ledger account category.
func (r *LedgerAccountCategoryService) Get(ctx context.Context, id string, query *requests.LedgerAccountCategoryGetParams, opts ...option.RequestOption) (res *responses.LedgerAccountCategory, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, query, &res, opts...)
	return
}

// Update the details of a ledger account category.
func (r *LedgerAccountCategoryService) Update(ctx context.Context, id string, body *requests.LedgerAccountCategoryUpdateParams, opts ...option.RequestOption) (res *responses.LedgerAccountCategory, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Get a list of ledger account categories.
func (r *LedgerAccountCategoryService) List(ctx context.Context, query *requests.LedgerAccountCategoryListParams, opts ...option.RequestOption) (res *responses.Page[responses.LedgerAccountCategory], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/ledger_account_categories"
	cfg, err := option.NewRequestConfig(ctx, "GET", path, query, &res, opts...)
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

// Get a list of ledger account categories.
func (r *LedgerAccountCategoryService) ListAutoPager(ctx context.Context, query *requests.LedgerAccountCategoryListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.LedgerAccountCategory] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a ledger account category.
func (r *LedgerAccountCategoryService) Delete(ctx context.Context, id string, query *requests.LedgerAccountCategoryDeleteParams, opts ...option.RequestOption) (res *responses.LedgerAccountCategory, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s", id)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, query, &res, opts...)
	return
}

// Add a ledger account category to an account.
func (r *LedgerAccountCategoryService) AddLedgerAccount(ctx context.Context, id string, ledger_account_id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s/ledger_accounts/%s", id, ledger_account_id)
	err = option.ExecuteNewRequest(ctx, "PUT", path, nil, nil, opts...)
	return
}

// Add a ledger account category to a ledger account category.
func (r *LedgerAccountCategoryService) AddNestedCategory(ctx context.Context, id string, sub_category_id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s/ledger_account_categories/%s", id, sub_category_id)
	err = option.ExecuteNewRequest(ctx, "PUT", path, nil, nil, opts...)
	return
}

// Delete a ledger account category from an account.
func (r *LedgerAccountCategoryService) RemoveLedgerAccount(ctx context.Context, id string, ledger_account_id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s/ledger_accounts/%s", id, ledger_account_id)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, nil, opts...)
	return
}

// Delete a ledger account category from a ledger account category.
func (r *LedgerAccountCategoryService) RemoveNestedCategory(ctx context.Context, id string, sub_category_id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/ledger_account_categories/%s/ledger_account_categories/%s", id, sub_category_id)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, nil, opts...)
	return
}
