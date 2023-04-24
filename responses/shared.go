package responses

import (
	"net/http"

	pjson "github.com/Modern-Treasury/modern-treasury-go/core/json"
	"github.com/Modern-Treasury/modern-treasury-go/option"
)

type Page[T any] struct {
	Items []T `json:"items"`
	JSON  PageJSON
	cfg   *option.RequestConfig
	res   *http.Response
}

type PageJSON struct {
	Items  pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Page[T] using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Page[T]) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *Page[T]) GetNextPage() (res *Page[T], err error) {
	next := r.res.Header.Get("X-After-Cursor")
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	cfg.Apply(option.WithQuery("after_cursor", next))
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *Page[T]) SetPageConfig(cfg *option.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type PageAutoPager[T any] struct {
	page *Page[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewPageAutoPager[T any](page *Page[T], err error) *PageAutoPager[T] {
	return &PageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *PageAutoPager[T]) Next() bool {
	if len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *PageAutoPager[T]) Current() T {
	return r.cur
}

func (r *PageAutoPager[T]) Err() error {
	return r.err
}

func (r *PageAutoPager[T]) Index() int {
	return r.run
}

type AsyncResponse struct {
	ID     string `json:"id,required" format:"uuid"`
	Object string `json:"object,required"`
	JSON   AsyncResponseJSON
}

type AsyncResponseJSON struct {
	ID     pjson.Metadata
	Object pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AsyncResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AsyncResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

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
