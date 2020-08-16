package model

type Records struct {
	ID int `json:"ID"`
	SERVER string `json:"SERVER"`
	NETWORK string `json:"NETWORK"`
	REQ_ID string `json:"REQ_ID"`
	REGION_IN string `json:"REGION_IN"`
	REGION_OUT string `json:"REGION_OUT"`
	CLOAD int `json:"CLOAD"`
	REQ_TIME string `json:"REQ_TIME"`
	HASH_ID string `json:"HASH_ID"`
}