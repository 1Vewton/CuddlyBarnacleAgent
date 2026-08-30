package workflows

import (
	"context"
)

// Workflow defines interface for workflow
type Workflow interface {
	Execute(ctx context.Context) error
}
