// Code generated by cdpgen. DO NOT EDIT.

package page

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Page domain. Actions and events related to the inspected page belong to the page domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Page domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Enable invokes the Page method. Enables page domain notifications.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Page.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "Enable", Err: err}
	}
	return
}

// Disable invokes the Page method. Disables page domain notifications.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Page.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "Disable", Err: err}
	}
	return
}

// AddScriptToEvaluateOnLoad invokes the Page method.
func (d *domainClient) AddScriptToEvaluateOnLoad(ctx context.Context, args *AddScriptToEvaluateOnLoadArgs) (reply *AddScriptToEvaluateOnLoadReply, err error) {
	reply = new(AddScriptToEvaluateOnLoadReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.addScriptToEvaluateOnLoad", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.addScriptToEvaluateOnLoad", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "AddScriptToEvaluateOnLoad", Err: err}
	}
	return
}

// RemoveScriptToEvaluateOnLoad invokes the Page method.
func (d *domainClient) RemoveScriptToEvaluateOnLoad(ctx context.Context, args *RemoveScriptToEvaluateOnLoadArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.removeScriptToEvaluateOnLoad", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.removeScriptToEvaluateOnLoad", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "RemoveScriptToEvaluateOnLoad", Err: err}
	}
	return
}

// SetAutoAttachToCreatedPages invokes the Page method. Controls whether browser will open a new inspector window for connected pages.
func (d *domainClient) SetAutoAttachToCreatedPages(ctx context.Context, args *SetAutoAttachToCreatedPagesArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.setAutoAttachToCreatedPages", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.setAutoAttachToCreatedPages", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "SetAutoAttachToCreatedPages", Err: err}
	}
	return
}

// Reload invokes the Page method. Reloads given page optionally ignoring the cache.
func (d *domainClient) Reload(ctx context.Context, args *ReloadArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.reload", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.reload", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "Reload", Err: err}
	}
	return
}

// Navigate invokes the Page method. Navigates current page to the given URL.
func (d *domainClient) Navigate(ctx context.Context, args *NavigateArgs) (reply *NavigateReply, err error) {
	reply = new(NavigateReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.navigate", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.navigate", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "Navigate", Err: err}
	}
	return
}

// StopLoading invokes the Page method. Force the page stop all navigations and pending resource fetches.
func (d *domainClient) StopLoading(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Page.stopLoading", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "StopLoading", Err: err}
	}
	return
}

// GetNavigationHistory invokes the Page method. Returns navigation history for the current page.
func (d *domainClient) GetNavigationHistory(ctx context.Context) (reply *GetNavigationHistoryReply, err error) {
	reply = new(GetNavigationHistoryReply)
	err = rpcc.Invoke(ctx, "Page.getNavigationHistory", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "GetNavigationHistory", Err: err}
	}
	return
}

// NavigateToHistoryEntry invokes the Page method. Navigates current page to the given history entry.
func (d *domainClient) NavigateToHistoryEntry(ctx context.Context, args *NavigateToHistoryEntryArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.navigateToHistoryEntry", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.navigateToHistoryEntry", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "NavigateToHistoryEntry", Err: err}
	}
	return
}

// GetCookies invokes the Page method. Returns all browser cookies. Depending on the backend support, will return detailed cookie information in the cookies field.
func (d *domainClient) GetCookies(ctx context.Context) (reply *GetCookiesReply, err error) {
	reply = new(GetCookiesReply)
	err = rpcc.Invoke(ctx, "Page.getCookies", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "GetCookies", Err: err}
	}
	return
}

// DeleteCookie invokes the Page method. Deletes browser cookie with given name, domain and path.
func (d *domainClient) DeleteCookie(ctx context.Context, args *DeleteCookieArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.deleteCookie", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.deleteCookie", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "DeleteCookie", Err: err}
	}
	return
}

// GetResourceTree invokes the Page method. Returns present frame / resource tree structure.
func (d *domainClient) GetResourceTree(ctx context.Context) (reply *GetResourceTreeReply, err error) {
	reply = new(GetResourceTreeReply)
	err = rpcc.Invoke(ctx, "Page.getResourceTree", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "GetResourceTree", Err: err}
	}
	return
}

// GetResourceContent invokes the Page method. Returns content of the given resource.
func (d *domainClient) GetResourceContent(ctx context.Context, args *GetResourceContentArgs) (reply *GetResourceContentReply, err error) {
	reply = new(GetResourceContentReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.getResourceContent", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.getResourceContent", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "GetResourceContent", Err: err}
	}
	return
}

// SearchInResource invokes the Page method. Searches for given string in resource content.
func (d *domainClient) SearchInResource(ctx context.Context, args *SearchInResourceArgs) (reply *SearchInResourceReply, err error) {
	reply = new(SearchInResourceReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.searchInResource", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.searchInResource", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "SearchInResource", Err: err}
	}
	return
}

