// Code generated by cdpgen. DO NOT EDIT.

package serviceworker

import (
	"github.com/mafredri/cdp/rpcc"
)

// WorkerRegistrationUpdatedClient receives WorkerRegistrationUpdated events.
type WorkerRegistrationUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WorkerRegistrationUpdatedReply, error)
	rpcc.Stream
}

// WorkerRegistrationUpdatedReply
type WorkerRegistrationUpdatedReply struct {
	Registrations []Registration `json:"registrations"` //
}

// WorkerVersionUpdatedClient receives WorkerVersionUpdated events.
type WorkerVersionUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WorkerVersionUpdatedReply, error)
	rpcc.Stream
}

// WorkerVersionUpdatedReply
type WorkerVersionUpdatedReply struct {
	Versions []Version `json:"versions"` //
}

// WorkerErrorReportedClient receives WorkerErrorReported events.
type WorkerErrorReportedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WorkerErrorReportedReply, error)
	rpcc.Stream
}

// WorkerErrorReportedReply
type WorkerErrorReportedReply struct {
	ErrorMessage ErrorMessage `json:"errorMessage"` //
}
