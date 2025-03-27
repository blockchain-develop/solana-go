// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupiter

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type TokenLedgerAccount struct {
	TokenAccount ag_solanago.PublicKey
	Amount       uint64
}

var TokenLedgerAccountDiscriminator = [8]byte{156, 247, 9, 188, 54, 108, 85, 77}

func (obj TokenLedgerAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TokenLedgerAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `TokenAccount` param:
	err = encoder.Encode(obj.TokenAccount)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TokenLedgerAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(TokenLedgerAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[156 247 9 188 54 108 85 77]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `TokenAccount`:
	err = decoder.Decode(&obj.TokenAccount)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
