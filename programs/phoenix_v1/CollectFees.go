// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package phoenix_v1

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CollectFees is the `CollectFees` instruction.
type CollectFees struct {

	// [0] = [] phoenixProgram
	//
	// [1] = [] logAuthority
	//
	// [2] = [WRITE] market
	//
	// [3] = [SIGNER] sweeper
	//
	// [4] = [WRITE] feeRecipient
	//
	// [5] = [WRITE] quoteVault
	//
	// [6] = [] tokenProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCollectFeesInstructionBuilder creates a new `CollectFees` instruction builder.
func NewCollectFeesInstructionBuilder() *CollectFees {
	nd := &CollectFees{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetPhoenixProgramAccount sets the "phoenixProgram" account.
func (inst *CollectFees) SetPhoenixProgramAccount(phoenixProgram ag_solanago.PublicKey) *CollectFees {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(phoenixProgram)
	return inst
}

// GetPhoenixProgramAccount gets the "phoenixProgram" account.
func (inst *CollectFees) GetPhoenixProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetLogAuthorityAccount sets the "logAuthority" account.
func (inst *CollectFees) SetLogAuthorityAccount(logAuthority ag_solanago.PublicKey) *CollectFees {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(logAuthority)
	return inst
}

// GetLogAuthorityAccount gets the "logAuthority" account.
func (inst *CollectFees) GetLogAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetMarketAccount sets the "market" account.
func (inst *CollectFees) SetMarketAccount(market ag_solanago.PublicKey) *CollectFees {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *CollectFees) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSweeperAccount sets the "sweeper" account.
func (inst *CollectFees) SetSweeperAccount(sweeper ag_solanago.PublicKey) *CollectFees {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(sweeper).SIGNER()
	return inst
}

// GetSweeperAccount gets the "sweeper" account.
func (inst *CollectFees) GetSweeperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetFeeRecipientAccount sets the "feeRecipient" account.
func (inst *CollectFees) SetFeeRecipientAccount(feeRecipient ag_solanago.PublicKey) *CollectFees {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(feeRecipient).WRITE()
	return inst
}

// GetFeeRecipientAccount gets the "feeRecipient" account.
func (inst *CollectFees) GetFeeRecipientAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetQuoteVaultAccount sets the "quoteVault" account.
func (inst *CollectFees) SetQuoteVaultAccount(quoteVault ag_solanago.PublicKey) *CollectFees {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(quoteVault).WRITE()
	return inst
}

// GetQuoteVaultAccount gets the "quoteVault" account.
func (inst *CollectFees) GetQuoteVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *CollectFees) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CollectFees {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *CollectFees) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst CollectFees) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_CollectFees),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CollectFees) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CollectFees) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.PhoenixProgram is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.LogAuthority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Sweeper is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.FeeRecipient is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.QuoteVault is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
	}
	return nil
}

func (inst *CollectFees) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CollectFees")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("phoenixProgram", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("  logAuthority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("        market", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       sweeper", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("  feeRecipient", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("    quoteVault", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("  tokenProgram", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj CollectFees) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *CollectFees) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCollectFeesInstruction declares a new CollectFees instruction with the provided parameters and accounts.
func NewCollectFeesInstruction(
	// Accounts:
	phoenixProgram ag_solanago.PublicKey,
	logAuthority ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	sweeper ag_solanago.PublicKey,
	feeRecipient ag_solanago.PublicKey,
	quoteVault ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey) *CollectFees {
	return NewCollectFeesInstructionBuilder().
		SetPhoenixProgramAccount(phoenixProgram).
		SetLogAuthorityAccount(logAuthority).
		SetMarketAccount(market).
		SetSweeperAccount(sweeper).
		SetFeeRecipientAccount(feeRecipient).
		SetQuoteVaultAccount(quoteVault).
		SetTokenProgramAccount(tokenProgram)
}
