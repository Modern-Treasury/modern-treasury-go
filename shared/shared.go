// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/Modern-Treasury/modern-treasury-go/v2"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/param"
)

type Accounting struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID string `json:"account_id,nullable" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID string         `json:"class_id,nullable" format:"uuid"`
	JSON    accountingJSON `json:"-"`
}

// accountingJSON contains the JSON metadata for the struct [Accounting]
type accountingJSON struct {
	AccountID   apijson.Field
	ClassID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Accounting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountingJSON) RawJSON() string {
	return r.raw
}

type AccountingParam struct {
	// The ID of one of your accounting categories. Note that these will only be
	// accessible if your accounting system has been connected.
	AccountID param.Field[string] `json:"account_id" format:"uuid"`
	// The ID of one of the class objects in your accounting system. Class objects
	// track segments of your business independent of client or project. Note that
	// these will only be accessible if your accounting system has been connected.
	ClassID param.Field[string] `json:"class_id" format:"uuid"`
}

func (r AccountingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountsType string

const (
	AccountsTypeExternalAccounts AccountsType = "external_accounts"
	AccountsTypeInternalAccounts AccountsType = "internal_accounts"
)

func (r AccountsType) IsKnown() bool {
	switch r {
	case AccountsTypeExternalAccounts, AccountsTypeInternalAccounts:
		return true
	}
	return false
}

type Address struct {
	ID string `json:"id,required" format:"uuid"`
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country   string    `json:"country,required,nullable"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	Line1     string    `json:"line1,required,nullable"`
	Line2     string    `json:"line2,required,nullable"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool `json:"live_mode,required"`
	// Locality or City.
	Locality string `json:"locality,required,nullable"`
	Object   string `json:"object,required"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,required,nullable"`
	// Region or State.
	Region    string      `json:"region,required,nullable"`
	UpdatedAt time.Time   `json:"updated_at,required" format:"date-time"`
	JSON      addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	ID          apijson.Field
	Country     apijson.Field
	CreatedAt   apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	LiveMode    apijson.Field
	Locality    apijson.Field
	Object      apijson.Field
	PostalCode  apijson.Field
	Region      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

type AddressRequest struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country string `json:"country,nullable"`
	Line1   string `json:"line1,nullable"`
	Line2   string `json:"line2,nullable"`
	// Locality or City.
	Locality string `json:"locality,nullable"`
	// The postal code of the address.
	PostalCode string `json:"postal_code,nullable"`
	// Region or State.
	Region string             `json:"region,nullable"`
	JSON   addressRequestJSON `json:"-"`
}

// addressRequestJSON contains the JSON metadata for the struct [AddressRequest]
type addressRequestJSON struct {
	Country     apijson.Field
	Line1       apijson.Field
	Line2       apijson.Field
	Locality    apijson.Field
	PostalCode  apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressRequestJSON) RawJSON() string {
	return r.raw
}

type AddressRequestParam struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country"`
	Line1   param.Field[string] `json:"line1"`
	Line2   param.Field[string] `json:"line2"`
	// Locality or City.
	Locality param.Field[string] `json:"locality"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code"`
	// Region or State.
	Region param.Field[string] `json:"region"`
}

