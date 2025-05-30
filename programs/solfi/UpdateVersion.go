// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package solfi

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateVersion is the `updateVersion` instruction.
type UpdateVersion struct {
	Version *uint8

	// [0] = [WRITE] tradingPair
	//
	// [1] = [] mintX
	//
	// [2] = [] mintY
	//
	// [3] = [] mintSslpX
	//
	// [4] = [] mintSslpY
	//
	// [5] = [SIGNER] user
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateVersionInstructionBuilder creates a new `UpdateVersion` instruction builder.
func NewUpdateVersionInstructionBuilder() *UpdateVersion {
	nd := &UpdateVersion{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetVersion sets the "version" parameter.
func (inst *UpdateVersion) SetVersion(version uint8) *UpdateVersion {
	inst.Version = &version
	return inst
}

// SetTradingPairAccount sets the "tradingPair" account.
func (inst *UpdateVersion) SetTradingPairAccount(tradingPair ag_solanago.PublicKey) *UpdateVersion {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(tradingPair).WRITE()
	return inst
}

// GetTradingPairAccount gets the "tradingPair" account.
func (inst *UpdateVersion) GetTradingPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintXAccount sets the "mintX" account.
func (inst *UpdateVersion) SetMintXAccount(mintX ag_solanago.PublicKey) *UpdateVersion {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mintX)
	return inst
}

// GetMintXAccount gets the "mintX" account.
func (inst *UpdateVersion) GetMintXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMintYAccount sets the "mintY" account.
func (inst *UpdateVersion) SetMintYAccount(mintY ag_solanago.PublicKey) *UpdateVersion {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(mintY)
	return inst
}

// GetMintYAccount gets the "mintY" account.
func (inst *UpdateVersion) GetMintYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetMintSslpXAccount sets the "mintSslpX" account.
func (inst *UpdateVersion) SetMintSslpXAccount(mintSslpX ag_solanago.PublicKey) *UpdateVersion {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mintSslpX)
	return inst
}

// GetMintSslpXAccount gets the "mintSslpX" account.
func (inst *UpdateVersion) GetMintSslpXAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMintSslpYAccount sets the "mintSslpY" account.
func (inst *UpdateVersion) SetMintSslpYAccount(mintSslpY ag_solanago.PublicKey) *UpdateVersion {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(mintSslpY)
	return inst
}

// GetMintSslpYAccount gets the "mintSslpY" account.
func (inst *UpdateVersion) GetMintSslpYAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetUserAccount sets the "user" account.
func (inst *UpdateVersion) SetUserAccount(user ag_solanago.PublicKey) *UpdateVersion {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(user).SIGNER()
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *UpdateVersion) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst UpdateVersion) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_UpdateVersion),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateVersion) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateVersion) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Version == nil {
			return errors.New("Version parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.TradingPair is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MintX is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.MintY is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MintSslpX is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MintSslpY is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.User is not set")
		}
	}
	return nil
}

func (inst *UpdateVersion) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateVersion")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Version", *inst.Version))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("tradingPair", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      mintX", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      mintY", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("  mintSslpX", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("  mintSslpY", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       user", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj UpdateVersion) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Version` param:
	err = encoder.Encode(obj.Version)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateVersion) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Version`:
	err = decoder.Decode(&obj.Version)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateVersionInstruction declares a new UpdateVersion instruction with the provided parameters and accounts.
func NewUpdateVersionInstruction(
	// Parameters:
	version uint8,
	// Accounts:
	tradingPair ag_solanago.PublicKey,
	mintX ag_solanago.PublicKey,
	mintY ag_solanago.PublicKey,
	mintSslpX ag_solanago.PublicKey,
	mintSslpY ag_solanago.PublicKey,
	user ag_solanago.PublicKey) *UpdateVersion {
	return NewUpdateVersionInstructionBuilder().
		SetVersion(version).
		SetTradingPairAccount(tradingPair).
		SetMintXAccount(mintX).
		SetMintYAccount(mintY).
		SetMintSslpXAccount(mintSslpX).
		SetMintSslpYAccount(mintSslpY).
		SetUserAccount(user)
}
