// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
)

type PingResponse struct {
	Ping string           `json:"ping,required"`
	JSON pingResponseJSON `json:"-"`
}

// pingResponseJSON contains the JSON metadata for the struct [PingResponse]
type pingResponseJSON struct {
	Ping        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pingResponseJSON) RawJSON() string {
	return r.raw
}
