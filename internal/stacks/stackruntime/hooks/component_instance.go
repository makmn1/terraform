package hooks

import (
	"github.com/hashicorp/terraform/internal/rpcapi/terraform1"
)

// ComponentInstanceStatus is a UI-focused description of the overall status
// for a given component instance undergoing a Terraform plan or apply
// operation. The "pending" and "errored" status are used for both operation
// types, and the others will be used only for one of plan or apply.
type ComponentInstanceStatus rune

//go:generate go run golang.org/x/tools/cmd/stringer -type=ComponentInstanceStatus component_instance.go

const (
	ComponentInstanceStatusInvalid ComponentInstanceStatus = 0
	ComponentInstancePending       ComponentInstanceStatus = '.'
	ComponentInstancePlanning      ComponentInstanceStatus = 'p'
	ComponentInstancePlanned       ComponentInstanceStatus = 'P'
	ComponentInstanceApplying      ComponentInstanceStatus = 'a'
	ComponentInstanceApplied       ComponentInstanceStatus = 'A'
	ComponentInstanceErrored       ComponentInstanceStatus = 'E'
)

// TODO: move this into the rpcapi package somewhere
func (s ComponentInstanceStatus) ForProtobuf() terraform1.ComponentInstanceStatus_Status {
	switch s {
	case ComponentInstancePending:
		return terraform1.ComponentInstanceStatus_PENDING
	case ComponentInstancePlanning:
		return terraform1.ComponentInstanceStatus_PLANNING
	case ComponentInstancePlanned:
		return terraform1.ComponentInstanceStatus_PLANNED
	case ComponentInstanceApplying:
		return terraform1.ComponentInstanceStatus_APPLYING
	case ComponentInstanceApplied:
		return terraform1.ComponentInstanceStatus_APPLIED
	case ComponentInstanceErrored:
		return terraform1.ComponentInstanceStatus_ERRORED
	default:
		return terraform1.ComponentInstanceStatus_INVALID
	}
}
