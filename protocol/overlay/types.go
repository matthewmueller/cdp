// Code generated by cdpgen. DO NOT EDIT.

package overlay

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mafredri/cdp/protocol/dom"
)

// HighlightConfig Configuration data for the highlighting of page elements.
type HighlightConfig struct {
	ShowInfo           *bool     `json:"showInfo,omitempty"`           // Whether the node info tooltip should be shown (default: false).
	ShowRulers         *bool     `json:"showRulers,omitempty"`         // Whether the rulers should be shown (default: false).
	ShowExtensionLines *bool     `json:"showExtensionLines,omitempty"` // Whether the extension lines from node to the rulers should be shown (default: false).
	DisplayAsMaterial  *bool     `json:"displayAsMaterial,omitempty"`  //
	ContentColor       *dom.RGBA `json:"contentColor,omitempty"`       // The content box highlight fill color (default: transparent).
	PaddingColor       *dom.RGBA `json:"paddingColor,omitempty"`       // The padding highlight fill color (default: transparent).
	BorderColor        *dom.RGBA `json:"borderColor,omitempty"`        // The border highlight fill color (default: transparent).
	MarginColor        *dom.RGBA `json:"marginColor,omitempty"`        // The margin highlight fill color (default: transparent).
	EventTargetColor   *dom.RGBA `json:"eventTargetColor,omitempty"`   // The event target element highlight fill color (default: transparent).
	ShapeColor         *dom.RGBA `json:"shapeColor,omitempty"`         // The shape outside fill color (default: transparent).
	ShapeMarginColor   *dom.RGBA `json:"shapeMarginColor,omitempty"`   // The shape margin fill color (default: transparent).
	SelectorList       *string   `json:"selectorList,omitempty"`       // Selectors to highlight relevant nodes.
}

// InspectMode
type InspectMode int

// InspectMode as enums.
const (
	InspectModeNotSet InspectMode = iota
	InspectModeSearchForNode
	InspectModeSearchForUAShadowDOM
	InspectModeNone
)

// Valid returns true if enum is set.
func (e InspectMode) Valid() bool {
	return e >= 1 && e <= 3
}

func (e InspectMode) String() string {
	switch e {
	case 0:
		return "InspectModeNotSet"
	case 1:
		return "searchForNode"
	case 2:
		return "searchForUAShadowDOM"
	case 3:
		return "none"
	}
	return fmt.Sprintf("InspectMode(%d)", e)
}

// MarshalJSON encodes enum into a string or null when not set.
func (e InspectMode) MarshalJSON() ([]byte, error) {
	if e == 0 {
		return []byte("null"), nil
	}
	if !e.Valid() {
		return nil, errors.New("overlay.InspectMode: MarshalJSON on bad enum value: " + e.String())
	}
	return json.Marshal(e.String())
}

// UnmarshalJSON decodes a string value into a enum.
func (e *InspectMode) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null":
		*e = 0
	case "\"searchForNode\"":
		*e = 1
	case "\"searchForUAShadowDOM\"":
		*e = 2
	case "\"none\"":
		*e = 3
	default:
		return fmt.Errorf("overlay.InspectMode: UnmarshalJSON on bad input: %s", data)
	}
	return nil
}
