// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package moderntreasury

import (
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apierror"
	"github.com/Modern-Treasury/modern-treasury-go/v2/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type AccountsType = shared.AccountsType

// This is an alias to an internal value.
const AccountsTypeExternalAccounts = shared.AccountsTypeExternalAccounts

// This is an alias to an internal value.
const AccountsTypeInternalAccounts = shared.AccountsTypeInternalAccounts

// This is an alias to an internal type.
type Address = shared.Address

// This is an alias to an internal type.
type AddressRequest = shared.AddressRequest

// This is an alias to an internal type.
type AddressRequestParam = shared.AddressRequestParam

// This is an alias to an internal type.
type AsyncResponse = shared.AsyncResponse

// This is an alias to an internal type.
type ChildLegalEntityCreateParam = shared.ChildLegalEntityCreateParam

// The type of legal entity.
//
// This is an alias to an internal type.
type ChildLegalEntityCreateLegalEntityType = shared.ChildLegalEntityCreateLegalEntityType

// This is an alias to an internal value.
const ChildLegalEntityCreateLegalEntityTypeBusiness = shared.ChildLegalEntityCreateLegalEntityTypeBusiness

// This is an alias to an internal value.
const ChildLegalEntityCreateLegalEntityTypeIndividual = shared.ChildLegalEntityCreateLegalEntityTypeIndividual

// The business's legal structure.
//
// This is an alias to an internal type.
type ChildLegalEntityCreateLegalStructure = shared.ChildLegalEntityCreateLegalStructure

// This is an alias to an internal value.
const ChildLegalEntityCreateLegalStructureCorporation = shared.ChildLegalEntityCreateLegalStructureCorporation

// This is an alias to an internal value.
const ChildLegalEntityCreateLegalStructureLlc = shared.ChildLegalEntityCreateLegalStructureLlc

// This is an alias to an internal value.
const ChildLegalEntityCreateLegalStructureNonProfit = shared.ChildLegalEntityCreateLegalStructureNonProfit

// This is an alias to an internal value.
const ChildLegalEntityCreateLegalStructurePartnership = shared.ChildLegalEntityCreateLegalStructurePartnership

// This is an alias to an internal value.
const ChildLegalEntityCreateLegalStructureSoleProprietorship = shared.ChildLegalEntityCreateLegalStructureSoleProprietorship

// This is an alias to an internal value.
const ChildLegalEntityCreateLegalStructureTrust = shared.ChildLegalEntityCreateLegalStructureTrust

// A list of phone numbers in E.164 format.
//
// This is an alias to an internal type.
type ChildLegalEntityCreatePhoneNumberParam = shared.ChildLegalEntityCreatePhoneNumberParam

// This is an alias to an internal type.
type ChildLegalEntityCreateRegulatorParam = shared.ChildLegalEntityCreateRegulatorParam

// The risk rating of the legal entity. One of low, medium, high.
//
// This is an alias to an internal type.
type ChildLegalEntityCreateRiskRating = shared.ChildLegalEntityCreateRiskRating

// This is an alias to an internal value.
const ChildLegalEntityCreateRiskRatingLow = shared.ChildLegalEntityCreateRiskRatingLow

// This is an alias to an internal value.
const ChildLegalEntityCreateRiskRatingMedium = shared.ChildLegalEntityCreateRiskRatingMedium

// This is an alias to an internal value.
const ChildLegalEntityCreateRiskRatingHigh = shared.ChildLegalEntityCreateRiskRatingHigh

// The activation status of the legal entity. One of pending, active, suspended, or
// closed.
//
// This is an alias to an internal type.
type ChildLegalEntityCreateStatus = shared.ChildLegalEntityCreateStatus

// This is an alias to an internal value.
const ChildLegalEntityCreateStatusActive = shared.ChildLegalEntityCreateStatusActive

// This is an alias to an internal value.
const ChildLegalEntityCreateStatusClosed = shared.ChildLegalEntityCreateStatusClosed

// This is an alias to an internal value.
const ChildLegalEntityCreateStatusPending = shared.ChildLegalEntityCreateStatusPending

// This is an alias to an internal value.
const ChildLegalEntityCreateStatusSuspended = shared.ChildLegalEntityCreateStatusSuspended

// Information describing a third-party verification run by an external vendor.
//
// This is an alias to an internal type.
type ChildLegalEntityCreateThirdPartyVerificationParam = shared.ChildLegalEntityCreateThirdPartyVerificationParam

// The vendor that performed the verification, e.g. `persona`.
//
// This is an alias to an internal type.
type ChildLegalEntityCreateThirdPartyVerificationVendor = shared.ChildLegalEntityCreateThirdPartyVerificationVendor

// This is an alias to an internal value.
const ChildLegalEntityCreateThirdPartyVerificationVendorPersona = shared.ChildLegalEntityCreateThirdPartyVerificationVendorPersona

// This is an alias to an internal type.
type ContactDetail = shared.ContactDetail

// This is an alias to an internal type.
type ContactDetailContactIdentifierType = shared.ContactDetailContactIdentifierType

// This is an alias to an internal value.
const ContactDetailContactIdentifierTypeEmail = shared.ContactDetailContactIdentifierTypeEmail

// This is an alias to an internal value.
const ContactDetailContactIdentifierTypePhoneNumber = shared.ContactDetailContactIdentifierTypePhoneNumber

// This is an alias to an internal value.
const ContactDetailContactIdentifierTypeWebsite = shared.ContactDetailContactIdentifierTypeWebsite

// This is an alias to an internal type.
type ContactDetailParam = shared.ContactDetailParam

// Three-letter ISO currency code.
//
// This is an alias to an internal type.
type Currency = shared.Currency

// This is an alias to an internal value.
const CurrencyAed = shared.CurrencyAed

// This is an alias to an internal value.
const CurrencyAfn = shared.CurrencyAfn

// This is an alias to an internal value.
const CurrencyAll = shared.CurrencyAll

// This is an alias to an internal value.
const CurrencyAmd = shared.CurrencyAmd

// This is an alias to an internal value.
const CurrencyAng = shared.CurrencyAng

// This is an alias to an internal value.
const CurrencyAoa = shared.CurrencyAoa

// This is an alias to an internal value.
const CurrencyArs = shared.CurrencyArs

// This is an alias to an internal value.
const CurrencyAud = shared.CurrencyAud

// This is an alias to an internal value.
const CurrencyAwg = shared.CurrencyAwg

// This is an alias to an internal value.
const CurrencyAzn = shared.CurrencyAzn

// This is an alias to an internal value.
const CurrencyBam = shared.CurrencyBam

// This is an alias to an internal value.
const CurrencyBbd = shared.CurrencyBbd

// This is an alias to an internal value.
const CurrencyBch = shared.CurrencyBch

// This is an alias to an internal value.
const CurrencyBdt = shared.CurrencyBdt

// This is an alias to an internal value.
const CurrencyBgn = shared.CurrencyBgn

// This is an alias to an internal value.
const CurrencyBhd = shared.CurrencyBhd

// This is an alias to an internal value.
const CurrencyBif = shared.CurrencyBif

// This is an alias to an internal value.
const CurrencyBmd = shared.CurrencyBmd

// This is an alias to an internal value.
const CurrencyBnd = shared.CurrencyBnd

// This is an alias to an internal value.
const CurrencyBob = shared.CurrencyBob

// This is an alias to an internal value.
const CurrencyBrl = shared.CurrencyBrl

// This is an alias to an internal value.
const CurrencyBsd = shared.CurrencyBsd

// This is an alias to an internal value.
const CurrencyBtc = shared.CurrencyBtc

// This is an alias to an internal value.
const CurrencyBtn = shared.CurrencyBtn

// This is an alias to an internal value.
const CurrencyBwp = shared.CurrencyBwp

// This is an alias to an internal value.
const CurrencyByn = shared.CurrencyByn

// This is an alias to an internal value.
const CurrencyByr = shared.CurrencyByr

// This is an alias to an internal value.
const CurrencyBzd = shared.CurrencyBzd

// This is an alias to an internal value.
const CurrencyCad = shared.CurrencyCad

// This is an alias to an internal value.
const CurrencyCdf = shared.CurrencyCdf

// This is an alias to an internal value.
const CurrencyChf = shared.CurrencyChf

// This is an alias to an internal value.
const CurrencyClf = shared.CurrencyClf

// This is an alias to an internal value.
const CurrencyClp = shared.CurrencyClp

// This is an alias to an internal value.
const CurrencyCnh = shared.CurrencyCnh

// This is an alias to an internal value.
const CurrencyCny = shared.CurrencyCny

// This is an alias to an internal value.
const CurrencyCop = shared.CurrencyCop

// This is an alias to an internal value.
const CurrencyCrc = shared.CurrencyCrc

// This is an alias to an internal value.
const CurrencyCuc = shared.CurrencyCuc

// This is an alias to an internal value.
const CurrencyCup = shared.CurrencyCup

// This is an alias to an internal value.
const CurrencyCve = shared.CurrencyCve

// This is an alias to an internal value.
const CurrencyCzk = shared.CurrencyCzk

// This is an alias to an internal value.
const CurrencyDjf = shared.CurrencyDjf

// This is an alias to an internal value.
const CurrencyDkk = shared.CurrencyDkk

// This is an alias to an internal value.
const CurrencyDop = shared.CurrencyDop

// This is an alias to an internal value.
const CurrencyDzd = shared.CurrencyDzd

// This is an alias to an internal value.
const CurrencyEek = shared.CurrencyEek

// This is an alias to an internal value.
const CurrencyEgp = shared.CurrencyEgp

// This is an alias to an internal value.
const CurrencyErn = shared.CurrencyErn

// This is an alias to an internal value.
const CurrencyEtb = shared.CurrencyEtb

// This is an alias to an internal value.
const CurrencyEur = shared.CurrencyEur

// This is an alias to an internal value.
const CurrencyEurc = shared.CurrencyEurc

// This is an alias to an internal value.
const CurrencyFjd = shared.CurrencyFjd

// This is an alias to an internal value.
const CurrencyFkp = shared.CurrencyFkp

// This is an alias to an internal value.
const CurrencyGbp = shared.CurrencyGbp

// This is an alias to an internal value.
const CurrencyGbx = shared.CurrencyGbx

// This is an alias to an internal value.
const CurrencyGel = shared.CurrencyGel

// This is an alias to an internal value.
const CurrencyGgp = shared.CurrencyGgp

// This is an alias to an internal value.
const CurrencyGhs = shared.CurrencyGhs

// This is an alias to an internal value.
const CurrencyGip = shared.CurrencyGip

// This is an alias to an internal value.
const CurrencyGmd = shared.CurrencyGmd

// This is an alias to an internal value.
const CurrencyGnf = shared.CurrencyGnf

// This is an alias to an internal value.
const CurrencyGtq = shared.CurrencyGtq

// This is an alias to an internal value.
const CurrencyGyd = shared.CurrencyGyd

// This is an alias to an internal value.
const CurrencyHkd = shared.CurrencyHkd

// This is an alias to an internal value.
const CurrencyHnl = shared.CurrencyHnl

// This is an alias to an internal value.
const CurrencyHrk = shared.CurrencyHrk

// This is an alias to an internal value.
const CurrencyHtg = shared.CurrencyHtg

// This is an alias to an internal value.
const CurrencyHuf = shared.CurrencyHuf

// This is an alias to an internal value.
const CurrencyIdr = shared.CurrencyIdr

// This is an alias to an internal value.
const CurrencyIls = shared.CurrencyIls

// This is an alias to an internal value.
const CurrencyImp = shared.CurrencyImp

// This is an alias to an internal value.
const CurrencyInr = shared.CurrencyInr

// This is an alias to an internal value.
const CurrencyIqd = shared.CurrencyIqd

// This is an alias to an internal value.
const CurrencyIrr = shared.CurrencyIrr

// This is an alias to an internal value.
const CurrencyIsk = shared.CurrencyIsk

// This is an alias to an internal value.
const CurrencyJep = shared.CurrencyJep

// This is an alias to an internal value.
const CurrencyJmd = shared.CurrencyJmd

// This is an alias to an internal value.
const CurrencyJod = shared.CurrencyJod

// This is an alias to an internal value.
const CurrencyJpy = shared.CurrencyJpy

// This is an alias to an internal value.
const CurrencyKes = shared.CurrencyKes

// This is an alias to an internal value.
const CurrencyKgs = shared.CurrencyKgs

// This is an alias to an internal value.
const CurrencyKhr = shared.CurrencyKhr

// This is an alias to an internal value.
const CurrencyKmf = shared.CurrencyKmf

// This is an alias to an internal value.
const CurrencyKpw = shared.CurrencyKpw

// This is an alias to an internal value.
const CurrencyKrw = shared.CurrencyKrw

// This is an alias to an internal value.
const CurrencyKwd = shared.CurrencyKwd

// This is an alias to an internal value.
const CurrencyKyd = shared.CurrencyKyd

// This is an alias to an internal value.
const CurrencyKzt = shared.CurrencyKzt

// This is an alias to an internal value.
const CurrencyLak = shared.CurrencyLak

// This is an alias to an internal value.
const CurrencyLbp = shared.CurrencyLbp

// This is an alias to an internal value.
const CurrencyLkr = shared.CurrencyLkr

// This is an alias to an internal value.
const CurrencyLrd = shared.CurrencyLrd

// This is an alias to an internal value.
const CurrencyLsl = shared.CurrencyLsl

// This is an alias to an internal value.
const CurrencyLtl = shared.CurrencyLtl

// This is an alias to an internal value.
const CurrencyLvl = shared.CurrencyLvl

// This is an alias to an internal value.
const CurrencyLyd = shared.CurrencyLyd

// This is an alias to an internal value.
const CurrencyMad = shared.CurrencyMad

// This is an alias to an internal value.
const CurrencyMdl = shared.CurrencyMdl

// This is an alias to an internal value.
const CurrencyMga = shared.CurrencyMga

// This is an alias to an internal value.
const CurrencyMkd = shared.CurrencyMkd

// This is an alias to an internal value.
const CurrencyMmk = shared.CurrencyMmk

// This is an alias to an internal value.
const CurrencyMnt = shared.CurrencyMnt

// This is an alias to an internal value.
const CurrencyMop = shared.CurrencyMop

// This is an alias to an internal value.
const CurrencyMro = shared.CurrencyMro

// This is an alias to an internal value.
const CurrencyMru = shared.CurrencyMru

// This is an alias to an internal value.
const CurrencyMtl = shared.CurrencyMtl

// This is an alias to an internal value.
const CurrencyMur = shared.CurrencyMur

// This is an alias to an internal value.
const CurrencyMvr = shared.CurrencyMvr

// This is an alias to an internal value.
const CurrencyMwk = shared.CurrencyMwk

// This is an alias to an internal value.
const CurrencyMxn = shared.CurrencyMxn

// This is an alias to an internal value.
const CurrencyMyr = shared.CurrencyMyr

// This is an alias to an internal value.
const CurrencyMzn = shared.CurrencyMzn

// This is an alias to an internal value.
const CurrencyNad = shared.CurrencyNad

// This is an alias to an internal value.
const CurrencyNgn = shared.CurrencyNgn

// This is an alias to an internal value.
const CurrencyNio = shared.CurrencyNio

// This is an alias to an internal value.
const CurrencyNok = shared.CurrencyNok

// This is an alias to an internal value.
const CurrencyNpr = shared.CurrencyNpr

// This is an alias to an internal value.
const CurrencyNzd = shared.CurrencyNzd

// This is an alias to an internal value.
const CurrencyOmr = shared.CurrencyOmr

// This is an alias to an internal value.
const CurrencyPab = shared.CurrencyPab

// This is an alias to an internal value.
const CurrencyPen = shared.CurrencyPen

// This is an alias to an internal value.
const CurrencyPgk = shared.CurrencyPgk

// This is an alias to an internal value.
const CurrencyPhp = shared.CurrencyPhp

// This is an alias to an internal value.
const CurrencyPkr = shared.CurrencyPkr

// This is an alias to an internal value.
const CurrencyPln = shared.CurrencyPln

// This is an alias to an internal value.
const CurrencyPyg = shared.CurrencyPyg

// This is an alias to an internal value.
const CurrencyPyusd = shared.CurrencyPyusd

// This is an alias to an internal value.
const CurrencyQar = shared.CurrencyQar

// This is an alias to an internal value.
const CurrencyRon = shared.CurrencyRon

// This is an alias to an internal value.
const CurrencyRsd = shared.CurrencyRsd

// This is an alias to an internal value.
const CurrencyRub = shared.CurrencyRub

// This is an alias to an internal value.
const CurrencyRwf = shared.CurrencyRwf

// This is an alias to an internal value.
const CurrencySar = shared.CurrencySar

// This is an alias to an internal value.
const CurrencySbc = shared.CurrencySbc

// This is an alias to an internal value.
const CurrencySbd = shared.CurrencySbd

// This is an alias to an internal value.
const CurrencyScr = shared.CurrencyScr

// This is an alias to an internal value.
const CurrencySdg = shared.CurrencySdg

// This is an alias to an internal value.
const CurrencySek = shared.CurrencySek

// This is an alias to an internal value.
const CurrencySgd = shared.CurrencySgd

// This is an alias to an internal value.
const CurrencyShp = shared.CurrencyShp

// This is an alias to an internal value.
const CurrencySkk = shared.CurrencySkk

// This is an alias to an internal value.
const CurrencySll = shared.CurrencySll

// This is an alias to an internal value.
const CurrencySos = shared.CurrencySos

// This is an alias to an internal value.
const CurrencySrd = shared.CurrencySrd

// This is an alias to an internal value.
const CurrencySsp = shared.CurrencySsp

// This is an alias to an internal value.
const CurrencyStd = shared.CurrencyStd

// This is an alias to an internal value.
const CurrencySvc = shared.CurrencySvc

// This is an alias to an internal value.
const CurrencySyp = shared.CurrencySyp

// This is an alias to an internal value.
const CurrencySzl = shared.CurrencySzl

// This is an alias to an internal value.
const CurrencyThb = shared.CurrencyThb

// This is an alias to an internal value.
const CurrencyTjs = shared.CurrencyTjs

// This is an alias to an internal value.
const CurrencyTmm = shared.CurrencyTmm

// This is an alias to an internal value.
const CurrencyTmt = shared.CurrencyTmt

// This is an alias to an internal value.
const CurrencyTnd = shared.CurrencyTnd

// This is an alias to an internal value.
const CurrencyTop = shared.CurrencyTop

// This is an alias to an internal value.
const CurrencyTry = shared.CurrencyTry

// This is an alias to an internal value.
const CurrencyTtd = shared.CurrencyTtd

// This is an alias to an internal value.
const CurrencyTwd = shared.CurrencyTwd

// This is an alias to an internal value.
const CurrencyTzs = shared.CurrencyTzs

// This is an alias to an internal value.
const CurrencyUah = shared.CurrencyUah

// This is an alias to an internal value.
const CurrencyUgx = shared.CurrencyUgx

// This is an alias to an internal value.
const CurrencyUsd = shared.CurrencyUsd

// This is an alias to an internal value.
const CurrencyUsdb = shared.CurrencyUsdb

// This is an alias to an internal value.
const CurrencyUsdc = shared.CurrencyUsdc

// This is an alias to an internal value.
const CurrencyUsdp = shared.CurrencyUsdp

// This is an alias to an internal value.
const CurrencyUsdt = shared.CurrencyUsdt

// This is an alias to an internal value.
const CurrencyUyu = shared.CurrencyUyu

// This is an alias to an internal value.
const CurrencyUzs = shared.CurrencyUzs

// This is an alias to an internal value.
const CurrencyVef = shared.CurrencyVef

// This is an alias to an internal value.
const CurrencyVes = shared.CurrencyVes

// This is an alias to an internal value.
const CurrencyVnd = shared.CurrencyVnd

// This is an alias to an internal value.
const CurrencyVuv = shared.CurrencyVuv

// This is an alias to an internal value.
const CurrencyWst = shared.CurrencyWst

// This is an alias to an internal value.
const CurrencyXaf = shared.CurrencyXaf

// This is an alias to an internal value.
const CurrencyXag = shared.CurrencyXag

// This is an alias to an internal value.
const CurrencyXau = shared.CurrencyXau

// This is an alias to an internal value.
const CurrencyXba = shared.CurrencyXba

// This is an alias to an internal value.
const CurrencyXbb = shared.CurrencyXbb

// This is an alias to an internal value.
const CurrencyXbc = shared.CurrencyXbc

// This is an alias to an internal value.
const CurrencyXbd = shared.CurrencyXbd

// This is an alias to an internal value.
const CurrencyXcd = shared.CurrencyXcd

// This is an alias to an internal value.
const CurrencyXdr = shared.CurrencyXdr

// This is an alias to an internal value.
const CurrencyXfu = shared.CurrencyXfu

// This is an alias to an internal value.
const CurrencyXof = shared.CurrencyXof

// This is an alias to an internal value.
const CurrencyXpd = shared.CurrencyXpd

// This is an alias to an internal value.
const CurrencyXpf = shared.CurrencyXpf

// This is an alias to an internal value.
const CurrencyXpt = shared.CurrencyXpt

// This is an alias to an internal value.
const CurrencyXts = shared.CurrencyXts

// This is an alias to an internal value.
const CurrencyYer = shared.CurrencyYer

// This is an alias to an internal value.
const CurrencyZar = shared.CurrencyZar

// This is an alias to an internal value.
const CurrencyZmk = shared.CurrencyZmk

// This is an alias to an internal value.
const CurrencyZmw = shared.CurrencyZmw

// This is an alias to an internal value.
const CurrencyZwd = shared.CurrencyZwd

// This is an alias to an internal value.
const CurrencyZwl = shared.CurrencyZwl

// This is an alias to an internal value.
const CurrencyZwn = shared.CurrencyZwn

// This is an alias to an internal value.
const CurrencyZwr = shared.CurrencyZwr

// This is an alias to an internal type.
type ForeignExchangeRate = shared.ForeignExchangeRate

// This is an alias to an internal type.
type IdentificationCreateRequestParam = shared.IdentificationCreateRequestParam

// The type of ID number.
//
// This is an alias to an internal type.
type IdentificationCreateRequestIDType = shared.IdentificationCreateRequestIDType

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeArCuil = shared.IdentificationCreateRequestIDTypeArCuil

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeArCuit = shared.IdentificationCreateRequestIDTypeArCuit

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeBrCnpj = shared.IdentificationCreateRequestIDTypeBrCnpj

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeBrCpf = shared.IdentificationCreateRequestIDTypeBrCpf

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeClRun = shared.IdentificationCreateRequestIDTypeClRun

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeClRut = shared.IdentificationCreateRequestIDTypeClRut

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeCoCedulas = shared.IdentificationCreateRequestIDTypeCoCedulas

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeCoNit = shared.IdentificationCreateRequestIDTypeCoNit

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeDriversLicense = shared.IdentificationCreateRequestIDTypeDriversLicense

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeHnID = shared.IdentificationCreateRequestIDTypeHnID

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeHnRtn = shared.IdentificationCreateRequestIDTypeHnRtn

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeInLei = shared.IdentificationCreateRequestIDTypeInLei

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeKrBrn = shared.IdentificationCreateRequestIDTypeKrBrn

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeKrCrn = shared.IdentificationCreateRequestIDTypeKrCrn

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeKrRrn = shared.IdentificationCreateRequestIDTypeKrRrn

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypePassport = shared.IdentificationCreateRequestIDTypePassport

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeSaTin = shared.IdentificationCreateRequestIDTypeSaTin

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeSaVat = shared.IdentificationCreateRequestIDTypeSaVat

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeUsEin = shared.IdentificationCreateRequestIDTypeUsEin

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeUsItin = shared.IdentificationCreateRequestIDTypeUsItin

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeUsSsn = shared.IdentificationCreateRequestIDTypeUsSsn

// This is an alias to an internal value.
const IdentificationCreateRequestIDTypeVnTin = shared.IdentificationCreateRequestIDTypeVnTin

// This is an alias to an internal type.
type LedgerAccountCreateRequestParam = shared.LedgerAccountCreateRequestParam

// If the ledger account links to another object in Modern Treasury, the type will
// be populated here, otherwise null. The value is one of internal_account or
// external_account.
//
// This is an alias to an internal type.
type LedgerAccountCreateRequestLedgerableType = shared.LedgerAccountCreateRequestLedgerableType

// This is an alias to an internal value.
const LedgerAccountCreateRequestLedgerableTypeCounterparty = shared.LedgerAccountCreateRequestLedgerableTypeCounterparty

// This is an alias to an internal value.
const LedgerAccountCreateRequestLedgerableTypeExternalAccount = shared.LedgerAccountCreateRequestLedgerableTypeExternalAccount

// This is an alias to an internal value.
const LedgerAccountCreateRequestLedgerableTypeInternalAccount = shared.LedgerAccountCreateRequestLedgerableTypeInternalAccount

// This is an alias to an internal value.
const LedgerAccountCreateRequestLedgerableTypeVirtualAccount = shared.LedgerAccountCreateRequestLedgerableTypeVirtualAccount

// This is an alias to an internal type.
type LedgerBalance = shared.LedgerBalance

// This is an alias to an internal type.
type LedgerBalances = shared.LedgerBalances

// This is an alias to an internal type.
type LedgerEntryCreateRequestParam = shared.LedgerEntryCreateRequestParam

// This is an alias to an internal type.
type LedgerTransactionCreateRequestParam = shared.LedgerTransactionCreateRequestParam

// If the ledger transaction can be reconciled to another object in Modern
// Treasury, the type will be populated here, otherwise null. This can be one of
// payment_order, incoming_payment_detail, expected_payment, return, or reversal.
//
// This is an alias to an internal type.
type LedgerTransactionCreateRequestLedgerableType = shared.LedgerTransactionCreateRequestLedgerableType

// This is an alias to an internal value.
const LedgerTransactionCreateRequestLedgerableTypeExpectedPayment = shared.LedgerTransactionCreateRequestLedgerableTypeExpectedPayment

// This is an alias to an internal value.
const LedgerTransactionCreateRequestLedgerableTypeIncomingPaymentDetail = shared.LedgerTransactionCreateRequestLedgerableTypeIncomingPaymentDetail

// This is an alias to an internal value.
const LedgerTransactionCreateRequestLedgerableTypePaymentOrder = shared.LedgerTransactionCreateRequestLedgerableTypePaymentOrder

// This is an alias to an internal value.
const LedgerTransactionCreateRequestLedgerableTypeReturn = shared.LedgerTransactionCreateRequestLedgerableTypeReturn

// This is an alias to an internal value.
const LedgerTransactionCreateRequestLedgerableTypeReversal = shared.LedgerTransactionCreateRequestLedgerableTypeReversal

// To post a ledger transaction at creation, use `posted`.
//
// This is an alias to an internal type.
type LedgerTransactionCreateRequestStatus = shared.LedgerTransactionCreateRequestStatus

// This is an alias to an internal value.
const LedgerTransactionCreateRequestStatusArchived = shared.LedgerTransactionCreateRequestStatusArchived

// This is an alias to an internal value.
const LedgerTransactionCreateRequestStatusPending = shared.LedgerTransactionCreateRequestStatusPending

// This is an alias to an internal value.
const LedgerTransactionCreateRequestStatusPosted = shared.LedgerTransactionCreateRequestStatusPosted

// This is an alias to an internal type.
type LegalEntityAddressCreateRequestParam = shared.LegalEntityAddressCreateRequestParam

// This is an alias to an internal type.
type LegalEntityAddressCreateRequestAddressType = shared.LegalEntityAddressCreateRequestAddressType

// This is an alias to an internal value.
const LegalEntityAddressCreateRequestAddressTypeBusiness = shared.LegalEntityAddressCreateRequestAddressTypeBusiness

// This is an alias to an internal value.
const LegalEntityAddressCreateRequestAddressTypeMailing = shared.LegalEntityAddressCreateRequestAddressTypeMailing

// This is an alias to an internal value.
const LegalEntityAddressCreateRequestAddressTypeOther = shared.LegalEntityAddressCreateRequestAddressTypeOther

// This is an alias to an internal value.
const LegalEntityAddressCreateRequestAddressTypePoBox = shared.LegalEntityAddressCreateRequestAddressTypePoBox

// This is an alias to an internal value.
const LegalEntityAddressCreateRequestAddressTypeResidential = shared.LegalEntityAddressCreateRequestAddressTypeResidential

// This is an alias to an internal type.
type LegalEntityAssociationInlineCreateParam = shared.LegalEntityAssociationInlineCreateParam

// A list of relationship types for how the child entity relates to parent entity.
//
// This is an alias to an internal type.
type LegalEntityAssociationInlineCreateRelationshipType = shared.LegalEntityAssociationInlineCreateRelationshipType

// This is an alias to an internal value.
const LegalEntityAssociationInlineCreateRelationshipTypeAuthorizedSigner = shared.LegalEntityAssociationInlineCreateRelationshipTypeAuthorizedSigner

// This is an alias to an internal value.
const LegalEntityAssociationInlineCreateRelationshipTypeBeneficialOwner = shared.LegalEntityAssociationInlineCreateRelationshipTypeBeneficialOwner

// This is an alias to an internal value.
const LegalEntityAssociationInlineCreateRelationshipTypeControlPerson = shared.LegalEntityAssociationInlineCreateRelationshipTypeControlPerson

// This is an alias to an internal type.
type LegalEntityBankSettings = shared.LegalEntityBankSettings

// This is an alias to an internal type.
type LegalEntityBankSettingsParam = shared.LegalEntityBankSettingsParam

// This is an alias to an internal type.
type LegalEntityIndustryClassification = shared.LegalEntityIndustryClassification

// The classification system of the classification codes.
//
// This is an alias to an internal type.
type LegalEntityIndustryClassificationClassificationType = shared.LegalEntityIndustryClassificationClassificationType

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeAnzsic = shared.LegalEntityIndustryClassificationClassificationTypeAnzsic

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeBics = shared.LegalEntityIndustryClassificationClassificationTypeBics

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeGics = shared.LegalEntityIndustryClassificationClassificationTypeGics

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeHsics = shared.LegalEntityIndustryClassificationClassificationTypeHsics

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeIcb = shared.LegalEntityIndustryClassificationClassificationTypeIcb

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeIsic = shared.LegalEntityIndustryClassificationClassificationTypeIsic

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeMgecs = shared.LegalEntityIndustryClassificationClassificationTypeMgecs

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeNace = shared.LegalEntityIndustryClassificationClassificationTypeNace

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeNaics = shared.LegalEntityIndustryClassificationClassificationTypeNaics

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeRbics = shared.LegalEntityIndustryClassificationClassificationTypeRbics

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeSic = shared.LegalEntityIndustryClassificationClassificationTypeSic

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeSni = shared.LegalEntityIndustryClassificationClassificationTypeSni

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeTrbc = shared.LegalEntityIndustryClassificationClassificationTypeTrbc

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeUksic = shared.LegalEntityIndustryClassificationClassificationTypeUksic

// This is an alias to an internal value.
const LegalEntityIndustryClassificationClassificationTypeUnspsc = shared.LegalEntityIndustryClassificationClassificationTypeUnspsc

// This is an alias to an internal type.
type LegalEntityIndustryClassificationParam = shared.LegalEntityIndustryClassificationParam

// This is an alias to an internal type.
type LegalEntityWealthEmploymentDetail = shared.LegalEntityWealthEmploymentDetail

// The employment status of the individual.
//
// This is an alias to an internal type.
type LegalEntityWealthEmploymentDetailEmploymentStatus = shared.LegalEntityWealthEmploymentDetailEmploymentStatus

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailEmploymentStatusEmployed = shared.LegalEntityWealthEmploymentDetailEmploymentStatusEmployed

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailEmploymentStatusRetired = shared.LegalEntityWealthEmploymentDetailEmploymentStatusRetired

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailEmploymentStatusSelfEmployed = shared.LegalEntityWealthEmploymentDetailEmploymentStatusSelfEmployed

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailEmploymentStatusStudent = shared.LegalEntityWealthEmploymentDetailEmploymentStatusStudent

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailEmploymentStatusUnemployed = shared.LegalEntityWealthEmploymentDetailEmploymentStatusUnemployed

// The source of the individual's income.
//
// This is an alias to an internal type.
type LegalEntityWealthEmploymentDetailIncomeSource = shared.LegalEntityWealthEmploymentDetailIncomeSource

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIncomeSourceFamilySupport = shared.LegalEntityWealthEmploymentDetailIncomeSourceFamilySupport

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIncomeSourceGovernmentBenefits = shared.LegalEntityWealthEmploymentDetailIncomeSourceGovernmentBenefits

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIncomeSourceInheritance = shared.LegalEntityWealthEmploymentDetailIncomeSourceInheritance

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIncomeSourceInvestments = shared.LegalEntityWealthEmploymentDetailIncomeSourceInvestments

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIncomeSourceRentalIncome = shared.LegalEntityWealthEmploymentDetailIncomeSourceRentalIncome

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIncomeSourceRetirement = shared.LegalEntityWealthEmploymentDetailIncomeSourceRetirement

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIncomeSourceSalary = shared.LegalEntityWealthEmploymentDetailIncomeSourceSalary

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIncomeSourceSelfEmployed = shared.LegalEntityWealthEmploymentDetailIncomeSourceSelfEmployed

// The industry of the individual.
//
// This is an alias to an internal type.
type LegalEntityWealthEmploymentDetailIndustry = shared.LegalEntityWealthEmploymentDetailIndustry

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryAccounting = shared.LegalEntityWealthEmploymentDetailIndustryAccounting

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryAgriculture = shared.LegalEntityWealthEmploymentDetailIndustryAgriculture

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryAutomotive = shared.LegalEntityWealthEmploymentDetailIndustryAutomotive

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryChemicalManufacturing = shared.LegalEntityWealthEmploymentDetailIndustryChemicalManufacturing

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryConstruction = shared.LegalEntityWealthEmploymentDetailIndustryConstruction

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryEducationalMedical = shared.LegalEntityWealthEmploymentDetailIndustryEducationalMedical

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryFoodService = shared.LegalEntityWealthEmploymentDetailIndustryFoodService

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryFinance = shared.LegalEntityWealthEmploymentDetailIndustryFinance

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryGasoline = shared.LegalEntityWealthEmploymentDetailIndustryGasoline

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryHealthStores = shared.LegalEntityWealthEmploymentDetailIndustryHealthStores

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryLaundry = shared.LegalEntityWealthEmploymentDetailIndustryLaundry

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryMaintenance = shared.LegalEntityWealthEmploymentDetailIndustryMaintenance

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryManufacturing = shared.LegalEntityWealthEmploymentDetailIndustryManufacturing

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryMerchantWholesale = shared.LegalEntityWealthEmploymentDetailIndustryMerchantWholesale

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryMining = shared.LegalEntityWealthEmploymentDetailIndustryMining

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryPerformingArts = shared.LegalEntityWealthEmploymentDetailIndustryPerformingArts

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryProfessionalNonLegal = shared.LegalEntityWealthEmploymentDetailIndustryProfessionalNonLegal

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryPublicAdministration = shared.LegalEntityWealthEmploymentDetailIndustryPublicAdministration

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryPublishing = shared.LegalEntityWealthEmploymentDetailIndustryPublishing

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryRealEstate = shared.LegalEntityWealthEmploymentDetailIndustryRealEstate

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryRecreationGambling = shared.LegalEntityWealthEmploymentDetailIndustryRecreationGambling

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryReligiousCharity = shared.LegalEntityWealthEmploymentDetailIndustryReligiousCharity

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryRentalServices = shared.LegalEntityWealthEmploymentDetailIndustryRentalServices

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryRetailClothing = shared.LegalEntityWealthEmploymentDetailIndustryRetailClothing

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryRetailElectronics = shared.LegalEntityWealthEmploymentDetailIndustryRetailElectronics

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryRetailFood = shared.LegalEntityWealthEmploymentDetailIndustryRetailFood

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryRetailFurnishing = shared.LegalEntityWealthEmploymentDetailIndustryRetailFurnishing

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryRetailHome = shared.LegalEntityWealthEmploymentDetailIndustryRetailHome

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryRetailNonStore = shared.LegalEntityWealthEmploymentDetailIndustryRetailNonStore

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryRetailSporting = shared.LegalEntityWealthEmploymentDetailIndustryRetailSporting

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryTransportation = shared.LegalEntityWealthEmploymentDetailIndustryTransportation

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryTravel = shared.LegalEntityWealthEmploymentDetailIndustryTravel

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailIndustryUtilities = shared.LegalEntityWealthEmploymentDetailIndustryUtilities

// The occupation of the individual.
//
// This is an alias to an internal type.
type LegalEntityWealthEmploymentDetailOccupation = shared.LegalEntityWealthEmploymentDetailOccupation

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailOccupationConsulting = shared.LegalEntityWealthEmploymentDetailOccupationConsulting

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailOccupationExecutive = shared.LegalEntityWealthEmploymentDetailOccupationExecutive

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailOccupationFinanceAccounting = shared.LegalEntityWealthEmploymentDetailOccupationFinanceAccounting

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailOccupationFoodServices = shared.LegalEntityWealthEmploymentDetailOccupationFoodServices

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailOccupationGovernment = shared.LegalEntityWealthEmploymentDetailOccupationGovernment

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailOccupationHealthcare = shared.LegalEntityWealthEmploymentDetailOccupationHealthcare

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailOccupationLegalServices = shared.LegalEntityWealthEmploymentDetailOccupationLegalServices

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailOccupationManufacturing = shared.LegalEntityWealthEmploymentDetailOccupationManufacturing

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailOccupationOther = shared.LegalEntityWealthEmploymentDetailOccupationOther

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailOccupationSales = shared.LegalEntityWealthEmploymentDetailOccupationSales

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailOccupationScienceEngineering = shared.LegalEntityWealthEmploymentDetailOccupationScienceEngineering

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailOccupationTechnology = shared.LegalEntityWealthEmploymentDetailOccupationTechnology

// The source of the individual's funds.
//
// This is an alias to an internal type.
type LegalEntityWealthEmploymentDetailSourceOfFunds = shared.LegalEntityWealthEmploymentDetailSourceOfFunds

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsAlimony = shared.LegalEntityWealthEmploymentDetailSourceOfFundsAlimony

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsAnnuity = shared.LegalEntityWealthEmploymentDetailSourceOfFundsAnnuity

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsBusinessOwner = shared.LegalEntityWealthEmploymentDetailSourceOfFundsBusinessOwner

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsBusinessRevenue = shared.LegalEntityWealthEmploymentDetailSourceOfFundsBusinessRevenue

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsDebtFinancing = shared.LegalEntityWealthEmploymentDetailSourceOfFundsDebtFinancing

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsGeneralEmployee = shared.LegalEntityWealthEmploymentDetailSourceOfFundsGeneralEmployee

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsGovernmentBenefits = shared.LegalEntityWealthEmploymentDetailSourceOfFundsGovernmentBenefits

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsHomemaker = shared.LegalEntityWealthEmploymentDetailSourceOfFundsHomemaker

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsInheritanceGift = shared.LegalEntityWealthEmploymentDetailSourceOfFundsInheritanceGift

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsIntercompanyLoan = shared.LegalEntityWealthEmploymentDetailSourceOfFundsIntercompanyLoan

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsInvestment = shared.LegalEntityWealthEmploymentDetailSourceOfFundsInvestment

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsInvestorFunding = shared.LegalEntityWealthEmploymentDetailSourceOfFundsInvestorFunding

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsLegalSettlement = shared.LegalEntityWealthEmploymentDetailSourceOfFundsLegalSettlement

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsLottery = shared.LegalEntityWealthEmploymentDetailSourceOfFundsLottery

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsRealEstate = shared.LegalEntityWealthEmploymentDetailSourceOfFundsRealEstate

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsRetainedEarningsOrSavings = shared.LegalEntityWealthEmploymentDetailSourceOfFundsRetainedEarningsOrSavings

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsRetired = shared.LegalEntityWealthEmploymentDetailSourceOfFundsRetired

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsRetirement = shared.LegalEntityWealthEmploymentDetailSourceOfFundsRetirement

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsSalary = shared.LegalEntityWealthEmploymentDetailSourceOfFundsSalary

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsSaleOfBusinessAssets = shared.LegalEntityWealthEmploymentDetailSourceOfFundsSaleOfBusinessAssets

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsSaleOfRealEstate = shared.LegalEntityWealthEmploymentDetailSourceOfFundsSaleOfRealEstate

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsSelfEmployed = shared.LegalEntityWealthEmploymentDetailSourceOfFundsSelfEmployed

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsSeniorExecutive = shared.LegalEntityWealthEmploymentDetailSourceOfFundsSeniorExecutive

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailSourceOfFundsTrustIncome = shared.LegalEntityWealthEmploymentDetailSourceOfFundsTrustIncome

// The source of the individual's wealth.
//
// This is an alias to an internal type.
type LegalEntityWealthEmploymentDetailWealthSource = shared.LegalEntityWealthEmploymentDetailWealthSource

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailWealthSourceBusinessSale = shared.LegalEntityWealthEmploymentDetailWealthSourceBusinessSale

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailWealthSourceFamilySupport = shared.LegalEntityWealthEmploymentDetailWealthSourceFamilySupport

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailWealthSourceGovernmentBenefits = shared.LegalEntityWealthEmploymentDetailWealthSourceGovernmentBenefits

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailWealthSourceInheritance = shared.LegalEntityWealthEmploymentDetailWealthSourceInheritance

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailWealthSourceInvestments = shared.LegalEntityWealthEmploymentDetailWealthSourceInvestments

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailWealthSourceOther = shared.LegalEntityWealthEmploymentDetailWealthSourceOther

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailWealthSourceRentalIncome = shared.LegalEntityWealthEmploymentDetailWealthSourceRentalIncome

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailWealthSourceRetirement = shared.LegalEntityWealthEmploymentDetailWealthSourceRetirement

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailWealthSourceSalary = shared.LegalEntityWealthEmploymentDetailWealthSourceSalary

// This is an alias to an internal value.
const LegalEntityWealthEmploymentDetailWealthSourceSelfEmployed = shared.LegalEntityWealthEmploymentDetailWealthSourceSelfEmployed

// This is an alias to an internal type.
type LegalEntityWealthEmploymentDetailParam = shared.LegalEntityWealthEmploymentDetailParam

// This is an alias to an internal type.
type TransactionDirection = shared.TransactionDirection

// This is an alias to an internal value.
const TransactionDirectionCredit = shared.TransactionDirectionCredit

// This is an alias to an internal value.
const TransactionDirectionDebit = shared.TransactionDirectionDebit
