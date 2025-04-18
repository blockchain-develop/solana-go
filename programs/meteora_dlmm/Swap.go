// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Swap is the `swap` instruction.
type Swap struct {
	AmountIn     *uint64
	MinAmountOut *uint64

	// [0] = [WRITE] lbPair
	//
	// [1] = [] binArrayBitmapExtension
	//
	// [2] = [WRITE] reserveX
	//
	// [3] = [WRITE] reserveY
	//
	// [4] = [WRITE] userTokenIn
	//
	// [5] = [WRITE] userTokenOut
	//
	// [6] = [] tokenXMint
	//
	// [7] = [] tokenYMint
	//
	// [8] = [WRITE] oracle
	//
	// [9] = [WRITE] hostFeeIn
	//
	// [10] = [SIGNER] user
	//
	// [11] = [] tokenXProgram
	//
	// [12] = [] tokenYProgram
	//
	// [13] = [] eventAuthority
	//
	// [14] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSwapInstructionBuilder creates a new `Swap` instruction builder.
func NewSwapInstructionBuilder() *Swap {
	nd := &Swap{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 15),
	}
	return nd
}

// SetAmountIn sets the "amountIn" parameter.
func (inst *Swap) SetAmountIn(amountIn uint64) *Swap {
	inst.AmountIn = &amountIn
	return inst
}

// SetMinAmountOut sets the "minAmountOut" parameter.
func (inst *Swap) SetMinAmountOut(minAmountOut uint64) *Swap {
	inst.MinAmountOut = &minAmountOut
	return inst
}

// SetLbPairAccount sets the "lbPair" account.
func (inst *Swap) SetLbPairAccount(lbPair ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(lbPair).WRITE()
	return inst
}

// GetLbPairAccount gets the "lbPair" account.
func (inst *Swap) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetBinArrayBitmapExtensionAccount sets the "binArrayBitmapExtension" account.
func (inst *Swap) SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(binArrayBitmapExtension)
	return inst
}

// GetBinArrayBitmapExtensionAccount gets the "binArrayBitmapExtension" account.
func (inst *Swap) GetBinArrayBitmapExtensionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetReserveXAccount sets the "reserveX" account.
func (inst *Swap) SetReserveXAccount(reserveX ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(reserveX).WRITE()
	return inst
}

// GetReserveXAccount gets the "reserveX" account.
func (inst *Swap) GetReserveXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetReserveYAccount sets the "reserveY" account.
func (inst *Swap) SetReserveYAccount(reserveY ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(reserveY).WRITE()
	return inst
}

// GetReserveYAccount gets the "reserveY" account.
func (inst *Swap) GetReserveYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetUserTokenInAccount sets the "userTokenIn" account.
func (inst *Swap) SetUserTokenInAccount(userTokenIn ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(userTokenIn).WRITE()
	return inst
}

// GetUserTokenInAccount gets the "userTokenIn" account.
func (inst *Swap) GetUserTokenInAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetUserTokenOutAccount sets the "userTokenOut" account.
func (inst *Swap) SetUserTokenOutAccount(userTokenOut ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(userTokenOut).WRITE()
	return inst
}

// GetUserTokenOutAccount gets the "userTokenOut" account.
func (inst *Swap) GetUserTokenOutAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenXMintAccount sets the "tokenXMint" account.
func (inst *Swap) SetTokenXMintAccount(tokenXMint ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenXMint)
	return inst
}

// GetTokenXMintAccount gets the "tokenXMint" account.
func (inst *Swap) GetTokenXMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenYMintAccount sets the "tokenYMint" account.
func (inst *Swap) SetTokenYMintAccount(tokenYMint ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenYMint)
	return inst
}

// GetTokenYMintAccount gets the "tokenYMint" account.
func (inst *Swap) GetTokenYMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetOracleAccount sets the "oracle" account.
func (inst *Swap) SetOracleAccount(oracle ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(oracle).WRITE()
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *Swap) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetHostFeeInAccount sets the "hostFeeIn" account.
func (inst *Swap) SetHostFeeInAccount(hostFeeIn ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(hostFeeIn).WRITE()
	return inst
}

