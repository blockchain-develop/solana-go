// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package stable_vault

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// TransferAdmin is the `transfer_admin` instruction.
type TransferAdmin struct {
	NewAdmin *ag_solanago.PublicKey

	// [0] = [] admin
	//
	// [1] = [] vault
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewTransferAdminInstructionBuilder creates a new `TransferAdmin` instruction builder.
func NewTransferAdminInstructionBuilder() *TransferAdmin {
	nd := &TransferAdmin{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetNewAdmin sets the "new_admin" parameter.
func (inst *TransferAdmin) SetNewAdmin(new_admin ag_solanago.PublicKey) *TransferAdmin {
	inst.NewAdmin = &new_admin
	return inst
}

// SetAdminAccount sets the "admin" account.
func (inst *TransferAdmin) SetAdminAccount(admin ag_solanago.PublicKey) *TransferAdmin {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(admin)
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *TransferAdmin) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetVaultAccount sets the "vault" account.
func (inst *TransferAdmin) SetVaultAccount(vault ag_solanago.PublicKey) *TransferAdmin {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(vault)
	return inst
}

// GetVaultAccount gets the "vault" account.
func (inst *TransferAdmin) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst TransferAdmin) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_TransferAdmin,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst TransferAdmin) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *TransferAdmin) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewAdmin == nil {
			return errors.New("NewAdmin parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Vault is not set")
		}
	}
	return nil
}

func (inst *TransferAdmin) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("TransferAdmin")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" NewAdmin", *inst.NewAdmin))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("admin", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("vault", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj TransferAdmin) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewAdmin` param:
	err = encoder.Encode(obj.NewAdmin)
	if err != nil {
		return err
	}
	return nil
}
func (obj *TransferAdmin) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewAdmin`:
	err = decoder.Decode(&obj.NewAdmin)
	if err != nil {
		return err
	}
	return nil
}

// NewTransferAdminInstruction declares a new TransferAdmin instruction with the provided parameters and accounts.
func NewTransferAdminInstruction(
	// Parameters:
	new_admin ag_solanago.PublicKey,
	// Accounts:
	admin ag_solanago.PublicKey,
	vault ag_solanago.PublicKey) *TransferAdmin {
	return NewTransferAdminInstructionBuilder().
		SetNewAdmin(new_admin).
		SetAdminAccount(admin).
		SetVaultAccount(vault)
}
