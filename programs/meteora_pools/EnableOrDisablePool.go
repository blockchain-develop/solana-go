// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_pools

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Enable or disable a pool. A disabled pool allow only remove balanced liquidity operation.
type EnableOrDisablePool struct {
	Enable *bool

	// [0] = [WRITE] pool
	// ··········· Pool account (PDA)
	//
	// [1] = [SIGNER] admin
	// ··········· Admin account. Must be owner of the pool.
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewEnableOrDisablePoolInstructionBuilder creates a new `EnableOrDisablePool` instruction builder.
func NewEnableOrDisablePoolInstructionBuilder() *EnableOrDisablePool {
	nd := &EnableOrDisablePool{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetEnable sets the "enable" parameter.
func (inst *EnableOrDisablePool) SetEnable(enable bool) *EnableOrDisablePool {
	inst.Enable = &enable
	return inst
}

// SetPoolAccount sets the "pool" account.
// Pool account (PDA)
func (inst *EnableOrDisablePool) SetPoolAccount(pool ag_solanago.PublicKey) *EnableOrDisablePool {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

// GetPoolAccount gets the "pool" account.
// Pool account (PDA)
func (inst *EnableOrDisablePool) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
// Admin account. Must be owner of the pool.
func (inst *EnableOrDisablePool) SetAdminAccount(admin ag_solanago.PublicKey) *EnableOrDisablePool {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
// Admin account. Must be owner of the pool.
func (inst *EnableOrDisablePool) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst EnableOrDisablePool) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_EnableOrDisablePool,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst EnableOrDisablePool) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *EnableOrDisablePool) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Enable == nil {
			return errors.New("Enable parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Pool is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *EnableOrDisablePool) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("EnableOrDisablePool")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Enable", *inst.Enable))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta(" pool", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("admin", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj EnableOrDisablePool) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Enable` param:
	err = encoder.Encode(obj.Enable)
	if err != nil {
		return err
	}
	return nil
}
func (obj *EnableOrDisablePool) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Enable`:
	err = decoder.Decode(&obj.Enable)
	if err != nil {
		return err
	}
	return nil
}

// NewEnableOrDisablePoolInstruction declares a new EnableOrDisablePool instruction with the provided parameters and accounts.
func NewEnableOrDisablePoolInstruction(
	// Parameters:
	enable bool,
	// Accounts:
	pool ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *EnableOrDisablePool {
	return NewEnableOrDisablePoolInstructionBuilder().
		SetEnable(enable).
		SetPoolAccount(pool).
		SetAdminAccount(admin)
}
