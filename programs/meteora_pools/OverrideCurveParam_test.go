// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_pools

import (
	"bytes"
	ag_gofuzz "github.com/gagliardetto/gofuzz"
	ag_require "github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestEncodeDecode_OverrideCurveParam(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("OverrideCurveParam"+strconv.Itoa(i), func(t *testing.T) {
			{
				{
					{
						params := new(OverrideCurveParam)
						fu.Fuzz(params)
						params.AccountMetaSlice = nil
						tmp := new(ConstantProduct)
						fu.Fuzz(tmp)
						params.SetCurveType(tmp)
						buf := new(bytes.Buffer)
						err := encodeT(*params, buf)
						ag_require.NoError(t, err)
						got := new(OverrideCurveParam)
						err = decodeT(got, buf.Bytes())
						got.AccountMetaSlice = nil
						ag_require.NoError(t, err)
						ag_require.Equal(t, params, got)
					}
					{
						params := new(OverrideCurveParam)
						fu.Fuzz(params)
						params.AccountMetaSlice = nil
						tmp := new(Stable)
						fu.Fuzz(tmp)
						params.SetCurveType(tmp)
						buf := new(bytes.Buffer)
						err := encodeT(*params, buf)
						ag_require.NoError(t, err)
						got := new(OverrideCurveParam)
						err = decodeT(got, buf.Bytes())
						got.AccountMetaSlice = nil
						ag_require.NoError(t, err)
						ag_require.Equal(t, params, got)
					}
				}
			}
		})
	}
}
