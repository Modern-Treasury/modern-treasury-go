package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type ExternalAccountService struct {
	Options []option.RequestOption
}

func NewExternalAccountService(opts ...option.RequestOption) (r *ExternalAccountService) {
	r = &ExternalAccountService{}
	r.Options = opts
	return
}

// create external account
func (r *ExternalAccountService) New(ctx context.Context, body *requests.ExternalAccountNewParams, opts ...option.RequestOption) (res *responses.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/external_accounts"
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// show external account
func (r *ExternalAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *responses.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/external_accounts/%s", id)
	err = option.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// update external account
func (r *ExternalAccountService) Update(ctx context.Context, id string, body *requests.ExternalAccountUpdateParams, opts ...option.RequestOption) (res *responses.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/external_accounts/%s", id)
	err = option.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// list external accounts
func (r *ExternalAccountService) List(ctx context.Context, query *requests.ExternalAccountListParams, opts ...option.RequestOption) (res *responses.Page[responses.ExternalAccount], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "api/external_accounts"
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

// list external accounts
func (r *ExternalAccountService) ListAutoPager(ctx context.Context, query *requests.ExternalAccountListParams, opts ...option.RequestOption) *responses.PageAutoPager[responses.ExternalAccount] {
	return responses.NewPageAutoPager(r.List(ctx, query, opts...))
}

// delete external account
func (r *ExternalAccountService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/external_accounts/%s", id)
	err = option.ExecuteNewRequest(ctx, "DELETE", path, nil, nil, opts...)
	return
}

// complete verification of external account
func (r *ExternalAccountService) CompleteVerification(ctx context.Context, id string, body *requests.ExternalAccountCompleteVerificationParams, opts ...option.RequestOption) (res *responses.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/external_accounts/%s/complete_verification", id)
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}

// verify external account
func (r *ExternalAccountService) Verify(ctx context.Context, id string, body *requests.ExternalAccountVerifyParams, opts ...option.RequestOption) (res *responses.ExternalAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/external_accounts/%s/verify", id)
	err = option.ExecuteNewRequest(ctx, "POST", path, body, &res, opts...)
	return
}
