// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package addresslookuptable

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ExtendLookupTable is the `extend_lookup_table` instruction.
type ExtendLookupTable struct {
	NewAddresses *[]ag_solanago.PublicKey

	// [0] = [WRITE] address_lookup_table
	//
	// [1] = [SIGNER] owner
	//
	// [2] = [WRITE, SIGNER] payer
	//
	// [3] = [] system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewExtendLookupTableInstructionBuilder creates a new `ExtendLookupTable` instruction builder.
func NewExtendLookupTableInstructionBuilder() *ExtendLookupTable {
	nd := &ExtendLookupTable{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetNewAddresses sets the "new_addresses" parameter.
func (inst *ExtendLookupTable) SetNewAddresses(new_addresses []ag_solanago.PublicKey) *ExtendLookupTable {
	inst.NewAddresses = &new_addresses
	return inst
}

// SetAddressLookupTableAccount sets the "address_lookup_table" account.
func (inst *ExtendLookupTable) SetAddressLookupTableAccount(addressLookupTable ag_solanago.PublicKey) *ExtendLookupTable {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(addressLookupTable).WRITE()
	return inst
}

// GetAddressLookupTableAccount gets the "address_lookup_table" account.
func (inst *ExtendLookupTable) GetAddressLookupTableAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *ExtendLookupTable) SetOwnerAccount(owner ag_solanago.PublicKey) *ExtendLookupTable {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *ExtendLookupTable) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
func (inst *ExtendLookupTable) SetPayerAccount(payer ag_solanago.PublicKey) *ExtendLookupTable {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *ExtendLookupTable) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *ExtendLookupTable) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ExtendLookupTable {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *ExtendLookupTable) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst ExtendLookupTable) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_ExtendLookupTable),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ExtendLookupTable) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ExtendLookupTable) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.NewAddresses == nil {
			return errors.New("NewAddresses parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AddressLookupTable is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *ExtendLookupTable) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ExtendLookupTable")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" NewAddresses", *inst.NewAddresses))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("address_lookup_table", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("               owner", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      system_program", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj ExtendLookupTable) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `NewAddresses` param:
	err = encoder.Encode(obj.NewAddresses)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ExtendLookupTable) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `NewAddresses`:
	err = decoder.Decode(&obj.NewAddresses)
	if err != nil {
		return err
	}
	return nil
}

// NewExtendLookupTableInstruction declares a new ExtendLookupTable instruction with the provided parameters and accounts.
func NewExtendLookupTableInstruction(
	// Parameters:
	new_addresses []ag_solanago.PublicKey,
	// Accounts:
	addressLookupTable ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *ExtendLookupTable {
	return NewExtendLookupTableInstructionBuilder().
		SetNewAddresses(new_addresses).
		SetAddressLookupTableAccount(addressLookupTable).
		SetOwnerAccount(owner).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram)
}