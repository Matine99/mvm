package rollup

import (
	"encoding/hex"

	"github.com/ethereum-optimism/optimism/l2geth/common"
	"github.com/ethereum-optimism/optimism/l2geth/core/types"
	"github.com/ethereum-optimism/optimism/l2geth/log"
)

const (
	updateSeqMethod  = ("12345678")
	updateSeqDataLen = 36
)

// RollupAdapter is the adapter for decentralized seqencers
// that is required by the SyncService
type RollupAdapter interface {
	// get tx seqencer for checking tx is valid
	GetTxSeqencer(tx *types.Transaction, expectIndex uint64) (common.Address, error)
	// check current seqencer is working
	CheckSeqencerIsWorking() bool
	//
	CheckPosLayerSynced() bool
}

// SeqAdapter is an adpater used by seqencer based RollupClient
type SeqAdapter struct {
	// posClient  // connnect to pos layer
	// l2client // connect to l2geth client
	l2SeqContract          common.Address // l2 seq contract address
	seqContractValidHeight uint64         // l2 seq contract valid height
}

func NewSeqAdapter(l2SeqContract common.Address, seqContractValidHeight uint64) *SeqAdapter {
	return &SeqAdapter{
		l2SeqContract:          l2SeqContract,
		seqContractValidHeight: seqContractValidHeight,
	}
}

func parseUpdateSeqData(data []byte) (bool, common.Address) {
	if len(data) < updateSeqDataLen {
		return false, common.HexToAddress("0x0")
	}
	method := hex.EncodeToString(data[0:4])
	address := common.BytesToAddress(data[16:36])
	if method == updateSeqMethod {
		return true, address
	}
	return false, common.HexToAddress("0x0")
}
func (s *SeqAdapter) GetTxSeqencer(tx *types.Transaction, expectIndex uint64) (common.Address, error) {
	// check is update seqencer operate
	seqOper, data := tx.IsSystemContractCall(s.l2SeqContract)
	if seqOper {
		updateSeq, newSeq := parseUpdateSeqData(data)
		if updateSeq {
			return newSeq, nil
		}
	}
	if expectIndex > s.seqContractValidHeight {
		log.Debug("Will get seqencer info from seq contract on L2")
		// get status from contract on height expectIndex - 1
		// return result ,err
	}
	// return default address
	return common.HexToAddress("0x0"), nil
}

// CheckSeqencerIsWorking check current seqencer is working
func (s *SeqAdapter) CheckSeqencerIsWorking() bool {
	// if
	return true
}

func (s *SeqAdapter) CheckPosLayerSynced() bool {
	// if
	return true
}
