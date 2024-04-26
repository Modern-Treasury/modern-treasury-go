// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/Modern-Treasury/modern-treasury-go/v2/internal/apijson"
)

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

// Three-letter ISO currency code.
type Currency string

const (
	CurrencyAed Currency = "AED"
	CurrencyAfn Currency = "AFN"
	CurrencyAll Currency = "ALL"
	CurrencyAmd Currency = "AMD"
	CurrencyAng Currency = "ANG"
	CurrencyAoa Currency = "AOA"
	CurrencyArs Currency = "ARS"
	CurrencyAud Currency = "AUD"
	CurrencyAwg Currency = "AWG"
	CurrencyAzn Currency = "AZN"
	CurrencyBam Currency = "BAM"
	CurrencyBbd Currency = "BBD"
	CurrencyBch Currency = "BCH"
	CurrencyBdt Currency = "BDT"
	CurrencyBgn Currency = "BGN"
	CurrencyBhd Currency = "BHD"
	CurrencyBif Currency = "BIF"
	CurrencyBmd Currency = "BMD"
	CurrencyBnd Currency = "BND"
	CurrencyBob Currency = "BOB"
	CurrencyBrl Currency = "BRL"
	CurrencyBsd Currency = "BSD"
	CurrencyBtc Currency = "BTC"
	CurrencyBtn Currency = "BTN"
	CurrencyBwp Currency = "BWP"
	CurrencyByn Currency = "BYN"
	CurrencyByr Currency = "BYR"
	CurrencyBzd Currency = "BZD"
	CurrencyCad Currency = "CAD"
	CurrencyCdf Currency = "CDF"
	CurrencyChf Currency = "CHF"
	CurrencyClf Currency = "CLF"
	CurrencyClp Currency = "CLP"
	CurrencyCnh Currency = "CNH"
	CurrencyCny Currency = "CNY"
	CurrencyCop Currency = "COP"
	CurrencyCrc Currency = "CRC"
	CurrencyCuc Currency = "CUC"
	CurrencyCup Currency = "CUP"
	CurrencyCve Currency = "CVE"
	CurrencyCzk Currency = "CZK"
	CurrencyDjf Currency = "DJF"
	CurrencyDkk Currency = "DKK"
	CurrencyDop Currency = "DOP"
	CurrencyDzd Currency = "DZD"
	CurrencyEek Currency = "EEK"
	CurrencyEgp Currency = "EGP"
	CurrencyErn Currency = "ERN"
	CurrencyEtb Currency = "ETB"
	CurrencyEur Currency = "EUR"
	CurrencyFjd Currency = "FJD"
	CurrencyFkp Currency = "FKP"
	CurrencyGbp Currency = "GBP"
	CurrencyGbx Currency = "GBX"
	CurrencyGel Currency = "GEL"
	CurrencyGgp Currency = "GGP"
	CurrencyGhs Currency = "GHS"
	CurrencyGip Currency = "GIP"
	CurrencyGmd Currency = "GMD"
	CurrencyGnf Currency = "GNF"
	CurrencyGtq Currency = "GTQ"
	CurrencyGyd Currency = "GYD"
	CurrencyHkd Currency = "HKD"
	CurrencyHnl Currency = "HNL"
	CurrencyHrk Currency = "HRK"
	CurrencyHtg Currency = "HTG"
	CurrencyHuf Currency = "HUF"
	CurrencyIdr Currency = "IDR"
	CurrencyIls Currency = "ILS"
	CurrencyImp Currency = "IMP"
	CurrencyInr Currency = "INR"
	CurrencyIqd Currency = "IQD"
	CurrencyIrr Currency = "IRR"
	CurrencyIsk Currency = "ISK"
	CurrencyJep Currency = "JEP"
	CurrencyJmd Currency = "JMD"
	CurrencyJod Currency = "JOD"
	CurrencyJpy Currency = "JPY"
	CurrencyKes Currency = "KES"
	CurrencyKgs Currency = "KGS"
	CurrencyKhr Currency = "KHR"
	CurrencyKmf Currency = "KMF"
	CurrencyKpw Currency = "KPW"
	CurrencyKrw Currency = "KRW"
	CurrencyKwd Currency = "KWD"
	CurrencyKyd Currency = "KYD"
	CurrencyKzt Currency = "KZT"
	CurrencyLak Currency = "LAK"
	CurrencyLbp Currency = "LBP"
	CurrencyLkr Currency = "LKR"
	CurrencyLrd Currency = "LRD"
	CurrencyLsl Currency = "LSL"
	CurrencyLtl Currency = "LTL"
	CurrencyLvl Currency = "LVL"
	CurrencyLyd Currency = "LYD"
	CurrencyMad Currency = "MAD"
	CurrencyMdl Currency = "MDL"
	CurrencyMga Currency = "MGA"
	CurrencyMkd Currency = "MKD"
	CurrencyMmk Currency = "MMK"
	CurrencyMnt Currency = "MNT"
	CurrencyMop Currency = "MOP"
	CurrencyMro Currency = "MRO"
	CurrencyMru Currency = "MRU"
	CurrencyMtl Currency = "MTL"
	CurrencyMur Currency = "MUR"
	CurrencyMvr Currency = "MVR"
	CurrencyMwk Currency = "MWK"
	CurrencyMxn Currency = "MXN"
	CurrencyMyr Currency = "MYR"
	CurrencyMzn Currency = "MZN"
	CurrencyNad Currency = "NAD"
	CurrencyNgn Currency = "NGN"
	CurrencyNio Currency = "NIO"
	CurrencyNok Currency = "NOK"
	CurrencyNpr Currency = "NPR"
	CurrencyNzd Currency = "NZD"
	CurrencyOmr Currency = "OMR"
	CurrencyPab Currency = "PAB"
	CurrencyPen Currency = "PEN"
	CurrencyPgk Currency = "PGK"
	CurrencyPhp Currency = "PHP"
	CurrencyPkr Currency = "PKR"
	CurrencyPln Currency = "PLN"
	CurrencyPyg Currency = "PYG"
	CurrencyQar Currency = "QAR"
	CurrencyRon Currency = "RON"
	CurrencyRsd Currency = "RSD"
	CurrencyRub Currency = "RUB"
	CurrencyRwf Currency = "RWF"
	CurrencySar Currency = "SAR"
	CurrencySbd Currency = "SBD"
	CurrencyScr Currency = "SCR"
	CurrencySdg Currency = "SDG"
	CurrencySek Currency = "SEK"
	CurrencySgd Currency = "SGD"
	CurrencyShp Currency = "SHP"
	CurrencySkk Currency = "SKK"
	CurrencySll Currency = "SLL"
	CurrencySos Currency = "SOS"
	CurrencySrd Currency = "SRD"
	CurrencySsp Currency = "SSP"
	CurrencyStd Currency = "STD"
	CurrencySvc Currency = "SVC"
	CurrencySyp Currency = "SYP"
	CurrencySzl Currency = "SZL"
	CurrencyThb Currency = "THB"
	CurrencyTjs Currency = "TJS"
	CurrencyTmm Currency = "TMM"
	CurrencyTmt Currency = "TMT"
	CurrencyTnd Currency = "TND"
	CurrencyTop Currency = "TOP"
	CurrencyTry Currency = "TRY"
	CurrencyTtd Currency = "TTD"
	CurrencyTwd Currency = "TWD"
	CurrencyTzs Currency = "TZS"
	CurrencyUah Currency = "UAH"
	CurrencyUgx Currency = "UGX"
	CurrencyUsd Currency = "USD"
	CurrencyUyu Currency = "UYU"
	CurrencyUzs Currency = "UZS"
	CurrencyVef Currency = "VEF"
	CurrencyVes Currency = "VES"
	CurrencyVnd Currency = "VND"
	CurrencyVuv Currency = "VUV"
	CurrencyWst Currency = "WST"
	CurrencyXaf Currency = "XAF"
	CurrencyXag Currency = "XAG"
	CurrencyXau Currency = "XAU"
	CurrencyXba Currency = "XBA"
	CurrencyXbb Currency = "XBB"
	CurrencyXbc Currency = "XBC"
	CurrencyXbd Currency = "XBD"
	CurrencyXcd Currency = "XCD"
	CurrencyXdr Currency = "XDR"
	CurrencyXfu Currency = "XFU"
	CurrencyXof Currency = "XOF"
	CurrencyXpd Currency = "XPD"
	CurrencyXpf Currency = "XPF"
	CurrencyXpt Currency = "XPT"
	CurrencyXts Currency = "XTS"
	CurrencyYer Currency = "YER"
	CurrencyZar Currency = "ZAR"
	CurrencyZmk Currency = "ZMK"
	CurrencyZmw Currency = "ZMW"
	CurrencyZwd Currency = "ZWD"
	CurrencyZwl Currency = "ZWL"
	CurrencyZwn Currency = "ZWN"
	CurrencyZwr Currency = "ZWR"
)

