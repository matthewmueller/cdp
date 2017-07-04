// Code generated by cdpgen. DO NOT EDIT.

package debugger

import (
	"encoding/json"

	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/rpcc"
)

// ScriptParsedClient receives ScriptParsed events.
type ScriptParsedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ScriptParsedReply, error)
	rpcc.Stream
}

// ScriptParsedReply fired when virtual machine parses script. This event is also fired for all known and uncollected scripts upon enabling debugger.
type ScriptParsedReply struct {
	ScriptID                runtime.ScriptID           `json:"scriptId"`                          // Identifier of the script parsed.
	URL                     string                     `json:"url"`                               // URL or name of the script parsed (if any).
	StartLine               int                        `json:"startLine"`                         // Line offset of the script within the resource with given URL (for script tags).
	StartColumn             int                        `json:"startColumn"`                       // Column offset of the script within the resource with given URL.
	EndLine                 int                        `json:"endLine"`                           // Last line of the script.
	EndColumn               int                        `json:"endColumn"`                         // Length of the last line of the script.
	ExecutionContextID      runtime.ExecutionContextID `json:"executionContextId"`                // Specifies script creation context.
	Hash                    string                     `json:"hash"`                              // Content hash of the script.
	ExecutionContextAuxData json.RawMessage            `json:"executionContextAuxData,omitempty"` // Embedder-specific auxiliary data.
	IsLiveEdit              *bool                      `json:"isLiveEdit,omitempty"`              // True, if this script is generated as a result of the live edit operation.
	SourceMapURL            *string                    `json:"sourceMapURL,omitempty"`            // URL of source map associated with script (if any).
	HasSourceURL            *bool                      `json:"hasSourceURL,omitempty"`            // True, if this script has sourceURL.
	IsModule                *bool                      `json:"isModule,omitempty"`                // True, if this script is ES6 module.
	Length                  *int                       `json:"length,omitempty"`                  // This script length.
	StackTrace              *runtime.StackTrace        `json:"stackTrace,omitempty"`              // JavaScript top stack frame of where the script parsed event was triggered if available.
}

// ScriptFailedToParseClient receives ScriptFailedToParse events.
type ScriptFailedToParseClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ScriptFailedToParseReply, error)
	rpcc.Stream
}

// ScriptFailedToParseReply fired when virtual machine fails to parse the script.
type ScriptFailedToParseReply struct {
	ScriptID                runtime.ScriptID           `json:"scriptId"`                          // Identifier of the script parsed.
	URL                     string                     `json:"url"`                               // URL or name of the script parsed (if any).
	StartLine               int                        `json:"startLine"`                         // Line offset of the script within the resource with given URL (for script tags).
	StartColumn             int                        `json:"startColumn"`                       // Column offset of the script within the resource with given URL.
	EndLine                 int                        `json:"endLine"`                           // Last line of the script.
	EndColumn               int                        `json:"endColumn"`                         // Length of the last line of the script.
	ExecutionContextID      runtime.ExecutionContextID `json:"executionContextId"`                // Specifies script creation context.
	Hash                    string                     `json:"hash"`                              // Content hash of the script.
	ExecutionContextAuxData json.RawMessage            `json:"executionContextAuxData,omitempty"` // Embedder-specific auxiliary data.
	SourceMapURL            *string                    `json:"sourceMapURL,omitempty"`            // URL of source map associated with script (if any).
	HasSourceURL            *bool                      `json:"hasSourceURL,omitempty"`            // True, if this script has sourceURL.
	IsModule                *bool                      `json:"isModule,omitempty"`                // True, if this script is ES6 module.
	Length                  *int                       `json:"length,omitempty"`                  // This script length.
	StackTrace              *runtime.StackTrace        `json:"stackTrace,omitempty"`              // JavaScript top stack frame of where the script parsed event was triggered if available.
}

// BreakpointResolvedClient receives BreakpointResolved events.
type BreakpointResolvedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*BreakpointResolvedReply, error)
	rpcc.Stream
}

// BreakpointResolvedReply fired when breakpoint is resolved to an actual script and location.
type BreakpointResolvedReply struct {
	BreakpointID BreakpointID `json:"breakpointId"` // Breakpoint unique identifier.
	Location     Location     `json:"location"`     // Actual breakpoint location.
}

// PausedClient receives Paused events.
type PausedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*PausedReply, error)
	rpcc.Stream
}

// PausedReply fired when the virtual machine stopped on breakpoint or exception or any other stop criteria.
type PausedReply struct {
	CallFrames      []CallFrame         `json:"callFrames"`                // Call stack the virtual machine stopped on.
	Reason          string              `json:"reason"`                    // Pause reason.
	Data            json.RawMessage     `json:"data,omitempty"`            // Object containing break-specific auxiliary properties.
	HitBreakpoints  []string            `json:"hitBreakpoints,omitempty"`  // Hit breakpoints IDs
	AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"` // Async stack trace, if any.
}

// ResumedClient receives Resumed events.
type ResumedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ResumedReply, error)
	rpcc.Stream
}

// ResumedReply fired when the virtual machine resumed execution.
type ResumedReply struct{}
