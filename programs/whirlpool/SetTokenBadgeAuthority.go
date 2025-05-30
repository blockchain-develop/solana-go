// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package whirlpool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetTokenBadgeAuthority is the `setTokenBadgeAuthority` instruction.
type SetTokenBadgeAuthority struct {

	// [0] = [] whirlpoolsConfig
	//
	// [1] = [WRITE] whirlpoolsConfigExtension
	//
	// [2] = [SIGNER] configExtensionAuthority
	//
	// [3] = [] newTokenBadgeAuthority
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetTokenBadgeAuthorityInstructionBuilder creates a new `SetTokenBadgeAuthority` instruction builder.
func NewSetTokenBadgeAuthorityInstructionBuilder() *SetTokenBadgeAuthority {
	nd := &SetTokenBadgeAuthority{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetWhirlpoolsConfigAccount sets the "whirlpoolsConfig" account.
func (inst *SetTokenBadgeAuthority) SetWhirlpoolsConfigAccount(whirlpoolsConfig ag_solanago.PublicKey) *SetTokenBadgeAuthority {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(whirlpoolsConfig)
	return inst
}

// GetWhirlpoolsConfigAccount gets the "whirlpoolsConfig" account.
func (inst *SetTokenBadgeAuthority) GetWhirlpoolsConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetWhirlpoolsConfigExtensionAccount sets the "whirlpoolsConfigExtension" account.
func (inst *SetTokenBadgeAuthority) SetWhirlpoolsConfigExtensionAccount(whirlpoolsConfigExtension ag_solanago.PublicKey) *SetTokenBadgeAuthority {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(whirlpoolsConfigExtension).WRITE()
	return inst
}

// GetWhirlpoolsConfigExtensionAccount gets the "whirlpoolsConfigExtension" account.
func (inst *SetTokenBadgeAuthority) GetWhirlpoolsConfigExtensionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetConfigExtensionAuthorityAccount sets the "configExtensionAuthority" account.
func (inst *SetTokenBadgeAuthority) SetConfigExtensionAuthorityAccount(configExtensionAuthority ag_solanago.PublicKey) *SetTokenBadgeAuthority {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(configExtensionAuthority).SIGNER()
	return inst
}

// GetConfigExtensionAuthorityAccount gets the "configExtensionAuthority" account.
func (inst *SetTokenBadgeAuthority) GetConfigExtensionAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetNewTokenBadgeAuthorityAccount sets the "newTokenBadgeAuthority" account.
func (inst *SetTokenBadgeAuthority) SetNewTokenBadgeAuthorityAccount(newTokenBadgeAuthority ag_solanago.PublicKey) *SetTokenBadgeAuthority {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(newTokenBadgeAuthority)
	return inst
}

// GetNewTokenBadgeAuthorityAccount gets the "newTokenBadgeAuthority" account.
func (inst *SetTokenBadgeAuthority) GetNewTokenBadgeAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst SetTokenBadgeAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetTokenBadgeAuthority,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetTokenBadgeAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetTokenBadgeAuthority) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.WhirlpoolsConfig is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.WhirlpoolsConfigExtension is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ConfigExtensionAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.NewTokenBadgeAuthority is not set")
		}
	}
	return nil
}

func (inst *SetTokenBadgeAuthority) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetTokenBadgeAuthority")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         whirlpoolsConfig", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("whirlpoolsConfigExtension", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" configExtensionAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("   newTokenBadgeAuthority", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj SetTokenBadgeAuthority) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SetTokenBadgeAuthority) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSetTokenBadgeAuthorityInstruction declares a new SetTokenBadgeAuthority instruction with the provided parameters and accounts.
func NewSetTokenBadgeAuthorityInstruction(
	// Accounts:
	whirlpoolsConfig ag_solanago.PublicKey,
	whirlpoolsConfigExtension ag_solanago.PublicKey,
	configExtensionAuthority ag_solanago.PublicKey,
	newTokenBadgeAuthority ag_solanago.PublicKey) *SetTokenBadgeAuthority {
	return NewSetTokenBadgeAuthorityInstructionBuilder().
		SetWhirlpoolsConfigAccount(whirlpoolsConfig).
		SetWhirlpoolsConfigExtensionAccount(whirlpoolsConfigExtension).
		SetConfigExtensionAuthorityAccount(configExtensionAuthority).
		SetNewTokenBadgeAuthorityAccount(newTokenBadgeAuthority)
}
