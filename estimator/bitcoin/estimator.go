package bitcoin

import "github.com/lightningnetwork/lnd/lnwallet/chainfee"

type BtcFeeEstimator interface {
	Start() error
	Stop() error
	EstimateFeePerKb() chainfee.SatPerKVByte
}
