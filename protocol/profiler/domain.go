// Code generated by cdpgen. DO NOT EDIT.

// Package profiler implements the Profiler domain.
package profiler

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Profiler domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Profiler domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Enable invokes the Profiler method.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "Enable", Err: err}
	}
	return
}

// Disable invokes the Profiler method.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "Disable", Err: err}
	}
	return
}

// SetSamplingInterval invokes the Profiler method. Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
func (d *domainClient) SetSamplingInterval(ctx context.Context, args *SetSamplingIntervalArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Profiler.setSamplingInterval", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Profiler.setSamplingInterval", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "SetSamplingInterval", Err: err}
	}
	return
}

// Start invokes the Profiler method.
func (d *domainClient) Start(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.start", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "Start", Err: err}
	}
	return
}

// Stop invokes the Profiler method.
func (d *domainClient) Stop(ctx context.Context) (reply *StopReply, err error) {
	reply = new(StopReply)
	err = rpcc.Invoke(ctx, "Profiler.stop", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "Stop", Err: err}
	}
	return
}

// StartPreciseCoverage invokes the Profiler method. Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code coverage may be incomplete. Enabling prevents running optimized code and resets execution counters.
func (d *domainClient) StartPreciseCoverage(ctx context.Context, args *StartPreciseCoverageArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Profiler.startPreciseCoverage", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Profiler.startPreciseCoverage", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "StartPreciseCoverage", Err: err}
	}
	return
}

// StopPreciseCoverage invokes the Profiler method. Disable precise code coverage. Disabling releases unnecessary execution count records and allows executing optimized code.
func (d *domainClient) StopPreciseCoverage(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Profiler.stopPreciseCoverage", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "StopPreciseCoverage", Err: err}
	}
	return
}

// TakePreciseCoverage invokes the Profiler method. Collect coverage data for the current isolate, and resets execution counters. Precise code coverage needs to have started.
func (d *domainClient) TakePreciseCoverage(ctx context.Context) (reply *TakePreciseCoverageReply, err error) {
	reply = new(TakePreciseCoverageReply)
	err = rpcc.Invoke(ctx, "Profiler.takePreciseCoverage", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "TakePreciseCoverage", Err: err}
	}
	return
}

// GetBestEffortCoverage invokes the Profiler method. Collect coverage data for the current isolate. The coverage data may be incomplete due to garbage collection.
func (d *domainClient) GetBestEffortCoverage(ctx context.Context) (reply *GetBestEffortCoverageReply, err error) {
	reply = new(GetBestEffortCoverageReply)
	err = rpcc.Invoke(ctx, "Profiler.getBestEffortCoverage", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Profiler", Op: "GetBestEffortCoverage", Err: err}
	}
	return
}

func (d *domainClient) ConsoleProfileStarted(ctx context.Context) (ConsoleProfileStartedClient, error) {
	s, err := rpcc.NewStream(ctx, "Profiler.consoleProfileStarted", d.conn)
	if err != nil {
		return nil, err
	}
	return &consoleProfileStartedClient{Stream: s}, nil
}

type consoleProfileStartedClient struct{ rpcc.Stream }

func (c *consoleProfileStartedClient) Recv() (*ConsoleProfileStartedReply, error) {
	event := new(ConsoleProfileStartedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Profiler", Op: "ConsoleProfileStarted Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) ConsoleProfileFinished(ctx context.Context) (ConsoleProfileFinishedClient, error) {
	s, err := rpcc.NewStream(ctx, "Profiler.consoleProfileFinished", d.conn)
	if err != nil {
		return nil, err
	}
	return &consoleProfileFinishedClient{Stream: s}, nil
}

type consoleProfileFinishedClient struct{ rpcc.Stream }

func (c *consoleProfileFinishedClient) Recv() (*ConsoleProfileFinishedReply, error) {
	event := new(ConsoleProfileFinishedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Profiler", Op: "ConsoleProfileFinished Recv", Err: err}
	}
	return event, nil
}
