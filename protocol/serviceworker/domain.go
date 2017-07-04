// Code generated by cdpgen. DO NOT EDIT.

package serviceworker

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the ServiceWorker domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the ServiceWorker domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Enable invokes the ServiceWorker method.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "ServiceWorker.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "ServiceWorker", Op: "Enable", Err: err}
	}
	return
}

// Disable invokes the ServiceWorker method.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "ServiceWorker.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "ServiceWorker", Op: "Disable", Err: err}
	}
	return
}

// Unregister invokes the ServiceWorker method.
func (d *domainClient) Unregister(ctx context.Context, args *UnregisterArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "ServiceWorker.unregister", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "ServiceWorker.unregister", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "ServiceWorker", Op: "Unregister", Err: err}
	}
	return
}

// UpdateRegistration invokes the ServiceWorker method.
func (d *domainClient) UpdateRegistration(ctx context.Context, args *UpdateRegistrationArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "ServiceWorker.updateRegistration", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "ServiceWorker.updateRegistration", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "ServiceWorker", Op: "UpdateRegistration", Err: err}
	}
	return
}

// StartWorker invokes the ServiceWorker method.
func (d *domainClient) StartWorker(ctx context.Context, args *StartWorkerArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "ServiceWorker.startWorker", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "ServiceWorker.startWorker", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "ServiceWorker", Op: "StartWorker", Err: err}
	}
	return
}

// SkipWaiting invokes the ServiceWorker method.
func (d *domainClient) SkipWaiting(ctx context.Context, args *SkipWaitingArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "ServiceWorker.skipWaiting", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "ServiceWorker.skipWaiting", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "ServiceWorker", Op: "SkipWaiting", Err: err}
	}
	return
}

// StopWorker invokes the ServiceWorker method.
func (d *domainClient) StopWorker(ctx context.Context, args *StopWorkerArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "ServiceWorker.stopWorker", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "ServiceWorker.stopWorker", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "ServiceWorker", Op: "StopWorker", Err: err}
	}
	return
}

// InspectWorker invokes the ServiceWorker method.
func (d *domainClient) InspectWorker(ctx context.Context, args *InspectWorkerArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "ServiceWorker.inspectWorker", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "ServiceWorker.inspectWorker", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "ServiceWorker", Op: "InspectWorker", Err: err}
	}
	return
}

// SetForceUpdateOnPageLoad invokes the ServiceWorker method.
func (d *domainClient) SetForceUpdateOnPageLoad(ctx context.Context, args *SetForceUpdateOnPageLoadArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "ServiceWorker.setForceUpdateOnPageLoad", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "ServiceWorker.setForceUpdateOnPageLoad", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "ServiceWorker", Op: "SetForceUpdateOnPageLoad", Err: err}
	}
	return
}

// DeliverPushMessage invokes the ServiceWorker method.
func (d *domainClient) DeliverPushMessage(ctx context.Context, args *DeliverPushMessageArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "ServiceWorker.deliverPushMessage", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "ServiceWorker.deliverPushMessage", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "ServiceWorker", Op: "DeliverPushMessage", Err: err}
	}
	return
}

// DispatchSyncEvent invokes the ServiceWorker method.
func (d *domainClient) DispatchSyncEvent(ctx context.Context, args *DispatchSyncEventArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "ServiceWorker.dispatchSyncEvent", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "ServiceWorker.dispatchSyncEvent", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "ServiceWorker", Op: "DispatchSyncEvent", Err: err}
	}
	return
}

func (d *domainClient) WorkerRegistrationUpdated(ctx context.Context) (WorkerRegistrationUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "ServiceWorker.workerRegistrationUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &workerRegistrationUpdatedClient{Stream: s}, nil
}

type workerRegistrationUpdatedClient struct{ rpcc.Stream }

func (c *workerRegistrationUpdatedClient) Recv() (*WorkerRegistrationUpdatedReply, error) {
	event := new(WorkerRegistrationUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "ServiceWorker", Op: "WorkerRegistrationUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) WorkerVersionUpdated(ctx context.Context) (WorkerVersionUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "ServiceWorker.workerVersionUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &workerVersionUpdatedClient{Stream: s}, nil
}

type workerVersionUpdatedClient struct{ rpcc.Stream }

func (c *workerVersionUpdatedClient) Recv() (*WorkerVersionUpdatedReply, error) {
	event := new(WorkerVersionUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "ServiceWorker", Op: "WorkerVersionUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) WorkerErrorReported(ctx context.Context) (WorkerErrorReportedClient, error) {
	s, err := rpcc.NewStream(ctx, "ServiceWorker.workerErrorReported", d.conn)
	if err != nil {
		return nil, err
	}
	return &workerErrorReportedClient{Stream: s}, nil
}

type workerErrorReportedClient struct{ rpcc.Stream }

func (c *workerErrorReportedClient) Recv() (*WorkerErrorReportedReply, error) {
	event := new(WorkerErrorReportedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "ServiceWorker", Op: "WorkerErrorReported Recv", Err: err}
	}
	return event, nil
}
