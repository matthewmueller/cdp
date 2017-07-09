// Code generated by cdpgen. DO NOT EDIT.

package domstorage

import (
	"github.com/mafredri/cdp/rpcc"
)

// ItemsClearedClient receives DOMStorageItemsCleared events.
type ItemsClearedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ItemsClearedReply, error)
	rpcc.Stream
}

// ItemsClearedReply
type ItemsClearedReply struct {
	StorageID StorageID `json:"storageId"` //
}

// ItemRemovedClient receives DOMStorageItemRemoved events.
type ItemRemovedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ItemRemovedReply, error)
	rpcc.Stream
}

// ItemRemovedReply
type ItemRemovedReply struct {
	StorageID StorageID `json:"storageId"` //
	Key       string    `json:"key"`       //
}

// ItemAddedClient receives DOMStorageItemAdded events.
type ItemAddedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ItemAddedReply, error)
	rpcc.Stream
}

// ItemAddedReply
type ItemAddedReply struct {
	StorageID StorageID `json:"storageId"` //
	Key       string    `json:"key"`       //
	NewValue  string    `json:"newValue"`  //
}

// ItemUpdatedClient receives DOMStorageItemUpdated events.
type ItemUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ItemUpdatedReply, error)
	rpcc.Stream
}

// ItemUpdatedReply
type ItemUpdatedReply struct {
	StorageID StorageID `json:"storageId"` //
	Key       string    `json:"key"`       //
	OldValue  string    `json:"oldValue"`  //
	NewValue  string    `json:"newValue"`  //
}