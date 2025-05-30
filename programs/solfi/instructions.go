// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package solfi

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey = ag_solanago.MustPublicKeyFromBase58("SoLFiHG9TfgtdUXUjWAxi3LtvYuFyDLVhBWxdMZxyCe")

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "Solfi"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

const (
	Instruction_CreatePair uint8 = iota

	Instruction_CreatePairV2

	Instruction_UpdateConcentration

	Instruction_UpdateVersion

	Instruction_UpdateFeeParams

	Instruction_UpdateOracles

	Instruction_WithdrawFees

	Instruction_Swap
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id uint8) string {
	switch id {
	case Instruction_CreatePair:
		return "CreatePair"
	case Instruction_CreatePairV2:
		return "CreatePairV2"
	case Instruction_UpdateConcentration:
		return "UpdateConcentration"
	case Instruction_UpdateVersion:
		return "UpdateVersion"
	case Instruction_UpdateFeeParams:
		return "UpdateFeeParams"
	case Instruction_UpdateOracles:
		return "UpdateOracles"
	case Instruction_WithdrawFees:
		return "WithdrawFees"
	case Instruction_Swap:
		return "Swap"
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
	ag_binary.Uint8TypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"CreatePair", (*CreatePair)(nil),
		},
		{
			"CreatePairV2", (*CreatePairV2)(nil),
		},
		{
			"UpdateConcentration", (*UpdateConcentration)(nil),
		},
		{
			"UpdateVersion", (*UpdateVersion)(nil),
		},
		{
			"UpdateFeeParams", (*UpdateFeeParams)(nil),
		},
		{
			"UpdateOracles", (*UpdateOracles)(nil),
		},
		{
			"WithdrawFees", (*WithdrawFees)(nil),
		},
		{
			"Swap", (*Swap)(nil),
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
	err := encoder.WriteUint8(inst.TypeID.Uint8())
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
