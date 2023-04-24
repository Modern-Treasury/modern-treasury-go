package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type ExternalAccount struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// Can be `checking`, `savings` or `other`.
	AccountType ExternalAccountType `json:"account_type,required"`
	// Either `individual` or `business`.
	PartyType ExternalAccountPartyType `json:"party_type,required,nullable"`
	// The address associated with the owner or `null`.
	PartyAddress ExternalAccountPartyAddress `json:"party_address,required,nullable"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           string          `json:"name,required,nullable"`
	CounterpartyID string          `json:"counterparty_id,required,nullable" format:"uuid"`
	AccountDetails []AccountDetail `json:"account_details,required"`
	RoutingDetails []RoutingDetail `json:"routing_details,required"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// The legal name of the entity which owns the account.
	PartyName          string                            `json:"party_name,required"`
	ContactDetails     []ExternalAccountContactDetails   `json:"contact_details,required"`
	VerificationStatus ExternalAccountVerificationStatus `json:"verification_status,required"`
	JSON               ExternalAccountJSON
}

type ExternalAccountJSON struct {
	ID                 pjson.Metadata
	Object             pjson.Metadata
	LiveMode           pjson.Metadata
	CreatedAt          pjson.Metadata
	UpdatedAt          pjson.Metadata
	DiscardedAt        pjson.Metadata
	AccountType        pjson.Metadata
	PartyType          pjson.Metadata
	PartyAddress       pjson.Metadata
	Name               pjson.Metadata
	CounterpartyID     pjson.Metadata
	AccountDetails     pjson.Metadata
	RoutingDetails     pjson.Metadata
	Metadata           pjson.Metadata
	PartyName          pjson.Metadata
	ContactDetails     pjson.Metadata
	VerificationStatus pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ExternalAccount using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ExternalAccount) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ExternalAccountType string

const (
	ExternalAccountTypeCash        ExternalAccountType = "cash"
	ExternalAccountTypeChecking    ExternalAccountType = "checking"
	ExternalAccountTypeLoan        ExternalAccountType = "loan"
	ExternalAccountTypeNonResident ExternalAccountType = "non_resident"
	ExternalAccountTypeOther       ExternalAccountType = "other"
	ExternalAccountTypeOverdraft   ExternalAccountType = "overdraft"
	ExternalAccountTypeSavings     ExternalAccountType = "savings"
)

type ExternalAccountPartyType string

const (
	ExternalAccountPartyTypeBusiness   ExternalAccountPartyType = "business"
	ExternalAccountPartyTypeIndividual ExternalAccountPartyType = "individual"
)

type ExternalAccountPartyAddress struct {
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
	JSON    ExternalAccountPartyAddressJSON
}

type ExternalAccountPartyAddressJSON struct {
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

// UnmarshalJSON deserializes the provided bytes into ExternalAccountPartyAddress
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ExternalAccountPartyAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ExternalAccountContactDetails struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              bool                                               `json:"live_mode,required"`
	CreatedAt             time.Time                                          `json:"created_at,required" format:"date-time"`
	UpdatedAt             time.Time                                          `json:"updated_at,required" format:"date-time"`
	DiscardedAt           time.Time                                          `json:"discarded_at,required,nullable" format:"date-time"`
	ContactIdentifier     string                                             `json:"contact_identifier,required"`
	ContactIdentifierType ExternalAccountContactDetailsContactIdentifierType `json:"contact_identifier_type,required"`
	JSON                  ExternalAccountContactDetailsJSON
}

type ExternalAccountContactDetailsJSON struct {
	ID                    pjson.Metadata
	Object                pjson.Metadata
	LiveMode              pjson.Metadata
	CreatedAt             pjson.Metadata
	UpdatedAt             pjson.Metadata
	DiscardedAt           pjson.Metadata
	ContactIdentifier     pjson.Metadata
	ContactIdentifierType pjson.Metadata
	Raw                   []byte
	Extras                map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ExternalAccountContactDetails
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ExternalAccountContactDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ExternalAccountContactDetailsContactIdentifierType string

const (
	ExternalAccountContactDetailsContactIdentifierTypeEmail       ExternalAccountContactDetailsContactIdentifierType = "email"
	ExternalAccountContactDetailsContactIdentifierTypePhoneNumber ExternalAccountContactDetailsContactIdentifierType = "phone_number"
	ExternalAccountContactDetailsContactIdentifierTypeWebsite     ExternalAccountContactDetailsContactIdentifierType = "website"
)

type ExternalAccountVerificationStatus string

const (
	ExternalAccountVerificationStatusPendingVerification ExternalAccountVerificationStatus = "pending_verification"
	ExternalAccountVerificationStatusUnverified          ExternalAccountVerificationStatus = "unverified"
	ExternalAccountVerificationStatusVerified            ExternalAccountVerificationStatus = "verified"
)
