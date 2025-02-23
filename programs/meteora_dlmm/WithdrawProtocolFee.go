// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// WithdrawProtocolFee is the `withdrawProtocolFee` instruction.
type WithdrawProtocolFee struct {
	AmountX *uint64
	AmountY *uint64

	// [0] = [WRITE] lbPair
	//
	// [1] = [WRITE] reserveX
	//
	// [2] = [WRITE] reserveY
	//
	// [3] = [] tokenXMint
	//
	// [4] = [] tokenYMint
	//
	// [5] = [WRITE] receiverTokenX
	//
	// [6] = [WRITE] receiverTokenY
	//
	// [7] = [] tokenXProgram
	//
	// [8] = [] tokenYProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewWithdrawProtocolFeeInstructionBuilder creates a new `WithdrawProtocolFee` instruction builder.
func NewWithdrawProtocolFeeInstructionBuilder() *WithdrawProtocolFee {
	nd := &WithdrawProtocolFee{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetAmountX sets the "amountX" parameter.
func (inst *WithdrawProtocolFee) SetAmountX(amountX uint64) *WithdrawProtocolFee {
	inst.AmountX = &amountX
	return inst
}

// SetAmountY sets the "amountY" parameter.
func (inst *WithdrawProtocolFee) SetAmountY(amountY uint64) *WithdrawProtocolFee {
	inst.AmountY = &amountY
	return inst
}

// SetLbPairAccount sets the "lbPair" account.
func (inst *WithdrawProtocolFee) SetLbPairAccount(lbPair ag_solanago.PublicKey) *WithdrawProtocolFee {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(lbPair).WRITE()
	return inst
}

// GetLbPairAccount gets the "lbPair" account.
func (inst *WithdrawProtocolFee) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetReserveXAccount sets the "reserveX" account.
func (inst *WithdrawProtocolFee) SetReserveXAccount(reserveX ag_solanago.PublicKey) *WithdrawProtocolFee {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(reserveX).WRITE()
	return inst
}

// GetReserveXAccount gets the "reserveX" account.
func (inst *WithdrawProtocolFee) GetReserveXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetReserveYAccount sets the "reserveY" account.
func (inst *WithdrawProtocolFee) SetReserveYAccount(reserveY ag_solanago.PublicKey) *WithdrawProtocolFee {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(reserveY).WRITE()
	return inst
}

// GetReserveYAccount gets the "reserveY" account.
func (inst *WithdrawProtocolFee) GetReserveYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenXMintAccount sets the "tokenXMint" account.
func (inst *WithdrawProtocolFee) SetTokenXMintAccount(tokenXMint ag_solanago.PublicKey) *WithdrawProtocolFee {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tokenXMint)
	return inst
}

// GetTokenXMintAccount gets the "tokenXMint" account.
func (inst *WithdrawProtocolFee) GetTokenXMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenYMintAccount sets the "tokenYMint" account.
func (inst *WithdrawProtocolFee) SetTokenYMintAccount(tokenYMint ag_solanago.PublicKey) *WithdrawProtocolFee {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenYMint)
	return inst
}

// GetTokenYMintAccount gets the "tokenYMint" account.
func (inst *WithdrawProtocolFee) GetTokenYMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetReceiverTokenXAccount sets the "receiverTokenX" account.
func (inst *WithdrawProtocolFee) SetReceiverTokenXAccount(receiverTokenX ag_solanago.PublicKey) *WithdrawProtocolFee {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(receiverTokenX).WRITE()
	return inst
}

// GetReceiverTokenXAccount gets the "receiverTokenX" account.
func (inst *WithdrawProtocolFee) GetReceiverTokenXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetReceiverTokenYAccount sets the "receiverTokenY" account.
func (inst *WithdrawProtocolFee) SetReceiverTokenYAccount(receiverTokenY ag_solanago.PublicKey) *WithdrawProtocolFee {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(receiverTokenY).WRITE()
	return inst
}

// GetReceiverTokenYAccount gets the "receiverTokenY" account.
func (inst *WithdrawProtocolFee) GetReceiverTokenYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenXProgramAccount sets the "tokenXProgram" account.
func (inst *WithdrawProtocolFee) SetTokenXProgramAccount(tokenXProgram ag_solanago.PublicKey) *WithdrawProtocolFee {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenXProgram)
	return inst
}

// GetTokenXProgramAccount gets the "tokenXProgram" account.
func (inst *WithdrawProtocolFee) GetTokenXProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenYProgramAccount sets the "tokenYProgram" account.
func (inst *WithdrawProtocolFee) SetTokenYProgramAccount(tokenYProgram ag_solanago.PublicKey) *WithdrawProtocolFee {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenYProgram)
	return inst
}

// GetTokenYProgramAccount gets the "tokenYProgram" account.
func (inst *WithdrawProtocolFee) GetTokenYProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst WithdrawProtocolFee) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_WithdrawProtocolFee,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WithdrawProtocolFee) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawProtocolFee) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.AmountX == nil {
			return errors.New("AmountX parameter is not set")
		}
		if inst.AmountY == nil {
			return errors.New("AmountY parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.LbPair is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ReserveX is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ReserveY is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TokenXMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenYMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ReceiverTokenX is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ReceiverTokenY is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenXProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenYProgram is not set")
		}
	}
	return nil
}

func (inst *WithdrawProtocolFee) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawProtocolFee")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("AmountX", *inst.AmountX))
						paramsBranch.Child(ag_format.Param("AmountY", *inst.AmountY))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        lbPair", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      reserveX", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      reserveY", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    tokenXMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("    tokenYMint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("receiverTokenX", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("receiverTokenY", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta(" tokenXProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta(" tokenYProgram", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj WithdrawProtocolFee) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AmountX` param:
	err = encoder.Encode(obj.AmountX)
	if err != nil {
		return err
	}
	// Serialize `AmountY` param:
	err = encoder.Encode(obj.AmountY)
	if err != nil {
		return err
	}
	return nil
}
func (obj *WithdrawProtocolFee) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AmountX`:
	err = decoder.Decode(&obj.AmountX)
	if err != nil {
		return err
	}
	// Deserialize `AmountY`:
	err = decoder.Decode(&obj.AmountY)
	if err != nil {
		return err
	}
	return nil
}

// NewWithdrawProtocolFeeInstruction declares a new WithdrawProtocolFee instruction with the provided parameters and accounts.
func NewWithdrawProtocolFeeInstruction(
	// Parameters:
	amountX uint64,
	amountY uint64,
	// Accounts:
	lbPair ag_solanago.PublicKey,
	reserveX ag_solanago.PublicKey,
	reserveY ag_solanago.PublicKey,
	tokenXMint ag_solanago.PublicKey,
	tokenYMint ag_solanago.PublicKey,
	receiverTokenX ag_solanago.PublicKey,
	receiverTokenY ag_solanago.PublicKey,
	tokenXProgram ag_solanago.PublicKey,
	tokenYProgram ag_solanago.PublicKey) *WithdrawProtocolFee {
	return NewWithdrawProtocolFeeInstructionBuilder().
		SetAmountX(amountX).
		SetAmountY(amountY).
		SetLbPairAccount(lbPair).
		SetReserveXAccount(reserveX).
		SetReserveYAccount(reserveY).
		SetTokenXMintAccount(tokenXMint).
		SetTokenYMintAccount(tokenYMint).
		SetReceiverTokenXAccount(receiverTokenX).
		SetReceiverTokenYAccount(receiverTokenY).
		SetTokenXProgramAccount(tokenXProgram).
		SetTokenYProgramAccount(tokenYProgram)
}
