// Code generated by cdpgen. DO NOT EDIT.

package heapprofiler

import (
	"github.com/mafredri/cdp/rpcc"
)

// AddHeapSnapshotChunkClient is a client for AddHeapSnapshotChunk events.
type AddHeapSnapshotChunkClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*AddHeapSnapshotChunkReply, error)
	rpcc.Stream
}

// AddHeapSnapshotChunkReply is the reply for AddHeapSnapshotChunk events.
type AddHeapSnapshotChunkReply struct {
	Chunk string `json:"chunk"` //
}

// ResetProfilesClient is a client for ResetProfiles events.
type ResetProfilesClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ResetProfilesReply, error)
	rpcc.Stream
}

// ResetProfilesReply is the reply for ResetProfiles events.
type ResetProfilesReply struct{}

// ReportHeapSnapshotProgressClient is a client for ReportHeapSnapshotProgress events.
type ReportHeapSnapshotProgressClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ReportHeapSnapshotProgressReply, error)
	rpcc.Stream
}

// ReportHeapSnapshotProgressReply is the reply for ReportHeapSnapshotProgress events.
type ReportHeapSnapshotProgressReply struct {
	Done     int   `json:"done"`               //
	Total    int   `json:"total"`              //
	Finished *bool `json:"finished,omitempty"` //
}

// LastSeenObjectIDClient is a client for LastSeenObjectID events. If heap objects tracking has been started then backend regularly sends a current value for last seen object id and corresponding timestamp. If the were changes in the heap since last event then one or more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
type LastSeenObjectIDClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LastSeenObjectIDReply, error)
	rpcc.Stream
}

// LastSeenObjectIDReply is the reply for LastSeenObjectID events.
type LastSeenObjectIDReply struct {
	LastSeenObjectID int     `json:"lastSeenObjectId"` //
	Timestamp        float64 `json:"timestamp"`        //
}

// HeapStatsUpdateClient is a client for HeapStatsUpdate events. If heap objects tracking has been started then backend may send update for one or more fragments
type HeapStatsUpdateClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*HeapStatsUpdateReply, error)
	rpcc.Stream
}

// HeapStatsUpdateReply is the reply for HeapStatsUpdate events.
type HeapStatsUpdateReply struct {
	StatsUpdate []int `json:"statsUpdate"` // An array of triplets. Each triplet describes a fragment. The first integer is the fragment index, the second integer is a total count of objects for the fragment, the third integer is a total size of the objects for the fragment.
}