// GetHostFeeInAccount gets the "hostFeeIn" account.
func (inst *Swap) GetHostFeeInAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetUserAccount sets the "user" account.
func (inst *Swap) SetUserAccount(user ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(user).SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *Swap) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenXProgramAccount sets the "tokenXProgram" account.
func (inst *Swap) SetTokenXProgramAccount(tokenXProgram ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenXProgram)
	return inst
}

// GetTokenXProgramAccount gets the "tokenXProgram" account.
func (inst *Swap) GetTokenXProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetTokenYProgramAccount sets the "tokenYProgram" account.
func (inst *Swap) SetTokenYProgramAccount(tokenYProgram ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tokenYProgram)
	return inst
}

// GetTokenYProgramAccount gets the "tokenYProgram" account.
func (inst *Swap) GetTokenYProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *Swap) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *Swap) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetProgramAccount sets the "program" account.
func (inst *Swap) SetProgramAccount(program ag_solanago.PublicKey) *Swap {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *Swap) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

func (inst Swap) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Swap,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Swap) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Swap) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.AmountIn == nil {
			return errors.New("AmountIn parameter is not set")
		}
		if inst.MinAmountOut == nil {
			return errors.New("MinAmountOut parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.LbPair is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.BinArrayBitmapExtension is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ReserveX is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ReserveY is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.UserTokenIn is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.UserTokenOut is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenXMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenYMint is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Oracle is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.HostFeeIn is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenXProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TokenYProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *Swap) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Swap")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("    AmountIn", *inst.AmountIn))
						paramsBranch.Child(ag_format.Param("MinAmountOut", *inst.MinAmountOut))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=15]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 lbPair", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("binArrayBitmapExtension", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               reserveX", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("               reserveY", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("            userTokenIn", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("           userTokenOut", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("             tokenXMint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("             tokenYMint", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                 oracle", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("              hostFeeIn", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                   user", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("          tokenXProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("          tokenYProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("         eventAuthority", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("                program", inst.AccountMetaSlice.Get(14)))
					})
				})
		})
}

func (obj Swap) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AmountIn` param:
	err = encoder.Encode(obj.AmountIn)
	if err != nil {
		return err
	}
	// Serialize `MinAmountOut` param:
	err = encoder.Encode(obj.MinAmountOut)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Swap) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AmountIn`:
	err = decoder.Decode(&obj.AmountIn)
	if err != nil {
		return err
	}
	// Deserialize `MinAmountOut`:
	err = decoder.Decode(&obj.MinAmountOut)
	if err != nil {
		return err
	}
	return nil
}

// NewSwapInstruction declares a new Swap instruction with the provided parameters and accounts.
func NewSwapInstruction(
	// Parameters:
	amountIn uint64,
	minAmountOut uint64,
	// Accounts:
	lbPair ag_solanago.PublicKey,
	binArrayBitmapExtension ag_solanago.PublicKey,
	reserveX ag_solanago.PublicKey,
	reserveY ag_solanago.PublicKey,
	userTokenIn ag_solanago.PublicKey,
	userTokenOut ag_solanago.PublicKey,
	tokenXMint ag_solanago.PublicKey,
	tokenYMint ag_solanago.PublicKey,
	oracle ag_solanago.PublicKey,
	hostFeeIn ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	tokenXProgram ag_solanago.PublicKey,
	tokenYProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *Swap {
	return NewSwapInstructionBuilder().
		SetAmountIn(amountIn).
		SetMinAmountOut(minAmountOut).
		SetLbPairAccount(lbPair).
		SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension).
		SetReserveXAccount(reserveX).
		SetReserveYAccount(reserveY).
		SetUserTokenInAccount(userTokenIn).
		SetUserTokenOutAccount(userTokenOut).
		SetTokenXMintAccount(tokenXMint).
		SetTokenYMintAccount(tokenYMint).
		SetOracleAccount(oracle).
		SetHostFeeInAccount(hostFeeIn).
		SetUserAccount(user).
		SetTokenXProgramAccount(tokenXProgram).
		SetTokenYProgramAccount(tokenYProgram).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
