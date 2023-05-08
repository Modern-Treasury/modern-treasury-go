package moderntreasury

import (
	"github.com/Modern-Treasury/modern-treasury-go/internal/apijson"
)

type PingResponse struct {
	Ping string `json:"ping,required"`
	JSON pingResponseJSON
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
