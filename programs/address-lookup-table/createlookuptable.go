// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package addresslookuptable

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateLookupTable is the `create_lookup_table` instruction.
type CreateLookupTable struct {
	Slot     *uint64
	BumpSeed *uint8

	// [0] = [WRITE] address_lookup_table
	//
	// [1] = [SIGNER] owner
	//
	// [2] = [WRITE, SIGNER] payer
	//
	// [3] = [] system_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateLookupTableInstructionBuilder creates a new `CreateLookupTable` instruction builder.
func NewCreateLookupTableInstructionBuilder() *CreateLookupTable {
	nd := &CreateLookupTable{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetSlot sets the "slot" parameter.
func (inst *CreateLookupTable) SetSlot(slot uint64) *CreateLookupTable {
	inst.Slot = &slot
	return inst
}

// SetBumpSeed sets the "bump_seed" parameter.
func (inst *CreateLookupTable) SetBumpSeed(bump_seed uint8) *CreateLookupTable {
	inst.BumpSeed = &bump_seed
	return inst
}

// SetAddressLookupTableAccount sets the "address_lookup_table" account.
func (inst *CreateLookupTable) SetAddressLookupTableAccount(addressLookupTable ag_solanago.PublicKey) *CreateLookupTable {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(addressLookupTable).WRITE()
	return inst
}

// GetAddressLookupTableAccount gets the "address_lookup_table" account.
func (inst *CreateLookupTable) GetAddressLookupTableAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAccount sets the "owner" account.
func (inst *CreateLookupTable) SetOwnerAccount(owner ag_solanago.PublicKey) *CreateLookupTable {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(owner).SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *CreateLookupTable) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
func (inst *CreateLookupTable) SetPayerAccount(payer ag_solanago.PublicKey) *CreateLookupTable {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *CreateLookupTable) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *CreateLookupTable) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateLookupTable {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *CreateLookupTable) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst CreateLookupTable) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_CreateLookupTable),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateLookupTable) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateLookupTable) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Slot == nil {
			return errors.New("Slot parameter is not set")
		}
		if inst.BumpSeed == nil {
			return errors.New("BumpSeed parameter is not set")
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

func (inst *CreateLookupTable) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateLookupTable")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("     Slot", *inst.Slot))
						paramsBranch.Child(ag_format.Param(" BumpSeed", *inst.BumpSeed))
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

func (obj CreateLookupTable) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Slot` param:
	err = encoder.Encode(obj.Slot)
	if err != nil {
		return err
	}
	// Serialize `BumpSeed` param:
	err = encoder.Encode(obj.BumpSeed)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateLookupTable) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Slot`:
	err = decoder.Decode(&obj.Slot)
	if err != nil {
		return err
	}
	// Deserialize `BumpSeed`:
	err = decoder.Decode(&obj.BumpSeed)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateLookupTableInstruction declares a new CreateLookupTable instruction with the provided parameters and accounts.
func NewCreateLookupTableInstruction(
	// Parameters:
	slot uint64,
	bump_seed uint8,
	// Accounts:
	addressLookupTable ag_solanago.PublicKey,
	owner ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *CreateLookupTable {
	return NewCreateLookupTableInstructionBuilder().
		SetSlot(slot).
		SetBumpSeed(bump_seed).
		SetAddressLookupTableAccount(addressLookupTable).
		SetOwnerAccount(owner).
		SetPayerAccount(payer).
		SetSystemProgramAccount(systemProgram)
}