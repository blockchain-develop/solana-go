// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package phoenix_v1

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RequestSeat is the `RequestSeat` instruction.
type RequestSeat struct {

	// [0] = [] phoenixProgram
	//
	// [1] = [] logAuthority
	//
	// [2] = [WRITE] market
	//
	// [3] = [WRITE, SIGNER] payer
	//
	// [4] = [WRITE] seat
	//
	// [5] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRequestSeatInstructionBuilder creates a new `RequestSeat` instruction builder.
func NewRequestSeatInstructionBuilder() *RequestSeat {
	nd := &RequestSeat{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetPhoenixProgramAccount sets the "phoenixProgram" account.
func (inst *RequestSeat) SetPhoenixProgramAccount(phoenixProgram ag_solanago.PublicKey) *RequestSeat {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(phoenixProgram)
	return inst
}

// GetPhoenixProgramAccount gets the "phoenixProgram" account.
func (inst *RequestSeat) GetPhoenixProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLogAuthorityAccount sets the "logAuthority" account.
func (inst *RequestSeat) SetLogAuthorityAccount(logAuthority ag_solanago.PublicKey) *RequestSeat {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(logAuthority)
	return inst
}

// GetLogAuthorityAccount gets the "logAuthority" account.
func (inst *RequestSeat) GetLogAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMarketAccount sets the "market" account.
func (inst *RequestSeat) SetMarketAccount(market ag_solanago.PublicKey) *RequestSeat {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *RequestSeat) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayerAccount sets the "payer" account.
func (inst *RequestSeat) SetPayerAccount(payer ag_solanago.PublicKey) *RequestSeat {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *RequestSeat) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSeatAccount sets the "seat" account.
func (inst *RequestSeat) SetSeatAccount(seat ag_solanago.PublicKey) *RequestSeat {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(seat).WRITE()
	return inst
}

// GetSeatAccount gets the "seat" account.
func (inst *RequestSeat) GetSeatAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *RequestSeat) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *RequestSeat {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *RequestSeat) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst RequestSeat) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_RequestSeat),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RequestSeat) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RequestSeat) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.PhoenixProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.LogAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Seat is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *RequestSeat) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RequestSeat")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("phoenixProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  logAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        market", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         payer", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("          seat", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta(" systemProgram", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj RequestSeat) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RequestSeat) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRequestSeatInstruction declares a new RequestSeat instruction with the provided parameters and accounts.
func NewRequestSeatInstruction(
	// Accounts:
	phoenixProgram ag_solanago.PublicKey,
	logAuthority ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	seat ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *RequestSeat {
	return NewRequestSeatInstructionBuilder().
		SetPhoenixProgramAccount(phoenixProgram).
		SetLogAuthorityAccount(logAuthority).
		SetMarketAccount(market).
		SetPayerAccount(payer).
		SetSeatAccount(seat).
		SetSystemProgramAccount(systemProgram)
}
