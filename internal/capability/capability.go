package capability

// Capability describes whether a backend feature can be used safely.
type Capability struct {
	ID        string `json:"id"`
	Supported bool   `json:"supported"`
	Reason    string `json:"reason,omitempty"`
}
