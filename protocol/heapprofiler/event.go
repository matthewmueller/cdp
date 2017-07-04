// Code generated by cdpgen. DO NOT EDIT.

package heapprofiler

import (
	"github.com/mafredri/cdp/protocol"
	"github.com/mafredri/cdp/rpcc"
)

// AddHeapSnapshotChunkClient receives AddHeapSnapshotChunk events.
type AddHeapSnapshotChunkClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*AddHeapSnapshotChunkReply, error)
	rpcc.Stream
}

// AddHeapSnapshotChunkReply
type AddHeapSnapshotChunkReply struct {
	Chunk string `json:"chunk"` //
}

// ResetProfilesClient receives ResetProfiles events.
type ResetProfilesClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ResetProfilesReply, error)
	rpcc.Stream
}

// ResetProfilesReply
type ResetProfilesReply struct{}

// ReportHeapSnapshotProgressClient receives ReportHeapSnapshotProgress events.
type ReportHeapSnapshotProgressClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ReportHeapSnapshotProgressReply, error)
	rpcc.Stream
}

// ReportHeapSnapshotProgressReply
type ReportHeapSnapshotProgressReply struct {
	Done     int   `json:"done"`               //
	Total    int   `json:"total"`              //
	Finished *bool `json:"finished,omitempty"` //
}

// LastSeenObjectIDClient receives LastSeenObjectID events.
type LastSeenObjectIDClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LastSeenObjectIDReply, error)
	rpcc.Stream
}

// LastSeenObjectIDReply if heap objects tracking has been started then backend regularly sends a current value for last seen object id and corresponding timestamp. If the were changes in the heap since last event then one or more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
type LastSeenObjectIDReply struct {
	LastSeenObjectID int                `json:"lastSeenObjectId"` //
	Timestamp        protocol.Timestamp `json:"timestamp"`        //
}

// HeapStatsUpdateClient receives HeapStatsUpdate events.
type HeapStatsUpdateClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*HeapStatsUpdateReply, error)
	rpcc.Stream
}

// HeapStatsUpdateReply if heap objects tracking has been started then backend may send update for one or more fragments
type HeapStatsUpdateReply struct {
	StatsUpdate []int `json:"statsUpdate"` // An array of triplets. Each triplet describes a fragment. The first integer is the fragment index, the second integer is a total count of objects for the fragment, the third integer is a total size of the objects for the fragment.
}
