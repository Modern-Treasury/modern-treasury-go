package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type InternalAccount struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Can be checking, savings or other.
	AccountType InternalAccountAccountType `json:"account_type,required,nullable"`
	// The legal name of the entity which owns the account.
	PartyName string `json:"party_name,required"`
	// Either individual or business.
	PartyType InternalAccountPartyType `json:"party_type,required,nullable"`
	// The address associated with the owner or null.
	PartyAddress InternalAccountPartyAddress `json:"party_address,required,nullable"`
	// A nickname for the account.
	Name string `json:"name,required,nullable"`
	// An array of account detail objects.
	AccountDetails []AccountDetail `json:"account_details,required"`
	// An array of routing detail objects.
	RoutingDetails []RoutingDetail `json:"routing_details,required"`
	// Specifies which financial institution the accounts belong to.
	Connection Connection `json:"connection,required"`
	// The currency of the account.
	Currency Currency `json:"currency,required,nullable"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The parent InternalAccount of this account.
	ParentAccountID string `json:"parent_account_id,required,nullable" format:"uuid"`
	// The Counterparty associated to this account.
	CounterpartyID string `json:"counterparty_id,required,nullable" format:"uuid"`
	JSON           InternalAccountJSON
}

type InternalAccountJSON struct {
	ID              pjson.Metadata
	Object          pjson.Metadata
	LiveMode        pjson.Metadata
	CreatedAt       pjson.Metadata
	UpdatedAt       pjson.Metadata
	AccountType     pjson.Metadata
	PartyName       pjson.Metadata
	PartyType       pjson.Metadata
	PartyAddress    pjson.Metadata
	Name            pjson.Metadata
	AccountDetails  pjson.Metadata
	RoutingDetails  pjson.Metadata
	Connection      pjson.Metadata
	Currency        pjson.Metadata
	Metadata        pjson.Metadata
	ParentAccountID pjson.Metadata
	CounterpartyID  pjson.Metadata
	Raw             []byte
	Extras          map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InternalAccount using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *InternalAccount) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type InternalAccountAccountType string

const (
	InternalAccountAccountTypeCash        InternalAccountAccountType = "cash"
	InternalAccountAccountTypeChecking    InternalAccountAccountType = "checking"
	InternalAccountAccountTypeLoan        InternalAccountAccountType = "loan"
	InternalAccountAccountTypeNonResident InternalAccountAccountType = "non_resident"
	InternalAccountAccountTypeOther       InternalAccountAccountType = "other"
	InternalAccountAccountTypeOverdraft   InternalAccountAccountType = "overdraft"
	InternalAccountAccountTypeSavings     InternalAccountAccountType = "savings"
)

type InternalAccountPartyType string

const (
	InternalAccountPartyTypeBusiness   InternalAccountPartyType = "business"
	InternalAccountPartyTypeIndividual InternalAccountPartyType = "individual"
)

type InternalAccountPartyAddress struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool      `json:"live_mode,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	Line1     string    `json:"line1,required,nullable"`
	Line2     string    `json:"line2,required,nullable"`
	// Locality or City.
	Locality string `json:"locality,required,nullable"`
	// Region or State.
	Region string `json:"region,required,nullable"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required,nullable"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,required,nullable"`
	JSON    InternalAccountPartyAddressJSON
}

type InternalAccountPartyAddressJSON struct {
	ID         pjson.Metadata
	Object     pjson.Metadata
	LiveMode   pjson.Metadata
	CreatedAt  pjson.Metadata
	UpdatedAt  pjson.Metadata
	Line1      pjson.Metadata
	Line2      pjson.Metadata
	Locality   pjson.Metadata
	Region     pjson.Metadata
	PostalCode pjson.Metadata
	Country    pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into InternalAccountPartyAddress
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *InternalAccountPartyAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
