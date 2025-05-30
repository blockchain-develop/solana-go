// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package lifinity_v2

import ag_binary "github.com/gagliardetto/binary"

type AmmFees struct {
	TradeFeeNumerator           uint64
	TradeFeeDenominator         uint64
	OwnerTradeFeeNumerator      uint64
	OwnerTradeFeeDenominator    uint64
	OwnerWithdrawFeeNumerator   uint64
	OwnerWithdrawFeeDenominator uint64
	HostFeeNumerator            uint64
	HostFeeDenominator          uint64
}

func (obj AmmFees) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `TradeFeeNumerator` param:
	err = encoder.Encode(obj.TradeFeeNumerator)
	if err != nil {
		return err
	}
	// Serialize `TradeFeeDenominator` param:
	err = encoder.Encode(obj.TradeFeeDenominator)
	if err != nil {
		return err
	}
	// Serialize `OwnerTradeFeeNumerator` param:
	err = encoder.Encode(obj.OwnerTradeFeeNumerator)
	if err != nil {
		return err
	}
	// Serialize `OwnerTradeFeeDenominator` param:
	err = encoder.Encode(obj.OwnerTradeFeeDenominator)
	if err != nil {
		return err
	}
	// Serialize `OwnerWithdrawFeeNumerator` param:
	err = encoder.Encode(obj.OwnerWithdrawFeeNumerator)
	if err != nil {
		return err
	}
	// Serialize `OwnerWithdrawFeeDenominator` param:
	err = encoder.Encode(obj.OwnerWithdrawFeeDenominator)
	if err != nil {
		return err
	}
	// Serialize `HostFeeNumerator` param:
	err = encoder.Encode(obj.HostFeeNumerator)
	if err != nil {
		return err
	}
	// Serialize `HostFeeDenominator` param:
	err = encoder.Encode(obj.HostFeeDenominator)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AmmFees) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `TradeFeeNumerator`:
	err = decoder.Decode(&obj.TradeFeeNumerator)
	if err != nil {
		return err
	}
	// Deserialize `TradeFeeDenominator`:
	err = decoder.Decode(&obj.TradeFeeDenominator)
	if err != nil {
		return err
	}
	// Deserialize `OwnerTradeFeeNumerator`:
	err = decoder.Decode(&obj.OwnerTradeFeeNumerator)
	if err != nil {
		return err
	}
	// Deserialize `OwnerTradeFeeDenominator`:
	err = decoder.Decode(&obj.OwnerTradeFeeDenominator)
	if err != nil {
		return err
	}
	// Deserialize `OwnerWithdrawFeeNumerator`:
	err = decoder.Decode(&obj.OwnerWithdrawFeeNumerator)
	if err != nil {
		return err
	}
	// Deserialize `OwnerWithdrawFeeDenominator`:
	err = decoder.Decode(&obj.OwnerWithdrawFeeDenominator)
	if err != nil {
		return err
	}
	// Deserialize `HostFeeNumerator`:
	err = decoder.Decode(&obj.HostFeeNumerator)
	if err != nil {
		return err
	}
	// Deserialize `HostFeeDenominator`:
	err = decoder.Decode(&obj.HostFeeDenominator)
	if err != nil {
		return err
	}
	return nil
}

type AmmCurve struct {
	CurveType       uint8
	CurveParameters uint64
}

func (obj AmmCurve) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CurveType` param:
	err = encoder.Encode(obj.CurveType)
	if err != nil {
		return err
	}
	// Serialize `CurveParameters` param:
	err = encoder.Encode(obj.CurveParameters)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AmmCurve) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CurveType`:
	err = decoder.Decode(&obj.CurveType)
	if err != nil {
		return err
	}
	// Deserialize `CurveParameters`:
	err = decoder.Decode(&obj.CurveParameters)
	if err != nil {
		return err
	}
	return nil
}

type AmmConfig struct {
	LastPrice                uint64
	LastBalancedPrice        uint64
	ConfigDenominator        uint64
	VolumeX                  uint64
	VolumeY                  uint64
	VolumeXInY               uint64
	DepositCap               uint64
	RegressionTarget         uint64
	OracleType               uint64
	OracleStatus             uint64
	OracleMainSlotLimit      uint64
	OracleSubConfidenceLimit uint64
	OracleSubSlotLimit       uint64
	OraclePcConfidenceLimit  uint64
	OraclePcSlotLimit        uint64
	StdSpread                uint64
	StdSpreadBuffer          uint64
	SpreadCoefficient        uint64
	PriceBufferCoin          int64
	PriceBufferPc            int64
	RebalanceRatio           uint64
	FeeTrade                 uint64
	FeePlatform              uint64
	OracleMainSlotBuffer     uint64
	ConfigTemp4              uint64
	ConfigTemp5              uint64
	ConfigTemp6              uint64
	ConfigTemp7              uint64
	ConfigTemp8              uint64
}

