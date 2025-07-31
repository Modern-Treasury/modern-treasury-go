// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"github.com/Modern-Treasury/modern-treasury-go/v2/option"
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
