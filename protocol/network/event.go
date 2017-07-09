// Code generated by cdpgen. DO NOT EDIT.

package network

import (
	"github.com/mafredri/cdp/rpcc"
)

// ResourceChangedPriorityClient receives ResourceChangedPriority events.
type ResourceChangedPriorityClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ResourceChangedPriorityReply, error)
	rpcc.Stream
}

// ResourceChangedPriorityReply fired when resource loading priority is changed
type ResourceChangedPriorityReply struct {
	RequestID   RequestID        `json:"requestId"`   // Request identifier.
	NewPriority ResourcePriority `json:"newPriority"` // New priority
	Timestamp   MonotonicTime    `json:"timestamp"`   // Timestamp.
}

// RequestWillBeSentClient receives RequestWillBeSent events.
type RequestWillBeSentClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RequestWillBeSentReply, error)
	rpcc.Stream
}

// RequestServedFromCacheReply fired if request ended up loading from cache.
type RequestServedFromCacheReply struct {
	RequestID RequestID `json:"requestId"` // Request identifier.
}

// ResponseReceivedClient receives ResponseReceived events.
type ResponseReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ResponseReceivedReply, error)
	rpcc.Stream
}

// DataReceivedReply fired when data chunk was received over the network.
type DataReceivedReply struct {
	RequestID         RequestID     `json:"requestId"`         // Request identifier.
	Timestamp         MonotonicTime `json:"timestamp"`         // Timestamp.
	DataLength        int           `json:"dataLength"`        // Data chunk length.
	EncodedDataLength int           `json:"encodedDataLength"` // Actual bytes received (might be less than dataLength for compressed encodings).
}

// LoadingFinishedClient receives LoadingFinished events.
type LoadingFinishedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LoadingFinishedReply, error)
	rpcc.Stream
}

// LoadingFinishedReply fired when HTTP request has finished loading.
type LoadingFinishedReply struct {
	RequestID         RequestID     `json:"requestId"`         // Request identifier.
	Timestamp         MonotonicTime `json:"timestamp"`         // Timestamp.
	EncodedDataLength float64       `json:"encodedDataLength"` // Total number of bytes received for this request.
}

// LoadingFailedClient receives LoadingFailed events.
type LoadingFailedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LoadingFailedReply, error)
	rpcc.Stream
}

// WebSocketWillSendHandshakeRequestReply fired when WebSocket is about to initiate handshake.
type WebSocketWillSendHandshakeRequestReply struct {
	RequestID RequestID        `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime    `json:"timestamp"` // Timestamp.
	WallTime  TimeSinceEpoch   `json:"wallTime"`  // UTC Timestamp.
	Request   WebSocketRequest `json:"request"`   // WebSocket request data.
}

// WebSocketHandshakeResponseReceivedClient receives WebSocketHandshakeResponseReceived events.
type WebSocketHandshakeResponseReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketHandshakeResponseReceivedReply, error)
	rpcc.Stream
}

// WebSocketHandshakeResponseReceivedReply fired when WebSocket handshake response becomes available.
type WebSocketHandshakeResponseReceivedReply struct {
	RequestID RequestID         `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime     `json:"timestamp"` // Timestamp.
	Response  WebSocketResponse `json:"response"`  // WebSocket response data.
}

// WebSocketCreatedClient receives WebSocketCreated events.
type WebSocketCreatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketCreatedReply, error)
	rpcc.Stream
}

// WebSocketCreatedReply fired upon WebSocket creation.
type WebSocketCreatedReply struct {
	RequestID RequestID  `json:"requestId"`           // Request identifier.
	URL       string     `json:"url"`                 // WebSocket request URL.
	Initiator *Initiator `json:"initiator,omitempty"` // Request initiator.
}

// WebSocketClosedClient receives WebSocketClosed events.
type WebSocketClosedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketClosedReply, error)
	rpcc.Stream
}

// WebSocketClosedReply fired when WebSocket is closed.
type WebSocketClosedReply struct {
	RequestID RequestID     `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime `json:"timestamp"` // Timestamp.
}

// WebSocketFrameReceivedClient receives WebSocketFrameReceived events.
type WebSocketFrameReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketFrameReceivedReply, error)
	rpcc.Stream
}

// WebSocketFrameReceivedReply fired when WebSocket frame is received.
type WebSocketFrameReceivedReply struct {
	RequestID RequestID      `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime  `json:"timestamp"` // Timestamp.
	Response  WebSocketFrame `json:"response"`  // WebSocket response data.
}

// WebSocketFrameErrorClient receives WebSocketFrameError events.
type WebSocketFrameErrorClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketFrameErrorReply, error)
	rpcc.Stream
}

// WebSocketFrameErrorReply fired when WebSocket frame error occurs.
type WebSocketFrameErrorReply struct {
	RequestID    RequestID     `json:"requestId"`    // Request identifier.
	Timestamp    MonotonicTime `json:"timestamp"`    // Timestamp.
	ErrorMessage string        `json:"errorMessage"` // WebSocket frame error message.
}

// WebSocketFrameSentClient receives WebSocketFrameSent events.
type WebSocketFrameSentClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketFrameSentReply, error)
	rpcc.Stream
}

// WebSocketFrameSentReply fired when WebSocket frame is sent.
type WebSocketFrameSentReply struct {
	RequestID RequestID      `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime  `json:"timestamp"` // Timestamp.
	Response  WebSocketFrame `json:"response"`  // WebSocket response data.
}

// EventSourceMessageReceivedClient receives EventSourceMessageReceived events.
type EventSourceMessageReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*EventSourceMessageReceivedReply, error)
	rpcc.Stream
}

// EventSourceMessageReceivedReply fired when EventSource message is received.
type EventSourceMessageReceivedReply struct {
	RequestID RequestID     `json:"requestId"` // Request identifier.
	Timestamp MonotonicTime `json:"timestamp"` // Timestamp.
	EventName string        `json:"eventName"` // Message type.
	EventID   string        `json:"eventId"`   // Message identifier.
	Data      string        `json:"data"`      // Message content.
}

// RequestInterceptedClient receives RequestIntercepted events.
type RequestInterceptedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RequestInterceptedReply, error)
	rpcc.Stream
}
