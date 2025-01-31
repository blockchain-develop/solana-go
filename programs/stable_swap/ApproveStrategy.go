// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package stable_swap

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ApproveStrategy is the `approve_strategy` instruction.
type ApproveStrategy struct {

	// ····· admin_only: [0] = [WRITE] pool
	//
	// [1] = [WRITE] strategy
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewApproveStrategyInstructionBuilder creates a new `ApproveStrategy` instruction builder.
func NewApproveStrategyInstructionBuilder() *ApproveStrategy {
	nd := &ApproveStrategy{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

type ApproveStrategyAdminOnlyAccountsBuilder struct {
	ag_solanago.AccountMetaSlice `bin:"-"`
}

func NewApproveStrategyAdminOnlyAccountsBuilder() *ApproveStrategyAdminOnlyAccountsBuilder {
	return &ApproveStrategyAdminOnlyAccountsBuilder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 1),
	}
}

func (inst *ApproveStrategy) SetAdminOnlyAccountsFromBuilder(approveStrategyAdminOnlyAccountsBuilder *ApproveStrategyAdminOnlyAccountsBuilder) *ApproveStrategy {
	inst.AccountMetaSlice[0] = approveStrategyAdminOnlyAccountsBuilder.GetPoolAccount()
	return inst
}

// SetPoolAccount sets the "pool" account.
func (inst *ApproveStrategyAdminOnlyAccountsBuilder) SetPoolAccount(pool ag_solanago.PublicKey) *ApproveStrategyAdminOnlyAccountsBuilder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(pool).WRITE()
	return inst
}

// GetPoolAccount gets the "pool" account.
func (inst *ApproveStrategyAdminOnlyAccountsBuilder) GetPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetStrategyAccount sets the "strategy" account.
func (inst *ApproveStrategy) SetStrategyAccount(strategy ag_solanago.PublicKey) *ApproveStrategy {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(strategy).WRITE()
	return inst
}

// GetStrategyAccount gets the "strategy" account.
func (inst *ApproveStrategy) GetStrategyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst ApproveStrategy) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ApproveStrategy,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ApproveStrategy) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ApproveStrategy) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AdminOnlyPool is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Strategy is not set")
		}
	}
	return nil
}

func (inst *ApproveStrategy) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ApproveStrategy")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("admin_only/pool", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       strategy", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj ApproveStrategy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ApproveStrategy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewApproveStrategyInstruction declares a new ApproveStrategy instruction with the provided parameters and accounts.
func NewApproveStrategyInstruction(
	// Accounts:
	adminOnlyPool ag_solanago.PublicKey,
	strategy ag_solanago.PublicKey) *ApproveStrategy {
	return NewApproveStrategyInstructionBuilder().
		SetAdminOnlyAccountsFromBuilder(
			NewApproveStrategyAdminOnlyAccountsBuilder().
				SetPoolAccount(adminOnlyPool)).
		SetStrategyAccount(strategy)
}