// SetDocumentContent invokes the Page method. Sets given markup as the document's HTML.
func (d *domainClient) SetDocumentContent(ctx context.Context, args *SetDocumentContentArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.setDocumentContent", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.setDocumentContent", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "SetDocumentContent", Err: err}
	}
	return
}

// SetDeviceMetricsOverride invokes the Page method. Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
func (d *domainClient) SetDeviceMetricsOverride(ctx context.Context, args *SetDeviceMetricsOverrideArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.setDeviceMetricsOverride", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.setDeviceMetricsOverride", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "SetDeviceMetricsOverride", Err: err}
	}
	return
}

// ClearDeviceMetricsOverride invokes the Page method. Clears the overridden device metrics.
func (d *domainClient) ClearDeviceMetricsOverride(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Page.clearDeviceMetricsOverride", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "ClearDeviceMetricsOverride", Err: err}
	}
	return
}

// SetGeolocationOverride invokes the Page method. Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
func (d *domainClient) SetGeolocationOverride(ctx context.Context, args *SetGeolocationOverrideArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.setGeolocationOverride", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.setGeolocationOverride", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "SetGeolocationOverride", Err: err}
	}
	return
}

// ClearGeolocationOverride invokes the Page method. Clears the overridden Geolocation Position and Error.
func (d *domainClient) ClearGeolocationOverride(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Page.clearGeolocationOverride", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "ClearGeolocationOverride", Err: err}
	}
	return
}

// SetDeviceOrientationOverride invokes the Page method. Overrides the Device Orientation.
func (d *domainClient) SetDeviceOrientationOverride(ctx context.Context, args *SetDeviceOrientationOverrideArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.setDeviceOrientationOverride", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.setDeviceOrientationOverride", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "SetDeviceOrientationOverride", Err: err}
	}
	return
}

// ClearDeviceOrientationOverride invokes the Page method. Clears the overridden Device Orientation.
func (d *domainClient) ClearDeviceOrientationOverride(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Page.clearDeviceOrientationOverride", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "ClearDeviceOrientationOverride", Err: err}
	}
	return
}

// SetTouchEmulationEnabled invokes the Page method. Toggles mouse event-based touch event emulation.
func (d *domainClient) SetTouchEmulationEnabled(ctx context.Context, args *SetTouchEmulationEnabledArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.setTouchEmulationEnabled", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.setTouchEmulationEnabled", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "SetTouchEmulationEnabled", Err: err}
	}
	return
}

// CaptureScreenshot invokes the Page method. Capture page screenshot.
func (d *domainClient) CaptureScreenshot(ctx context.Context, args *CaptureScreenshotArgs) (reply *CaptureScreenshotReply, err error) {
	reply = new(CaptureScreenshotReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.captureScreenshot", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.captureScreenshot", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "CaptureScreenshot", Err: err}
	}
	return
}

// PrintToPDF invokes the Page method. Print page as PDF.
func (d *domainClient) PrintToPDF(ctx context.Context, args *PrintToPDFArgs) (reply *PrintToPDFReply, err error) {
	reply = new(PrintToPDFReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.printToPDF", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.printToPDF", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "PrintToPDF", Err: err}
	}
	return
}

// StartScreencast invokes the Page method. Starts sending each frame using the screencastFrame event.
func (d *domainClient) StartScreencast(ctx context.Context, args *StartScreencastArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.startScreencast", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.startScreencast", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "StartScreencast", Err: err}
	}
	return
}

// StopScreencast invokes the Page method. Stops sending each frame in the screencastFrame.
func (d *domainClient) StopScreencast(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Page.stopScreencast", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "StopScreencast", Err: err}
	}
	return
}

// ScreencastFrameAck invokes the Page method. Acknowledges that a screencast frame has been received by the frontend.
func (d *domainClient) ScreencastFrameAck(ctx context.Context, args *ScreencastFrameAckArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.screencastFrameAck", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.screencastFrameAck", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "ScreencastFrameAck", Err: err}
	}
	return
}

// HandleJavaScriptDialog invokes the Page method. Accepts or dismisses a JavaScript initiated dialog (alert, confirm, prompt, or onbeforeunload).
func (d *domainClient) HandleJavaScriptDialog(ctx context.Context, args *HandleJavaScriptDialogArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.handleJavaScriptDialog", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.handleJavaScriptDialog", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "HandleJavaScriptDialog", Err: err}
	}
	return
}

// GetAppManifest invokes the Page method.
func (d *domainClient) GetAppManifest(ctx context.Context) (reply *GetAppManifestReply, err error) {
	reply = new(GetAppManifestReply)
	err = rpcc.Invoke(ctx, "Page.getAppManifest", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "GetAppManifest", Err: err}
	}
	return
}

