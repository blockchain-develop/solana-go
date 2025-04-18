// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package raydium_cp

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Swap the tokens in the pool base output amount
//
// # Arguments
//
// * `ctx`- The context of accounts
// * `max_amount_in` -  input amount prevents excessive slippage
// * `amount_out` -  amount of output token
//
type SwapBaseOutput struct {
	MaxAmountIn *uint64
	AmountOut   *uint64

	// [0] = [SIGNER] payer
	// ··········· The user performing the swap
	//
	// [1] = [] authority
	//
	// [2] = [] ammConfig
	// ··········· The factory state to read protocol fees
	//
	// [3] = [WRITE] poolState
	// ··········· The program account of the pool in which the swap will be performed
	//
	// [4] = [WRITE] inputTokenAccount
	// ··········· The user token account for input token
	//
	// [5] = [WRITE] outputTokenAccount
	// ··········· The user token account for output token
	//
	// [6] = [WRITE] inputVault
	// ··········· The vault token account for input token
	//
	// [7] = [WRITE] outputVault
	// ··········· The vault token account for output token
	//
	// [8] = [] inputTokenProgram
	// ··········· SPL program for input token transfers
	//
	// [9] = [] outputTokenProgram
	// ··········· SPL program for output token transfers
	//
	// [10] = [] inputTokenMint
	// ··········· The mint of input token
	//
	// [11] = [] outputTokenMint
	// ··········· The mint of output token
	//
	// [12] = [WRITE] observationState
	// ··········· The program account for the most recent oracle observation
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSwapBaseOutputInstructionBuilder creates a new `SwapBaseOutput` instruction builder.
func NewSwapBaseOutputInstructionBuilder() *SwapBaseOutput {
	nd := &SwapBaseOutput{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetMaxAmountIn sets the "maxAmountIn" parameter.
func (inst *SwapBaseOutput) SetMaxAmountIn(maxAmountIn uint64) *SwapBaseOutput {
	inst.MaxAmountIn = &maxAmountIn
	return inst
}

// SetAmountOut sets the "amountOut" parameter.
func (inst *SwapBaseOutput) SetAmountOut(amountOut uint64) *SwapBaseOutput {
	inst.AmountOut = &amountOut
	return inst
}

// SetPayerAccount sets the "payer" account.
// The user performing the swap
func (inst *SwapBaseOutput) SetPayerAccount(payer ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// The user performing the swap
func (inst *SwapBaseOutput) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *SwapBaseOutput) SetAuthorityAccount(authority ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(authority)
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *SwapBaseOutput) GetAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAmmConfigAccount sets the "ammConfig" account.
// The factory state to read protocol fees
func (inst *SwapBaseOutput) SetAmmConfigAccount(ammConfig ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(ammConfig)
	return inst
}

// GetAmmConfigAccount gets the "ammConfig" account.
// The factory state to read protocol fees
func (inst *SwapBaseOutput) GetAmmConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPoolStateAccount sets the "poolState" account.
// The program account of the pool in which the swap will be performed
func (inst *SwapBaseOutput) SetPoolStateAccount(poolState ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(poolState).WRITE()
	return inst
}

// GetPoolStateAccount gets the "poolState" account.
// The program account of the pool in which the swap will be performed
func (inst *SwapBaseOutput) GetPoolStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetInputTokenAccountAccount sets the "inputTokenAccount" account.
// The user token account for input token
func (inst *SwapBaseOutput) SetInputTokenAccountAccount(inputTokenAccount ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(inputTokenAccount).WRITE()
	return inst
}

// GetInputTokenAccountAccount gets the "inputTokenAccount" account.
// The user token account for input token
func (inst *SwapBaseOutput) GetInputTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetOutputTokenAccountAccount sets the "outputTokenAccount" account.
// The user token account for output token
func (inst *SwapBaseOutput) SetOutputTokenAccountAccount(outputTokenAccount ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(outputTokenAccount).WRITE()
	return inst
}

// GetOutputTokenAccountAccount gets the "outputTokenAccount" account.
// The user token account for output token
func (inst *SwapBaseOutput) GetOutputTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetInputVaultAccount sets the "inputVault" account.
// The vault token account for input token
func (inst *SwapBaseOutput) SetInputVaultAccount(inputVault ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(inputVault).WRITE()
	return inst
}

// GetInputVaultAccount gets the "inputVault" account.
// The vault token account for input token
func (inst *SwapBaseOutput) GetInputVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetOutputVaultAccount sets the "outputVault" account.
// The vault token account for output token
func (inst *SwapBaseOutput) SetOutputVaultAccount(outputVault ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(outputVault).WRITE()
	return inst
}

// GetOutputVaultAccount gets the "outputVault" account.
// The vault token account for output token
func (inst *SwapBaseOutput) GetOutputVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetInputTokenProgramAccount sets the "inputTokenProgram" account.
// SPL program for input token transfers
func (inst *SwapBaseOutput) SetInputTokenProgramAccount(inputTokenProgram ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(inputTokenProgram)
	return inst
}

// GetInputTokenProgramAccount gets the "inputTokenProgram" account.
// SPL program for input token transfers
func (inst *SwapBaseOutput) GetInputTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetOutputTokenProgramAccount sets the "outputTokenProgram" account.
// SPL program for output token transfers
func (inst *SwapBaseOutput) SetOutputTokenProgramAccount(outputTokenProgram ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(outputTokenProgram)
	return inst
}

// GetOutputTokenProgramAccount gets the "outputTokenProgram" account.
// SPL program for output token transfers
func (inst *SwapBaseOutput) GetOutputTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetInputTokenMintAccount sets the "inputTokenMint" account.
// The mint of input token
func (inst *SwapBaseOutput) SetInputTokenMintAccount(inputTokenMint ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(inputTokenMint)
	return inst
}

// GetInputTokenMintAccount gets the "inputTokenMint" account.
// The mint of input token
func (inst *SwapBaseOutput) GetInputTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetOutputTokenMintAccount sets the "outputTokenMint" account.
// The mint of output token
func (inst *SwapBaseOutput) SetOutputTokenMintAccount(outputTokenMint ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(outputTokenMint)
	return inst
}

// GetOutputTokenMintAccount gets the "outputTokenMint" account.
// The mint of output token
func (inst *SwapBaseOutput) GetOutputTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetObservationStateAccount sets the "observationState" account.
// The program account for the most recent oracle observation
func (inst *SwapBaseOutput) SetObservationStateAccount(observationState ag_solanago.PublicKey) *SwapBaseOutput {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(observationState).WRITE()
	return inst
}

// GetObservationStateAccount gets the "observationState" account.
// The program account for the most recent oracle observation
func (inst *SwapBaseOutput) GetObservationStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

func (inst SwapBaseOutput) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SwapBaseOutput,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SwapBaseOutput) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SwapBaseOutput) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.MaxAmountIn == nil {
			return errors.New("MaxAmountIn parameter is not set")
		}
		if inst.AmountOut == nil {
			return errors.New("AmountOut parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AmmConfig is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PoolState is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.InputTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.OutputTokenAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.InputVault is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.OutputVault is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.InputTokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.OutputTokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.InputTokenMint is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.OutputTokenMint is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.ObservationState is not set")
		}
	}
	return nil
}

