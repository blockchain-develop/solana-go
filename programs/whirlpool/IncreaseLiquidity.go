// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package generated

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// IncreaseLiquidity is the `increaseLiquidity` instruction.
type IncreaseLiquidity struct {
	LiquidityAmount *ag_binary.Uint128
	TokenMaxA       *uint64
	TokenMaxB       *uint64

	// [0] = [WRITE] whirlpool
	//
	// [1] = [] tokenProgram
	//
	// [2] = [SIGNER] positionAuthority
	//
	// [3] = [WRITE] position
	//
	// [4] = [] positionTokenAccount
	//
	// [5] = [WRITE] tokenOwnerAccountA
	//
	// [6] = [WRITE] tokenOwnerAccountB
	//
	// [7] = [WRITE] tokenVaultA
	//
	// [8] = [WRITE] tokenVaultB
	//
	// [9] = [WRITE] tickArrayLower
	//
	// [10] = [WRITE] tickArrayUpper
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewIncreaseLiquidityInstructionBuilder creates a new `IncreaseLiquidity` instruction builder.
func NewIncreaseLiquidityInstructionBuilder() *IncreaseLiquidity {
	nd := &IncreaseLiquidity{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	return nd
}

// SetLiquidityAmount sets the "liquidityAmount" parameter.
func (inst *IncreaseLiquidity) SetLiquidityAmount(liquidityAmount ag_binary.Uint128) *IncreaseLiquidity {
	inst.LiquidityAmount = &liquidityAmount
	return inst
}

// SetTokenMaxA sets the "tokenMaxA" parameter.
func (inst *IncreaseLiquidity) SetTokenMaxA(tokenMaxA uint64) *IncreaseLiquidity {
	inst.TokenMaxA = &tokenMaxA
	return inst
}

// SetTokenMaxB sets the "tokenMaxB" parameter.
func (inst *IncreaseLiquidity) SetTokenMaxB(tokenMaxB uint64) *IncreaseLiquidity {
	inst.TokenMaxB = &tokenMaxB
	return inst
}

// SetWhirlpoolAccount sets the "whirlpool" account.
func (inst *IncreaseLiquidity) SetWhirlpoolAccount(whirlpool ag_solanago.PublicKey) *IncreaseLiquidity {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(whirlpool).WRITE()
	return inst
}

// GetWhirlpoolAccount gets the "whirlpool" account.
func (inst *IncreaseLiquidity) GetWhirlpoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *IncreaseLiquidity) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *IncreaseLiquidity {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *IncreaseLiquidity) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPositionAuthorityAccount sets the "positionAuthority" account.
func (inst *IncreaseLiquidity) SetPositionAuthorityAccount(positionAuthority ag_solanago.PublicKey) *IncreaseLiquidity {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(positionAuthority).SIGNER()
	return inst
}

// GetPositionAuthorityAccount gets the "positionAuthority" account.
func (inst *IncreaseLiquidity) GetPositionAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPositionAccount sets the "position" account.
func (inst *IncreaseLiquidity) SetPositionAccount(position ag_solanago.PublicKey) *IncreaseLiquidity {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(position).WRITE()
	return inst
}

// GetPositionAccount gets the "position" account.
func (inst *IncreaseLiquidity) GetPositionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPositionTokenAccountAccount sets the "positionTokenAccount" account.
func (inst *IncreaseLiquidity) SetPositionTokenAccountAccount(positionTokenAccount ag_solanago.PublicKey) *IncreaseLiquidity {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(positionTokenAccount)
	return inst
}

// GetPositionTokenAccountAccount gets the "positionTokenAccount" account.
func (inst *IncreaseLiquidity) GetPositionTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenOwnerAccountAAccount sets the "tokenOwnerAccountA" account.
func (inst *IncreaseLiquidity) SetTokenOwnerAccountAAccount(tokenOwnerAccountA ag_solanago.PublicKey) *IncreaseLiquidity {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenOwnerAccountA).WRITE()
	return inst
}

// GetTokenOwnerAccountAAccount gets the "tokenOwnerAccountA" account.
func (inst *IncreaseLiquidity) GetTokenOwnerAccountAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetTokenOwnerAccountBAccount sets the "tokenOwnerAccountB" account.
func (inst *IncreaseLiquidity) SetTokenOwnerAccountBAccount(tokenOwnerAccountB ag_solanago.PublicKey) *IncreaseLiquidity {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenOwnerAccountB).WRITE()
	return inst
}

// GetTokenOwnerAccountBAccount gets the "tokenOwnerAccountB" account.
func (inst *IncreaseLiquidity) GetTokenOwnerAccountBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenVaultAAccount sets the "tokenVaultA" account.
func (inst *IncreaseLiquidity) SetTokenVaultAAccount(tokenVaultA ag_solanago.PublicKey) *IncreaseLiquidity {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenVaultA).WRITE()
	return inst
}

// GetTokenVaultAAccount gets the "tokenVaultA" account.
func (inst *IncreaseLiquidity) GetTokenVaultAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenVaultBAccount sets the "tokenVaultB" account.
func (inst *IncreaseLiquidity) SetTokenVaultBAccount(tokenVaultB ag_solanago.PublicKey) *IncreaseLiquidity {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenVaultB).WRITE()
	return inst
}

