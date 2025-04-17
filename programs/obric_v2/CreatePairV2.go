// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package obric_v2

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreatePairV2 is the `createPairV2` instruction.
type CreatePairV2 struct {
	Concentration              *uint64
	FeeMillionth               *uint64
	RebatePercentage           *uint64
	ProtocolFeeShareThousandth *uint64

	// [0] = [WRITE] tradingPair
	//
	// [1] = [] mintX
	//
	// [2] = [] mintY
	//
	// [3] = [WRITE] mintSslpX
	//
	// [4] = [WRITE] mintSslpY
	//
	// [5] = [WRITE] reserveX
	//
	// [6] = [WRITE] reserveY
	//
	// [7] = [WRITE] protocolFeeX
	//
	// [8] = [WRITE] protocolFeeY
	//
	// [9] = [] xPriceFeed
	//
	// [10] = [] yPriceFeed
	//
	// [11] = [WRITE, SIGNER] user
	//
	// [12] = [] systemProgram
	//
	// [13] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreatePairV2InstructionBuilder creates a new `CreatePairV2` instruction builder.
func NewCreatePairV2InstructionBuilder() *CreatePairV2 {
	nd := &CreatePairV2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 14),
	}
	return nd
}

// SetConcentration sets the "concentration" parameter.
func (inst *CreatePairV2) SetConcentration(concentration uint64) *CreatePairV2 {
	inst.Concentration = &concentration
	return inst
}

// SetFeeMillionth sets the "feeMillionth" parameter.
func (inst *CreatePairV2) SetFeeMillionth(feeMillionth uint64) *CreatePairV2 {
	inst.FeeMillionth = &feeMillionth
	return inst
}

// SetRebatePercentage sets the "rebatePercentage" parameter.
func (inst *CreatePairV2) SetRebatePercentage(rebatePercentage uint64) *CreatePairV2 {
	inst.RebatePercentage = &rebatePercentage
	return inst
}

// SetProtocolFeeShareThousandth sets the "protocolFeeShareThousandth" parameter.
func (inst *CreatePairV2) SetProtocolFeeShareThousandth(protocolFeeShareThousandth uint64) *CreatePairV2 {
	inst.ProtocolFeeShareThousandth = &protocolFeeShareThousandth
	return inst
}

// SetTradingPairAccount sets the "tradingPair" account.
func (inst *CreatePairV2) SetTradingPairAccount(tradingPair ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tradingPair).WRITE()
	return inst
}

// GetTradingPairAccount gets the "tradingPair" account.
func (inst *CreatePairV2) GetTradingPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintXAccount sets the "mintX" account.
func (inst *CreatePairV2) SetMintXAccount(mintX ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mintX)
	return inst
}

// GetMintXAccount gets the "mintX" account.
func (inst *CreatePairV2) GetMintXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintYAccount sets the "mintY" account.
func (inst *CreatePairV2) SetMintYAccount(mintY ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mintY)
	return inst
}

// GetMintYAccount gets the "mintY" account.
func (inst *CreatePairV2) GetMintYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintSslpXAccount sets the "mintSslpX" account.
func (inst *CreatePairV2) SetMintSslpXAccount(mintSslpX ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mintSslpX).WRITE()
	return inst
}

// GetMintSslpXAccount gets the "mintSslpX" account.
func (inst *CreatePairV2) GetMintSslpXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMintSslpYAccount sets the "mintSslpY" account.
func (inst *CreatePairV2) SetMintSslpYAccount(mintSslpY ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(mintSslpY).WRITE()
	return inst
}

// GetMintSslpYAccount gets the "mintSslpY" account.
func (inst *CreatePairV2) GetMintSslpYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetReserveXAccount sets the "reserveX" account.
func (inst *CreatePairV2) SetReserveXAccount(reserveX ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(reserveX).WRITE()
	return inst
}

// GetReserveXAccount gets the "reserveX" account.
func (inst *CreatePairV2) GetReserveXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetReserveYAccount sets the "reserveY" account.
func (inst *CreatePairV2) SetReserveYAccount(reserveY ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(reserveY).WRITE()
	return inst
}

// GetReserveYAccount gets the "reserveY" account.
func (inst *CreatePairV2) GetReserveYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetProtocolFeeXAccount sets the "protocolFeeX" account.
func (inst *CreatePairV2) SetProtocolFeeXAccount(protocolFeeX ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(protocolFeeX).WRITE()
	return inst
}

// GetProtocolFeeXAccount gets the "protocolFeeX" account.
func (inst *CreatePairV2) GetProtocolFeeXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetProtocolFeeYAccount sets the "protocolFeeY" account.
func (inst *CreatePairV2) SetProtocolFeeYAccount(protocolFeeY ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(protocolFeeY).WRITE()
	return inst
}

// GetProtocolFeeYAccount gets the "protocolFeeY" account.
func (inst *CreatePairV2) GetProtocolFeeYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetXPriceFeedAccount sets the "xPriceFeed" account.
func (inst *CreatePairV2) SetXPriceFeedAccount(xPriceFeed ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(xPriceFeed)
	return inst
}

