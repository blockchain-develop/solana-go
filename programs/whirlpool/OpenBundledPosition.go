// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package whirlpool

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// OpenBundledPosition is the `openBundledPosition` instruction.
type OpenBundledPosition struct {
	BundleIndex    *uint16
	TickLowerIndex *int32
	TickUpperIndex *int32

	// [0] = [WRITE] bundledPosition
	//
	// [1] = [WRITE] positionBundle
	//
	// [2] = [] positionBundleTokenAccount
	//
	// [3] = [SIGNER] positionBundleAuthority
	//
	// [4] = [] whirlpool
	//
	// [5] = [WRITE, SIGNER] funder
	//
	// [6] = [] systemProgram
	//
	// [7] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewOpenBundledPositionInstructionBuilder creates a new `OpenBundledPosition` instruction builder.
func NewOpenBundledPositionInstructionBuilder() *OpenBundledPosition {
	nd := &OpenBundledPosition{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetBundleIndex sets the "bundleIndex" parameter.
func (inst *OpenBundledPosition) SetBundleIndex(bundleIndex uint16) *OpenBundledPosition {
	inst.BundleIndex = &bundleIndex
	return inst
}

// SetTickLowerIndex sets the "tickLowerIndex" parameter.
func (inst *OpenBundledPosition) SetTickLowerIndex(tickLowerIndex int32) *OpenBundledPosition {
	inst.TickLowerIndex = &tickLowerIndex
	return inst
}

// SetTickUpperIndex sets the "tickUpperIndex" parameter.
func (inst *OpenBundledPosition) SetTickUpperIndex(tickUpperIndex int32) *OpenBundledPosition {
	inst.TickUpperIndex = &tickUpperIndex
	return inst
}

// SetBundledPositionAccount sets the "bundledPosition" account.
func (inst *OpenBundledPosition) SetBundledPositionAccount(bundledPosition ag_solanago.PublicKey) *OpenBundledPosition {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(bundledPosition).WRITE()
	return inst
}

// GetBundledPositionAccount gets the "bundledPosition" account.
func (inst *OpenBundledPosition) GetBundledPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPositionBundleAccount sets the "positionBundle" account.
func (inst *OpenBundledPosition) SetPositionBundleAccount(positionBundle ag_solanago.PublicKey) *OpenBundledPosition {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(positionBundle).WRITE()
	return inst
}

// GetPositionBundleAccount gets the "positionBundle" account.
func (inst *OpenBundledPosition) GetPositionBundleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPositionBundleTokenAccountAccount sets the "positionBundleTokenAccount" account.
func (inst *OpenBundledPosition) SetPositionBundleTokenAccountAccount(positionBundleTokenAccount ag_solanago.PublicKey) *OpenBundledPosition {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(positionBundleTokenAccount)
	return inst
}

// GetPositionBundleTokenAccountAccount gets the "positionBundleTokenAccount" account.
func (inst *OpenBundledPosition) GetPositionBundleTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPositionBundleAuthorityAccount sets the "positionBundleAuthority" account.
func (inst *OpenBundledPosition) SetPositionBundleAuthorityAccount(positionBundleAuthority ag_solanago.PublicKey) *OpenBundledPosition {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(positionBundleAuthority).SIGNER()
	return inst
}

// GetPositionBundleAuthorityAccount gets the "positionBundleAuthority" account.
func (inst *OpenBundledPosition) GetPositionBundleAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetWhirlpoolAccount sets the "whirlpool" account.
func (inst *OpenBundledPosition) SetWhirlpoolAccount(whirlpool ag_solanago.PublicKey) *OpenBundledPosition {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(whirlpool)
	return inst
}

// GetWhirlpoolAccount gets the "whirlpool" account.
func (inst *OpenBundledPosition) GetWhirlpoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetFunderAccount sets the "funder" account.
func (inst *OpenBundledPosition) SetFunderAccount(funder ag_solanago.PublicKey) *OpenBundledPosition {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(funder).WRITE().SIGNER()
	return inst
}

// GetFunderAccount gets the "funder" account.
func (inst *OpenBundledPosition) GetFunderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *OpenBundledPosition) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *OpenBundledPosition {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *OpenBundledPosition) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetRentAccount sets the "rent" account.
func (inst *OpenBundledPosition) SetRentAccount(rent ag_solanago.PublicKey) *OpenBundledPosition {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *OpenBundledPosition) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst OpenBundledPosition) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_OpenBundledPosition,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst OpenBundledPosition) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *OpenBundledPosition) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.BundleIndex == nil {
			return errors.New("BundleIndex parameter is not set")
		}
		if inst.TickLowerIndex == nil {
			return errors.New("TickLowerIndex parameter is not set")
		}
		if inst.TickUpperIndex == nil {
			return errors.New("TickUpperIndex parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.BundledPosition is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.PositionBundle is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PositionBundleTokenAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PositionBundleAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Whirlpool is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Funder is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *OpenBundledPosition) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("OpenBundledPosition")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   BundleIndex", *inst.BundleIndex))
						paramsBranch.Child(ag_format.Param("TickLowerIndex", *inst.TickLowerIndex))
						paramsBranch.Child(ag_format.Param("TickUpperIndex", *inst.TickUpperIndex))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        bundledPosition", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         positionBundle", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    positionBundleToken", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("positionBundleAuthority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("              whirlpool", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                 funder", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("          systemProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                   rent", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj OpenBundledPosition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `BundleIndex` param:
	err = encoder.Encode(obj.BundleIndex)
	if err != nil {
		return err
	}
	// Serialize `TickLowerIndex` param:
	err = encoder.Encode(obj.TickLowerIndex)
	if err != nil {
		return err
	}
	// Serialize `TickUpperIndex` param:
	err = encoder.Encode(obj.TickUpperIndex)
	if err != nil {
		return err
	}
	return nil
}
func (obj *OpenBundledPosition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `BundleIndex`:
	err = decoder.Decode(&obj.BundleIndex)
	if err != nil {
		return err
	}
	// Deserialize `TickLowerIndex`:
	err = decoder.Decode(&obj.TickLowerIndex)
	if err != nil {
		return err
	}
	// Deserialize `TickUpperIndex`:
	err = decoder.Decode(&obj.TickUpperIndex)
	if err != nil {
		return err
	}
	return nil
}

// NewOpenBundledPositionInstruction declares a new OpenBundledPosition instruction with the provided parameters and accounts.
func NewOpenBundledPositionInstruction(
	// Parameters:
	bundleIndex uint16,
	tickLowerIndex int32,
	tickUpperIndex int32,
	// Accounts:
	bundledPosition ag_solanago.PublicKey,
	positionBundle ag_solanago.PublicKey,
	positionBundleTokenAccount ag_solanago.PublicKey,
	positionBundleAuthority ag_solanago.PublicKey,
	whirlpool ag_solanago.PublicKey,
	funder ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *OpenBundledPosition {
	return NewOpenBundledPositionInstructionBuilder().
		SetBundleIndex(bundleIndex).
		SetTickLowerIndex(tickLowerIndex).
		SetTickUpperIndex(tickUpperIndex).
		SetBundledPositionAccount(bundledPosition).
		SetPositionBundleAccount(positionBundle).
		SetPositionBundleTokenAccountAccount(positionBundleTokenAccount).
		SetPositionBundleAuthorityAccount(positionBundleAuthority).
		SetWhirlpoolAccount(whirlpool).
		SetFunderAccount(funder).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
