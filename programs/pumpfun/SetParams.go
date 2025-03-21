// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package pumpfun

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Sets the global state parameters.
type SetParams struct {
	FeeRecipient                *ag_solanago.PublicKey
	InitialVirtualTokenReserves *uint64
	InitialVirtualSolReserves   *uint64
	InitialRealTokenReserves    *uint64
	TokenTotalSupply            *uint64
	FeeBasisPoints              *uint64

	// [0] = [WRITE] global
	//
	// [1] = [WRITE, SIGNER] user
	//
	// [2] = [] systemProgram
	//
	// [3] = [] eventAuthority
	//
	// [4] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetParamsInstructionBuilder creates a new `SetParams` instruction builder.
func NewSetParamsInstructionBuilder() *SetParams {
	nd := &SetParams{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetFeeRecipient sets the "feeRecipient" parameter.
func (inst *SetParams) SetFeeRecipient(feeRecipient ag_solanago.PublicKey) *SetParams {
	inst.FeeRecipient = &feeRecipient
	return inst
}

// SetInitialVirtualTokenReserves sets the "initialVirtualTokenReserves" parameter.
func (inst *SetParams) SetInitialVirtualTokenReserves(initialVirtualTokenReserves uint64) *SetParams {
	inst.InitialVirtualTokenReserves = &initialVirtualTokenReserves
	return inst
}

// SetInitialVirtualSolReserves sets the "initialVirtualSolReserves" parameter.
func (inst *SetParams) SetInitialVirtualSolReserves(initialVirtualSolReserves uint64) *SetParams {
	inst.InitialVirtualSolReserves = &initialVirtualSolReserves
	return inst
}

// SetInitialRealTokenReserves sets the "initialRealTokenReserves" parameter.
func (inst *SetParams) SetInitialRealTokenReserves(initialRealTokenReserves uint64) *SetParams {
	inst.InitialRealTokenReserves = &initialRealTokenReserves
	return inst
}

// SetTokenTotalSupply sets the "tokenTotalSupply" parameter.
func (inst *SetParams) SetTokenTotalSupply(tokenTotalSupply uint64) *SetParams {
	inst.TokenTotalSupply = &tokenTotalSupply
	return inst
}

// SetFeeBasisPoints sets the "feeBasisPoints" parameter.
func (inst *SetParams) SetFeeBasisPoints(feeBasisPoints uint64) *SetParams {
	inst.FeeBasisPoints = &feeBasisPoints
	return inst
}

// SetGlobalAccount sets the "global" account.
func (inst *SetParams) SetGlobalAccount(global ag_solanago.PublicKey) *SetParams {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(global).WRITE()
	return inst
}

// GetGlobalAccount gets the "global" account.
func (inst *SetParams) GetGlobalAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetUserAccount sets the "user" account.
func (inst *SetParams) SetUserAccount(user ag_solanago.PublicKey) *SetParams {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(user).WRITE().SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *SetParams) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *SetParams) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *SetParams {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *SetParams) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *SetParams) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *SetParams {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *SetParams) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetProgramAccount sets the "program" account.
func (inst *SetParams) SetProgramAccount(program ag_solanago.PublicKey) *SetParams {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *SetParams) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst SetParams) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetParams,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetParams) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetParams) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.FeeRecipient == nil {
			return errors.New("FeeRecipient parameter is not set")
		}
		if inst.InitialVirtualTokenReserves == nil {
			return errors.New("InitialVirtualTokenReserves parameter is not set")
		}
		if inst.InitialVirtualSolReserves == nil {
			return errors.New("InitialVirtualSolReserves parameter is not set")
		}
		if inst.InitialRealTokenReserves == nil {
			return errors.New("InitialRealTokenReserves parameter is not set")
		}
		if inst.TokenTotalSupply == nil {
			return errors.New("TokenTotalSupply parameter is not set")
		}
		if inst.FeeBasisPoints == nil {
			return errors.New("FeeBasisPoints parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Global is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *SetParams) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetParams")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=6]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("               FeeRecipient", *inst.FeeRecipient))
						paramsBranch.Child(ag_format.Param("InitialVirtualTokenReserves", *inst.InitialVirtualTokenReserves))
						paramsBranch.Child(ag_format.Param("  InitialVirtualSolReserves", *inst.InitialVirtualSolReserves))
						paramsBranch.Child(ag_format.Param("   InitialRealTokenReserves", *inst.InitialRealTokenReserves))
						paramsBranch.Child(ag_format.Param("           TokenTotalSupply", *inst.TokenTotalSupply))
						paramsBranch.Child(ag_format.Param("             FeeBasisPoints", *inst.FeeBasisPoints))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        global", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          user", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" systemProgram", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("eventAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       program", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj SetParams) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `FeeRecipient` param:
	err = encoder.Encode(obj.FeeRecipient)
	if err != nil {
		return err
	}
	// Serialize `InitialVirtualTokenReserves` param:
	err = encoder.Encode(obj.InitialVirtualTokenReserves)
	if err != nil {
		return err
	}
	// Serialize `InitialVirtualSolReserves` param:
	err = encoder.Encode(obj.InitialVirtualSolReserves)
	if err != nil {
		return err
	}
	// Serialize `InitialRealTokenReserves` param:
	err = encoder.Encode(obj.InitialRealTokenReserves)
	if err != nil {
		return err
	}
	// Serialize `TokenTotalSupply` param:
	err = encoder.Encode(obj.TokenTotalSupply)
	if err != nil {
		return err
	}
	// Serialize `FeeBasisPoints` param:
	err = encoder.Encode(obj.FeeBasisPoints)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SetParams) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `FeeRecipient`:
	err = decoder.Decode(&obj.FeeRecipient)
	if err != nil {
		return err
	}
	// Deserialize `InitialVirtualTokenReserves`:
	err = decoder.Decode(&obj.InitialVirtualTokenReserves)
	if err != nil {
		return err
	}
	// Deserialize `InitialVirtualSolReserves`:
	err = decoder.Decode(&obj.InitialVirtualSolReserves)
	if err != nil {
		return err
	}
	// Deserialize `InitialRealTokenReserves`:
	err = decoder.Decode(&obj.InitialRealTokenReserves)
	if err != nil {
		return err
	}
	// Deserialize `TokenTotalSupply`:
	err = decoder.Decode(&obj.TokenTotalSupply)
	if err != nil {
		return err
	}
	// Deserialize `FeeBasisPoints`:
	err = decoder.Decode(&obj.FeeBasisPoints)
	if err != nil {
		return err
	}
	return nil
}

// NewSetParamsInstruction declares a new SetParams instruction with the provided parameters and accounts.
func NewSetParamsInstruction(
	// Parameters:
	feeRecipient ag_solanago.PublicKey,
	initialVirtualTokenReserves uint64,
	initialVirtualSolReserves uint64,
	initialRealTokenReserves uint64,
	tokenTotalSupply uint64,
	feeBasisPoints uint64,
	// Accounts:
	global ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *SetParams {
	return NewSetParamsInstructionBuilder().
		SetFeeRecipient(feeRecipient).
		SetInitialVirtualTokenReserves(initialVirtualTokenReserves).
		SetInitialVirtualSolReserves(initialVirtualSolReserves).
		SetInitialRealTokenReserves(initialRealTokenReserves).
		SetTokenTotalSupply(tokenTotalSupply).
		SetFeeBasisPoints(feeBasisPoints).
		SetGlobalAccount(global).
		SetUserAccount(user).
		SetSystemProgramAccount(systemProgram).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
