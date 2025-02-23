// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package stable_swap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ChangeSwapFee is the `change_swap_fee` instruction.
type ChangeSwapFee struct {
	NewSwapFee *uint64

	// [0] = [SIGNER] owner
	//
	// [1] = [WRITE] pool
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewChangeSwapFeeInstructionBuilder creates a new `ChangeSwapFee` instruction builder.
func NewChangeSwapFeeInstructionBuilder() *ChangeSwapFee {
	nd := &ChangeSwapFee{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetNewSwapFee sets the "new_swap_fee" parameter.
func (inst *ChangeSwapFee) SetNewSwapFee(new_swap_fee uint64) *ChangeSwapFee {
	inst.NewSwapFee = &new_swap_fee
	return inst
}

// SetOwnerAccount sets the "owner" account.
func (inst *ChangeSwapFee) SetOwnerAccount(owner ag_solanago.PublicKey) *ChangeSwapFee {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *ChangeSwapFee) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPoolAccount sets the "pool" account.
func (inst *ChangeSwapFee) SetPoolAccount(pool ag_solanago.PublicKey) *ChangeSwapFee {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(pool).WRITE()
	return inst
}

// GetPoolAccount gets the "pool" account.
func (inst *ChangeSwapFee) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst ChangeSwapFee) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ChangeSwapFee,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ChangeSwapFee) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ChangeSwapFee) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewSwapFee == nil {
			return errors.New("NewSwapFee parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Pool is not set")
		}
	}
	return nil
}

func (inst *ChangeSwapFee) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ChangeSwapFee")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("  NewSwapFee", *inst.NewSwapFee))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("owner", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" pool", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj ChangeSwapFee) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewSwapFee` param:
	err = encoder.Encode(obj.NewSwapFee)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ChangeSwapFee) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewSwapFee`:
	err = decoder.Decode(&obj.NewSwapFee)
	if err != nil {
		return err
	}
	return nil
}

// NewChangeSwapFeeInstruction declares a new ChangeSwapFee instruction with the provided parameters and accounts.
func NewChangeSwapFeeInstruction(
	// Parameters:
	new_swap_fee uint64,
	// Accounts:
	owner ag_solanago.PublicKey,
	pool ag_solanago.PublicKey) *ChangeSwapFee {
	return NewChangeSwapFeeInstructionBuilder().
		SetNewSwapFee(new_swap_fee).
		SetOwnerAccount(owner).
		SetPoolAccount(pool)
}
