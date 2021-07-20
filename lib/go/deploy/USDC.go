package deploy

import (
	"encoding/hex"

	util "github.com/flow-usdc/flow-usdc"
	"github.com/onflow/cadence"

	"github.com/bjartek/go-with-the-flow/gwtf"
)

func DeployFiatTokenContract(
	g *gwtf.GoWithTheFlow,
	ownerAcct string) (err error) {

	contractCode := util.ParseCadenceTemplate("../../contracts/FiatToken.cdc")
	txFilename := "../../transactions/deploy_contract_with_auth.cdc"
	code := util.ParseCadenceTemplate(txFilename)
	encodedStr := hex.EncodeToString(contractCode)
	g.CreateAccountPrintEvents(
		"vaulted-account",
		"non-vaulted-account",
		"pauser",
		"non-pauser",
		"blocklister",
		"non-blocklister",
		"allowance",
		"non-allowance",
		"minter",
		"non-minter",
		"minterController1",
		"minterController2",
	)

	err = g.TransactionFromFile(txFilename, code).
		SignProposeAndPayAs(ownerAcct).
		StringArgument("FiatToken").
		StringArgument(encodedStr).
		// Vault
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCVault"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCVaultBalance"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCVaultUUID"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCVaultAllowance"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCVaultReceiver"}).
		// Blocklist executor
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCBlocklistExe"}).
		Argument(cadence.Path{Domain: "private", Identifier: "USDCBlocklistExeCap"}).
		// Blocklister
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCBlocklister"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCBlocklisterCapReceiver"}).
		// Pause executor
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCPauseExe"}).
		Argument(cadence.Path{Domain: "private", Identifier: "USDCPauseExeCap"}).
		// Pauser
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCPauser"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCPauserCapReceiver"}).
		// Owner
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCOwner"}).
		Argument(cadence.Path{Domain: "private", Identifier: "USDCOwnerCap"}).
		// Masterminter
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCMasterMinter"}).
		Argument(cadence.Path{Domain: "private", Identifier: "USDCMasterMinterCap"}).
		// Minter Controller
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCMinterController"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCMinterControllerUUID"}).
		// Minter
		Argument(cadence.Path{Domain: "storage", Identifier: "USDCMinter"}).
		Argument(cadence.Path{Domain: "public", Identifier: "USDCMinterUUID"}).
		StringArgument("USDC").
		UFix64Argument("10000.0").
		BooleanArgument(false).
		RunPrintEventsFull()
	return
}