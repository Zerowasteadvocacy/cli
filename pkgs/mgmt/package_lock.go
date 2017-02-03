package mgmt

type Lock struct {
	LockVersion       string                  `json: "lockfile_version"`
	PackageName       string                  `json: "package_name"`
	PackageVersion    string                  `json: "version"`
	Sources           string                  `json:"sources"`
	ContractTypes     map[string]ContractType `json:"contract_types"`
	BuildDependencies map[string]string       `json:"build_dependencies"`
	//Deployments       map[string]map[string]ContractInstance `json:"deployments"`
}

type ContractType struct {
	Name            string `json:"contract_name"`
	Bytecode        string `json:"bytecode"`
	RuntimeBytecode string `json:"runtime_bytecode"`
	Abi             string `json:"abi"`
	Natspec         string `json:"natspec"`
	//Compiler        Compiler `json:"compiler"`
}

/*type ContractInstance struct {
	Type            string      `json:"contract_type"`
	Address         string      `json:"address"`
	Transaction     string      `json:"transaction"`
	Block           string      `json:"block"`
	RuntimeBytecode string      `json:"runtime_bytecode"`
	Links           []LinkValue `json:"link_dependencies"`
	//Compiler        Compiler    `json:"compiler"`
}

type LinkValue struct {
	Offset uint   `json:"offset"`
	Value  string `json:"value"`
}
*/