func (r Currency) IsKnown() bool {
	switch r {
	case CurrencyAed, CurrencyAfn, CurrencyAll, CurrencyAmd, CurrencyAng, CurrencyAoa, CurrencyArs, CurrencyAud, CurrencyAwg, CurrencyAzn, CurrencyBam, CurrencyBbd, CurrencyBch, CurrencyBdt, CurrencyBgn, CurrencyBhd, CurrencyBif, CurrencyBmd, CurrencyBnd, CurrencyBob, CurrencyBrl, CurrencyBsd, CurrencyBtc, CurrencyBtn, CurrencyBwp, CurrencyByn, CurrencyByr, CurrencyBzd, CurrencyCad, CurrencyCdf, CurrencyChf, CurrencyClf, CurrencyClp, CurrencyCnh, CurrencyCny, CurrencyCop, CurrencyCrc, CurrencyCuc, CurrencyCup, CurrencyCve, CurrencyCzk, CurrencyDjf, CurrencyDkk, CurrencyDop, CurrencyDzd, CurrencyEek, CurrencyEgp, CurrencyErn, CurrencyEtb, CurrencyEur, CurrencyFjd, CurrencyFkp, CurrencyGbp, CurrencyGbx, CurrencyGel, CurrencyGgp, CurrencyGhs, CurrencyGip, CurrencyGmd, CurrencyGnf, CurrencyGtq, CurrencyGyd, CurrencyHkd, CurrencyHnl, CurrencyHrk, CurrencyHtg, CurrencyHuf, CurrencyIdr, CurrencyIls, CurrencyImp, CurrencyInr, CurrencyIqd, CurrencyIrr, CurrencyIsk, CurrencyJep, CurrencyJmd, CurrencyJod, CurrencyJpy, CurrencyKes, CurrencyKgs, CurrencyKhr, CurrencyKmf, CurrencyKpw, CurrencyKrw, CurrencyKwd, CurrencyKyd, CurrencyKzt, CurrencyLak, CurrencyLbp, CurrencyLkr, CurrencyLrd, CurrencyLsl, CurrencyLtl, CurrencyLvl, CurrencyLyd, CurrencyMad, CurrencyMdl, CurrencyMga, CurrencyMkd, CurrencyMmk, CurrencyMnt, CurrencyMop, CurrencyMro, CurrencyMru, CurrencyMtl, CurrencyMur, CurrencyMvr, CurrencyMwk, CurrencyMxn, CurrencyMyr, CurrencyMzn, CurrencyNad, CurrencyNgn, CurrencyNio, CurrencyNok, CurrencyNpr, CurrencyNzd, CurrencyOmr, CurrencyPab, CurrencyPen, CurrencyPgk, CurrencyPhp, CurrencyPkr, CurrencyPln, CurrencyPyg, CurrencyQar, CurrencyRon, CurrencyRsd, CurrencyRub, CurrencyRwf, CurrencySar, CurrencySbd, CurrencyScr, CurrencySdg, CurrencySek, CurrencySgd, CurrencyShp, CurrencySkk, CurrencySll, CurrencySos, CurrencySrd, CurrencySsp, CurrencyStd, CurrencySvc, CurrencySyp, CurrencySzl, CurrencyThb, CurrencyTjs, CurrencyTmm, CurrencyTmt, CurrencyTnd, CurrencyTop, CurrencyTry, CurrencyTtd, CurrencyTwd, CurrencyTzs, CurrencyUah, CurrencyUgx, CurrencyUsd, CurrencyUyu, CurrencyUzs, CurrencyVef, CurrencyVes, CurrencyVnd, CurrencyVuv, CurrencyWst, CurrencyXaf, CurrencyXag, CurrencyXau, CurrencyXba, CurrencyXbb, CurrencyXbc, CurrencyXbd, CurrencyXcd, CurrencyXdr, CurrencyXfu, CurrencyXof, CurrencyXpd, CurrencyXpf, CurrencyXpt, CurrencyXts, CurrencyYer, CurrencyZar, CurrencyZmk, CurrencyZmw, CurrencyZwd, CurrencyZwl, CurrencyZwn, CurrencyZwr:
		return true
	}
	return false
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