// RequestAppBanner invokes the Page method.
func (d *domainClient) RequestAppBanner(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Page.requestAppBanner", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "RequestAppBanner", Err: err}
	}
	return
}

// SetControlNavigations invokes the Page method. Toggles navigation throttling which allows programatic control over navigation and redirect response.
func (d *domainClient) SetControlNavigations(ctx context.Context, args *SetControlNavigationsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.setControlNavigations", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.setControlNavigations", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "SetControlNavigations", Err: err}
	}
	return
}

// ProcessNavigation invokes the Page method. Should be sent in response to a navigationRequested or a redirectRequested event, telling the browser how to handle the navigation.
func (d *domainClient) ProcessNavigation(ctx context.Context, args *ProcessNavigationArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.processNavigation", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.processNavigation", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "ProcessNavigation", Err: err}
	}
	return
}

// GetLayoutMetrics invokes the Page method. Returns metrics relating to the layouting of the page, such as viewport bounds/scale.
func (d *domainClient) GetLayoutMetrics(ctx context.Context) (reply *GetLayoutMetricsReply, err error) {
	reply = new(GetLayoutMetricsReply)
	err = rpcc.Invoke(ctx, "Page.getLayoutMetrics", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "GetLayoutMetrics", Err: err}
	}
	return
}

// CreateIsolatedWorld invokes the Page method. Creates an isolated world for the given frame.
func (d *domainClient) CreateIsolatedWorld(ctx context.Context, args *CreateIsolatedWorldArgs) (reply *CreateIsolatedWorldReply, err error) {
	reply = new(CreateIsolatedWorldReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Page.createIsolatedWorld", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Page.createIsolatedWorld", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Page", Op: "CreateIsolatedWorld", Err: err}
	}
	return
}

func (d *domainClient) DOMContentEventFired(ctx context.Context) (DOMContentEventFiredClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.domContentEventFired", d.conn)
	if err != nil {
		return nil, err
	}
	return &dOMContentEventFiredClient{Stream: s}, nil
}

type dOMContentEventFiredClient struct{ rpcc.Stream }

func (c *dOMContentEventFiredClient) Recv() (*DOMContentEventFiredReply, error) {
	event := new(DOMContentEventFiredReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "DOMContentEventFired Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) LoadEventFired(ctx context.Context) (LoadEventFiredClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.loadEventFired", d.conn)
	if err != nil {
		return nil, err
	}
	return &loadEventFiredClient{Stream: s}, nil
}

type loadEventFiredClient struct{ rpcc.Stream }

func (c *loadEventFiredClient) Recv() (*LoadEventFiredReply, error) {
	event := new(LoadEventFiredReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "LoadEventFired Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) FrameAttached(ctx context.Context) (FrameAttachedClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.frameAttached", d.conn)
	if err != nil {
		return nil, err
	}
	return &frameAttachedClient{Stream: s}, nil
}

type frameAttachedClient struct{ rpcc.Stream }

func (c *frameAttachedClient) Recv() (*FrameAttachedReply, error) {
	event := new(FrameAttachedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "FrameAttached Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) FrameNavigated(ctx context.Context) (FrameNavigatedClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.frameNavigated", d.conn)
	if err != nil {
		return nil, err
	}
	return &frameNavigatedClient{Stream: s}, nil
}

type frameNavigatedClient struct{ rpcc.Stream }

func (c *frameNavigatedClient) Recv() (*FrameNavigatedReply, error) {
	event := new(FrameNavigatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "FrameNavigated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) FrameDetached(ctx context.Context) (FrameDetachedClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.frameDetached", d.conn)
	if err != nil {
		return nil, err
	}
	return &frameDetachedClient{Stream: s}, nil
}

type frameDetachedClient struct{ rpcc.Stream }

func (c *frameDetachedClient) Recv() (*FrameDetachedReply, error) {
	event := new(FrameDetachedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "FrameDetached Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) FrameStartedLoading(ctx context.Context) (FrameStartedLoadingClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.frameStartedLoading", d.conn)
	if err != nil {
		return nil, err
	}
	return &frameStartedLoadingClient{Stream: s}, nil
}

type frameStartedLoadingClient struct{ rpcc.Stream }

func (c *frameStartedLoadingClient) Recv() (*FrameStartedLoadingReply, error) {
	event := new(FrameStartedLoadingReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "FrameStartedLoading Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) FrameStoppedLoading(ctx context.Context) (FrameStoppedLoadingClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.frameStoppedLoading", d.conn)
	if err != nil {
		return nil, err
	}
	return &frameStoppedLoadingClient{Stream: s}, nil
}

type frameStoppedLoadingClient struct{ rpcc.Stream }

