// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupiter

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// CreateProgramOpenOrders is the `create_program_open_orders` instruction.
type CreateProgramOpenOrders struct {
	Id *uint8

	// [0] = [WRITE] open_orders
	//
	// [1] = [WRITE, SIGNER] payer
	//
	// [2] = [] program_authority
	//
	// [3] = [] dex_program
	//
	// [4] = [] system_program
	//
	// [5] = [] rent
	//
	// [6] = [] market
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCreateProgramOpenOrdersInstructionBuilder creates a new `CreateProgramOpenOrders` instruction builder.
func NewCreateProgramOpenOrdersInstructionBuilder() *CreateProgramOpenOrders {
	nd := &CreateProgramOpenOrders{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetId sets the "id" parameter.
func (inst *CreateProgramOpenOrders) SetId(id uint8) *CreateProgramOpenOrders {
	inst.Id = &id
	return inst
}

// SetOpenOrdersAccount sets the "open_orders" account.
func (inst *CreateProgramOpenOrders) SetOpenOrdersAccount(openOrders ag_solanago.PublicKey) *CreateProgramOpenOrders {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(openOrders).WRITE()
	return inst
}

// GetOpenOrdersAccount gets the "open_orders" account.
func (inst *CreateProgramOpenOrders) GetOpenOrdersAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPayerAccount sets the "payer" account.
func (inst *CreateProgramOpenOrders) SetPayerAccount(payer ag_solanago.PublicKey) *CreateProgramOpenOrders {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *CreateProgramOpenOrders) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetProgramAuthorityAccount sets the "program_authority" account.
func (inst *CreateProgramOpenOrders) SetProgramAuthorityAccount(programAuthority ag_solanago.PublicKey) *CreateProgramOpenOrders {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(programAuthority)
	return inst
}

// GetProgramAuthorityAccount gets the "program_authority" account.
func (inst *CreateProgramOpenOrders) GetProgramAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetDexProgramAccount sets the "dex_program" account.
func (inst *CreateProgramOpenOrders) SetDexProgramAccount(dexProgram ag_solanago.PublicKey) *CreateProgramOpenOrders {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(dexProgram)
	return inst
}

// GetDexProgramAccount gets the "dex_program" account.
func (inst *CreateProgramOpenOrders) GetDexProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *CreateProgramOpenOrders) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateProgramOpenOrders {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *CreateProgramOpenOrders) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetRentAccount sets the "rent" account.
func (inst *CreateProgramOpenOrders) SetRentAccount(rent ag_solanago.PublicKey) *CreateProgramOpenOrders {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *CreateProgramOpenOrders) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMarketAccount sets the "market" account.
func (inst *CreateProgramOpenOrders) SetMarketAccount(market ag_solanago.PublicKey) *CreateProgramOpenOrders {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(market)
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *CreateProgramOpenOrders) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst CreateProgramOpenOrders) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_CreateProgramOpenOrders,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateProgramOpenOrders) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateProgramOpenOrders) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Id == nil {
			return errors.New("Id parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.OpenOrders is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ProgramAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.DexProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Market is not set")
		}
	}
	return nil
}

func (inst *CreateProgramOpenOrders) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateProgramOpenOrders")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Id", *inst.Id))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      open_orders", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("            payer", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("program_authority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      dex_program", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   system_program", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("             rent", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("           market", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj CreateProgramOpenOrders) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Id` param:
	err = encoder.Encode(obj.Id)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateProgramOpenOrders) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Id`:
	err = decoder.Decode(&obj.Id)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateProgramOpenOrdersInstruction declares a new CreateProgramOpenOrders instruction with the provided parameters and accounts.
func NewCreateProgramOpenOrdersInstruction(
	// Parameters:
	id uint8,
	// Accounts:
	openOrders ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	programAuthority ag_solanago.PublicKey,
	dexProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	market ag_solanago.PublicKey) *CreateProgramOpenOrders {
	return NewCreateProgramOpenOrdersInstructionBuilder().
		SetId(id).
		SetOpenOrdersAccount(openOrders).
		SetPayerAccount(payer).
		SetProgramAuthorityAccount(programAuthority).
		SetDexProgramAccount(dexProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetMarketAccount(market)
}
