package service

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
)

type TxSender struct {
	client *web3go.Client

	sender common.Address

	chainId    uint64
	chainIdBig *hexutil.Big

	mu sync.Mutex
}

func NewTxSender(client *web3go.Client) (*TxSender, error) {
	// get the default signer
	sm, err := client.GetSignerManager()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get signer manager from RPC client")
	}

	signers := sm.List()
	if len(signers) == 0 {
		return nil, errors.New("No signer found")
	}

	sender := signers[0].Address()

	// check sender balance
	balance, err := client.Eth.Balance(sender, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to retrieve sender balance")
	}

	if balance.Sign() == 0 {
		return nil, errors.Errorf("Sender balance is 0, address = %v", sender)
	}

	// retrieve chain ID
	chainId, err := client.Eth.ChainId()
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to retrieve chain ID")
	}

	if chainId == nil {
		return nil, errors.New("Chain ID unavailable on fullnode")
	}

	return &TxSender{
		client:     client,
		sender:     sender,
		chainId:    *chainId,
		chainIdBig: (*hexutil.Big)(new(big.Int).SetUint64(*chainId)),
	}, nil
}

func (s *TxSender) Send(tx types.TransactionArgs) (common.Hash, error) {
	if tx.From == nil {
		tx.From = &s.sender
	}

	if tx.ChainID == nil {
		tx.ChainID = s.chainIdBig
	}

	// use mutex to ensure correct tx nonce, and need to enhance for high QPS case
	s.mu.Lock()
	defer s.mu.Unlock()

	txHash, err := s.client.Eth.SendTransactionByArgs(tx)
	if err != nil {
		return common.Hash{}, NewRPCError(err, "Failed to send transaction")
	}

	return txHash, nil
}
