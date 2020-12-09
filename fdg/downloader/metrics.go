// Copyright 2015 The go-fdg Authors
// This file is part of the go-fdg library.
//
// The go-fdg library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-fdg library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-fdg library. If not, see <http://www.gnu.org/licenses/>.

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/liuji3978/fdg-chain/metrics"
)

var (
	headerInMeter      = metrics.NewRegisteredMeter("fdg/downloader/headers/in", nil)
	headerReqTimer     = metrics.NewRegisteredTimer("fdg/downloader/headers/req", nil)
	headerDropMeter    = metrics.NewRegisteredMeter("fdg/downloader/headers/drop", nil)
	headerTimeoutMeter = metrics.NewRegisteredMeter("fdg/downloader/headers/timeout", nil)

	bodyInMeter      = metrics.NewRegisteredMeter("fdg/downloader/bodies/in", nil)
	bodyReqTimer     = metrics.NewRegisteredTimer("fdg/downloader/bodies/req", nil)
	bodyDropMeter    = metrics.NewRegisteredMeter("fdg/downloader/bodies/drop", nil)
	bodyTimeoutMeter = metrics.NewRegisteredMeter("fdg/downloader/bodies/timeout", nil)

	receiptInMeter      = metrics.NewRegisteredMeter("fdg/downloader/receipts/in", nil)
	receiptReqTimer     = metrics.NewRegisteredTimer("fdg/downloader/receipts/req", nil)
	receiptDropMeter    = metrics.NewRegisteredMeter("fdg/downloader/receipts/drop", nil)
	receiptTimeoutMeter = metrics.NewRegisteredMeter("fdg/downloader/receipts/timeout", nil)

	stateInMeter   = metrics.NewRegisteredMeter("fdg/downloader/states/in", nil)
	stateDropMeter = metrics.NewRegisteredMeter("fdg/downloader/states/drop", nil)

	throttleCounter = metrics.NewRegisteredCounter("eth/downloader/throttle", nil)
)
