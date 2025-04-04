// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package pumpswap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateConfig is the `create_config` instruction.
type CreateConfig struct {
	LpFeeBasisPoints       *uint64
	ProtocolFeeBasisPoints *uint64
	ProtocolFeeRecipients  *[8]ag_solanago.PublicKey

	// [0] = [WRITE, SIGNER] admin
	//
	// [1] = [WRITE] global_config
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateConfigInstructionBuilder creates a new `CreateConfig` instruction builder.
func NewCreateConfigInstructionBuilder() *CreateConfig {
	nd := &CreateConfig{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetLpFeeBasisPoints sets the "lp_fee_basis_points" parameter.
func (inst *CreateConfig) SetLpFeeBasisPoints(lp_fee_basis_points uint64) *CreateConfig {
	inst.LpFeeBasisPoints = &lp_fee_basis_points
	return inst
}

// SetProtocolFeeBasisPoints sets the "protocol_fee_basis_points" parameter.
func (inst *CreateConfig) SetProtocolFeeBasisPoints(protocol_fee_basis_points uint64) *CreateConfig {
	inst.ProtocolFeeBasisPoints = &protocol_fee_basis_points
	return inst
}

// SetProtocolFeeRecipients sets the "protocol_fee_recipients" parameter.
func (inst *CreateConfig) SetProtocolFeeRecipients(protocol_fee_recipients [8]ag_solanago.PublicKey) *CreateConfig {
	inst.ProtocolFeeRecipients = &protocol_fee_recipients
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *CreateConfig) SetAdminAccount(admin ag_solanago.PublicKey) *CreateConfig {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *CreateConfig) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetGlobalConfigAccount sets the "global_config" account.
func (inst *CreateConfig) SetGlobalConfigAccount(globalConfig ag_solanago.PublicKey) *CreateConfig {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(globalConfig).WRITE()
	return inst
}

// GetGlobalConfigAccount gets the "global_config" account.
func (inst *CreateConfig) GetGlobalConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
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
		if inst.LpFeeBasisPoints == nil {
			return errors.New("LpFeeBasisPoints parameter is not set")
		}
		if inst.ProtocolFeeBasisPoints == nil {
			return errors.New("ProtocolFeeBasisPoints parameter is not set")
		}
		if inst.ProtocolFeeRecipients == nil {
			return errors.New("ProtocolFeeRecipients parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.GlobalConfig is not set")
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
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("         LpFeeBasisPoints", *inst.LpFeeBasisPoints))
						paramsBranch.Child(ag_format.Param("   ProtocolFeeBasisPoints", *inst.ProtocolFeeBasisPoints))
						paramsBranch.Child(ag_format.Param("    ProtocolFeeRecipients", *inst.ProtocolFeeRecipients))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        admin", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("global_config", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj CreateConfig) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LpFeeBasisPoints` param:
	err = encoder.Encode(obj.LpFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `ProtocolFeeBasisPoints` param:
	err = encoder.Encode(obj.ProtocolFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `ProtocolFeeRecipients` param:
	err = encoder.Encode(obj.ProtocolFeeRecipients)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateConfig) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LpFeeBasisPoints`:
	err = decoder.Decode(&obj.LpFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolFeeBasisPoints`:
	err = decoder.Decode(&obj.ProtocolFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `ProtocolFeeRecipients`:
	err = decoder.Decode(&obj.ProtocolFeeRecipients)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateConfigInstruction declares a new CreateConfig instruction with the provided parameters and accounts.
func NewCreateConfigInstruction(
	// Parameters:
	lp_fee_basis_points uint64,
	protocol_fee_basis_points uint64,
	protocol_fee_recipients [8]ag_solanago.PublicKey,
	// Accounts:
	admin ag_solanago.PublicKey,
	globalConfig ag_solanago.PublicKey) *CreateConfig {
	return NewCreateConfigInstructionBuilder().
		SetLpFeeBasisPoints(lp_fee_basis_points).
		SetProtocolFeeBasisPoints(protocol_fee_basis_points).
		SetProtocolFeeRecipients(protocol_fee_recipients).
		SetAdminAccount(admin).
		SetGlobalConfigAccount(globalConfig)
}
