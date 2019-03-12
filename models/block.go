package models

import (
	"time"
)

type Block struct {
	ID                  uint64                `json:"id"`
	TotalTxs            uint64                `json:"txTotal" sql:"default:0"`
	Size                uint64                `json:"size"`
	ProposerValidatorID uint64                `json:"proposer_validator_id"`
	NumTxs              uint32                `json:"txCount" sql:"default:0"`
	BlockTime           uint64                `json:"blockTime"`
	CreatedAt           time.Time             `json:"createdAt"`
	UpdatedAt           time.Time             `json:"updatedAt"`
	BlockReward         string                `json:"blockReward" sql:"type:numeric(70)"`
	Hash                string                `json:"hash"`
	Proposer            *Validator            `json:"proposer" pg:"fk:proposer_validator_id"`    //relation has one to Validators
	Validators          []*Validator          `json:"validators" pg:"many2many:block_validator"` //relation has many to Validators
	Transactions        []*Transaction        `json:"transactions"`                              //relation has many to Transactions
	InvalidTransactions []*InvalidTransaction `json:"invalid_transactions"`                      //relation has many to InvalidTransactions
	Rewards             []*Reward             `json:"rewards"`                                   //relation has many to Rewards
	Slashes             []*Slash              `json:"slashes"`                                   //relation has many to Slashes
}

//Return block hash with prefix
func (t *Block) GetHash() string {
	return `Mh` + t.Hash
}

type BlockAddresses struct {
	Height    uint64
	Addresses []string
}
