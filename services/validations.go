package services

import (
	"context"

	"github.com/Modern-Treasury/modern-treasury-go/option"
	"github.com/Modern-Treasury/modern-treasury-go/requests"
	"github.com/Modern-Treasury/modern-treasury-go/responses"
)

type ValidationService struct {
	Options []option.RequestOption
}

func NewValidationService(opts ...option.RequestOption) (r *ValidationService) {
	r = &ValidationService{}
	r.Options = opts
	return
}

// Validates the routing number information supplied without creating a routing
// detail
func (r *ValidationService) ValidateRoutingNumber(ctx context.Context, query *requests.ValidationValidateRoutingNumberParams, opts ...option.RequestOption) (res *responses.RoutingNumberLookupRequest, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/validations/routing_numbers"
	err = option.ExecuteNewRequest(ctx, "GET", path, query, &res, opts...)
	return
}
