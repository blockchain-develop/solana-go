// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateRewardFunder is the `updateRewardFunder` instruction.
type UpdateRewardFunder struct {
	RewardIndex *uint64
	NewFunder   *ag_solanago.PublicKey

	// [0] = [WRITE] lbPair
	//
	// [1] = [SIGNER] admin
	//
	// [2] = [] eventAuthority
	//
	// [3] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateRewardFunderInstructionBuilder creates a new `UpdateRewardFunder` instruction builder.
func NewUpdateRewardFunderInstructionBuilder() *UpdateRewardFunder {
	nd := &UpdateRewardFunder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetRewardIndex sets the "rewardIndex" parameter.
func (inst *UpdateRewardFunder) SetRewardIndex(rewardIndex uint64) *UpdateRewardFunder {
	inst.RewardIndex = &rewardIndex
	return inst
}

// SetNewFunder sets the "newFunder" parameter.
func (inst *UpdateRewardFunder) SetNewFunder(newFunder ag_solanago.PublicKey) *UpdateRewardFunder {
	inst.NewFunder = &newFunder
	return inst
}

// SetLbPairAccount sets the "lbPair" account.
func (inst *UpdateRewardFunder) SetLbPairAccount(lbPair ag_solanago.PublicKey) *UpdateRewardFunder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(lbPair).WRITE()
	return inst
}

// GetLbPairAccount gets the "lbPair" account.
func (inst *UpdateRewardFunder) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdateRewardFunder) SetAdminAccount(admin ag_solanago.PublicKey) *UpdateRewardFunder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdateRewardFunder) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *UpdateRewardFunder) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *UpdateRewardFunder {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *UpdateRewardFunder) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetProgramAccount sets the "program" account.
func (inst *UpdateRewardFunder) SetProgramAccount(program ag_solanago.PublicKey) *UpdateRewardFunder {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *UpdateRewardFunder) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst UpdateRewardFunder) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateRewardFunder,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateRewardFunder) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateRewardFunder) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RewardIndex == nil {
			return errors.New("RewardIndex parameter is not set")
		}
		if inst.NewFunder == nil {
			return errors.New("NewFunder parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.LbPair is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *UpdateRewardFunder) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateRewardFunder")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("RewardIndex", *inst.RewardIndex))
						paramsBranch.Child(ag_format.Param("  NewFunder", *inst.NewFunder))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        lbPair", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("eventAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       program", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj UpdateRewardFunder) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RewardIndex` param:
	err = encoder.Encode(obj.RewardIndex)
	if err != nil {
		return err
	}
	// Serialize `NewFunder` param:
	err = encoder.Encode(obj.NewFunder)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateRewardFunder) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RewardIndex`:
	err = decoder.Decode(&obj.RewardIndex)
	if err != nil {
		return err
	}
	// Deserialize `NewFunder`:
	err = decoder.Decode(&obj.NewFunder)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateRewardFunderInstruction declares a new UpdateRewardFunder instruction with the provided parameters and accounts.
func NewUpdateRewardFunderInstruction(
	// Parameters:
	rewardIndex uint64,
	newFunder ag_solanago.PublicKey,
	// Accounts:
	lbPair ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *UpdateRewardFunder {
	return NewUpdateRewardFunderInstructionBuilder().
		SetRewardIndex(rewardIndex).
		SetNewFunder(newFunder).
		SetLbPairAccount(lbPair).
		SetAdminAccount(admin).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
