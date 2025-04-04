// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_dlmm

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// FundReward is the `fundReward` instruction.
type FundReward struct {
	RewardIndex  *uint64
	Amount       *uint64
	CarryForward *bool

	// [0] = [WRITE] lbPair
	//
	// [1] = [WRITE] rewardVault
	//
	// [2] = [] rewardMint
	//
	// [3] = [WRITE] funderTokenAccount
	//
	// [4] = [SIGNER] funder
	//
	// [5] = [WRITE] binArray
	//
	// [6] = [] tokenProgram
	//
	// [7] = [] eventAuthority
	//
	// [8] = [] program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewFundRewardInstructionBuilder creates a new `FundReward` instruction builder.
func NewFundRewardInstructionBuilder() *FundReward {
	nd := &FundReward{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetRewardIndex sets the "rewardIndex" parameter.
func (inst *FundReward) SetRewardIndex(rewardIndex uint64) *FundReward {
	inst.RewardIndex = &rewardIndex
	return inst
}

// SetAmount sets the "amount" parameter.
func (inst *FundReward) SetAmount(amount uint64) *FundReward {
	inst.Amount = &amount
	return inst
}

// SetCarryForward sets the "carryForward" parameter.
func (inst *FundReward) SetCarryForward(carryForward bool) *FundReward {
	inst.CarryForward = &carryForward
	return inst
}

// SetLbPairAccount sets the "lbPair" account.
func (inst *FundReward) SetLbPairAccount(lbPair ag_solanago.PublicKey) *FundReward {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(lbPair).WRITE()
	return inst
}

// GetLbPairAccount gets the "lbPair" account.
func (inst *FundReward) GetLbPairAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetRewardVaultAccount sets the "rewardVault" account.
func (inst *FundReward) SetRewardVaultAccount(rewardVault ag_solanago.PublicKey) *FundReward {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(rewardVault).WRITE()
	return inst
}

// GetRewardVaultAccount gets the "rewardVault" account.
func (inst *FundReward) GetRewardVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetRewardMintAccount sets the "rewardMint" account.
func (inst *FundReward) SetRewardMintAccount(rewardMint ag_solanago.PublicKey) *FundReward {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(rewardMint)
	return inst
}

// GetRewardMintAccount gets the "rewardMint" account.
func (inst *FundReward) GetRewardMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetFunderTokenAccountAccount sets the "funderTokenAccount" account.
func (inst *FundReward) SetFunderTokenAccountAccount(funderTokenAccount ag_solanago.PublicKey) *FundReward {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(funderTokenAccount).WRITE()
	return inst
}

// GetFunderTokenAccountAccount gets the "funderTokenAccount" account.
func (inst *FundReward) GetFunderTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetFunderAccount sets the "funder" account.
func (inst *FundReward) SetFunderAccount(funder ag_solanago.PublicKey) *FundReward {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(funder).SIGNER()
	return inst
}

// GetFunderAccount gets the "funder" account.
func (inst *FundReward) GetFunderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetBinArrayAccount sets the "binArray" account.
func (inst *FundReward) SetBinArrayAccount(binArray ag_solanago.PublicKey) *FundReward {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(binArray).WRITE()
	return inst
}

// GetBinArrayAccount gets the "binArray" account.
func (inst *FundReward) GetBinArrayAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *FundReward) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *FundReward {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *FundReward) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetEventAuthorityAccount sets the "eventAuthority" account.
func (inst *FundReward) SetEventAuthorityAccount(eventAuthority ag_solanago.PublicKey) *FundReward {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(eventAuthority)
	return inst
}

// GetEventAuthorityAccount gets the "eventAuthority" account.
func (inst *FundReward) GetEventAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetProgramAccount sets the "program" account.
func (inst *FundReward) SetProgramAccount(program ag_solanago.PublicKey) *FundReward {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(program)
	return inst
}

// GetProgramAccount gets the "program" account.
func (inst *FundReward) GetProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst FundReward) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_FundReward,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst FundReward) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *FundReward) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RewardIndex == nil {
			return errors.New("RewardIndex parameter is not set")
		}
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
		if inst.CarryForward == nil {
			return errors.New("CarryForward parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.LbPair is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.RewardVault is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.RewardMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.FunderTokenAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Funder is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.BinArray is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.EventAuthority is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Program is not set")
		}
	}
	return nil
}

func (inst *FundReward) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("FundReward")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" RewardIndex", *inst.RewardIndex))
						paramsBranch.Child(ag_format.Param("      Amount", *inst.Amount))
						paramsBranch.Child(ag_format.Param("CarryForward", *inst.CarryForward))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        lbPair", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   rewardVault", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    rewardMint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("   funderToken", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        funder", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("      binArray", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("  tokenProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("eventAuthority", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("       program", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj FundReward) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RewardIndex` param:
	err = encoder.Encode(obj.RewardIndex)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `CarryForward` param:
	err = encoder.Encode(obj.CarryForward)
	if err != nil {
		return err
	}
	return nil
}
func (obj *FundReward) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RewardIndex`:
	err = decoder.Decode(&obj.RewardIndex)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `CarryForward`:
	err = decoder.Decode(&obj.CarryForward)
	if err != nil {
		return err
	}
	return nil
}

// NewFundRewardInstruction declares a new FundReward instruction with the provided parameters and accounts.
func NewFundRewardInstruction(
	// Parameters:
	rewardIndex uint64,
	amount uint64,
	carryForward bool,
	// Accounts:
	lbPair ag_solanago.PublicKey,
	rewardVault ag_solanago.PublicKey,
	rewardMint ag_solanago.PublicKey,
	funderTokenAccount ag_solanago.PublicKey,
	funder ag_solanago.PublicKey,
	binArray ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	eventAuthority ag_solanago.PublicKey,
	program ag_solanago.PublicKey) *FundReward {
	return NewFundRewardInstructionBuilder().
		SetRewardIndex(rewardIndex).
		SetAmount(amount).
		SetCarryForward(carryForward).
		SetLbPairAccount(lbPair).
		SetRewardVaultAccount(rewardVault).
		SetRewardMintAccount(rewardMint).
		SetFunderTokenAccountAccount(funderTokenAccount).
		SetFunderAccount(funder).
		SetBinArrayAccount(binArray).
		SetTokenProgramAccount(tokenProgram).
		SetEventAuthorityAccount(eventAuthority).
		SetProgramAccount(program)
}
