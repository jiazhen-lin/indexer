package model

import (
	domain "github.com/jiazhen-lin/indexer/application/domain/block"
)

// pk: id
// unique index:
//	- uni_hash
// index:
//	- idx_number_isRemoved
type Block struct {
	ID         uint64 `gorm:"column:id;bigint(20);PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Number     uint64 `gorm:"column:number;index:idx_number_isRemoved;bigint(20);DEFAULT:0;NOT NULL"`
	Hash       string `gorm:"column:hash;uniqueIndex:uni_hash;varchar(100);NOT NULL"`
	Timestamp  uint64 `gorm:"column:timestamp;bigint(10);DEFAULT:0;NOT NULL"`
	ParentHash string `gorm:"column:parentHash;varchar(100);NOT NULL"`
	IsRemoved  int    `gorm:"column:isRemoved;index:idx_number_isRemoved;int(1);DEFAULT:0;NOT NULL"`
}

func (b Block) TableName() string {
	return "Block"
}

func (b Block) ToBlockEntity() domain.Block {
	return domain.Block{
		Number:     b.Number,
		Hash:       b.Hash,
		Time:       b.Timestamp,
		ParentHash: b.ParentHash,
	}
}

func (b Block) ToBlockDetail() domain.BlockDetail {
	return domain.BlockDetail{
		Number:     b.Number,
		Hash:       b.Hash,
		Time:       b.Timestamp,
		ParentHash: b.ParentHash,
	}
}

// pk: id
// index:
// 	- idx_txHash
// 	- idx_blockHash_idx
type Transaction struct {
	ID        uint64 `gorm:"column:id;bigint(20);PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	TxHash    string `gorm:"column:txHash;index:idx_txHash;varchar(100);NOT NULL"`
	BlockHash string `gorm:"column:blockHash;index:idx_blockHash_idx;varchar(100);NOT NULL"`
	Idx       int    `gorm:"column:idx;index:idx_blockHash_idx;bigint(10);DEFAULT:0;NOT NULL"`
	From      string `gorm:"column:from;varchar(100);NOT NULL"`
	To        string `gorm:"column:to;varchar(100);NOT NULL"`
	Nonce     uint64 `gorm:"column:nonce;bigint(10);DEFAULT:0;NOT NULL"`
	Value     string `gorm:"column:value;varchar(100);NOT NULL"`
	Data      string `gorm:"column:data;type:text"`
}

func (t Transaction) TableName() string {
	return "Transaction"
}

func (t Transaction) ToTranactionEntity() domain.Transaction {
	return domain.Transaction{
		Hash:  t.TxHash,
		From:  t.From,
		To:    t.To,
		Nonce: t.Nonce,
		Value: t.Value,
		Data:  t.Data,
	}
}

// pk: id
// index:
// 	- idx_txHash_idx
type TransactionLog struct {
	ID     uint64 `gorm:"column:id;bigint(20);PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	TxHash string `gorm:"column:txHash;index:idx_txHash_idx;varchar(100);NOT NULL"`
	Idx    int    `gorm:"column:idx;index:idx_txHash_idx;bigint(10);DEFAULT:0;NOT NULL"`
	Data   string `gorm:"column:data;text"`
}

func (t TransactionLog) TableName() string {
	return "TransactionLog"
}
