// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package zeta

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeWhitelistTradingFeesAccount is the `initializeWhitelistTradingFeesAccount` instruction.
type InitializeWhitelistTradingFeesAccount struct {
	Nonce *uint8

	// [0] = [WRITE] whitelistTradingFeesAccount
	//
	// [1] = [WRITE, SIGNER] admin
	//
	// [2] = [] user
	//
	// [3] = [] systemProgram
	//
	// [4] = [] state
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeWhitelistTradingFeesAccountInstructionBuilder creates a new `InitializeWhitelistTradingFeesAccount` instruction builder.
func NewInitializeWhitelistTradingFeesAccountInstructionBuilder() *InitializeWhitelistTradingFeesAccount {
	nd := &InitializeWhitelistTradingFeesAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	return nd
}

// SetNonce sets the "nonce" parameter.
func (inst *InitializeWhitelistTradingFeesAccount) SetNonce(nonce uint8) *InitializeWhitelistTradingFeesAccount {
	inst.Nonce = &nonce
	return inst
}

// SetWhitelistTradingFeesAccountAccount sets the "whitelistTradingFeesAccount" account.
func (inst *InitializeWhitelistTradingFeesAccount) SetWhitelistTradingFeesAccountAccount(whitelistTradingFeesAccount ag_solanago.PublicKey) *InitializeWhitelistTradingFeesAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(whitelistTradingFeesAccount).WRITE()
	return inst
}

// GetWhitelistTradingFeesAccountAccount gets the "whitelistTradingFeesAccount" account.
func (inst *InitializeWhitelistTradingFeesAccount) GetWhitelistTradingFeesAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *InitializeWhitelistTradingFeesAccount) SetAdminAccount(admin ag_solanago.PublicKey) *InitializeWhitelistTradingFeesAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *InitializeWhitelistTradingFeesAccount) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetUserAccount sets the "user" account.
func (inst *InitializeWhitelistTradingFeesAccount) SetUserAccount(user ag_solanago.PublicKey) *InitializeWhitelistTradingFeesAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(user)
	return inst
}

// GetUserAccount gets the "user" account.
func (inst *InitializeWhitelistTradingFeesAccount) GetUserAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializeWhitelistTradingFeesAccount) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeWhitelistTradingFeesAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializeWhitelistTradingFeesAccount) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetStateAccount sets the "state" account.
func (inst *InitializeWhitelistTradingFeesAccount) SetStateAccount(state ag_solanago.PublicKey) *InitializeWhitelistTradingFeesAccount {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *InitializeWhitelistTradingFeesAccount) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst InitializeWhitelistTradingFeesAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeWhitelistTradingFeesAccount,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeWhitelistTradingFeesAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeWhitelistTradingFeesAccount) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Nonce == nil {
			return errors.New("Nonce parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.WhitelistTradingFeesAccount is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.User is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.State is not set")
		}
	}
	return nil
}

func (inst *InitializeWhitelistTradingFeesAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeWhitelistTradingFeesAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Nonce", *inst.Nonce))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("whitelistTradingFees", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("               admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                user", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("       systemProgram", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("               state", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj InitializeWhitelistTradingFeesAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Nonce` param:
	err = encoder.Encode(obj.Nonce)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeWhitelistTradingFeesAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Nonce`:
	err = decoder.Decode(&obj.Nonce)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeWhitelistTradingFeesAccountInstruction declares a new InitializeWhitelistTradingFeesAccount instruction with the provided parameters and accounts.
func NewInitializeWhitelistTradingFeesAccountInstruction(
	// Parameters:
	nonce uint8,
	// Accounts:
	whitelistTradingFeesAccount ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	user ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	state ag_solanago.PublicKey) *InitializeWhitelistTradingFeesAccount {
	return NewInitializeWhitelistTradingFeesAccountInstructionBuilder().
		SetNonce(nonce).
		SetWhitelistTradingFeesAccountAccount(whitelistTradingFeesAccount).
		SetAdminAccount(admin).
		SetUserAccount(user).
		SetSystemProgramAccount(systemProgram).
		SetStateAccount(state)
}
