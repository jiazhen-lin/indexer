package block

import (
	"context"
	"fmt"
)

var (
	ErrBlockNotFound       = fmt.Errorf("block not found")
	ErrTransactionNotFound = fmt.Errorf("transaction not found")
)

type Repository interface {
	GetLatestBlocks(ctx context.Context, count int) ([]Block, error)
	GetBlockDetail(ctx context.Context, number uint64) (*BlockDetail, error)
	GetTransaction(ctx context.Context, txHash string) (*Transaction, error)
	UpdateBlock(ctx context.Context, number uint64) error
}
