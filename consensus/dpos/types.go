// Copyright (C) 2017-2019 go-dacc authors
//
// This file is part of the go-dacc library.
//
// the go-dacc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-dacc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dacc library.  If not, see <http://www.gnu.org/licenses/>.
//

package dpos

import (
	"errors"
	"strconv"
	"time"

	"github.com/daccproject/go-dacc/metrics"
)

// const
const (
	DefaultMaxUnlockDuration time.Duration = 1<<63 - 1
	GenesisDynasty                         = 1
	GenesisDynastySerial                   = 0
)

// Consensus Related Constants
const (
	SecondInMs               = int64(1000)
	BlockIntervalInMs        = int64(15000)
	AcceptedNetWorkDelayInMs = int64(3750)
	MaxMintDurationInMs      = int64(5250)
	MinMintDurationInMs      = int64(2250)
	DynastyIntervalInMs      = int64(3150000)
	DynastySize              = 21
	ConsensusSize            = DynastySize*2/3 + 1
)

// Errors in PoW Consensus
var (
	ErrInvalidBlockTimestamp      = errors.New("invalid block timestamp, should be same as consensus's timestamp")
	ErrInvalidBlockInterval       = errors.New("invalid block interval")
	ErrMissingConfigForDpos       = errors.New("missing configuration for Dpos")
	ErrInvalidBlockProposer       = errors.New("invalid block proposer")
	ErrCannotMintWhenPending      = errors.New("cannot mint block now, waiting for cancel pending again")
	ErrCannotMintWhenDisable      = errors.New("cannot mint block now, waiting for enable it again")
	ErrWaitingBlockInLastSlot     = errors.New("cannot mint block now, waiting for last block")
	ErrBlockMintedInNextSlot      = errors.New("cannot mint block now, there is a block minted in current slot")
	ErrGenerateNextConsensusState = errors.New("Failed to generate next consensus state")
	ErrDoubleBlockMinted          = errors.New("double block minted")
	ErrAppendNewBlockFailed       = errors.New("failed to append new block to real chain")
	ErrInvalidArgument            = errors.New("invalid argument")
)

// Errors in dpos state
var (
	ErrTooFewCandidates        = errors.New("the size of candidates in consensus is un-safe, should be greater than or equal " + strconv.Itoa(ConsensusSize))
	ErrInitialDynastyNotEnough = errors.New("the size of initial dynasty in genesis block is un-safe, should be greater than or equal " + strconv.Itoa(ConsensusSize))
	ErrInvalidDynasty          = errors.New("the size of initial dynasty in genesis block is invalid, should be equal " + strconv.Itoa(DynastySize))
	ErrCloneDynastyTrie        = errors.New("Failed to clone dynasty trie")
	ErrCloneNextDynastyTrie    = errors.New("Failed to clone next dynasty trie")
	ErrCloneDelegateTrie       = errors.New("Failed to clone delegate trie")
	ErrCloneCandidatesTrie     = errors.New("Failed to clone candidates trie")
	ErrCloneVoteTrie           = errors.New("Failed to clone vote trie")
	ErrCloneMintCntTrie        = errors.New("Failed to clone mint count trie")
	ErrNotBlockForgTime        = errors.New("now is not time to forg block")
	ErrFoundNilProposer        = errors.New("found a nil proposer")
)

// Error in dynasty
var (
	ErrFailedToLoadDynasty     = errors.New("Failed to load dynasty file.")
	ErrFailedToParseDynasty    = errors.New("Failed to parse dynasty.conf.")
	ErrCheckDynastyChainID     = errors.New("ChainId in dynasty.conf differs from that in genesis.conf.")
	ErrCheckDynastyMinersCount = errors.New("Miners count in dynasty.conf differs from that in genesis.conf.")
)

// Metrics
var (
	metricsBlockPackingTime = metrics.NewGauge("neb.block.packing")
	metricsBlockWaitingTime = metrics.NewGauge("neb.block.waiting")
	metricsLruPoolSlotBlock = metrics.NewGauge("neb.block.lru.poolslot")
	metricsMintBlock        = metrics.NewCounter("neb.block.mint")
)
