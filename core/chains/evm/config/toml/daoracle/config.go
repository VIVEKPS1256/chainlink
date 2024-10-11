package daoracle

import "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"

type OracleType string

const (
	OPStack  = OracleType("opstack")
	Arbitrum = OracleType("arbitrum")
	ZKSync   = OracleType("zksync")
)

type DAOracle struct {
	OracleType             OracleType
	OracleAddress          *types.EIP55Address
	CustomGasPriceCalldata string
}

func (d *DAOracle) SetFrom(f *DAOracle) {
	d.OracleType = f.OracleType
	if v := f.OracleAddress; v != nil {
		d.OracleAddress = v
	}
	d.CustomGasPriceCalldata = f.CustomGasPriceCalldata
}