func (c *frameStoppedLoadingClient) Recv() (*FrameStoppedLoadingReply, error) {
	event := new(FrameStoppedLoadingReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "FrameStoppedLoading Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) FrameScheduledNavigation(ctx context.Context) (FrameScheduledNavigationClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.frameScheduledNavigation", d.conn)
	if err != nil {
		return nil, err
	}
	return &frameScheduledNavigationClient{Stream: s}, nil
}

type frameScheduledNavigationClient struct{ rpcc.Stream }

func (c *frameScheduledNavigationClient) Recv() (*FrameScheduledNavigationReply, error) {
	event := new(FrameScheduledNavigationReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "FrameScheduledNavigation Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) FrameClearedScheduledNavigation(ctx context.Context) (FrameClearedScheduledNavigationClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.frameClearedScheduledNavigation", d.conn)
	if err != nil {
		return nil, err
	}
	return &frameClearedScheduledNavigationClient{Stream: s}, nil
}

type frameClearedScheduledNavigationClient struct{ rpcc.Stream }

func (c *frameClearedScheduledNavigationClient) Recv() (*FrameClearedScheduledNavigationReply, error) {
	event := new(FrameClearedScheduledNavigationReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "FrameClearedScheduledNavigation Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) FrameResized(ctx context.Context) (FrameResizedClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.frameResized", d.conn)
	if err != nil {
		return nil, err
	}
	return &frameResizedClient{Stream: s}, nil
}

type frameResizedClient struct{ rpcc.Stream }

func (c *frameResizedClient) Recv() (*FrameResizedReply, error) {
	event := new(FrameResizedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "FrameResized Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) JavascriptDialogOpening(ctx context.Context) (JavascriptDialogOpeningClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.javascriptDialogOpening", d.conn)
	if err != nil {
		return nil, err
	}
	return &javascriptDialogOpeningClient{Stream: s}, nil
}

type javascriptDialogOpeningClient struct{ rpcc.Stream }

func (c *javascriptDialogOpeningClient) Recv() (*JavascriptDialogOpeningReply, error) {
	event := new(JavascriptDialogOpeningReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "JavascriptDialogOpening Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) JavascriptDialogClosed(ctx context.Context) (JavascriptDialogClosedClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.javascriptDialogClosed", d.conn)
	if err != nil {
		return nil, err
	}
	return &javascriptDialogClosedClient{Stream: s}, nil
}

type javascriptDialogClosedClient struct{ rpcc.Stream }

func (c *javascriptDialogClosedClient) Recv() (*JavascriptDialogClosedReply, error) {
	event := new(JavascriptDialogClosedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "JavascriptDialogClosed Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) ScreencastFrame(ctx context.Context) (ScreencastFrameClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.screencastFrame", d.conn)
	if err != nil {
		return nil, err
	}
	return &screencastFrameClient{Stream: s}, nil
}

type screencastFrameClient struct{ rpcc.Stream }

func (c *screencastFrameClient) Recv() (*ScreencastFrameReply, error) {
	event := new(ScreencastFrameReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "ScreencastFrame Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) ScreencastVisibilityChanged(ctx context.Context) (ScreencastVisibilityChangedClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.screencastVisibilityChanged", d.conn)
	if err != nil {
		return nil, err
	}
	return &screencastVisibilityChangedClient{Stream: s}, nil
}

type screencastVisibilityChangedClient struct{ rpcc.Stream }

func (c *screencastVisibilityChangedClient) Recv() (*ScreencastVisibilityChangedReply, error) {
	event := new(ScreencastVisibilityChangedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "ScreencastVisibilityChanged Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) InterstitialShown(ctx context.Context) (InterstitialShownClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.interstitialShown", d.conn)
	if err != nil {
		return nil, err
	}
	return &interstitialShownClient{Stream: s}, nil
}

type interstitialShownClient struct{ rpcc.Stream }

func (c *interstitialShownClient) Recv() (*InterstitialShownReply, error) {
	event := new(InterstitialShownReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "InterstitialShown Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) InterstitialHidden(ctx context.Context) (InterstitialHiddenClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.interstitialHidden", d.conn)
	if err != nil {
		return nil, err
	}
	return &interstitialHiddenClient{Stream: s}, nil
}

type interstitialHiddenClient struct{ rpcc.Stream }

func (c *interstitialHiddenClient) Recv() (*InterstitialHiddenReply, error) {
	event := new(InterstitialHiddenReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "InterstitialHidden Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) NavigationRequested(ctx context.Context) (NavigationRequestedClient, error) {
	s, err := rpcc.NewStream(ctx, "Page.navigationRequested", d.conn)
	if err != nil {
		return nil, err
	}
	return &navigationRequestedClient{Stream: s}, nil
}

type navigationRequestedClient struct{ rpcc.Stream }

func (c *navigationRequestedClient) Recv() (*NavigationRequestedReply, error) {
	event := new(NavigationRequestedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Page", Op: "NavigationRequested Recv", Err: err}
	}
	return event, nil
}
