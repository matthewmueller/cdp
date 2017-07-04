// Code generated by cdpgen. DO NOT EDIT.

package input

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Input domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Input domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// SetIgnoreInputEvents invokes the Input method. Ignores input events (useful while auditing page).
func (d *domainClient) SetIgnoreInputEvents(ctx context.Context, args *SetIgnoreInputEventsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.setIgnoreInputEvents", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.setIgnoreInputEvents", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "SetIgnoreInputEvents", Err: err}
	}
	return
}

// DispatchKeyEvent invokes the Input method. Dispatches a key event to the page.
func (d *domainClient) DispatchKeyEvent(ctx context.Context, args *DispatchKeyEventArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.dispatchKeyEvent", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.dispatchKeyEvent", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "DispatchKeyEvent", Err: err}
	}
	return
}

// DispatchMouseEvent invokes the Input method. Dispatches a mouse event to the page.
func (d *domainClient) DispatchMouseEvent(ctx context.Context, args *DispatchMouseEventArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.dispatchMouseEvent", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.dispatchMouseEvent", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "DispatchMouseEvent", Err: err}
	}
	return
}

// DispatchTouchEvent invokes the Input method. Dispatches a touch event to the page.
func (d *domainClient) DispatchTouchEvent(ctx context.Context, args *DispatchTouchEventArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.dispatchTouchEvent", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.dispatchTouchEvent", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "DispatchTouchEvent", Err: err}
	}
	return
}

// EmulateTouchFromMouseEvent invokes the Input method. Emulates touch event from the mouse event parameters.
func (d *domainClient) EmulateTouchFromMouseEvent(ctx context.Context, args *EmulateTouchFromMouseEventArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.emulateTouchFromMouseEvent", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.emulateTouchFromMouseEvent", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "EmulateTouchFromMouseEvent", Err: err}
	}
	return
}

// SynthesizePinchGesture invokes the Input method. Synthesizes a pinch gesture over a time period by issuing appropriate touch events.
func (d *domainClient) SynthesizePinchGesture(ctx context.Context, args *SynthesizePinchGestureArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.synthesizePinchGesture", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.synthesizePinchGesture", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "SynthesizePinchGesture", Err: err}
	}
	return
}

// SynthesizeScrollGesture invokes the Input method. Synthesizes a scroll gesture over a time period by issuing appropriate touch events.
func (d *domainClient) SynthesizeScrollGesture(ctx context.Context, args *SynthesizeScrollGestureArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.synthesizeScrollGesture", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.synthesizeScrollGesture", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "SynthesizeScrollGesture", Err: err}
	}
	return
}

// SynthesizeTapGesture invokes the Input method. Synthesizes a tap gesture over a time period by issuing appropriate touch events.
func (d *domainClient) SynthesizeTapGesture(ctx context.Context, args *SynthesizeTapGestureArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Input.synthesizeTapGesture", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Input.synthesizeTapGesture", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Input", Op: "SynthesizeTapGesture", Err: err}
	}
	return
}
