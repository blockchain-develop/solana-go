// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package stable_swap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// initializeapool
type Initialize struct {
	AmpFactor *uint16
	SwapFee   *uint64
	MaxCaps   *[]uint64

	// [0] = [SIGNER] owner
	//
	// [1] = [] mint
	//
	// [2] = [WRITE] pool
	//
	// [3] = [] pool_authority
	//
	// [4] = [] withdraw_authority
	//
	// [5] = [] vault
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetAmpFactor sets the "amp_factor" parameter.
func (inst *Initialize) SetAmpFactor(amp_factor uint16) *Initialize {
	inst.AmpFactor = &amp_factor
	return inst
}

// SetSwapFee sets the "swap_fee" parameter.
func (inst *Initialize) SetSwapFee(swap_fee uint64) *Initialize {
	inst.SwapFee = &swap_fee
	return inst
}

// SetMaxCaps sets the "max_caps" parameter.
func (inst *Initialize) SetMaxCaps(max_caps []uint64) *Initialize {
	inst.MaxCaps = &max_caps
	return inst
}

// SetOwnerAccount sets the "owner" account.
func (inst *Initialize) SetOwnerAccount(owner ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *Initialize) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" account.
func (inst *Initialize) SetMintAccount(mint ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *Initialize) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPoolAccount sets the "pool" account.
func (inst *Initialize) SetPoolAccount(pool ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(pool).WRITE()
	return inst
}

// GetPoolAccount gets the "pool" account.
func (inst *Initialize) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPoolAuthorityAccount sets the "pool_authority" account.
func (inst *Initialize) SetPoolAuthorityAccount(poolAuthority ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(poolAuthority)
	return inst
}

// GetPoolAuthorityAccount gets the "pool_authority" account.
func (inst *Initialize) GetPoolAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetWithdrawAuthorityAccount sets the "withdraw_authority" account.
func (inst *Initialize) SetWithdrawAuthorityAccount(withdrawAuthority ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(withdrawAuthority)
	return inst
}

// GetWithdrawAuthorityAccount gets the "withdraw_authority" account.
func (inst *Initialize) GetWithdrawAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetVaultAccount sets the "vault" account.
func (inst *Initialize) SetVaultAccount(vault ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(vault)
	return inst
}

// GetVaultAccount gets the "vault" account.
func (inst *Initialize) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Initialize,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.AmpFactor == nil {
			return errors.New("AmpFactor parameter is not set")
		}
		if inst.SwapFee == nil {
			return errors.New("SwapFee parameter is not set")
		}
		if inst.MaxCaps == nil {
			return errors.New("MaxCaps parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Pool is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PoolAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.WithdrawAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Vault is not set")
		}
	}
	return nil
}

func (inst *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" AmpFactor", *inst.AmpFactor))
						paramsBranch.Child(ag_format.Param("   SwapFee", *inst.SwapFee))
						paramsBranch.Child(ag_format.Param("   MaxCaps", *inst.MaxCaps))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             owner", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("              mint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("              pool", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    pool_authority", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("withdraw_authority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("             vault", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `AmpFactor` param:
	err = encoder.Encode(obj.AmpFactor)
	if err != nil {
		return err
	}
	// Serialize `SwapFee` param:
	err = encoder.Encode(obj.SwapFee)
	if err != nil {
		return err
	}
	// Serialize `MaxCaps` param:
	err = encoder.Encode(obj.MaxCaps)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `AmpFactor`:
	err = decoder.Decode(&obj.AmpFactor)
	if err != nil {
		return err
	}
	// Deserialize `SwapFee`:
	err = decoder.Decode(&obj.SwapFee)
	if err != nil {
		return err
	}
	// Deserialize `MaxCaps`:
	err = decoder.Decode(&obj.MaxCaps)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Parameters:
	amp_factor uint16,
	swap_fee uint64,
	max_caps []uint64,
	// Accounts:
	owner ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	pool ag_solanago.PublicKey,
	poolAuthority ag_solanago.PublicKey,
	withdrawAuthority ag_solanago.PublicKey,
	vault ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetAmpFactor(amp_factor).
		SetSwapFee(swap_fee).
		SetMaxCaps(max_caps).
		SetOwnerAccount(owner).
		SetMintAccount(mint).
		SetPoolAccount(pool).
		SetPoolAuthorityAccount(poolAuthority).
		SetWithdrawAuthorityAccount(withdrawAuthority).
		SetVaultAccount(vault)
}
