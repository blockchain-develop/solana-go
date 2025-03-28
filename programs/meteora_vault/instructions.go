// Program for vault
// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package meteora_vault

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey = ag_solanago.MustPublicKeyFromBase58("24Uqj9JCLxUeoC3hGfh5W3s9FM9uCHDS2SG3LYwBpyTi")

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "MeteoraVault"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	// initialize new vault
	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	// enable vault
	Instruction_EnableVault = ag_binary.TypeID([8]byte{145, 82, 241, 156, 26, 154, 233, 211})

	// set new operator
	Instruction_SetOperator = ag_binary.TypeID([8]byte{238, 153, 101, 169, 243, 131, 36, 1})

	// Initialize a strategy and add strategy to vault.strategies index
	Instruction_InitializeStrategy = ag_binary.TypeID([8]byte{208, 119, 144, 145, 178, 57, 105, 252})

	// remove a strategy
	Instruction_RemoveStrategy = ag_binary.TypeID([8]byte{185, 238, 33, 91, 134, 210, 97, 26})

	// remove a strategy by advance payment
	Instruction_RemoveStrategy2 = ag_binary.TypeID([8]byte{138, 104, 208, 148, 126, 35, 195, 14})

	// collect token, that someone send wrongly
	// also help in case Mango reimbursement
	Instruction_CollectDust = ag_binary.TypeID([8]byte{246, 149, 21, 82, 160, 74, 254, 240})

	// add a strategy
	Instruction_AddStrategy = ag_binary.TypeID([8]byte{64, 123, 127, 227, 192, 234, 198, 20})

	// deposit liquidity to a strategy
	Instruction_DepositStrategy = ag_binary.TypeID([8]byte{246, 82, 57, 226, 131, 222, 253, 249})

	// withdraw liquidity from a strategy
	Instruction_WithdrawStrategy = ag_binary.TypeID([8]byte{31, 45, 162, 5, 193, 217, 134, 188})

	// Withdraw v2. Withdraw from token vault if no remaining accounts are available. Else, it will attempt to withdraw from strategy and token vault. This method just proxy between 2 methods. Protocol integration should be using withdraw instead of this function.
	Instruction_Withdraw2 = ag_binary.TypeID([8]byte{80, 6, 111, 73, 174, 211, 66, 132})

	// user deposit liquidity to vault
	Instruction_Deposit = ag_binary.TypeID([8]byte{242, 35, 198, 137, 82, 225, 242, 182})

	// user withdraw liquidity from vault
	Instruction_Withdraw = ag_binary.TypeID([8]byte{183, 18, 70, 156, 148, 109, 161, 34})

	// user withdraw liquidity from vault, if vault reserve doesn't have enough liquidity, it will withdraw from the strategy firstly
	Instruction_WithdrawDirectlyFromStrategy = ag_binary.TypeID([8]byte{201, 141, 146, 46, 173, 116, 198, 22})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_EnableVault:
		return "EnableVault"
	case Instruction_SetOperator:
		return "SetOperator"
	case Instruction_InitializeStrategy:
		return "InitializeStrategy"
	case Instruction_RemoveStrategy:
		return "RemoveStrategy"
	case Instruction_RemoveStrategy2:
		return "RemoveStrategy2"
	case Instruction_CollectDust:
		return "CollectDust"
	case Instruction_AddStrategy:
		return "AddStrategy"
	case Instruction_DepositStrategy:
		return "DepositStrategy"
	case Instruction_WithdrawStrategy:
		return "WithdrawStrategy"
	case Instruction_Withdraw2:
		return "Withdraw2"
	case Instruction_Deposit:
		return "Deposit"
	case Instruction_Withdraw:
		return "Withdraw"
	case Instruction_WithdrawDirectlyFromStrategy:
		return "WithdrawDirectlyFromStrategy"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"initialize", (*Initialize)(nil),
		},
		{
			"enable_vault", (*EnableVault)(nil),
		},
		{
			"set_operator", (*SetOperator)(nil),
		},
		{
			"initialize_strategy", (*InitializeStrategy)(nil),
		},
		{
			"remove_strategy", (*RemoveStrategy)(nil),
		},
		{
			"remove_strategy2", (*RemoveStrategy2)(nil),
		},
		{
			"collect_dust", (*CollectDust)(nil),
		},
		{
			"add_strategy", (*AddStrategy)(nil),
		},
		{
			"deposit_strategy", (*DepositStrategy)(nil),
		},
		{
			"withdraw_strategy", (*WithdrawStrategy)(nil),
		},
		{
			"withdraw2", (*Withdraw2)(nil),
		},
		{
			"deposit", (*Deposit)(nil),
		},
		{
			"withdraw", (*Withdraw)(nil),
		},
		{
			"withdraw_directly_from_strategy", (*WithdrawDirectlyFromStrategy)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
