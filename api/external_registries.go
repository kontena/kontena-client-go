package api

type ExternalRegistryIndex struct {
	ExternalRegistries []ExternalRegistry `json:"external_registries"`
}

type ExternalRegistry struct {
	ID       string
	Name     string
	URL      string
	Username string
	Email    string
}

type ExternalRegistryPOST struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}
