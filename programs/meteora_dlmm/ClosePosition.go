// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ClosePosition is the `closePosition` instruction.
type ClosePosition struct {

	// [0] = [WRITE] position
	//
	// [1] = [WRITE] lbPair
	//
	// [2] = [WRITE] binArrayLower
	//
	// [3] = [WRITE] binArrayUpper
	//
	// [4] = [SIGNER] sender
	//
	// [5] = [WRITE] rentReceiver
	//
	// [6] = [] eventAuthority
	//
	// [7] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewClosePositionInstructionBuilder creates a new `ClosePosition` instruction builder.
func NewClosePositionInstructionBuilder() *ClosePosition {
	nd := &ClosePosition{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetPositionAccount sets the "position" account.
func (inst *ClosePosition) SetPositionAccount(position ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(position).WRITE()
	return inst
}

// GetPositionAccount gets the "position" account.
func (inst *ClosePosition) GetPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLbPairAccount sets the "lbPair" account.
func (inst *ClosePosition) SetLbPairAccount(lbPair ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(lbPair).WRITE()
	return inst
}

// GetLbPairAccount gets the "lbPair" account.
func (inst *ClosePosition) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetBinArrayLowerAccount sets the "binArrayLower" account.
func (inst *ClosePosition) SetBinArrayLowerAccount(binArrayLower ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(binArrayLower).WRITE()
	return inst
}

// GetBinArrayLowerAccount gets the "binArrayLower" account.
func (inst *ClosePosition) GetBinArrayLowerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetBinArrayUpperAccount sets the "binArrayUpper" account.
func (inst *ClosePosition) SetBinArrayUpperAccount(binArrayUpper ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(binArrayUpper).WRITE()
	return inst
}

// GetBinArrayUpperAccount gets the "binArrayUpper" account.
func (inst *ClosePosition) GetBinArrayUpperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSenderAccount sets the "sender" account.
func (inst *ClosePosition) SetSenderAccount(sender ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(sender).SIGNER()
	return inst
}

// GetSenderAccount gets the "sender" account.
func (inst *ClosePosition) GetSenderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetRentReceiverAccount sets the "rentReceiver" account.
func (inst *ClosePosition) SetRentReceiverAccount(rentReceiver ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(rentReceiver).WRITE()
	return inst
}

// GetRentReceiverAccount gets the "rentReceiver" account.
func (inst *ClosePosition) GetRentReceiverAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *ClosePosition) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *ClosePosition) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetProgramAccount sets the "program" account.
func (inst *ClosePosition) SetProgramAccount(program ag_solanago.PublicKey) *ClosePosition {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *ClosePosition) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst ClosePosition) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ClosePosition,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClosePosition) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClosePosition) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Position is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.LbPair is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.BinArrayLower is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.BinArrayUpper is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Sender is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.RentReceiver is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *ClosePosition) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ClosePosition")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      position", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        lbPair", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" binArrayLower", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta(" binArrayUpper", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        sender", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("  rentReceiver", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("eventAuthority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       program", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj ClosePosition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ClosePosition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewClosePositionInstruction declares a new ClosePosition instruction with the provided parameters and accounts.
func NewClosePositionInstruction(
	// Accounts:
	position ag_solanago.PublicKey,
	lbPair ag_solanago.PublicKey,
	binArrayLower ag_solanago.PublicKey,
	binArrayUpper ag_solanago.PublicKey,
	sender ag_solanago.PublicKey,
	rentReceiver ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *ClosePosition {
	return NewClosePositionInstructionBuilder().
		SetPositionAccount(position).
		SetLbPairAccount(lbPair).
		SetBinArrayLowerAccount(binArrayLower).
		SetBinArrayUpperAccount(binArrayUpper).
		SetSenderAccount(sender).
		SetRentReceiverAccount(rentReceiver).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
