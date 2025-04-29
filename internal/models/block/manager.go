package block

import (
	"context"
	"explorer/provider/chain"
	"explorer/provider/db"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

var (
	client *ethclient.Client
)

type BlockManager struct {
	Logger *log.Helper
}

func NewBlockManager(logger log.Logger) (*BlockManager, error) {

	if chain.GetEthereumWSClient() == nil {
		return nil, errors.New("ethereum websocket client is nil")
	} else {
		client = chain.GetEthereumWSClient()
	}

	return &BlockManager{
		Logger: log.NewHelper(log.With(logger, "models", "user")),
	}, nil
}

func (m *BlockManager) GetBlockHeight(ctx context.Context) (uint64, error) {
	return client.BlockNumber(ctx)
}

func (m *BlockManager) GetNetworkId(ctx context.Context) (*big.Int, error) {
	return client.NetworkID(ctx)
}

func (m *BlockManager) GetBalance(ctx context.Context, address string) (*big.Int, error) {
	return client.BalanceAt(ctx, common.HexToAddress(address), nil)
}

func (m *BlockManager) GetBlockByNumber(ctx context.Context, number uint64) (*Block, error) {
	blc, err := client.BlockByNumber(ctx, big.NewInt(int64(number)))
	if err != nil {
		return nil, errors.Wrap(err, "get block by number failed")
	}
	return FromTypesBlock(blc), nil
}

func (m *BlockManager) GetBlockByHash(ctx context.Context, hash string) (*Block, error) {
	blc, err := client.BlockByHash(ctx, common.HexToHash(hash))
	if err != nil {
		return nil, errors.Wrap(err, "get block by hash failed")
	}
	return FromTypesBlock(blc), nil
}

func (m *BlockManager) WatcherBlock(ctx context.Context) {
	var bnDB uint64
	if err := db.Get().Raw("SELECT MAX(number) FROM block").Scan(&bnDB); err != nil {
		m.Logger.Error(err)
	}
	bnBlock, err := client.BlockNumber(ctx)
	if err != nil {
		m.Logger.Error(err)
	}
	if bnBlock > bnDB {
		go m.parseHistory(ctx, bnDB+1, bnBlock+1)
	}

	go m.subscribeNewHead(ctx)
}

func (m *BlockManager) subscribeNewHead(ctx context.Context) {
	// TODO
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(ctx, headers)
	if err != nil {
		m.Logger.Error(err)
	}

	for {
		select {
		case err := <-sub.Err():
			m.Logger.Error(err)
		case header := <-headers:
			blc, err := client.BlockByNumber(ctx, header.Number)
			if err != nil {
				m.Logger.Error(err)
			}
			db.Get().Create(FromTypesBlock(blc))
		}
	}
}

func (m *BlockManager) HeaderByNumber(ctx context.Context) (*types.Header, error) {
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "get header by number failed")
	}
	return header, nil
}

func (m *BlockManager) parseHistory(ctx context.Context, start, end uint64) {
	for i := start; i < end; i++ {
		blc, err := client.BlockByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			m.Logger.Error(err)
		}
		db.Get().Save(FromTypesBlock(blc))
	}
}
