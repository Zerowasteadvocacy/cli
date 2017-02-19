package mgmt

type Manifest struct {
	GroupId      string         `json:"groupId"`
	BundleId     string         `json:"bundleId"`
	Name         string         `json:"name"`
	Version      string         `json:"version"`
	Description  string         `json:"description"`
	Dependencies []Dependencies `json:"dependencies"`
}

type Dependencies struct {
	GroupId  string `json:"groupId"`
	BundleId string `json:"bundleId"`
	Version  string `json:"version"`
}
