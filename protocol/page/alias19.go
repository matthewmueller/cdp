// +build go1.9

package page

import (
	"github.com/mafredri/cdp/protocol/internal"
)

// ResourceType Resource type as it was perceived by the rendering engine.
//
// Provided as an alias to prevent circular dependencies.
type ResourceType = internal.PageResourceType

// FrameID Unique frame identifier.
//
// Provided as an alias to prevent circular dependencies.
type FrameID = internal.PageFrameID
