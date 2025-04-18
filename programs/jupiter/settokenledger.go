// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupiter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetTokenLedger is the `set_token_ledger` instruction.
type SetTokenLedger struct {

	// [0] = [WRITE] token_ledger
	//
	// [1] = [] token_account
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetTokenLedgerInstructionBuilder creates a new `SetTokenLedger` instruction builder.
func NewSetTokenLedgerInstructionBuilder() *SetTokenLedger {
	nd := &SetTokenLedger{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetTokenLedgerAccount sets the "token_ledger" account.
func (inst *SetTokenLedger) SetTokenLedgerAccount(tokenLedger ag_solanago.PublicKey) *SetTokenLedger {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tokenLedger).WRITE()
	return inst
}

// GetTokenLedgerAccount gets the "token_ledger" account.
func (inst *SetTokenLedger) GetTokenLedgerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenAccountAccount sets the "token_account" account.
func (inst *SetTokenLedger) SetTokenAccountAccount(tokenAccount ag_solanago.PublicKey) *SetTokenLedger {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenAccount)
	return inst
}

// GetTokenAccountAccount gets the "token_account" account.
func (inst *SetTokenLedger) GetTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst SetTokenLedger) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetTokenLedger,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetTokenLedger) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetTokenLedger) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TokenLedger is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TokenAccount is not set")
		}
	}
	return nil
}

func (inst *SetTokenLedger) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetTokenLedger")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("token_ledger", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      token_", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj SetTokenLedger) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SetTokenLedger) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSetTokenLedgerInstruction declares a new SetTokenLedger instruction with the provided parameters and accounts.
func NewSetTokenLedgerInstruction(
	// Accounts:
	tokenLedger ag_solanago.PublicKey,
	tokenAccount ag_solanago.PublicKey) *SetTokenLedger {
	return NewSetTokenLedgerInstructionBuilder().
		SetTokenLedgerAccount(tokenLedger).
		SetTokenAccountAccount(tokenAccount)
}
