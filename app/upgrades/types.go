package upgrades

import (
	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	cronkeeper "github.com/neutron-org/neutron/x/cron/keeper"
	feeburnerkeeper "github.com/neutron-org/neutron/x/feeburner/keeper"
	icqkeeper "github.com/neutron-org/neutron/x/interchainqueries/keeper"
	tokenfactorykeeper "github.com/neutron-org/neutron/x/tokenfactory/keeper"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
// must have written, in order for the state migration to go smoothly.
// An upgrade must implement this struct, and then set it in the app.go.
// The app.go will then define the handler.
type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v7`
	UpgradeName string

	// CreateUpgradeHandler defines the function that creates an upgrade handler
	CreateUpgradeHandler func(*module.Manager, module.Configurator, *UpgradeKeepers) upgradetypes.UpgradeHandler

	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades store.StoreUpgrades
}

type UpgradeKeepers struct {
	// keepers
	IcqKeeper          icqkeeper.Keeper
	CronKeeper         cronkeeper.Keeper
	TokenFactoryKeeper *tokenfactorykeeper.Keeper
	FeeBurnerKeeper    *feeburnerkeeper.Keeper
	SlashingKeeper     slashingkeeper.Keeper
	ParamsKeeper       paramskeeper.Keeper
	// subspaces
	GlobalFeeSubspace paramtypes.Subspace
}
