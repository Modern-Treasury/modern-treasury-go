package moderntreasury

import (
	"context"
	"net/http"

	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/internal/requestconfig"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

type ModerntreasuryService struct {
	Options []option.RequestOption
}

func NewModerntreasuryService(opts ...option.RequestOption) (r *ModerntreasuryService) {
	r = &ModerntreasuryService{}
	r.Options = opts
	return
}

// A test endpoint often used to confirm credentials and headers are being passed
// in correctly.
func (r *ModerntreasuryService) Ping(ctx context.Context, opts ...option.RequestOption) (res *PingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PingResponse struct {
	Ping string `json:"ping,required"`
	JSON PingResponseJSON
}

type PingResponseJSON struct {
	Ping   apijson.Metadata
	raw    string
	Extras map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PingResponse using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
