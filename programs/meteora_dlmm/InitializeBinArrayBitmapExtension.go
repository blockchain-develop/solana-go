// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeBinArrayBitmapExtension is the `initializeBinArrayBitmapExtension` instruction.
type InitializeBinArrayBitmapExtension struct {

	// [0] = [] lbPair
	//
	// [1] = [WRITE] binArrayBitmapExtension
	// ··········· Initialize an account to store if a bin array is initialized.
	//
	// [2] = [WRITE, SIGNER] funder
	//
	// [3] = [] systemProgram
	//
	// [4] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeBinArrayBitmapExtensionInstructionBuilder creates a new `InitializeBinArrayBitmapExtension` instruction builder.
func NewInitializeBinArrayBitmapExtensionInstructionBuilder() *InitializeBinArrayBitmapExtension {
	nd := &InitializeBinArrayBitmapExtension{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetLbPairAccount sets the "lbPair" account.
func (inst *InitializeBinArrayBitmapExtension) SetLbPairAccount(lbPair ag_solanago.PublicKey) *InitializeBinArrayBitmapExtension {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(lbPair)
	return inst
}

// GetLbPairAccount gets the "lbPair" account.
func (inst *InitializeBinArrayBitmapExtension) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetBinArrayBitmapExtensionAccount sets the "binArrayBitmapExtension" account.
// Initialize an account to store if a bin array is initialized.
func (inst *InitializeBinArrayBitmapExtension) SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension ag_solanago.PublicKey) *InitializeBinArrayBitmapExtension {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(binArrayBitmapExtension).WRITE()
	return inst
}

// GetBinArrayBitmapExtensionAccount gets the "binArrayBitmapExtension" account.
// Initialize an account to store if a bin array is initialized.
func (inst *InitializeBinArrayBitmapExtension) GetBinArrayBitmapExtensionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetFunderAccount sets the "funder" account.
func (inst *InitializeBinArrayBitmapExtension) SetFunderAccount(funder ag_solanago.PublicKey) *InitializeBinArrayBitmapExtension {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(funder).WRITE().SIGNER()
	return inst
}

// GetFunderAccount gets the "funder" account.
func (inst *InitializeBinArrayBitmapExtension) GetFunderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializeBinArrayBitmapExtension) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeBinArrayBitmapExtension {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializeBinArrayBitmapExtension) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetRentAccount sets the "rent" account.
func (inst *InitializeBinArrayBitmapExtension) SetRentAccount(rent ag_solanago.PublicKey) *InitializeBinArrayBitmapExtension {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *InitializeBinArrayBitmapExtension) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst InitializeBinArrayBitmapExtension) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeBinArrayBitmapExtension,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeBinArrayBitmapExtension) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeBinArrayBitmapExtension) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.LbPair is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.BinArrayBitmapExtension is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Funder is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *InitializeBinArrayBitmapExtension) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeBinArrayBitmapExtension")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 lbPair", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("binArrayBitmapExtension", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                 funder", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          systemProgram", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                   rent", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj InitializeBinArrayBitmapExtension) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *InitializeBinArrayBitmapExtension) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewInitializeBinArrayBitmapExtensionInstruction declares a new InitializeBinArrayBitmapExtension instruction with the provided parameters and accounts.
func NewInitializeBinArrayBitmapExtensionInstruction(
	// Accounts:
	lbPair ag_solanago.PublicKey,
	binArrayBitmapExtension ag_solanago.PublicKey,
	funder ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *InitializeBinArrayBitmapExtension {
	return NewInitializeBinArrayBitmapExtensionInstructionBuilder().
		SetLbPairAccount(lbPair).
		SetBinArrayBitmapExtensionAccount(binArrayBitmapExtension).
		SetFunderAccount(funder).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
