package adapter

import (
	"context"

	"github.com/Iverysterog1/Linux-Desktop-Customizer/internal/capability"
)

// Adapter is the narrow mutation boundary used by the transaction engine.
// Implementations must make Set/Unset idempotent and must not execute
// arbitrary profile-provided commands.
type Adapter interface {
	Name() string
	Capabilities(context.Context) []capability.Capability
	Read(context.Context, string) (value string, exists bool, err error)
	Set(context.Context, string, string) error
	Unset(context.Context, string) error
}
