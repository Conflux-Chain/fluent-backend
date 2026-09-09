package worker

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
)

const errNarrowDownPattern = "narrow down"

// getLogs retrieves logs from the blockchain for a given contract address and optional topic within a specified block range.
//
// Now, its implementation depends on the Confura that do not limit the block number range, and support to index by address + topic0.
func getLogs(client *web3go.Client, contract common.Address, blockFrom, blockTo uint64, topic0 ...common.Hash) ([]types.Log, error) {
	// address filter
	filter := types.FilterQuery{
		Addresses: []common.Address{contract},
	}

	if len(topic0) > 0 {
		filter.Topics = [][]common.Hash{{topic0[0]}}
	}

	from, to := blockFrom, blockTo
	var result []types.Log

	blockNumberConverter := func(number uint64) *types.BlockNumber {
		bn := types.NewBlockNumber(int64(number))
		return &bn
	}

	for from <= to {
		// set block number range in log filter
		filter.FromBlock = blockNumberConverter(from)
		filter.ToBlock = blockNumberConverter(to)

		logs, err := client.Eth.Logs(filter)
		if err == nil {
			// success and move forward
			result = append(result, logs...)

			from = to + 1
			to = blockTo
		} else if strings.Contains(err.Error(), errNarrowDownPattern) {
			// narrow down the block number range
			if from == to {
				return nil, fmt.Errorf("Failed to narrow down block number range, from == to == %v", from)
			}

			to = from + (to-from)/2
		} else {
			// other error
			return nil, errors.WithMessagef(err, "Failed to retrieve event logs, from = %v, to = %v", from, to)
		}
	}

	return result, nil
}