// GetTokenVaultBAccount gets the "tokenVaultB" account.
func (inst *IncreaseLiquidity) GetTokenVaultBAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTickArrayLowerAccount sets the "tickArrayLower" account.
func (inst *IncreaseLiquidity) SetTickArrayLowerAccount(tickArrayLower ag_solanago.PublicKey) *IncreaseLiquidity {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tickArrayLower).WRITE()
	return inst
}

// GetTickArrayLowerAccount gets the "tickArrayLower" account.
func (inst *IncreaseLiquidity) GetTickArrayLowerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTickArrayUpperAccount sets the "tickArrayUpper" account.
func (inst *IncreaseLiquidity) SetTickArrayUpperAccount(tickArrayUpper ag_solanago.PublicKey) *IncreaseLiquidity {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tickArrayUpper).WRITE()
	return inst
}

// GetTickArrayUpperAccount gets the "tickArrayUpper" account.
func (inst *IncreaseLiquidity) GetTickArrayUpperAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst IncreaseLiquidity) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_IncreaseLiquidity,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst IncreaseLiquidity) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *IncreaseLiquidity) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.LiquidityAmount == nil {
			return errors.New("LiquidityAmount parameter is not set")
		}
		if inst.TokenMaxA == nil {
			return errors.New("TokenMaxA parameter is not set")
		}
		if inst.TokenMaxB == nil {
			return errors.New("TokenMaxB parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Whirlpool is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PositionAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Position is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.PositionTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenOwnerAccountA is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenOwnerAccountB is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenVaultA is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenVaultB is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TickArrayLower is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TickArrayUpper is not set")
		}
	}
	return nil
}

func (inst *IncreaseLiquidity) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("IncreaseLiquidity")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("LiquidityAmount", *inst.LiquidityAmount))
						paramsBranch.Child(ag_format.Param("      TokenMaxA", *inst.TokenMaxA))
						paramsBranch.Child(ag_format.Param("      TokenMaxB", *inst.TokenMaxB))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         whirlpool", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("      tokenProgram", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta(" positionAuthority", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          position", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     positionToken", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("tokenOwnerAccountA", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("tokenOwnerAccountB", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       tokenVaultA", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("       tokenVaultB", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("    tickArrayLower", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("    tickArrayUpper", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj IncreaseLiquidity) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `LiquidityAmount` param:
	err = encoder.Encode(obj.LiquidityAmount)
	if err != nil {
		return err
	}
	// Serialize `TokenMaxA` param:
	err = encoder.Encode(obj.TokenMaxA)
	if err != nil {
		return err
	}
	// Serialize `TokenMaxB` param:
	err = encoder.Encode(obj.TokenMaxB)
	if err != nil {
		return err
	}
	return nil
}
func (obj *IncreaseLiquidity) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `LiquidityAmount`:
	err = decoder.Decode(&obj.LiquidityAmount)
	if err != nil {
		return err
	}
	// Deserialize `TokenMaxA`:
	err = decoder.Decode(&obj.TokenMaxA)
	if err != nil {
		return err
	}
	// Deserialize `TokenMaxB`:
	err = decoder.Decode(&obj.TokenMaxB)
	if err != nil {
		return err
	}
	return nil
}

// NewIncreaseLiquidityInstruction declares a new IncreaseLiquidity instruction with the provided parameters and accounts.
func NewIncreaseLiquidityInstruction(
	// Parameters:
	liquidityAmount ag_binary.Uint128,
	tokenMaxA uint64,
	tokenMaxB uint64,
	// Accounts:
	whirlpool ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	positionAuthority ag_solanago.PublicKey,
	position ag_solanago.PublicKey,
	positionTokenAccount ag_solanago.PublicKey,
	tokenOwnerAccountA ag_solanago.PublicKey,
	tokenOwnerAccountB ag_solanago.PublicKey,
	tokenVaultA ag_solanago.PublicKey,
	tokenVaultB ag_solanago.PublicKey,
	tickArrayLower ag_solanago.PublicKey,
	tickArrayUpper ag_solanago.PublicKey) *IncreaseLiquidity {
	return NewIncreaseLiquidityInstructionBuilder().
		SetLiquidityAmount(liquidityAmount).
		SetTokenMaxA(tokenMaxA).
		SetTokenMaxB(tokenMaxB).
		SetWhirlpoolAccount(whirlpool).
		SetTokenProgramAccount(tokenProgram).
		SetPositionAuthorityAccount(positionAuthority).
		SetPositionAccount(position).
		SetPositionTokenAccountAccount(positionTokenAccount).
		SetTokenOwnerAccountAAccount(tokenOwnerAccountA).
		SetTokenOwnerAccountBAccount(tokenOwnerAccountB).
		SetTokenVaultAAccount(tokenVaultA).
		SetTokenVaultBAccount(tokenVaultB).
		SetTickArrayLowerAccount(tickArrayLower).
		SetTickArrayUpperAccount(tickArrayUpper)
}