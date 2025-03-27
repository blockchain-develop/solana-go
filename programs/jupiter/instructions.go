// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package jupiter

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey = ag_solanago.MustPublicKeyFromBase58("JUP6LkbZbjS1jKKwapdHNy74zcZ3tLUZoi5QNyVTaV4")

func SetProgramID(PublicKey ag_solanago.PublicKey) {
	ProgramID = PublicKey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "Jupiter"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	// route_plan Topologically sorted trade DAG
	Instruction_Route = ag_binary.TypeID([8]byte{229, 23, 203, 151, 122, 227, 173, 42})

	Instruction_RouteWithTokenLedger = ag_binary.TypeID([8]byte{150, 86, 71, 116, 167, 93, 14, 104})

	Instruction_ExactOutRoute = ag_binary.TypeID([8]byte{208, 51, 239, 151, 123, 43, 237, 92})

	// Route by using program owned token accounts and open orders accounts.
	Instruction_SharedAccountsRoute = ag_binary.TypeID([8]byte{193, 32, 155, 51, 65, 214, 156, 129})

	Instruction_SharedAccountsRouteWithTokenLedger = ag_binary.TypeID([8]byte{230, 121, 143, 80, 119, 159, 106, 170})

	// Route by using program owned token accounts and open orders accounts.
	Instruction_SharedAccountsExactOutRoute = ag_binary.TypeID([8]byte{176, 209, 105, 168, 154, 125, 69, 62})

	Instruction_SetTokenLedger = ag_binary.TypeID([8]byte{228, 85, 185, 112, 78, 79, 77, 2})

	Instruction_CreateOpenOrders = ag_binary.TypeID([8]byte{229, 194, 212, 172, 8, 10, 134, 147})

	Instruction_CreateTokenAccount = ag_binary.TypeID([8]byte{147, 241, 123, 100, 244, 132, 174, 118})

	Instruction_CreateProgramOpenOrders = ag_binary.TypeID([8]byte{28, 226, 32, 148, 188, 136, 113, 171})

	Instruction_Claim = ag_binary.TypeID([8]byte{62, 198, 214, 193, 213, 159, 108, 210})

	Instruction_ClaimToken = ag_binary.TypeID([8]byte{116, 206, 27, 191, 166, 19, 0, 73})

	Instruction_CreateTokenLedger = ag_binary.TypeID([8]byte{232, 242, 197, 253, 240, 143, 129, 52})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_Route:
		return "Route"
	case Instruction_RouteWithTokenLedger:
		return "RouteWithTokenLedger"
	case Instruction_ExactOutRoute:
		return "ExactOutRoute"
	case Instruction_SharedAccountsRoute:
		return "SharedAccountsRoute"
	case Instruction_SharedAccountsRouteWithTokenLedger:
		return "SharedAccountsRouteWithTokenLedger"
	case Instruction_SharedAccountsExactOutRoute:
		return "SharedAccountsExactOutRoute"
	case Instruction_SetTokenLedger:
		return "SetTokenLedger"
	case Instruction_CreateOpenOrders:
		return "CreateOpenOrders"
	case Instruction_CreateTokenAccount:
		return "CreateTokenAccount"
	case Instruction_CreateProgramOpenOrders:
		return "CreateProgramOpenOrders"
	case Instruction_Claim:
		return "Claim"
	case Instruction_ClaimToken:
		return "ClaimToken"
	case Instruction_CreateTokenLedger:
		return "CreateTokenLedger"
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
			Name: "route", Type: (*Route)(nil),
		},
		{
			Name: "route_with_token_ledger", Type: (*RouteWithTokenLedger)(nil),
		},
		{
			Name: "exact_out_route", Type: (*ExactOutRoute)(nil),
		},
		{
			Name: "shared_accounts_route", Type: (*SharedAccountsRoute)(nil),
		},
		{
			Name: "shared_accounts_route_with_token_ledger", Type: (*SharedAccountsRouteWithTokenLedger)(nil),
		},
		{
			Name: "shared_accounts_exact_out_route", Type: (*SharedAccountsExactOutRoute)(nil),
		},
		{
			Name: "set_token_ledger", Type: (*SetTokenLedger)(nil),
		},
		{
			Name: "create_open_orders", Type: (*CreateOpenOrders)(nil),
		},
		{
			Name: "create_token_account", Type: (*CreateTokenAccount)(nil),
		},
		{
			Name: "create_program_open_orders", Type: (*CreateProgramOpenOrders)(nil),
		},
		{
			Name: "claim", Type: (*Claim)(nil),
		},
		{
			Name: "claim_token", Type: (*ClaimToken)(nil),
		},
		{
			Name: "create_token_ledger", Type: (*CreateTokenLedger)(nil),
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

func DecodeInstructions(message *ag_solanago.Message) (instructions []*Instruction, err error) {
	for _, ins := range message.Instructions {
		var programID ag_solanago.PublicKey
		if programID, err = message.Program(ins.ProgramIDIndex); err != nil {
			return
		}
		if !programID.Equals(ProgramID) {
			continue
		}
		var accounts []*ag_solanago.AccountMeta
		if accounts, err = ins.ResolveInstructionAccounts(message); err != nil {
			return
		}
		var insDecoded *Instruction
		if insDecoded, err = DecodeInstruction(accounts, ins.Data); err != nil {
			return
		}
		instructions = append(instructions, insDecoded)
	}
	return
}
