package pkg

type Project struct {
	ID          string `json:"id"`
	Namespace   string `json:"namespace"`
	Name        string `json:"name"`
	VaultTopDir string `json:"vaultTopDir"`
}

type UserInfo struct {
	Email    string    `json:"email,omitempty"`
	ID       string    `json:"idep,omitempty"`
	Name     string    `json:"nomComplet,omitempty"`
	IP       string    `json:"ip,omitempty"`
	Groups   []string  `json:"groups,omitempty"`
	Projects []Project `json:"projects"`
}
