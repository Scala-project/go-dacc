// Copyright (C) 2017 go-dacc authors
//
// This file is part of the go-dacc library.
//
// the go-dacc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-dacc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; witho
// ut even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dacc library.  If not, see <http://www.gnu.org/licenses/>.
//

package rpc

import (
	"github.com/daccproject/go-dacc/metrics"
)

// Metrics for rpc
var (
	metricsRPCCounter = metrics.NewMeter("neb.rpc.request")

	metricsAccountStateSuccess = metrics.NewMeter("neb.rpc.account.success")
	metricsAccountStateFailed  = metrics.NewMeter("neb.rpc.account.failed")

	metricsSendTxSuccess = metrics.NewMeter("neb.rpc.sendTx.success")
	metricsSendTxFailed  = metrics.NewMeter("neb.rpc.sendTx.failed")

	metricsSignTxSuccess = metrics.NewMeter("neb.rpc.signTx.success")
	metricsSignTxFailed  = metrics.NewMeter("neb.rpc.signTx.failed")

	metricsUnlockSuccess = metrics.NewMeter("neb.rpc.unlock.success")
	metricsUnlockFailed  = metrics.NewMeter("neb.rpc.unlock.failed")
)