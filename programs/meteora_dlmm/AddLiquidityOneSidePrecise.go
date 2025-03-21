// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AddLiquidityOneSidePrecise is the `addLiquidityOneSidePrecise` instruction.
type AddLiquidityOneSidePrecise struct {
	Parameter *AddLiquiditySingleSidePreciseParameter

	// [0] = [WRITE] position
	//
	// [1] = [WRITE] lbPair
	//
	// [2] = [WRITE] binArrayBitmapExtension
	//
	// [3] = [WRITE] userToken
	//
	// [4] = [WRITE] reserve
	//
	// [5] = [] tokenMint
	//
	// [6] = [WRITE] binArrayLower
	//
	// [7] = [WRITE] binArrayUpper
	//
	// [8] = [SIGNER] sender
	//
	// [9] = [] tokenProgram
	//
	// [10] = [] eventAuthority
	//
	// [11] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAddLiquidityOneSidePreciseInstructionBuilder creates a new `AddLiquidityOneSidePrecise` instruction builder.
func NewAddLiquidityOneSidePreciseInstructionBuilder() *AddLiquidityOneSidePrecise {
	nd := &AddLiquidityOneSidePrecise{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 12),
	}
	return nd
}

// SetParameter sets the "parameter" parameter.
func (inst *AddLiquidityOneSidePrecise) SetParameter(parameter AddLiquiditySingleSidePreciseParameter) *AddLiquidityOneSidePrecise {
	inst.Parameter = &parameter
	return inst
}

// SetPositionAccount sets the "position" account.
func (inst *AddLiquidityOneSidePrecise) SetPositionAccount(position ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(position).WRITE()
	return inst
}

// GetPositionAccount gets the "position" account.
func (inst *AddLiquidityOneSidePrecise) GetPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLbPairAccount sets the "lbPair" account.
func (inst *AddLiquidityOneSidePrecise) SetLbPairAccount(lbPair ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(lbPair).WRITE()
	return inst
}

// GetLbPairAccount gets the "lbPair" account.
func (inst *AddLiquidityOneSidePrecise) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetBinArrayBitmapExtensionAccount sets the "binArrayBitmapExtension" account.
func (inst *AddLiquidityOneSidePrecise) SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(binArrayBitmapExtension).WRITE()
	return inst
}

// GetBinArrayBitmapExtensionAccount gets the "binArrayBitmapExtension" account.
func (inst *AddLiquidityOneSidePrecise) GetBinArrayBitmapExtensionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetUserTokenAccount sets the "userToken" account.
func (inst *AddLiquidityOneSidePrecise) SetUserTokenAccount(userToken ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(userToken).WRITE()
	return inst
}

// GetUserTokenAccount gets the "userToken" account.
func (inst *AddLiquidityOneSidePrecise) GetUserTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReserveAccount sets the "reserve" account.
func (inst *AddLiquidityOneSidePrecise) SetReserveAccount(reserve ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(reserve).WRITE()
	return inst
}

// GetReserveAccount gets the "reserve" account.
func (inst *AddLiquidityOneSidePrecise) GetReserveAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenMintAccount sets the "tokenMint" account.
func (inst *AddLiquidityOneSidePrecise) SetTokenMintAccount(tokenMint ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenMint)
	return inst
}

// GetTokenMintAccount gets the "tokenMint" account.
func (inst *AddLiquidityOneSidePrecise) GetTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetBinArrayLowerAccount sets the "binArrayLower" account.
func (inst *AddLiquidityOneSidePrecise) SetBinArrayLowerAccount(binArrayLower ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(binArrayLower).WRITE()
	return inst
}

// GetBinArrayLowerAccount gets the "binArrayLower" account.
func (inst *AddLiquidityOneSidePrecise) GetBinArrayLowerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetBinArrayUpperAccount sets the "binArrayUpper" account.
func (inst *AddLiquidityOneSidePrecise) SetBinArrayUpperAccount(binArrayUpper ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(binArrayUpper).WRITE()
	return inst
}

// GetBinArrayUpperAccount gets the "binArrayUpper" account.
func (inst *AddLiquidityOneSidePrecise) GetBinArrayUpperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSenderAccount sets the "sender" account.
func (inst *AddLiquidityOneSidePrecise) SetSenderAccount(sender ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(sender).SIGNER()
	return inst
}

// GetSenderAccount gets the "sender" account.
func (inst *AddLiquidityOneSidePrecise) GetSenderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *AddLiquidityOneSidePrecise) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *AddLiquidityOneSidePrecise) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *AddLiquidityOneSidePrecise) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *AddLiquidityOneSidePrecise) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetProgramAccount sets the "program" account.
func (inst *AddLiquidityOneSidePrecise) SetProgramAccount(program ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *AddLiquidityOneSidePrecise) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

func (inst AddLiquidityOneSidePrecise) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AddLiquidityOneSidePrecise,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddLiquidityOneSidePrecise) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddLiquidityOneSidePrecise) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Parameter == nil {
			return errors.New("Parameter parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Position is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.LbPair is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.BinArrayBitmapExtension is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.UserToken is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Reserve is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenMint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.BinArrayLower is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.BinArrayUpper is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Sender is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *AddLiquidityOneSidePrecise) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddLiquidityOneSidePrecise")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Parameter", *inst.Parameter))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=12]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               position", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                 lbPair", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("binArrayBitmapExtension", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("              userToken", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                reserve", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("              tokenMint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("          binArrayLower", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("          binArrayUpper", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                 sender", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("           tokenProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("         eventAuthority", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                program", inst.AccountMetaSlice.Get(11)))
					})
				})
		})
}

func (obj AddLiquidityOneSidePrecise) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Parameter` param:
	err = encoder.Encode(obj.Parameter)
	if err != nil {
		return err
	}
	return nil
}
func (obj *AddLiquidityOneSidePrecise) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Parameter`:
	err = decoder.Decode(&obj.Parameter)
	if err != nil {
		return err
	}
	return nil
}

// NewAddLiquidityOneSidePreciseInstruction declares a new AddLiquidityOneSidePrecise instruction with the provided parameters and accounts.
func NewAddLiquidityOneSidePreciseInstruction(
	// Parameters:
	parameter AddLiquiditySingleSidePreciseParameter,
	// Accounts:
	position ag_solanago.PublicKey,
	lbPair ag_solanago.PublicKey,
	binArrayBitmapExtension ag_solanago.PublicKey,
	userToken ag_solanago.PublicKey,
	reserve ag_solanago.PublicKey,
	tokenMint ag_solanago.PublicKey,
	binArrayLower ag_solanago.PublicKey,
	binArrayUpper ag_solanago.PublicKey,
	sender ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *AddLiquidityOneSidePrecise {
	return NewAddLiquidityOneSidePreciseInstructionBuilder().
		SetParameter(parameter).
		SetPositionAccount(position).
		SetLbPairAccount(lbPair).
		SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension).
		SetUserTokenAccount(userToken).
		SetReserveAccount(reserve).
		SetTokenMintAccount(tokenMint).
		SetBinArrayLowerAccount(binArrayLower).
		SetBinArrayUpperAccount(binArrayUpper).
		SetSenderAccount(sender).
		SetTokenProgramAccount(tokenProgram).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
