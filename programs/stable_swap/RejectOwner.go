// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package stable_swap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RejectOwner is the `reject_owner` instruction.
type RejectOwner struct {

	// [0] = [SIGNER] pending_owner
	//
	// [1] = [WRITE] pool
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRejectOwnerInstructionBuilder creates a new `RejectOwner` instruction builder.
func NewRejectOwnerInstructionBuilder() *RejectOwner {
	nd := &RejectOwner{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetPendingOwnerAccount sets the "pending_owner" account.
func (inst *RejectOwner) SetPendingOwnerAccount(pendingOwner ag_solanago.PublicKey) *RejectOwner {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(pendingOwner).SIGNER()
	return inst
}

// GetPendingOwnerAccount gets the "pending_owner" account.
func (inst *RejectOwner) GetPendingOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPoolAccount sets the "pool" account.
func (inst *RejectOwner) SetPoolAccount(pool ag_solanago.PublicKey) *RejectOwner {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(pool).WRITE()
	return inst
}

// GetPoolAccount gets the "pool" account.
func (inst *RejectOwner) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst RejectOwner) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RejectOwner,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RejectOwner) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RejectOwner) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.PendingOwner is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Pool is not set")
		}
	}
	return nil
}

func (inst *RejectOwner) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RejectOwner")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("pending_owner", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         pool", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj RejectOwner) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RejectOwner) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRejectOwnerInstruction declares a new RejectOwner instruction with the provided parameters and accounts.
func NewRejectOwnerInstruction(
	// Accounts:
	pendingOwner ag_solanago.PublicKey,
	pool ag_solanago.PublicKey) *RejectOwner {
	return NewRejectOwnerInstructionBuilder().
		SetPendingOwnerAccount(pendingOwner).
		SetPoolAccount(pool)
}
