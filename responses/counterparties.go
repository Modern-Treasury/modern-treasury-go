package responses

import (
	"time"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
)

type Counterparty struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode,required"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// A human friendly name for this counterparty.
	Name string `json:"name,required,nullable"`
	// The accounts for this counterparty.
	Accounts []CounterpartyAccounts `json:"accounts,required"`
	// The counterparty's email.
	Email string `json:"email,required,nullable" format:"email"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata,required"`
	// Send an email to the counterparty whenever an associated payment order is sent
	// to the bank.
	SendRemittanceAdvice bool `json:"send_remittance_advice,required"`
	JSON                 CounterpartyJSON
}

type CounterpartyJSON struct {
	ID                   pjson.Metadata
	Object               pjson.Metadata
	LiveMode             pjson.Metadata
	CreatedAt            pjson.Metadata
	UpdatedAt            pjson.Metadata
	DiscardedAt          pjson.Metadata
	Name                 pjson.Metadata
	Accounts             pjson.Metadata
	Email                pjson.Metadata
	Metadata             pjson.Metadata
	SendRemittanceAdvice pjson.Metadata
	Raw                  []byte
	Extras               map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Counterparty using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Counterparty) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CounterpartyAccounts struct {
	ID     string `json:"id" format:"uuid"`
	Object string `json:"object"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode    bool      `json:"live_mode"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	UpdatedAt   time.Time `json:"updated_at" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,nullable" format:"date-time"`
	// Can be `checking`, `savings` or `other`.
	AccountType ExternalAccountType `json:"account_type"`
	// Either `individual` or `business`.
	PartyType CounterpartyAccountsPartyType `json:"party_type,nullable"`
	// The address associated with the owner or `null`.
	PartyAddress CounterpartyAccountsPartyAddress `json:"party_address,nullable"`
	// A nickname for the external account. This is only for internal usage and won't
	// affect any payments
	Name           string          `json:"name,nullable"`
	AccountDetails []AccountDetail `json:"account_details"`
	RoutingDetails []RoutingDetail `json:"routing_details"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata map[string]string `json:"metadata"`
	// The legal name of the entity which owns the account.
	PartyName          string                                 `json:"party_name"`
	ContactDetails     []CounterpartyAccountsContactDetails   `json:"contact_details"`
	VerificationStatus CounterpartyAccountsVerificationStatus `json:"verification_status"`
	JSON               CounterpartyAccountsJSON
}

type CounterpartyAccountsJSON struct {
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
	AccountDetails     pjson.Metadata
	RoutingDetails     pjson.Metadata
	Metadata           pjson.Metadata
	PartyName          pjson.Metadata
	ContactDetails     pjson.Metadata
	VerificationStatus pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CounterpartyAccounts using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CounterpartyAccounts) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CounterpartyAccountsPartyType string

const (
	CounterpartyAccountsPartyTypeBusiness   CounterpartyAccountsPartyType = "business"
	CounterpartyAccountsPartyTypeIndividual CounterpartyAccountsPartyType = "individual"
)

type CounterpartyAccountsPartyAddress struct {
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
	JSON    CounterpartyAccountsPartyAddressJSON
}

type CounterpartyAccountsPartyAddressJSON struct {
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

// UnmarshalJSON deserializes the provided bytes into
// CounterpartyAccountsPartyAddress using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *CounterpartyAccountsPartyAddress) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CounterpartyAccountsContactDetails struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode              bool                                                    `json:"live_mode,required"`
	CreatedAt             time.Time                                               `json:"created_at,required" format:"date-time"`
	UpdatedAt             time.Time                                               `json:"updated_at,required" format:"date-time"`
	DiscardedAt           time.Time                                               `json:"discarded_at,required,nullable" format:"date-time"`
	ContactIdentifier     string                                                  `json:"contact_identifier,required"`
	ContactIdentifierType CounterpartyAccountsContactDetailsContactIdentifierType `json:"contact_identifier_type,required"`
	JSON                  CounterpartyAccountsContactDetailsJSON
}

type CounterpartyAccountsContactDetailsJSON struct {
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

// UnmarshalJSON deserializes the provided bytes into
// CounterpartyAccountsContactDetails using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CounterpartyAccountsContactDetails) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CounterpartyAccountsContactDetailsContactIdentifierType string

const (
	CounterpartyAccountsContactDetailsContactIdentifierTypeEmail       CounterpartyAccountsContactDetailsContactIdentifierType = "email"
	CounterpartyAccountsContactDetailsContactIdentifierTypePhoneNumber CounterpartyAccountsContactDetailsContactIdentifierType = "phone_number"
	CounterpartyAccountsContactDetailsContactIdentifierTypeWebsite     CounterpartyAccountsContactDetailsContactIdentifierType = "website"
)

type CounterpartyAccountsVerificationStatus string

const (
	CounterpartyAccountsVerificationStatusPendingVerification CounterpartyAccountsVerificationStatus = "pending_verification"
	CounterpartyAccountsVerificationStatusUnverified          CounterpartyAccountsVerificationStatus = "unverified"
	CounterpartyAccountsVerificationStatusVerified            CounterpartyAccountsVerificationStatus = "verified"
)

type CounterpartyCollectAccountResponse struct {
	// The id of the existing counterparty.
	ID string `json:"id,required"`
	// This field will be `true` if an email requesting account details has already
	// been sent to this counterparty.
	IsResend bool `json:"is_resend,required"`
	// This is the link to the secure Modern Treasury form. By default, Modern Treasury
	// will send an email to your counterparty that includes a link to this form.
	// However, if `send_email` is passed as `false` in the body then Modern Treasury
	// will not send the email and you can send it to the counterparty directly.
	FormLink string `json:"form_link,required" format:"uri"`
	JSON     CounterpartyCollectAccountResponseJSON
}

type CounterpartyCollectAccountResponseJSON struct {
	ID       pjson.Metadata
	IsResend pjson.Metadata
	FormLink pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// CounterpartyCollectAccountResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *CounterpartyCollectAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
