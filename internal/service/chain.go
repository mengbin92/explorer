package service

import (
	"context"
	"math/big"
	"net/http"

	pb "explorer/api/explorer/v1"
	"explorer/provider/chain"

	"github.com/bytedance/sonic"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChainService struct {
	pb.UnimplementedChainServer
	Logger *log.Helper
}

func NewChainService(logger log.Logger) *ChainService {
	return &ChainService{
		Logger: log.NewHelper(logger),
	}
}

func (s *ChainService) GetBlockNumer(ctx context.Context, req *emptypb.Empty) (*pb.GetBlockNumerReply, error) {
	bn, err := chain.GetEthereumHttpClient().BlockNumber(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get block number failed")
	}
	return &pb.GetBlockNumerReply{
		Status: &pb.Status{
			Code:    http.StatusOK,
			Message: "get block number success",
		},
		BlockNumber: bn,
	}, nil
}
func (s *ChainService) GetNetworkId(ctx context.Context, req *emptypb.Empty) (*pb.GetNetworkIdReply, error) {
	id, err := chain.GetEthereumHttpClient().NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get network id failed")
	}
	return &pb.GetNetworkIdReply{
		Status: &pb.Status{
			Code:    http.StatusOK,
			Message: "get network id success",
		},
		NetworkId: id.Uint64(),
	}, nil
}
func (s *ChainService) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceReply, error) {
	balance, err := chain.GetEthereumHttpClient().BalanceAt(ctx, common.HexToAddress(req.Address), nil)
	if err != nil {
		return nil, errors.Wrap(err, "get balance failed")
	}
	return &pb.GetBalanceReply{
		Status: &pb.Status{
			Code:    http.StatusOK,
			Message: "get balance success",
		},
		Balance: balance.Uint64(),
	}, nil
}
func (s *ChainService) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionReply, error) {
	tx, _, err := chain.GetEthereumHttpClient().TransactionByHash(ctx, common.HexToHash(req.TransactionHash))
	if err != nil {
		return nil, errors.Wrap(err, "get transaction failed")
	}
	txBytes, err := tx.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "marshal transaction failed")
	}
	return &pb.GetTransactionReply{
		Status: &pb.Status{
			Code:    http.StatusOK,
			Message: "get transaction success",
		},
		Transaction: string(txBytes),
	}, nil
}
func (s *ChainService) GetTransactionReceipt(ctx context.Context, req *pb.GetTransactionReceiptRequest) (*pb.GetTransactionReceiptReply, error) {
	recipt, err := chain.GetEthereumHttpClient().TransactionReceipt(ctx, common.HexToHash(req.TransactionHash))
	if err != nil {
		return nil, errors.Wrap(err, "get transaction receipt failed")
	}
	reciptBytes, err := recipt.MarshalJSON()
	if err != nil {
		return nil, errors.Wrap(err, "marshal transaction receipt failed")
	}
	return &pb.GetTransactionReceiptReply{
		Status: &pb.Status{
			Code:    http.StatusOK,
			Message: "get transaction receipt success",
		},
		TransactionReceipt: string(reciptBytes),
	}, nil
}
func (s *ChainService) GetBlockByNumber(ctx context.Context, req *pb.GetBlockByNumberRequest) (*pb.GetBlockReply, error) {
	block, err := chain.GetEthereumHttpClient().BlockByNumber(ctx, big.NewInt(int64(req.BlockNumber)))
	if err != nil {
		return nil, errors.Wrap(err, "get block by number failed")
	}
	blockBytes, err := sonic.Marshal(FromTypesBlock(block))
	if err != nil {
		return nil, errors.Wrap(err, "marshal block failed")
	}
	return &pb.GetBlockReply{
		Status: &pb.Status{
			Code:    http.StatusOK,
			Message: "get block by number success",
		},
		Block: string(blockBytes),
	}, nil
}
func (s *ChainService) GetBlockByHash(ctx context.Context, req *pb.GetBlockByHashRequest) (*pb.GetBlockReply, error) {
	block, err := chain.GetEthereumHttpClient().BlockByHash(ctx, common.HexToHash(req.BlockHash))
	if err != nil {
		return nil, errors.Wrap(err, "get block by hash failed")
	}
	blockBytes, err := sonic.Marshal(FromTypesBlock(block))
	if err != nil {
		return nil, errors.Wrap(err, "marshal block failed")
	}
	return &pb.GetBlockReply{
		Status: &pb.Status{
			Code:    http.StatusOK,
			Message: "get block by hash success",
		},
		Block: string(blockBytes),
	}, nil
}
