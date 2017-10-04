package v20170930

// PoolUpgradeProfile contains pool properties:
//  - orchestrator type and version
//  - pool name (for agent pool)
//  - OS type of the VMs in the pool
//  - list of applicable upgrades
type PoolUpgradeProfile struct {
	OrchestratorProfile
	Name     string                 `json:"name,omitempty"`
	OSType   OSType                 `json:"osType,omitempty"`
	Upgrades []*OrchestratorProfile `json:"upgrades,omitempty"`
}

// UpgradeProfile contains cluster properties:
//  - orchestrator type and version for the cluster
//  - list of pool profiles, constituting the cluster
type UpgradeProfile struct {
	ControlPlaneProfile *PoolUpgradeProfile   `json:"controlPlaneProfile"`
	AgentPoolProfiles   []*PoolUpgradeProfile `json:"agentPoolProfiles"`
}
