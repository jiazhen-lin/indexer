package block

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/jiazhen-lin/indexer/application/domain/block"
)

func NewBlockImpl(blockRepo block.Repository) *Impl {
	return &Impl{
		logger:    logrus.NewEntry(logrus.StandardLogger()),
		blockRepo: blockRepo,
	}
}

type Impl struct {
	logger    *logrus.Entry
	blockRepo block.Repository
}

func (im *Impl) GetRecentBlocks(ctx context.Context, count int) ([]block.Block, error) {
	if count <= 0 {
		return nil, fmt.Errorf("invalid count")
	}

	blocks, err := im.blockRepo.GetLatestBlocks(ctx, count)
	if err != nil {
		im.logger.WithFields(logrus.Fields{
			"err":   err,
			"count": count,
		}).Error("blockRepo.GetLatestBlocks error")
		return nil, err
	}

	return blocks, err
}

func (im *Impl) GetBlockDetail(ctx context.Context, number uint64) (*block.BlockDetail, error) {
	block, err := im.blockRepo.GetBlockDetail(ctx, number)
	if err != nil {
		im.logger.WithFields(logrus.Fields{
			"err":    err,
			"number": number,
		}).Error("blockRepo.GetBlockDetail error")
		return nil, err
	}

	return block, nil
}

func (im *Impl) GetTransaction(ctx context.Context, txHash string) (*block.Transaction, error) {
	tx, err := im.blockRepo.GetTransaction(ctx, txHash)
	if err != nil {
		im.logger.WithFields(logrus.Fields{
			"err":    err,
			"txHash": txHash,
		}).Error("blockRepo.GetTransaction error")
		return nil, err
	}

	return tx, nil
}