func (inst *SwapBaseOutput) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SwapBaseOutput")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("MaxAmountIn", *inst.MaxAmountIn))
						paramsBranch.Child(ag_format.Param("  AmountOut", *inst.AmountOut))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("             payer", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         authority", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         ammConfig", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         poolState", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        inputToken", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("       outputToken", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("        inputVault", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("       outputVault", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta(" inputTokenProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("outputTokenProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("    inputTokenMint", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("   outputTokenMint", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("  observationState", inst.AccountMetaSlice.Get(12)))
					})
				})
		})
}

func (obj SwapBaseOutput) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `MaxAmountIn` param:
	err = encoder.Encode(obj.MaxAmountIn)
	if err != nil {
		return err
	}
	// Serialize `AmountOut` param:
	err = encoder.Encode(obj.AmountOut)
	if err != nil {
		return err
	}
	return nil
}
func (obj *SwapBaseOutput) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `MaxAmountIn`:
	err = decoder.Decode(&obj.MaxAmountIn)
	if err != nil {
		return err
	}
	// Deserialize `AmountOut`:
	err = decoder.Decode(&obj.AmountOut)
	if err != nil {
		return err
	}
	return nil
}

// NewSwapBaseOutputInstruction declares a new SwapBaseOutput instruction with the provided parameters and accounts.
func NewSwapBaseOutputInstruction(
	// Parameters:
	maxAmountIn uint64,
	amountOut uint64,
	// Accounts:
	payer ag_solanago.PublicKey,
	authority ag_solanago.PublicKey,
	ammConfig ag_solanago.PublicKey,
	poolState ag_solanago.PublicKey,
	inputTokenAccount ag_solanago.PublicKey,
	outputTokenAccount ag_solanago.PublicKey,
	inputVault ag_solanago.PublicKey,
	outputVault ag_solanago.PublicKey,
	inputTokenProgram ag_solanago.PublicKey,
	outputTokenProgram ag_solanago.PublicKey,
	inputTokenMint ag_solanago.PublicKey,
	outputTokenMint ag_solanago.PublicKey,
	observationState ag_solanago.PublicKey) *SwapBaseOutput {
	return NewSwapBaseOutputInstructionBuilder().
		SetMaxAmountIn(maxAmountIn).
		SetAmountOut(amountOut).
		SetPayerAccount(payer).
		SetAuthorityAccount(authority).
		SetAmmConfigAccount(ammConfig).
		SetPoolStateAccount(poolState).
		SetInputTokenAccountAccount(inputTokenAccount).
		SetOutputTokenAccountAccount(outputTokenAccount).
		SetInputVaultAccount(inputVault).
		SetOutputVaultAccount(outputVault).
		SetInputTokenProgramAccount(inputTokenProgram).
		SetOutputTokenProgramAccount(outputTokenProgram).
		SetInputTokenMintAccount(inputTokenMint).
		SetOutputTokenMintAccount(outputTokenMint).
		SetObservationStateAccount(observationState)
}
