// Code generated by cdpgen. DO NOT EDIT.

package log

import (
	"github.com/mafredri/cdp/rpcc"
)

// EntryAddedClient receives EntryAdded events.
type EntryAddedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*EntryAddedReply, error)
	rpcc.Stream
}

// EntryAddedReply issued when new message was logged.
type EntryAddedReply struct {
	Entry Entry `json:"entry"` // The entry.
}