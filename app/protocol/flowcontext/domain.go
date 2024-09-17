package flowcontext

import (
	"github.com/astrix-network/astrixd/domain"
)

// Domain returns the Domain object associated to the flow context.
func (f *FlowContext) Domain() domain.Domain {
	return f.domain
}
