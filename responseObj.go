package main

type ResponseObject struct {
	GoldUsd           float64 `json:"gold_usd"`
	GoldEur           float64 `json:"gold_eur"`
	SilberUsd         float64 `json:"silber_usd"`
	SilberEur         float64 `json:"silber_eur"`
	PlatinUsd         float64 `json:"platin_usd"`
	PlatinEur         float64 `json:"platin_eur"`
	PalladiumUsd      float64 `json:"palladium_usd"`
	PalladiumEur      float64 `json:"palladium_eur"`
	Timestamp         int     `json:"timestamp"`
	WechselkursUsdEur float64 `json:"wechselkurs_usd_eur"`
}
