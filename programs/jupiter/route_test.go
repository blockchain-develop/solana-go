// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupiter

import (
	"bytes"
	"encoding/hex"
	"fmt"
	ag_gofuzz "github.com/gagliardetto/gofuzz"
	ag_require "github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestEncodeDecode_Route(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("Route"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(Route)
				fu.Fuzz(params)
				params.AccountMetaSlice = nil
				params.RoutePlan = nil
				params.InAmount = nil
				params.QuotedOutAmount = nil
				params.SlippageBps = nil
				params.PlatformFeeBps = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(Route)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}

func TestEncodeDecode_RoutePlan(t *testing.T) {
	var routePlan *RoutePlanStep
	s := SwapTokenSwapTuple{}
	routePlan = &RoutePlanStep{
		Swap:        &Swap{
			Value: s,
		},
		Percent:     100,
		InputIndex:  1,
		OutputIndex: 2,
	}
	buf := new(bytes.Buffer)
	err := encodeT(*routePlan, buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", hex.EncodeToString(buf.Bytes()))
}