func (r AddressRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AsyncResponse struct {
	ID     string            `json:"id,required" format:"uuid"`
	Object string            `json:"object,required"`
	JSON   asyncResponseJSON `json:"-"`
}

// asyncResponseJSON contains the JSON metadata for the struct [AsyncResponse]
type asyncResponseJSON struct {
	ID          apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AsyncResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncResponseJSON) RawJSON() string {
	return r.raw
}

type ChildLegalEntityCreateParam struct {
	// A list of addresses for the entity.
	Addresses    param.Field[[]LegalEntityAddressCreateRequestParam] `json:"addresses"`
	BankSettings param.Field[moderntreasury.BankSettingsParam]       `json:"bank_settings"`
	// The business's legal business name.
	BusinessName param.Field[string] `json:"business_name"`
	// The country of citizenship for an individual.
	CitizenshipCountry param.Field[string]                           `json:"citizenship_country"`
	ComplianceDetails  param.Field[LegalEntityComplianceDetailParam] `json:"compliance_details"`
	// A business's formation date (YYYY-MM-DD).
	DateFormed param.Field[time.Time] `json:"date_formed" format:"date"`
	// An individual's date of birth (YYYY-MM-DD).
	DateOfBirth          param.Field[time.Time] `json:"date_of_birth" format:"date"`
	DoingBusinessAsNames param.Field[[]string]  `json:"doing_business_as_names"`
	// The entity's primary email.
	Email param.Field[string] `json:"email"`
	// An individual's first name.
	FirstName param.Field[string] `json:"first_name"`
	// A list of identifications for the legal entity.
	Identifications param.Field[[]IdentificationCreateRequestParam] `json:"identifications"`
	// A list of industry classifications for the legal entity.
	IndustryClassifications param.Field[[]LegalEntityIndustryClassificationParam] `json:"industry_classifications"`
	// An individual's last name.
	LastName param.Field[string] `json:"last_name"`
	// The type of legal entity.
	LegalEntityType param.Field[ChildLegalEntityCreateLegalEntityType] `json:"legal_entity_type"`
	// The business's legal structure.
	LegalStructure param.Field[ChildLegalEntityCreateLegalStructure] `json:"legal_structure"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// An individual's middle name.
	MiddleName   param.Field[string]                                   `json:"middle_name"`
	PhoneNumbers param.Field[[]ChildLegalEntityCreatePhoneNumberParam] `json:"phone_numbers"`
	// Whether the individual is a politically exposed person.
	PoliticallyExposedPerson param.Field[bool] `json:"politically_exposed_person"`
	// An individual's preferred name.
	PreferredName param.Field[string] `json:"preferred_name"`
	// An individual's prefix.
	Prefix param.Field[string] `json:"prefix"`
	// The risk rating of the legal entity. One of low, medium, high.
	RiskRating param.Field[ChildLegalEntityCreateRiskRating] `json:"risk_rating"`
	// An individual's suffix.
	Suffix                     param.Field[string]                                         `json:"suffix"`
	WealthAndEmploymentDetails param.Field[moderntreasury.WealthAndEmploymentDetailsParam] `json:"wealth_and_employment_details"`
	// The entity's primary website URL.
	Website param.Field[string] `json:"website"`
}

func (r ChildLegalEntityCreateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of legal entity.
type ChildLegalEntityCreateLegalEntityType string

const (
	ChildLegalEntityCreateLegalEntityTypeBusiness   ChildLegalEntityCreateLegalEntityType = "business"
	ChildLegalEntityCreateLegalEntityTypeIndividual ChildLegalEntityCreateLegalEntityType = "individual"
)

func (r ChildLegalEntityCreateLegalEntityType) IsKnown() bool {
	switch r {
	case ChildLegalEntityCreateLegalEntityTypeBusiness, ChildLegalEntityCreateLegalEntityTypeIndividual:
		return true
	}
	return false
}

// The business's legal structure.
type ChildLegalEntityCreateLegalStructure string

const (
	ChildLegalEntityCreateLegalStructureCorporation        ChildLegalEntityCreateLegalStructure = "corporation"
	ChildLegalEntityCreateLegalStructureLlc                ChildLegalEntityCreateLegalStructure = "llc"
	ChildLegalEntityCreateLegalStructureNonProfit          ChildLegalEntityCreateLegalStructure = "non_profit"
	ChildLegalEntityCreateLegalStructurePartnership        ChildLegalEntityCreateLegalStructure = "partnership"
	ChildLegalEntityCreateLegalStructureSoleProprietorship ChildLegalEntityCreateLegalStructure = "sole_proprietorship"
	ChildLegalEntityCreateLegalStructureTrust              ChildLegalEntityCreateLegalStructure = "trust"
)

func (r ChildLegalEntityCreateLegalStructure) IsKnown() bool {
	switch r {
	case ChildLegalEntityCreateLegalStructureCorporation, ChildLegalEntityCreateLegalStructureLlc, ChildLegalEntityCreateLegalStructureNonProfit, ChildLegalEntityCreateLegalStructurePartnership, ChildLegalEntityCreateLegalStructureSoleProprietorship, ChildLegalEntityCreateLegalStructureTrust:
		return true
	}
	return false
}

// A list of phone numbers in E.164 format.
type ChildLegalEntityCreatePhoneNumberParam struct {
	PhoneNumber param.Field[string] `json:"phone_number"`
}

func (r ChildLegalEntityCreatePhoneNumberParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The risk rating of the legal entity. One of low, medium, high.
type ChildLegalEntityCreateRiskRating string

const (
	ChildLegalEntityCreateRiskRatingLow    ChildLegalEntityCreateRiskRating = "low"
	ChildLegalEntityCreateRiskRatingMedium ChildLegalEntityCreateRiskRating = "medium"
	ChildLegalEntityCreateRiskRatingHigh   ChildLegalEntityCreateRiskRating = "high"
)

func (r ChildLegalEntityCreateRiskRating) IsKnown() bool {
	switch r {
	case ChildLegalEntityCreateRiskRatingLow, ChildLegalEntityCreateRiskRatingMedium, ChildLegalEntityCreateRiskRatingHigh:
		return true
	}
	return false
}

type ContactDetail struct {
	ID                    string                             `json:"id,required" format:"uuid"`
	ContactIdentifier     string                             `json:"contact_identifier,required"`
	ContactIdentifierType ContactDetailContactIdentifierType `json:"contact_identifier_type,required"`
	CreatedAt             time.Time                          `json:"created_at,required" format:"date-time"`
	DiscardedAt           time.Time                          `json:"discarded_at,required,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool              `json:"live_mode,required"`
	Object    string            `json:"object,required"`
	UpdatedAt time.Time         `json:"updated_at,required" format:"date-time"`
	JSON      contactDetailJSON `json:"-"`
}

// contactDetailJSON contains the JSON metadata for the struct [ContactDetail]
type contactDetailJSON struct {
	ID                    apijson.Field
	ContactIdentifier     apijson.Field
	ContactIdentifierType apijson.Field
	CreatedAt             apijson.Field
	DiscardedAt           apijson.Field
	LiveMode              apijson.Field
	Object                apijson.Field
	UpdatedAt             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContactDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contactDetailJSON) RawJSON() string {
	return r.raw
}

type ContactDetailContactIdentifierType string

const (
	ContactDetailContactIdentifierTypeEmail       ContactDetailContactIdentifierType = "email"
	ContactDetailContactIdentifierTypePhoneNumber ContactDetailContactIdentifierType = "phone_number"
	ContactDetailContactIdentifierTypeWebsite     ContactDetailContactIdentifierType = "website"
)

func (r ContactDetailContactIdentifierType) IsKnown() bool {
	switch r {
	case ContactDetailContactIdentifierTypeEmail, ContactDetailContactIdentifierTypePhoneNumber, ContactDetailContactIdentifierTypeWebsite:
		return true
	}
	return false
}

type ContactDetailParam struct {
	ID                    param.Field[string]                             `json:"id,required" format:"uuid"`
	ContactIdentifier     param.Field[string]                             `json:"contact_identifier,required"`
	ContactIdentifierType param.Field[ContactDetailContactIdentifierType] `json:"contact_identifier_type,required"`
	CreatedAt             param.Field[time.Time]                          `json:"created_at,required" format:"date-time"`
	DiscardedAt           param.Field[time.Time]                          `json:"discarded_at,required" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  param.Field[bool]      `json:"live_mode,required"`
	Object    param.Field[string]    `json:"object,required"`
	UpdatedAt param.Field[time.Time] `json:"updated_at,required" format:"date-time"`
}

func (r ContactDetailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Three-letter ISO currency code.
type Currency string

const (
	CurrencyAed   Currency = "AED"
	CurrencyAfn   Currency = "AFN"
	CurrencyAll   Currency = "ALL"
	CurrencyAmd   Currency = "AMD"
	CurrencyAng   Currency = "ANG"
	CurrencyAoa   Currency = "AOA"
	CurrencyArs   Currency = "ARS"
	CurrencyAud   Currency = "AUD"
	CurrencyAwg   Currency = "AWG"
	CurrencyAzn   Currency = "AZN"
	CurrencyBam   Currency = "BAM"
	CurrencyBbd   Currency = "BBD"
	CurrencyBch   Currency = "BCH"
	CurrencyBdt   Currency = "BDT"
	CurrencyBgn   Currency = "BGN"
	CurrencyBhd   Currency = "BHD"
	CurrencyBif   Currency = "BIF"
	CurrencyBmd   Currency = "BMD"
	CurrencyBnd   Currency = "BND"
	CurrencyBob   Currency = "BOB"
	CurrencyBrl   Currency = "BRL"
	CurrencyBsd   Currency = "BSD"
	CurrencyBtc   Currency = "BTC"
	CurrencyBtn   Currency = "BTN"
	CurrencyBwp   Currency = "BWP"
	CurrencyByn   Currency = "BYN"
	CurrencyByr   Currency = "BYR"
	CurrencyBzd   Currency = "BZD"
	CurrencyCad   Currency = "CAD"
	CurrencyCdf   Currency = "CDF"
	CurrencyChf   Currency = "CHF"
	CurrencyClf   Currency = "CLF"
	CurrencyClp   Currency = "CLP"
	CurrencyCnh   Currency = "CNH"
	CurrencyCny   Currency = "CNY"
	CurrencyCop   Currency = "COP"
	CurrencyCrc   Currency = "CRC"
	CurrencyCuc   Currency = "CUC"
	CurrencyCup   Currency = "CUP"
	CurrencyCve   Currency = "CVE"
	CurrencyCzk   Currency = "CZK"
	CurrencyDjf   Currency = "DJF"
	CurrencyDkk   Currency = "DKK"
	CurrencyDop   Currency = "DOP"
	CurrencyDzd   Currency = "DZD"
	CurrencyEek   Currency = "EEK"
	CurrencyEgp   Currency = "EGP"
	CurrencyErn   Currency = "ERN"
	CurrencyEtb   Currency = "ETB"
	CurrencyEur   Currency = "EUR"
	CurrencyEurc  Currency = "EURC"
	CurrencyFjd   Currency = "FJD"
	CurrencyFkp   Currency = "FKP"
	CurrencyGbp   Currency = "GBP"
	CurrencyGbx   Currency = "GBX"
	CurrencyGel   Currency = "GEL"
	CurrencyGgp   Currency = "GGP"
	CurrencyGhs   Currency = "GHS"
	CurrencyGip   Currency = "GIP"
	CurrencyGmd   Currency = "GMD"
	CurrencyGnf   Currency = "GNF"
	CurrencyGtq   Currency = "GTQ"
	CurrencyGyd   Currency = "GYD"
	CurrencyHkd   Currency = "HKD"
	CurrencyHnl   Currency = "HNL"
	CurrencyHrk   Currency = "HRK"
	CurrencyHtg   Currency = "HTG"
	CurrencyHuf   Currency = "HUF"
	CurrencyIdr   Currency = "IDR"
	CurrencyIls   Currency = "ILS"
	CurrencyImp   Currency = "IMP"
	CurrencyInr   Currency = "INR"
	CurrencyIqd   Currency = "IQD"
	CurrencyIrr   Currency = "IRR"
	CurrencyIsk   Currency = "ISK"
	CurrencyJep   Currency = "JEP"
	CurrencyJmd   Currency = "JMD"
	CurrencyJod   Currency = "JOD"
	CurrencyJpy   Currency = "JPY"
	CurrencyKes   Currency = "KES"
	CurrencyKgs   Currency = "KGS"
	CurrencyKhr   Currency = "KHR"
	CurrencyKmf   Currency = "KMF"
	CurrencyKpw   Currency = "KPW"
	CurrencyKrw   Currency = "KRW"
	CurrencyKwd   Currency = "KWD"
	CurrencyKyd   Currency = "KYD"
	CurrencyKzt   Currency = "KZT"
	CurrencyLak   Currency = "LAK"
	CurrencyLbp   Currency = "LBP"
	CurrencyLkr   Currency = "LKR"
	CurrencyLrd   Currency = "LRD"
	CurrencyLsl   Currency = "LSL"
	CurrencyLtl   Currency = "LTL"
	CurrencyLvl   Currency = "LVL"
	CurrencyLyd   Currency = "LYD"
	CurrencyMad   Currency = "MAD"
	CurrencyMdl   Currency = "MDL"
	CurrencyMga   Currency = "MGA"
	CurrencyMkd   Currency = "MKD"
	CurrencyMmk   Currency = "MMK"
	CurrencyMnt   Currency = "MNT"
	CurrencyMop   Currency = "MOP"
	CurrencyMro   Currency = "MRO"
	CurrencyMru   Currency = "MRU"
	CurrencyMtl   Currency = "MTL"
	CurrencyMur   Currency = "MUR"
	CurrencyMvr   Currency = "MVR"
	CurrencyMwk   Currency = "MWK"
	CurrencyMxn   Currency = "MXN"
	CurrencyMyr   Currency = "MYR"
	CurrencyMzn   Currency = "MZN"
	CurrencyNad   Currency = "NAD"
	CurrencyNgn   Currency = "NGN"
	CurrencyNio   Currency = "NIO"
	CurrencyNok   Currency = "NOK"
	CurrencyNpr   Currency = "NPR"
	CurrencyNzd   Currency = "NZD"
	CurrencyOmr   Currency = "OMR"
	CurrencyPab   Currency = "PAB"
	CurrencyPen   Currency = "PEN"
	CurrencyPgk   Currency = "PGK"
	CurrencyPhp   Currency = "PHP"
	CurrencyPkr   Currency = "PKR"
	CurrencyPln   Currency = "PLN"
	CurrencyPyg   Currency = "PYG"
	CurrencyPyusd Currency = "PYUSD"
	CurrencyQar   Currency = "QAR"
	CurrencyRon   Currency = "RON"
	CurrencyRsd   Currency = "RSD"
	CurrencyRub   Currency = "RUB"
	CurrencyRwf   Currency = "RWF"
	CurrencySar   Currency = "SAR"
	CurrencySbc   Currency = "SBC"
	CurrencySbd   Currency = "SBD"
	CurrencyScr   Currency = "SCR"
	CurrencySdg   Currency = "SDG"
	CurrencySek   Currency = "SEK"
	CurrencySgd   Currency = "SGD"
	CurrencyShp   Currency = "SHP"
	CurrencySkk   Currency = "SKK"
	CurrencySll   Currency = "SLL"
	CurrencySos   Currency = "SOS"
	CurrencySrd   Currency = "SRD"
	CurrencySsp   Currency = "SSP"
	CurrencyStd   Currency = "STD"
	CurrencySvc   Currency = "SVC"
	CurrencySyp   Currency = "SYP"
	CurrencySzl   Currency = "SZL"
	CurrencyThb   Currency = "THB"
	CurrencyTjs   Currency = "TJS"
	CurrencyTmm   Currency = "TMM"
	CurrencyTmt   Currency = "TMT"
	CurrencyTnd   Currency = "TND"
	CurrencyTop   Currency = "TOP"
	CurrencyTry   Currency = "TRY"
	CurrencyTtd   Currency = "TTD"
	CurrencyTwd   Currency = "TWD"
	CurrencyTzs   Currency = "TZS"
	CurrencyUah   Currency = "UAH"
	CurrencyUgx   Currency = "UGX"
	CurrencyUsd   Currency = "USD"
	CurrencyUsdb  Currency = "USDB"
	CurrencyUsdc  Currency = "USDC"
	CurrencyUsdp  Currency = "USDP"
	CurrencyUsdt  Currency = "USDT"
	CurrencyUyu   Currency = "UYU"
	CurrencyUzs   Currency = "UZS"
	CurrencyVef   Currency = "VEF"
	CurrencyVes   Currency = "VES"
	CurrencyVnd   Currency = "VND"
	CurrencyVuv   Currency = "VUV"
	CurrencyWst   Currency = "WST"
	CurrencyXaf   Currency = "XAF"
	CurrencyXag   Currency = "XAG"
	CurrencyXau   Currency = "XAU"
	CurrencyXba   Currency = "XBA"
	CurrencyXbb   Currency = "XBB"
	CurrencyXbc   Currency = "XBC"
	CurrencyXbd   Currency = "XBD"
	CurrencyXcd   Currency = "XCD"
	CurrencyXdr   Currency = "XDR"
	CurrencyXfu   Currency = "XFU"
	CurrencyXof   Currency = "XOF"
	CurrencyXpd   Currency = "XPD"
	CurrencyXpf   Currency = "XPF"
	CurrencyXpt   Currency = "XPT"
	CurrencyXts   Currency = "XTS"
	CurrencyYer   Currency = "YER"
	CurrencyZar   Currency = "ZAR"
	CurrencyZmk   Currency = "ZMK"
	CurrencyZmw   Currency = "ZMW"
	CurrencyZwd   Currency = "ZWD"
	CurrencyZwl   Currency = "ZWL"
	CurrencyZwn   Currency = "ZWN"
	CurrencyZwr   Currency = "ZWR"
)

func (r Currency) IsKnown() bool {
	switch r {
	case CurrencyAed, CurrencyAfn, CurrencyAll, CurrencyAmd, CurrencyAng, CurrencyAoa, CurrencyArs, CurrencyAud, CurrencyAwg, CurrencyAzn, CurrencyBam, CurrencyBbd, CurrencyBch, CurrencyBdt, CurrencyBgn, CurrencyBhd, CurrencyBif, CurrencyBmd, CurrencyBnd, CurrencyBob, CurrencyBrl, CurrencyBsd, CurrencyBtc, CurrencyBtn, CurrencyBwp, CurrencyByn, CurrencyByr, CurrencyBzd, CurrencyCad, CurrencyCdf, CurrencyChf, CurrencyClf, CurrencyClp, CurrencyCnh, CurrencyCny, CurrencyCop, CurrencyCrc, CurrencyCuc, CurrencyCup, CurrencyCve, CurrencyCzk, CurrencyDjf, CurrencyDkk, CurrencyDop, CurrencyDzd, CurrencyEek, CurrencyEgp, CurrencyErn, CurrencyEtb, CurrencyEur, CurrencyEurc, CurrencyFjd, CurrencyFkp, CurrencyGbp, CurrencyGbx, CurrencyGel, CurrencyGgp, CurrencyGhs, CurrencyGip, CurrencyGmd, CurrencyGnf, CurrencyGtq, CurrencyGyd, CurrencyHkd, CurrencyHnl, CurrencyHrk, CurrencyHtg, CurrencyHuf, CurrencyIdr, CurrencyIls, CurrencyImp, CurrencyInr, CurrencyIqd, CurrencyIrr, CurrencyIsk, CurrencyJep, CurrencyJmd, CurrencyJod, CurrencyJpy, CurrencyKes, CurrencyKgs, CurrencyKhr, CurrencyKmf, CurrencyKpw, CurrencyKrw, CurrencyKwd, CurrencyKyd, CurrencyKzt, CurrencyLak, CurrencyLbp, CurrencyLkr, CurrencyLrd, CurrencyLsl, CurrencyLtl, CurrencyLvl, CurrencyLyd, CurrencyMad, CurrencyMdl, CurrencyMga, CurrencyMkd, CurrencyMmk, CurrencyMnt, CurrencyMop, CurrencyMro, CurrencyMru, CurrencyMtl, CurrencyMur, CurrencyMvr, CurrencyMwk, CurrencyMxn, CurrencyMyr, CurrencyMzn, CurrencyNad, CurrencyNgn, CurrencyNio, CurrencyNok, CurrencyNpr, CurrencyNzd, CurrencyOmr, CurrencyPab, CurrencyPen, CurrencyPgk, CurrencyPhp, CurrencyPkr, CurrencyPln, CurrencyPyg, CurrencyPyusd, CurrencyQar, CurrencyRon, CurrencyRsd, CurrencyRub, CurrencyRwf, CurrencySar, CurrencySbc, CurrencySbd, CurrencyScr, CurrencySdg, CurrencySek, CurrencySgd, CurrencyShp, CurrencySkk, CurrencySll, CurrencySos, CurrencySrd, CurrencySsp, CurrencyStd, CurrencySvc, CurrencySyp, CurrencySzl, CurrencyThb, CurrencyTjs, CurrencyTmm, CurrencyTmt, CurrencyTnd, CurrencyTop, CurrencyTry, CurrencyTtd, CurrencyTwd, CurrencyTzs, CurrencyUah, CurrencyUgx, CurrencyUsd, CurrencyUsdb, CurrencyUsdc, CurrencyUsdp, CurrencyUsdt, CurrencyUyu, CurrencyUzs, CurrencyVef, CurrencyVes, CurrencyVnd, CurrencyVuv, CurrencyWst, CurrencyXaf, CurrencyXag, CurrencyXau, CurrencyXba, CurrencyXbb, CurrencyXbc, CurrencyXbd, CurrencyXcd, CurrencyXdr, CurrencyXfu, CurrencyXof, CurrencyXpd, CurrencyXpf, CurrencyXpt, CurrencyXts, CurrencyYer, CurrencyZar, CurrencyZmk, CurrencyZmw, CurrencyZwd, CurrencyZwl, CurrencyZwn, CurrencyZwr:
		return true
	}
	return false
}

type ForeignExchangeRate struct {
	// Amount in the lowest denomination of the `base_currency` to convert, often
	// called the "sell" amount.
	BaseAmount int64 `json:"base_amount,required"`
	// Currency to convert, often called the "sell" currency.
	BaseCurrency Currency `json:"base_currency,required"`
	// The exponent component of the rate. The decimal is calculated as `value` / (10 ^
	// `exponent`).
	Exponent int64 `json:"exponent,required"`
	// A string representation of the rate.
	RateString string `json:"rate_string,required"`
	// Amount in the lowest denomination of the `target_currency`, often called the
	// "buy" amount.
	TargetAmount int64 `json:"target_amount,required"`
	// Currency to convert the `base_currency` to, often called the "buy" currency.
	TargetCurrency Currency `json:"target_currency,required"`
	// The whole number component of the rate. The decimal is calculated as `value` /
	// (10 ^ `exponent`).
	Value int64                   `json:"value,required"`
	JSON  foreignExchangeRateJSON `json:"-"`
}

// foreignExchangeRateJSON contains the JSON metadata for the struct
// [ForeignExchangeRate]
type foreignExchangeRateJSON struct {
	BaseAmount     apijson.Field
	BaseCurrency   apijson.Field
	Exponent       apijson.Field
	RateString     apijson.Field
	TargetAmount   apijson.Field
	TargetCurrency apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ForeignExchangeRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r foreignExchangeRateJSON) RawJSON() string {
	return r.raw
}

type IdentificationCreateRequestParam struct {
	// The ID number of identification document.
	IDNumber param.Field[string] `json:"id_number,required"`
	// The type of ID number.
	IDType param.Field[IdentificationCreateRequestIDType] `json:"id_type,required"`
	// The date when the Identification is no longer considered valid by the issuing
	// authority.
	ExpirationDate param.Field[time.Time] `json:"expiration_date" format:"date"`
	// The ISO 3166-1 alpha-2 country code of the country that issued the
	// identification
	IssuingCountry param.Field[string] `json:"issuing_country"`
	// The region in which the identifcation was issued.
	IssuingRegion param.Field[string] `json:"issuing_region"`
}

func (r IdentificationCreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of ID number.
type IdentificationCreateRequestIDType string

const (
	IdentificationCreateRequestIDTypeArCuil         IdentificationCreateRequestIDType = "ar_cuil"
	IdentificationCreateRequestIDTypeArCuit         IdentificationCreateRequestIDType = "ar_cuit"
	IdentificationCreateRequestIDTypeBrCnpj         IdentificationCreateRequestIDType = "br_cnpj"
	IdentificationCreateRequestIDTypeBrCpf          IdentificationCreateRequestIDType = "br_cpf"
	IdentificationCreateRequestIDTypeClRun          IdentificationCreateRequestIDType = "cl_run"
	IdentificationCreateRequestIDTypeClRut          IdentificationCreateRequestIDType = "cl_rut"
	IdentificationCreateRequestIDTypeCoCedulas      IdentificationCreateRequestIDType = "co_cedulas"
	IdentificationCreateRequestIDTypeCoNit          IdentificationCreateRequestIDType = "co_nit"
	IdentificationCreateRequestIDTypeDriversLicense IdentificationCreateRequestIDType = "drivers_license"
	IdentificationCreateRequestIDTypeHnID           IdentificationCreateRequestIDType = "hn_id"
	IdentificationCreateRequestIDTypeHnRtn          IdentificationCreateRequestIDType = "hn_rtn"
	IdentificationCreateRequestIDTypeInLei          IdentificationCreateRequestIDType = "in_lei"
	IdentificationCreateRequestIDTypeKrBrn          IdentificationCreateRequestIDType = "kr_brn"
	IdentificationCreateRequestIDTypeKrCrn          IdentificationCreateRequestIDType = "kr_crn"
	IdentificationCreateRequestIDTypeKrRrn          IdentificationCreateRequestIDType = "kr_rrn"
	IdentificationCreateRequestIDTypePassport       IdentificationCreateRequestIDType = "passport"
	IdentificationCreateRequestIDTypeSaTin          IdentificationCreateRequestIDType = "sa_tin"
	IdentificationCreateRequestIDTypeSaVat          IdentificationCreateRequestIDType = "sa_vat"
	IdentificationCreateRequestIDTypeUsEin          IdentificationCreateRequestIDType = "us_ein"
	IdentificationCreateRequestIDTypeUsItin         IdentificationCreateRequestIDType = "us_itin"
	IdentificationCreateRequestIDTypeUsSsn          IdentificationCreateRequestIDType = "us_ssn"
	IdentificationCreateRequestIDTypeVnTin          IdentificationCreateRequestIDType = "vn_tin"
)

func (r IdentificationCreateRequestIDType) IsKnown() bool {
	switch r {
	case IdentificationCreateRequestIDTypeArCuil, IdentificationCreateRequestIDTypeArCuit, IdentificationCreateRequestIDTypeBrCnpj, IdentificationCreateRequestIDTypeBrCpf, IdentificationCreateRequestIDTypeClRun, IdentificationCreateRequestIDTypeClRut, IdentificationCreateRequestIDTypeCoCedulas, IdentificationCreateRequestIDTypeCoNit, IdentificationCreateRequestIDTypeDriversLicense, IdentificationCreateRequestIDTypeHnID, IdentificationCreateRequestIDTypeHnRtn, IdentificationCreateRequestIDTypeInLei, IdentificationCreateRequestIDTypeKrBrn, IdentificationCreateRequestIDTypeKrCrn, IdentificationCreateRequestIDTypeKrRrn, IdentificationCreateRequestIDTypePassport, IdentificationCreateRequestIDTypeSaTin, IdentificationCreateRequestIDTypeSaVat, IdentificationCreateRequestIDTypeUsEin, IdentificationCreateRequestIDTypeUsItin, IdentificationCreateRequestIDTypeUsSsn, IdentificationCreateRequestIDTypeVnTin:
		return true
	}
	return false
}

type LedgerAccountCreateRequestParam struct {
	// The currency of the ledger account.
	Currency param.Field[string] `json:"currency,required"`
	// The id of the ledger that this account belongs to.
	LedgerID param.Field[string] `json:"ledger_id,required" format:"uuid"`
	// The name of the ledger account.
	Name param.Field[string] `json:"name,required"`
	// The normal balance of the ledger account.
	NormalBalance param.Field[TransactionDirection] `json:"normal_balance,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent param.Field[int64] `json:"currency_exponent"`
	// The description of the ledger account.
	Description param.Field[string] `json:"description"`
	// The array of ledger account category ids that this ledger account should be a
	// child of.
	LedgerAccountCategoryIDs param.Field[[]string] `json:"ledger_account_category_ids" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the id will be
	// populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger account links to another object in Modern Treasury, the type will
	// be populated here, otherwise null. The value is one of internal_account or
	// external_account.
	LedgerableType param.Field[LedgerAccountCreateRequestLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
}

func (r LedgerAccountCreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LedgerAccountCreateRequestParam) ImplementsBulkRequestNewParamsResourceUnion() {}

// If the ledger account links to another object in Modern Treasury, the type will
// be populated here, otherwise null. The value is one of internal_account or
// external_account.
type LedgerAccountCreateRequestLedgerableType string

const (
	LedgerAccountCreateRequestLedgerableTypeCounterparty    LedgerAccountCreateRequestLedgerableType = "counterparty"
	LedgerAccountCreateRequestLedgerableTypeExternalAccount LedgerAccountCreateRequestLedgerableType = "external_account"
	LedgerAccountCreateRequestLedgerableTypeInternalAccount LedgerAccountCreateRequestLedgerableType = "internal_account"
	LedgerAccountCreateRequestLedgerableTypeVirtualAccount  LedgerAccountCreateRequestLedgerableType = "virtual_account"
)

func (r LedgerAccountCreateRequestLedgerableType) IsKnown() bool {
	switch r {
	case LedgerAccountCreateRequestLedgerableTypeCounterparty, LedgerAccountCreateRequestLedgerableTypeExternalAccount, LedgerAccountCreateRequestLedgerableTypeInternalAccount, LedgerAccountCreateRequestLedgerableTypeVirtualAccount:
		return true
	}
	return false
}

type LedgerBalance struct {
	Amount  int64 `json:"amount,required"`
	Credits int64 `json:"credits,required"`
	// The currency of the ledger account.
	Currency string `json:"currency,required"`
	// The currency exponent of the ledger account.
	CurrencyExponent int64             `json:"currency_exponent,required"`
	Debits           int64             `json:"debits,required"`
	JSON             ledgerBalanceJSON `json:"-"`
}

// ledgerBalanceJSON contains the JSON metadata for the struct [LedgerBalance]
type ledgerBalanceJSON struct {
	Amount           apijson.Field
	Credits          apijson.Field
	Currency         apijson.Field
	CurrencyExponent apijson.Field
	Debits           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerBalanceJSON) RawJSON() string {
	return r.raw
}

type LedgerBalances struct {
	// The available_balance is the sum of all posted inbound entries and pending
	// outbound entries. For credit normal, available_amount = posted_credits -
	// pending_debits; for debit normal, available_amount = posted_debits -
	// pending_credits.
	AvailableBalance LedgerBalance `json:"available_balance,required"`
	// The pending_balance is the sum of all pending and posted entries.
	PendingBalance LedgerBalance `json:"pending_balance,required"`
	// The posted_balance is the sum of all posted entries.
	PostedBalance LedgerBalance      `json:"posted_balance,required"`
	JSON          ledgerBalancesJSON `json:"-"`
}

// ledgerBalancesJSON contains the JSON metadata for the struct [LedgerBalances]
type ledgerBalancesJSON struct {
	AvailableBalance apijson.Field
	PendingBalance   apijson.Field
	PostedBalance    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LedgerBalances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ledgerBalancesJSON) RawJSON() string {
	return r.raw
}

type LedgerEntryCreateRequestParam struct {
	// Value in specified currency's smallest unit. e.g. $10 would be represented
	// as 1000. Can be any integer up to 36 digits.
	Amount param.Field[int64] `json:"amount,required"`
	// One of `credit`, `debit`. Describes the direction money is flowing in the
	// transaction. A `credit` moves money from your account to someone else's. A
	// `debit` pulls money from someone else's account to your own. Note that wire,
	// rtp, and check payments will always be `credit`.
	Direction param.Field[TransactionDirection] `json:"direction,required"`
	// The ledger account that this ledger entry is associated with.
	LedgerAccountID param.Field[string] `json:"ledger_account_id,required" format:"uuid"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s available balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	AvailableBalanceAmount param.Field[map[string]int64] `json:"available_balance_amount"`
	// Lock version of the ledger account. This can be passed when creating a ledger
	// transaction to only succeed if no ledger transactions have posted since the
	// given version. See our post about Designing the Ledgers API with Optimistic
	// Locking for more details.
	LockVersion param.Field[int64] `json:"lock_version"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s pending balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PendingBalanceAmount param.Field[map[string]int64] `json:"pending_balance_amount"`
	// Use `gt` (>), `gte` (>=), `lt` (<), `lte` (<=), or `eq` (=) to lock on the
	// account’s posted balance. If any of these conditions would be false after the
	// transaction is created, the entire call will fail with error code 422.
	PostedBalanceAmount param.Field[map[string]int64] `json:"posted_balance_amount"`
	// If true, response will include the balance of the associated ledger account for
	// the entry.
	ShowResultingLedgerAccountBalances param.Field[bool] `json:"show_resulting_ledger_account_balances"`
}

func (r LedgerEntryCreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LedgerTransactionCreateRequestParam struct {
	// An array of ledger entry objects.
	LedgerEntries param.Field[[]LedgerEntryCreateRequestParam] `json:"ledger_entries,required"`
	// An optional description for internal use.
	Description param.Field[string] `json:"description"`
	// The timestamp (ISO8601 format) at which the ledger transaction happened for
	// reporting purposes.
	EffectiveAt param.Field[time.Time] `json:"effective_at" format:"date-time"`
	// The date (YYYY-MM-DD) on which the ledger transaction happened for reporting
	// purposes.
	EffectiveDate param.Field[time.Time] `json:"effective_date" format:"date"`
	// A unique string to represent the ledger transaction. Only one pending or posted
	// ledger transaction may have this ID in the ledger.
	ExternalID param.Field[string] `json:"external_id"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the id will be populated here, otherwise null.
	LedgerableID param.Field[string] `json:"ledgerable_id" format:"uuid"`
	// If the ledger transaction can be reconciled to another object in Modern
	// Treasury, the type will be populated here, otherwise null. This can be one of
	// payment_order, incoming_payment_detail, expected_payment, return, paper_item, or
	// reversal.
	LedgerableType param.Field[LedgerTransactionCreateRequestLedgerableType] `json:"ledgerable_type"`
	// Additional data represented as key-value pairs. Both the key and value must be
	// strings.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// To post a ledger transaction at creation, use `posted`.
	Status param.Field[LedgerTransactionCreateRequestStatus] `json:"status"`
}

func (r LedgerTransactionCreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LedgerTransactionCreateRequestParam) ImplementsBulkRequestNewParamsResourceUnion() {}

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, paper_item, or
// reversal.
type LedgerTransactionCreateRequestLedgerableType string

const (
	LedgerTransactionCreateRequestLedgerableTypeExpectedPayment       LedgerTransactionCreateRequestLedgerableType = "expected_payment"
	LedgerTransactionCreateRequestLedgerableTypeIncomingPaymentDetail LedgerTransactionCreateRequestLedgerableType = "incoming_payment_detail"
	LedgerTransactionCreateRequestLedgerableTypePaperItem             LedgerTransactionCreateRequestLedgerableType = "paper_item"
	LedgerTransactionCreateRequestLedgerableTypePaymentOrder          LedgerTransactionCreateRequestLedgerableType = "payment_order"
	LedgerTransactionCreateRequestLedgerableTypeReturn                LedgerTransactionCreateRequestLedgerableType = "return"
	LedgerTransactionCreateRequestLedgerableTypeReversal              LedgerTransactionCreateRequestLedgerableType = "reversal"
)

func (r LedgerTransactionCreateRequestLedgerableType) IsKnown() bool {
	switch r {
	case LedgerTransactionCreateRequestLedgerableTypeExpectedPayment, LedgerTransactionCreateRequestLedgerableTypeIncomingPaymentDetail, LedgerTransactionCreateRequestLedgerableTypePaperItem, LedgerTransactionCreateRequestLedgerableTypePaymentOrder, LedgerTransactionCreateRequestLedgerableTypeReturn, LedgerTransactionCreateRequestLedgerableTypeReversal:
		return true
	}
	return false
}

// To post a ledger transaction at creation, use `posted`.
type LedgerTransactionCreateRequestStatus string

const (
	LedgerTransactionCreateRequestStatusArchived LedgerTransactionCreateRequestStatus = "archived"
	LedgerTransactionCreateRequestStatusPending  LedgerTransactionCreateRequestStatus = "pending"
	LedgerTransactionCreateRequestStatusPosted   LedgerTransactionCreateRequestStatus = "posted"
)

func (r LedgerTransactionCreateRequestStatus) IsKnown() bool {
	switch r {
	case LedgerTransactionCreateRequestStatusArchived, LedgerTransactionCreateRequestStatusPending, LedgerTransactionCreateRequestStatusPosted:
		return true
	}
	return false
}

type LegalEntityAddressCreateRequestParam struct {
	// Country code conforms to [ISO 3166-1 alpha-2]
	Country param.Field[string] `json:"country,required"`
	Line1   param.Field[string] `json:"line1,required"`
	// Locality or City.
	Locality param.Field[string] `json:"locality,required"`
	// The postal code of the address.
	PostalCode param.Field[string] `json:"postal_code,required"`
	// Region or State.
	Region param.Field[string] `json:"region,required"`
	// The types of this address.
	AddressTypes param.Field[[]LegalEntityAddressCreateRequestAddressType] `json:"address_types"`
	Line2        param.Field[string]                                       `json:"line2"`
}

func (r LegalEntityAddressCreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityAddressCreateRequestAddressType string

const (
	LegalEntityAddressCreateRequestAddressTypeBusiness    LegalEntityAddressCreateRequestAddressType = "business"
	LegalEntityAddressCreateRequestAddressTypeMailing     LegalEntityAddressCreateRequestAddressType = "mailing"
	LegalEntityAddressCreateRequestAddressTypeOther       LegalEntityAddressCreateRequestAddressType = "other"
	LegalEntityAddressCreateRequestAddressTypePoBox       LegalEntityAddressCreateRequestAddressType = "po_box"
	LegalEntityAddressCreateRequestAddressTypeResidential LegalEntityAddressCreateRequestAddressType = "residential"
)

func (r LegalEntityAddressCreateRequestAddressType) IsKnown() bool {
	switch r {
	case LegalEntityAddressCreateRequestAddressTypeBusiness, LegalEntityAddressCreateRequestAddressTypeMailing, LegalEntityAddressCreateRequestAddressTypeOther, LegalEntityAddressCreateRequestAddressTypePoBox, LegalEntityAddressCreateRequestAddressTypeResidential:
		return true
	}
	return false
}

type LegalEntityComplianceDetail struct {
	ID          string    `json:"id,required" format:"uuid"`
	CreatedAt   time.Time `json:"created_at,required" format:"date-time"`
	DiscardedAt time.Time `json:"discarded_at,required,nullable" format:"date-time"`
	// The issuer of the compliance token.
	Issuer string `json:"issuer,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode bool   `json:"live_mode,required"`
	Object   string `json:"object,required"`
	// The timestamp when the compliance token expires.
	TokenExpiresAt time.Time `json:"token_expires_at,required,nullable" format:"date-time"`
	// The timestamp when the compliance token was issued.
	TokenIssuedAt time.Time `json:"token_issued_at,required,nullable" format:"date-time"`
	// The URL to the compliance token. (ex. provider portal URL)
	TokenURL  string    `json:"token_url,required,nullable"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// Whether entity corresponding to the compliance token has been validated.
	Validated bool `json:"validated,required"`
	// The timestamp when the entity was validated.
	ValidatedAt time.Time                       `json:"validated_at,required,nullable" format:"date-time"`
	JSON        legalEntityComplianceDetailJSON `json:"-"`
}

// legalEntityComplianceDetailJSON contains the JSON metadata for the struct
// [LegalEntityComplianceDetail]
type legalEntityComplianceDetailJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	DiscardedAt    apijson.Field
	Issuer         apijson.Field
	LiveMode       apijson.Field
	Object         apijson.Field
	TokenExpiresAt apijson.Field
	TokenIssuedAt  apijson.Field
	TokenURL       apijson.Field
	UpdatedAt      apijson.Field
	Validated      apijson.Field
	ValidatedAt    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LegalEntityComplianceDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityComplianceDetailJSON) RawJSON() string {
	return r.raw
}

type LegalEntityComplianceDetailParam struct {
	ID          param.Field[string]    `json:"id,required" format:"uuid"`
	CreatedAt   param.Field[time.Time] `json:"created_at,required" format:"date-time"`
	DiscardedAt param.Field[time.Time] `json:"discarded_at,required" format:"date-time"`
	// The issuer of the compliance token.
	Issuer param.Field[string] `json:"issuer,required"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode param.Field[bool]   `json:"live_mode,required"`
	Object   param.Field[string] `json:"object,required"`
	// The timestamp when the compliance token expires.
	TokenExpiresAt param.Field[time.Time] `json:"token_expires_at,required" format:"date-time"`
	// The timestamp when the compliance token was issued.
	TokenIssuedAt param.Field[time.Time] `json:"token_issued_at,required" format:"date-time"`
	// The URL to the compliance token. (ex. provider portal URL)
	TokenURL  param.Field[string]    `json:"token_url,required"`
	UpdatedAt param.Field[time.Time] `json:"updated_at,required" format:"date-time"`
	// Whether entity corresponding to the compliance token has been validated.
	Validated param.Field[bool] `json:"validated,required"`
	// The timestamp when the entity was validated.
	ValidatedAt param.Field[time.Time] `json:"validated_at,required" format:"date-time"`
}

func (r LegalEntityComplianceDetailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LegalEntityIndustryClassification struct {
	ID string `json:"id,required" format:"uuid"`
	// The industry classification codes for the legal entity.
	ClassificationCodes []string `json:"classification_codes,required"`
	// The classification system of the classification codes.
	ClassificationType LegalEntityIndustryClassificationClassificationType `json:"classification_type,required"`
	CreatedAt          time.Time                                           `json:"created_at,required" format:"date-time"`
	DiscardedAt        time.Time                                           `json:"discarded_at,required,nullable" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  bool                                  `json:"live_mode,required"`
	Object    string                                `json:"object,required"`
	UpdatedAt time.Time                             `json:"updated_at,required" format:"date-time"`
	JSON      legalEntityIndustryClassificationJSON `json:"-"`
}

// legalEntityIndustryClassificationJSON contains the JSON metadata for the struct
// [LegalEntityIndustryClassification]
type legalEntityIndustryClassificationJSON struct {
	ID                  apijson.Field
	ClassificationCodes apijson.Field
	ClassificationType  apijson.Field
	CreatedAt           apijson.Field
	DiscardedAt         apijson.Field
	LiveMode            apijson.Field
	Object              apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LegalEntityIndustryClassification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legalEntityIndustryClassificationJSON) RawJSON() string {
	return r.raw
}

// The classification system of the classification codes.
type LegalEntityIndustryClassificationClassificationType string

const (
	LegalEntityIndustryClassificationClassificationTypeAnzsic LegalEntityIndustryClassificationClassificationType = "anzsic"
	LegalEntityIndustryClassificationClassificationTypeBics   LegalEntityIndustryClassificationClassificationType = "bics"
	LegalEntityIndustryClassificationClassificationTypeGics   LegalEntityIndustryClassificationClassificationType = "gics"
	LegalEntityIndustryClassificationClassificationTypeHsics  LegalEntityIndustryClassificationClassificationType = "hsics"
	LegalEntityIndustryClassificationClassificationTypeIcb    LegalEntityIndustryClassificationClassificationType = "icb"
	LegalEntityIndustryClassificationClassificationTypeIsic   LegalEntityIndustryClassificationClassificationType = "isic"
	LegalEntityIndustryClassificationClassificationTypeMgecs  LegalEntityIndustryClassificationClassificationType = "mgecs"
	LegalEntityIndustryClassificationClassificationTypeNace   LegalEntityIndustryClassificationClassificationType = "nace"
	LegalEntityIndustryClassificationClassificationTypeNaics  LegalEntityIndustryClassificationClassificationType = "naics"
	LegalEntityIndustryClassificationClassificationTypeRbics  LegalEntityIndustryClassificationClassificationType = "rbics"
	LegalEntityIndustryClassificationClassificationTypeSic    LegalEntityIndustryClassificationClassificationType = "sic"
	LegalEntityIndustryClassificationClassificationTypeSni    LegalEntityIndustryClassificationClassificationType = "sni"
	LegalEntityIndustryClassificationClassificationTypeTrbc   LegalEntityIndustryClassificationClassificationType = "trbc"
	LegalEntityIndustryClassificationClassificationTypeUksic  LegalEntityIndustryClassificationClassificationType = "uksic"
	LegalEntityIndustryClassificationClassificationTypeUnspsc LegalEntityIndustryClassificationClassificationType = "unspsc"
)

func (r LegalEntityIndustryClassificationClassificationType) IsKnown() bool {
	switch r {
	case LegalEntityIndustryClassificationClassificationTypeAnzsic, LegalEntityIndustryClassificationClassificationTypeBics, LegalEntityIndustryClassificationClassificationTypeGics, LegalEntityIndustryClassificationClassificationTypeHsics, LegalEntityIndustryClassificationClassificationTypeIcb, LegalEntityIndustryClassificationClassificationTypeIsic, LegalEntityIndustryClassificationClassificationTypeMgecs, LegalEntityIndustryClassificationClassificationTypeNace, LegalEntityIndustryClassificationClassificationTypeNaics, LegalEntityIndustryClassificationClassificationTypeRbics, LegalEntityIndustryClassificationClassificationTypeSic, LegalEntityIndustryClassificationClassificationTypeSni, LegalEntityIndustryClassificationClassificationTypeTrbc, LegalEntityIndustryClassificationClassificationTypeUksic, LegalEntityIndustryClassificationClassificationTypeUnspsc:
		return true
	}
	return false
}

type LegalEntityIndustryClassificationParam struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// The industry classification codes for the legal entity.
	ClassificationCodes param.Field[[]string] `json:"classification_codes,required"`
	// The classification system of the classification codes.
	ClassificationType param.Field[LegalEntityIndustryClassificationClassificationType] `json:"classification_type,required"`
	CreatedAt          param.Field[time.Time]                                           `json:"created_at,required" format:"date-time"`
	DiscardedAt        param.Field[time.Time]                                           `json:"discarded_at,required" format:"date-time"`
	// This field will be true if this object exists in the live environment or false
	// if it exists in the test environment.
	LiveMode  param.Field[bool]      `json:"live_mode,required"`
	Object    param.Field[string]    `json:"object,required"`
	UpdatedAt param.Field[time.Time] `json:"updated_at,required" format:"date-time"`
}

func (r LegalEntityIndustryClassificationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionDirection string

const (
	TransactionDirectionCredit TransactionDirection = "credit"
	TransactionDirectionDebit  TransactionDirection = "debit"
)

func (r TransactionDirection) IsKnown() bool {
	switch r {
	case TransactionDirectionCredit, TransactionDirectionDebit:
		return true
	}
	return false
}