func (obj AmmConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LastPrice` param:
	err = encoder.Encode(obj.LastPrice)
	if err != nil {
		return err
	}
	// Serialize `LastBalancedPrice` param:
	err = encoder.Encode(obj.LastBalancedPrice)
	if err != nil {
		return err
	}
	// Serialize `ConfigDenominator` param:
	err = encoder.Encode(obj.ConfigDenominator)
	if err != nil {
		return err
	}
	// Serialize `VolumeX` param:
	err = encoder.Encode(obj.VolumeX)
	if err != nil {
		return err
	}
	// Serialize `VolumeY` param:
	err = encoder.Encode(obj.VolumeY)
	if err != nil {
		return err
	}
	// Serialize `VolumeXInY` param:
	err = encoder.Encode(obj.VolumeXInY)
	if err != nil {
		return err
	}
	// Serialize `DepositCap` param:
	err = encoder.Encode(obj.DepositCap)
	if err != nil {
		return err
	}
	// Serialize `RegressionTarget` param:
	err = encoder.Encode(obj.RegressionTarget)
	if err != nil {
		return err
	}
	// Serialize `OracleType` param:
	err = encoder.Encode(obj.OracleType)
	if err != nil {
		return err
	}
	// Serialize `OracleStatus` param:
	err = encoder.Encode(obj.OracleStatus)
	if err != nil {
		return err
	}
	// Serialize `OracleMainSlotLimit` param:
	err = encoder.Encode(obj.OracleMainSlotLimit)
	if err != nil {
		return err
	}
	// Serialize `OracleSubConfidenceLimit` param:
	err = encoder.Encode(obj.OracleSubConfidenceLimit)
	if err != nil {
		return err
	}
	// Serialize `OracleSubSlotLimit` param:
	err = encoder.Encode(obj.OracleSubSlotLimit)
	if err != nil {
		return err
	}
	// Serialize `OraclePcConfidenceLimit` param:
	err = encoder.Encode(obj.OraclePcConfidenceLimit)
	if err != nil {
		return err
	}
	// Serialize `OraclePcSlotLimit` param:
	err = encoder.Encode(obj.OraclePcSlotLimit)
	if err != nil {
		return err
	}
	// Serialize `StdSpread` param:
	err = encoder.Encode(obj.StdSpread)
	if err != nil {
		return err
	}
	// Serialize `StdSpreadBuffer` param:
	err = encoder.Encode(obj.StdSpreadBuffer)
	if err != nil {
		return err
	}
	// Serialize `SpreadCoefficient` param:
	err = encoder.Encode(obj.SpreadCoefficient)
	if err != nil {
		return err
	}
	// Serialize `PriceBufferCoin` param:
	err = encoder.Encode(obj.PriceBufferCoin)
	if err != nil {
		return err
	}
	// Serialize `PriceBufferPc` param:
	err = encoder.Encode(obj.PriceBufferPc)
	if err != nil {
		return err
	}
	// Serialize `RebalanceRatio` param:
	err = encoder.Encode(obj.RebalanceRatio)
	if err != nil {
		return err
	}
	// Serialize `FeeTrade` param:
	err = encoder.Encode(obj.FeeTrade)
	if err != nil {
		return err
	}
	// Serialize `FeePlatform` param:
	err = encoder.Encode(obj.FeePlatform)
	if err != nil {
		return err
	}
	// Serialize `OracleMainSlotBuffer` param:
	err = encoder.Encode(obj.OracleMainSlotBuffer)
	if err != nil {
		return err
	}
	// Serialize `ConfigTemp4` param:
	err = encoder.Encode(obj.ConfigTemp4)
	if err != nil {
		return err
	}
	// Serialize `ConfigTemp5` param:
	err = encoder.Encode(obj.ConfigTemp5)
	if err != nil {
		return err
	}
	// Serialize `ConfigTemp6` param:
	err = encoder.Encode(obj.ConfigTemp6)
	if err != nil {
		return err
	}
	// Serialize `ConfigTemp7` param:
	err = encoder.Encode(obj.ConfigTemp7)
	if err != nil {
		return err
	}
	// Serialize `ConfigTemp8` param:
	err = encoder.Encode(obj.ConfigTemp8)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AmmConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LastPrice`:
	err = decoder.Decode(&obj.LastPrice)
	if err != nil {
		return err
	}
	// Deserialize `LastBalancedPrice`:
	err = decoder.Decode(&obj.LastBalancedPrice)
	if err != nil {
		return err
	}
	// Deserialize `ConfigDenominator`:
	err = decoder.Decode(&obj.ConfigDenominator)
	if err != nil {
		return err
	}
	// Deserialize `VolumeX`:
	err = decoder.Decode(&obj.VolumeX)
	if err != nil {
		return err
	}
	// Deserialize `VolumeY`:
	err = decoder.Decode(&obj.VolumeY)
	if err != nil {
		return err
	}
	// Deserialize `VolumeXInY`:
	err = decoder.Decode(&obj.VolumeXInY)
	if err != nil {
		return err
	}
	// Deserialize `DepositCap`:
	err = decoder.Decode(&obj.DepositCap)
	if err != nil {
		return err
	}
	// Deserialize `RegressionTarget`:
	err = decoder.Decode(&obj.RegressionTarget)
	if err != nil {
		return err
	}
	// Deserialize `OracleType`:
	err = decoder.Decode(&obj.OracleType)
	if err != nil {
		return err
	}
	// Deserialize `OracleStatus`:
	err = decoder.Decode(&obj.OracleStatus)
	if err != nil {
		return err
	}
	// Deserialize `OracleMainSlotLimit`:
	err = decoder.Decode(&obj.OracleMainSlotLimit)
	if err != nil {
		return err
	}
	// Deserialize `OracleSubConfidenceLimit`:
	err = decoder.Decode(&obj.OracleSubConfidenceLimit)
	if err != nil {
		return err
	}
	// Deserialize `OracleSubSlotLimit`:
	err = decoder.Decode(&obj.OracleSubSlotLimit)
	if err != nil {
		return err
	}
	// Deserialize `OraclePcConfidenceLimit`:
	err = decoder.Decode(&obj.OraclePcConfidenceLimit)
	if err != nil {
		return err
	}
	// Deserialize `OraclePcSlotLimit`:
	err = decoder.Decode(&obj.OraclePcSlotLimit)
	if err != nil {
		return err
	}
	// Deserialize `StdSpread`:
	err = decoder.Decode(&obj.StdSpread)
	if err != nil {
		return err
	}
	// Deserialize `StdSpreadBuffer`:
	err = decoder.Decode(&obj.StdSpreadBuffer)
	if err != nil {
		return err
	}
	// Deserialize `SpreadCoefficient`:
	err = decoder.Decode(&obj.SpreadCoefficient)
	if err != nil {
		return err
	}
	// Deserialize `PriceBufferCoin`:
	err = decoder.Decode(&obj.PriceBufferCoin)
	if err != nil {
		return err
	}
	// Deserialize `PriceBufferPc`:
	err = decoder.Decode(&obj.PriceBufferPc)
	if err != nil {
		return err
	}
	// Deserialize `RebalanceRatio`:
	err = decoder.Decode(&obj.RebalanceRatio)
	if err != nil {
		return err
	}
	// Deserialize `FeeTrade`:
	err = decoder.Decode(&obj.FeeTrade)
	if err != nil {
		return err
	}
	// Deserialize `FeePlatform`:
	err = decoder.Decode(&obj.FeePlatform)
	if err != nil {
		return err
	}
	// Deserialize `OracleMainSlotBuffer`:
	err = decoder.Decode(&obj.OracleMainSlotBuffer)
	if err != nil {
		return err
	}
	// Deserialize `ConfigTemp4`:
	err = decoder.Decode(&obj.ConfigTemp4)
	if err != nil {
		return err
	}
	// Deserialize `ConfigTemp5`:
	err = decoder.Decode(&obj.ConfigTemp5)
	if err != nil {
		return err
	}
	// Deserialize `ConfigTemp6`:
	err = decoder.Decode(&obj.ConfigTemp6)
	if err != nil {
		return err
	}
	// Deserialize `ConfigTemp7`:
	err = decoder.Decode(&obj.ConfigTemp7)
	if err != nil {
		return err
	}
	// Deserialize `ConfigTemp8`:
	err = decoder.Decode(&obj.ConfigTemp8)
	if err != nil {
		return err
	}
	return nil
}

type CurveType ag_binary.BorshEnum

const (
	CurveTypeStandard CurveType = iota
	CurveTypeConstantProduct
)

func (value CurveType) String() string {
	switch value {
	case CurveTypeStandard:
		return "Standard"
	case CurveTypeConstantProduct:
		return "ConstantProduct"
	default:
		return ""
	}
}
