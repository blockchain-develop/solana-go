// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupiter

import (
	"encoding/json"
	"errors"
	"fmt"
	ag_jsonrpc "github.com/gagliardetto/solana-go/rpc/jsonrpc"
)

var (
	_ *json.Encoder        = nil
	_ *ag_jsonrpc.RPCError = nil
	_ fmt.Formatter        = nil
	_                      = errors.ErrUnsupported
)
var (
	ErrEmptyRoute = &customErrorDef{
		code: 6000,
		msg:  "Empty route",
		name: "EmptyRoute",
	}
	ErrSlippageToleranceExceeded = &customErrorDef{
		code: 6001,
		msg:  "Slippage tolerance exceeded",
		name: "SlippageToleranceExceeded",
	}
	ErrInvalidCalculation = &customErrorDef{
		code: 6002,
		msg:  "Invalid calculation",
		name: "InvalidCalculation",
	}
	ErrMissingPlatformFeeAccount = &customErrorDef{
		code: 6003,
		msg:  "Missing platform fee account",
		name: "MissingPlatformFeeAccount",
	}
	ErrInvalidSlippage = &customErrorDef{
		code: 6004,
		msg:  "Invalid slippage",
		name: "InvalidSlippage",
	}
	ErrNotEnoughPercent = &customErrorDef{
		code: 6005,
		msg:  "Not enough percent to 100",
		name: "NotEnoughPercent",
	}
	ErrInvalidInputIndex = &customErrorDef{
		code: 6006,
		msg:  "Token input index is invalid",
		name: "InvalidInputIndex",
	}
	ErrInvalidOutputIndex = &customErrorDef{
		code: 6007,
		msg:  "Token output index is invalid",
		name: "InvalidOutputIndex",
	}
	ErrNotEnoughAccountKeys = &customErrorDef{
		code: 6008,
		msg:  "Not Enough Account keys",
		name: "NotEnoughAccountKeys",
	}
	ErrNonZeroMinimumOutAmountNotSupported = &customErrorDef{
		code: 6009,
		msg:  "Non zero minimum out amount not supported",
		name: "NonZeroMinimumOutAmountNotSupported",
	}
	ErrInvalidRoutePlan = &customErrorDef{
		code: 6010,
		msg:  "Invalid route plan",
		name: "InvalidRoutePlan",
	}
	ErrInvalidReferralAuthority = &customErrorDef{
		code: 6011,
		msg:  "Invalid referral authority",
		name: "InvalidReferralAuthority",
	}
	ErrLedgerTokenAccountDoesNotMatch = &customErrorDef{
		code: 6012,
		msg:  "Token account doesn't match the ledger",
		name: "LedgerTokenAccountDoesNotMatch",
	}
	ErrInvalidTokenLedger = &customErrorDef{
		code: 6013,
		msg:  "Invalid token ledger",
		name: "InvalidTokenLedger",
	}
	ErrIncorrectTokenProgramID = &customErrorDef{
		code: 6014,
		msg:  "Token program ID is invalid",
		name: "IncorrectTokenProgramID",
	}
	ErrTokenProgramNotProvided = &customErrorDef{
		code: 6015,
		msg:  "Token program not provided",
		name: "TokenProgramNotProvided",
	}
	ErrSwapNotSupported = &customErrorDef{
		code: 6016,
		msg:  "Swap not supported",
		name: "SwapNotSupported",
	}
	ErrExactOutAmountNotMatched = &customErrorDef{
		code: 6017,
		msg:  "Exact out amount doesn't match",
		name: "ExactOutAmountNotMatched",
	}
	ErrSourceAndDestinationMintCannotBeTheSame = &customErrorDef{
		code: 6018,
		msg:  "Source mint and destination mint cannot the same",
		name: "SourceAndDestinationMintCannotBeTheSame",
	}
	Errors = map[int]CustomError{
		6000: ErrEmptyRoute,
		6001: ErrSlippageToleranceExceeded,
		6002: ErrInvalidCalculation,
		6003: ErrMissingPlatformFeeAccount,
		6004: ErrInvalidSlippage,
		6005: ErrNotEnoughPercent,
		6006: ErrInvalidInputIndex,
		6007: ErrInvalidOutputIndex,
		6008: ErrNotEnoughAccountKeys,
		6009: ErrNonZeroMinimumOutAmountNotSupported,
		6010: ErrInvalidRoutePlan,
		6011: ErrInvalidReferralAuthority,
		6012: ErrLedgerTokenAccountDoesNotMatch,
		6013: ErrInvalidTokenLedger,
		6014: ErrIncorrectTokenProgramID,
		6015: ErrTokenProgramNotProvided,
		6016: ErrSwapNotSupported,
		6017: ErrExactOutAmountNotMatched,
		6018: ErrSourceAndDestinationMintCannotBeTheSame,
	}
)

type CustomError interface {
	Code() int
	Name() string
	Error() string
}

type customErrorDef struct {
	code int
	name string
	msg  string
}

func (e *customErrorDef) Code() int {
	return e.code
}

func (e *customErrorDef) Name() string {
	return e.name
}

func (e *customErrorDef) Error() string {
	return fmt.Sprintf("%s(%d): %s", e.name, e.code, e.msg)
}

func DecodeCustomError(rpcErr error) (err error, ok bool) {
	if errCode, o := decodeErrorCode(rpcErr); o {
		if customErr, o := Errors[errCode]; o {
			err = customErr
			ok = true
			return
		}
	}
	return
}

func decodeErrorCode(rpcErr error) (errorCode int, ok bool) {
	var jErr *ag_jsonrpc.RPCError
	if errors.As(rpcErr, &jErr) && jErr.Data != nil {
		if root, o := jErr.Data.(map[string]interface{}); o {
			if rootErr, o := root["err"].(map[string]interface{}); o {
				if rootErrInstructionError, o := rootErr["InstructionError"]; o {
					if rootErrInstructionErrorItems, o := rootErrInstructionError.([]interface{}); o {
						if len(rootErrInstructionErrorItems) == 2 {
							if v, o := rootErrInstructionErrorItems[1].(map[string]interface{}); o {
								if v2, o := v["Custom"].(json.Number); o {
									if code, err := v2.Int64(); err == nil {
										ok = true
										errorCode = int(code)
									}
								} else if v2, o := v["Custom"].(float64); o {
									ok = true
									errorCode = int(v2)
								}
							}
						}
					}
				}
			}
		}
	}
	return
}
