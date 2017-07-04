// Code generated by cdpgen. DO NOT EDIT.

package database

import (
	"github.com/mafredri/cdp/rpcc"
)

// AddDatabaseClient receives AddDatabase events.
type AddDatabaseClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*AddDatabaseReply, error)
	rpcc.Stream
}

// AddDatabaseReply
type AddDatabaseReply struct {
	Database Database `json:"database"` //
}
