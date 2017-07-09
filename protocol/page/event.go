// Code generated by cdpgen. DO NOT EDIT.

package page

import (
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/rpcc"
)

// DOMContentEventFiredClient receives DOMContentEventFired events.
type DOMContentEventFiredClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*DOMContentEventFiredReply, error)
	rpcc.Stream
}

// DOMContentEventFiredReply
type DOMContentEventFiredReply struct {
	Timestamp network.MonotonicTime `json:"timestamp"` //
}

// LoadEventFiredClient receives LoadEventFired events.
type LoadEventFiredClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LoadEventFiredReply, error)
	rpcc.Stream
}

// LoadEventFiredReply
type LoadEventFiredReply struct {
	Timestamp network.MonotonicTime `json:"timestamp"` //
}

// FrameAttachedClient receives FrameAttached events.
type FrameAttachedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameAttachedReply, error)
	rpcc.Stream
}

// FrameNavigatedReply fired once navigation of the frame has completed. Frame is now associated with the new loader.
type FrameNavigatedReply struct {
	Frame Frame `json:"frame"` // Frame object.
}

// FrameDetachedClient receives FrameDetached events.
type FrameDetachedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameDetachedReply, error)
	rpcc.Stream
}

// FrameResizedReply
type FrameResizedReply struct{}

// JavascriptDialogOpeningClient receives JavascriptDialogOpening events.
type JavascriptDialogOpeningClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*JavascriptDialogOpeningReply, error)
	rpcc.Stream
}

// JavascriptDialogOpeningReply fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) is about to open.
type JavascriptDialogOpeningReply struct {
	Message string     `json:"message"` // Message that will be displayed by the dialog.
	Type    DialogType `json:"type"`    // Dialog type.
}

// JavascriptDialogClosedClient receives JavascriptDialogClosed events.
type JavascriptDialogClosedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*JavascriptDialogClosedReply, error)
	rpcc.Stream
}

// JavascriptDialogClosedReply fired when a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload) has been closed.
type JavascriptDialogClosedReply struct {
	Result bool `json:"result"` // Whether dialog was confirmed.
}

// ScreencastFrameClient receives ScreencastFrame events.
type ScreencastFrameClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ScreencastFrameReply, error)
	rpcc.Stream
}

// ScreencastFrameReply compressed image data requested by the startScreencast.
type ScreencastFrameReply struct {
	Data      []byte                  `json:"data"`      // Base64-encoded compressed image.
	Metadata  ScreencastFrameMetadata `json:"metadata"`  // Screencast frame metadata.
	SessionID int                     `json:"sessionId"` // Frame number.
}

// ScreencastVisibilityChangedClient receives ScreencastVisibilityChanged events.
type ScreencastVisibilityChangedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ScreencastVisibilityChangedReply, error)
	rpcc.Stream
}

// ScreencastVisibilityChangedReply fired when the page with currently enabled screencast was shown or hidden .
type ScreencastVisibilityChangedReply struct {
	Visible bool `json:"visible"` // True if the page is visible.
}

// InterstitialShownClient receives InterstitialShown events.
type InterstitialShownClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*InterstitialShownReply, error)
	rpcc.Stream
}

// InterstitialShownReply fired when interstitial page was shown
type InterstitialShownReply struct{}

// InterstitialHiddenClient receives InterstitialHidden events.
type InterstitialHiddenClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*InterstitialHiddenReply, error)
	rpcc.Stream
}

// InterstitialHiddenReply fired when interstitial page was hidden
type InterstitialHiddenReply struct{}

// NavigationRequestedClient receives NavigationRequested events.
type NavigationRequestedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*NavigationRequestedReply, error)
	rpcc.Stream
}

// NavigationRequestedReply fired when a navigation is started if navigation throttles are enabled.  The navigation will be deferred until processNavigation is called.
type NavigationRequestedReply struct {
	IsInMainFrame bool   `json:"isInMainFrame"` // Whether the navigation is taking place in the main frame or in a subframe.
	IsRedirect    bool   `json:"isRedirect"`    // Whether the navigation has encountered a server redirect or not.
	NavigationID  int    `json:"navigationId"`  //
	URL           string `json:"url"`           // URL of requested navigation.
}
