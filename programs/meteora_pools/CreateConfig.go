// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_pools

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Create config
type CreateConfig struct {
	ConfigParameters *ConfigParameters

	// [0] = [WRITE] config
	//
	// [1] = [WRITE, SIGNER] admin
	//
	// [2] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateConfigInstructionBuilder creates a new `CreateConfig` instruction builder.
func NewCreateConfigInstructionBuilder() *CreateConfig {
	nd := &CreateConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetConfigParameters sets the "configParameters" parameter.
func (inst *CreateConfig) SetConfigParameters(configParameters ConfigParameters) *CreateConfig {
	inst.ConfigParameters = &configParameters
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *CreateConfig) SetConfigAccount(config ag_solanago.PublicKey) *CreateConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config).WRITE()
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *CreateConfig) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *CreateConfig) SetAdminAccount(admin ag_solanago.PublicKey) *CreateConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *CreateConfig) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *CreateConfig) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateConfig {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *CreateConfig) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst CreateConfig) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateConfig,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateConfig) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateConfig) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.ConfigParameters == nil {
			return errors.New("ConfigParameters parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *CreateConfig) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateConfig")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("ConfigParameters", *inst.ConfigParameters))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("systemProgram", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj CreateConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ConfigParameters` param:
	err = encoder.Encode(obj.ConfigParameters)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ConfigParameters`:
	err = decoder.Decode(&obj.ConfigParameters)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateConfigInstruction declares a new CreateConfig instruction with the provided parameters and accounts.
func NewCreateConfigInstruction(
	// Parameters:
	configParameters ConfigParameters,
	// Accounts:
	config ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *CreateConfig {
	return NewCreateConfigInstructionBuilder().
		SetConfigParameters(configParameters).
		SetConfigAccount(config).
		SetAdminAccount(admin).
		SetSystemProgramAccount(systemProgram)
}
