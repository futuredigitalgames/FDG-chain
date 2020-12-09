
package fdgchain

import (
	"context"
	"errors"
	"math/big"

	"github.com/liuji3978/fdg-chain/common"
	"github.com/liuji3978/fdg-chain/core/types"
)


var NotFound = errors.New("not found")

// TODO: move subscription to package event

type Subscription interface {

	Unsubscribe()

	Err() <-chan error
}


type ChainReader interface {
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error)
	TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error)

	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (Subscription, error)
}


type TransactionReader interface {

	TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error)

	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}


type ChainStateReader interface {

	BalanceAt(ctx context.Context, account string, blockNumber *big.Int) (*big.Int, error)
	StorageAt(ctx context.Context, account string, key common.Hash, blockNumber *big.Int) ([]byte, error)
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
}


type SyncProgress struct {
	StartingBlock uint64 // Block number where sync began
	CurrentBlock  uint64 // Current block number where sync is at
	HighestBlock  uint64 // Highest alleged block number in the chain
	PulledStates  uint64 // Number of state trie entries already downloaded
	KnownStates   uint64 // Total number of state trie entries known about
}


type ChainSyncReader interface {
	SyncProgress(ctx context.Context) (*SyncProgress, error)
}

// CallMsg contains parameters for contract calls.
type CallMsg struct {
	From     common.Address  // the sender of the 'transaction'
	To       *common.Address // the destination contract (nil for contract creation)
	Gas      uint64          // if 0, the call executes with near-infinite gas
	GasPrice *big.Int        // wei <-> gas exchange ratio
	Value    *big.Int        // amount of wei sent along with the call
	Data     []byte          // input data, usually an ABI-encoded contract method invocation
}


type ContractCaller interface {
	CallContract(ctx context.Context, call CallMsg, blockNumber *big.Int) ([]byte, error)
}


type FilterQuery struct {
	BlockHash *common.Hash     // used by eth_getLogs, return logs only from block with this hash
	FromBlock *big.Int         // beginning of the queried range, nil means genesis block
	ToBlock   *big.Int         // end of the range, nil means latest block
	Addresses []common.Address // restricts matches to events created by specific contracts


	Topics [][]common.Hash
}


type LogFilterer interface {
	FilterLogs(ctx context.Context, q FilterQuery) ([]types.Log, error)
	SubscribeFilterLogs(ctx context.Context, q FilterQuery, ch chan<- types.Log) (Subscription, error)
}


type TransactionSender interface {
	SendTransaction(ctx context.Context, tx *types.Transaction) error
}

type GasPricer interface {
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
}

type PendingStateReader interface {
	PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error)
	PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error)
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	PendingTransactionCount(ctx context.Context) (uint, error)
}

type PendingContractCaller interface {
	PendingCallContract(ctx context.Context, call CallMsg) ([]byte, error)
}

type GasEstimator interface {
	EstimateGas(ctx context.Context, call CallMsg) (uint64, error)
}


type PendingStateEventer interface {
	SubscribePendingTransactions(ctx context.Context, ch chan<- *types.Transaction) (Subscription, error)
}
