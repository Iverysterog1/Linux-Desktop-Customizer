package transaction

import "time"

type Action string

const (
	ActionSet   Action = "set"
	ActionUnset Action = "unset"
)

type Change struct {
	Adapter string `json:"adapter"`
	Key     string `json:"key"`
	Action  Action `json:"action"`
	Value   string `json:"value,omitempty"`
}

type Plan struct {
	Changes []Change `json:"changes"`
}

type Snapshot struct {
	Exists bool   `json:"exists"`
	Value  string `json:"value,omitempty"`
}

type PreviewChange struct {
	Index     int      `json:"index"`
	Change    Change   `json:"change"`
	Before    Snapshot `json:"before"`
	After     Snapshot `json:"after"`
	Noop      bool     `json:"noop"`
	Supported bool     `json:"supported"`
	Reason    string   `json:"reason,omitempty"`
}

type Preview struct {
	Supported bool            `json:"supported"`
	Changes   []PreviewChange `json:"changes"`
}

type Status string

const (
	StatusApplying       Status = "applying"
	StatusApplied        Status = "applied"
	StatusRolledBack     Status = "rolled_back"
	StatusRollbackFailed Status = "rollback_failed"
)

type RecordedChange struct {
	Change Change   `json:"change"`
	Before Snapshot `json:"before"`
	Noop   bool     `json:"noop"`
}

type Transaction struct {
	ID        string           `json:"id"`
	CreatedAt time.Time        `json:"created_at"`
	Status    Status           `json:"status"`
	Changes   []RecordedChange `json:"changes"`
	Error     string           `json:"error,omitempty"`
}