// GetXPriceFeedAccount gets the "xPriceFeed" account.
func (inst *CreatePairV2) GetXPriceFeedAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetYPriceFeedAccount sets the "yPriceFeed" account.
func (inst *CreatePairV2) SetYPriceFeedAccount(yPriceFeed ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(yPriceFeed)
	return inst
}

// GetYPriceFeedAccount gets the "yPriceFeed" account.
func (inst *CreatePairV2) GetYPriceFeedAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetUserAccount sets the "user" account.
func (inst *CreatePairV2) SetUserAccount(user ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *CreatePairV2) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreatePairV2) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreatePairV2) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CreatePairV2) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CreatePairV2 {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CreatePairV2) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

func (inst CreatePairV2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreatePairV2,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreatePairV2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreatePairV2) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Concentration == nil {
			return errors.New("Concentration parameter is not set")
		}
		if inst.FeeMillionth == nil {
			return errors.New("FeeMillionth parameter is not set")
		}
		if inst.RebatePercentage == nil {
			return errors.New("RebatePercentage parameter is not set")
		}
		if inst.ProtocolFeeShareThousandth == nil {
			return errors.New("ProtocolFeeShareThousandth parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TradingPair is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MintX is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MintY is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MintSslpX is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MintSslpY is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ReserveX is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ReserveY is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.ProtocolFeeX is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.ProtocolFeeY is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.XPriceFeed is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.YPriceFeed is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *CreatePairV2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreatePairV2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("             Concentration", *inst.Concentration))
						paramsBranch.Child(ag_format.Param("              FeeMillionth", *inst.FeeMillionth))
						paramsBranch.Child(ag_format.Param("          RebatePercentage", *inst.RebatePercentage))
						paramsBranch.Child(ag_format.Param("ProtocolFeeShareThousandth", *inst.ProtocolFeeShareThousandth))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=14]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  tradingPair", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        mintX", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        mintY", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    mintSslpX", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    mintSslpY", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("     reserveX", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     reserveY", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta(" protocolFeeX", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta(" protocolFeeY", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("   xPriceFeed", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("   yPriceFeed", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("         user", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta(" tokenProgram", inst.AccountMetaSlice.Get(13)))
					})
				})
		})
}

func (obj CreatePairV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Concentration` param:
	err = encoder.Encode(obj.Concentration)
	if err != nil {
		return err
	}
	// Serialize `FeeMillionth` param:
	err = encoder.Encode(obj.FeeMillionth)
	if err != nil {
		return err
	}
	// Serialize `RebatePercentage` param:
	err = encoder.Encode(obj.RebatePercentage)
	if err != nil {
		return err
	}
	// Serialize `ProtocolFeeShareThousandth` param:
	err = encoder.Encode(obj.ProtocolFeeShareThousandth)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreatePairV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Concentration`:
	err = decoder.Decode(&obj.Concentration)
	if err != nil {
		return err
	}
	// Deserialize `FeeMillionth`:
	err = decoder.Decode(&obj.FeeMillionth)
	if err != nil {
		return err
	}
	// Deserialize `RebatePercentage`:
	err = decoder.Decode(&obj.RebatePercentage)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolFeeShareThousandth`:
	err = decoder.Decode(&obj.ProtocolFeeShareThousandth)
	if err != nil {
		return err
	}
	return nil
}

// NewCreatePairV2Instruction declares a new CreatePairV2 instruction with the provided parameters and accounts.
func NewCreatePairV2Instruction(
	// Parameters:
	concentration uint64,
	feeMillionth uint64,
	rebatePercentage uint64,
	protocolFeeShareThousandth uint64,
	// Accounts:
	tradingPair ag_solanago.PublicKey,
	mintX ag_solanago.PublicKey,
	mintY ag_solanago.PublicKey,
	mintSslpX ag_solanago.PublicKey,
	mintSslpY ag_solanago.PublicKey,
	reserveX ag_solanago.PublicKey,
	reserveY ag_solanago.PublicKey,
	protocolFeeX ag_solanago.PublicKey,
	protocolFeeY ag_solanago.PublicKey,
	xPriceFeed ag_solanago.PublicKey,
	yPriceFeed ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *CreatePairV2 {
	return NewCreatePairV2InstructionBuilder().
		SetConcentration(concentration).
		SetFeeMillionth(feeMillionth).
		SetRebatePercentage(rebatePercentage).
		SetProtocolFeeShareThousandth(protocolFeeShareThousandth).
		SetTradingPairAccount(tradingPair).
		SetMintXAccount(mintX).
		SetMintYAccount(mintY).
		SetMintSslpXAccount(mintSslpX).
		SetMintSslpYAccount(mintSslpY).
		SetReserveXAccount(reserveX).
		SetReserveYAccount(reserveY).
		SetProtocolFeeXAccount(protocolFeeX).
		SetProtocolFeeYAccount(protocolFeeY).
		SetXPriceFeedAccount(xPriceFeed).
		SetYPriceFeedAccount(yPriceFeed).
		SetUserAccount(user).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram)
}
