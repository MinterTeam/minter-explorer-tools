package models

import (
	"time"
)

type Coin struct {
	ID                    uint64     `json:"id" pg:",pk"`
	CreationAddressID     *uint64    `json:"creation_address_id"`
	CreationTransactionID *uint64    `json:"creation_transaction_id"`
	Crr                   uint64     `json:"crr"`
	MaxSupply             string     `json:"max_supply"      pg:"type:numeric(70)"`
	Volume                string     `json:"volume"          pg:"type:numeric(70)"`
	ReserveBalance        string     `json:"reserve_balance" pg:"type:numeric(70)"`
	Name                  string     `json:"name"            pg:"type:varchar(255)"`
	Symbol                string     `json:"symbol"          pg:"type:varchar(20)"`
	UpdatedAt             time.Time  `json:"updated_at"`
	DeletedAt             *time.Time `json:"deleted_at"      pg:",soft_delete"`
}
