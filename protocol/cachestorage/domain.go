// Code generated by cdpgen. DO NOT EDIT.

package cachestorage

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the CacheStorage domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the CacheStorage domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// RequestCacheNames invokes the CacheStorage method. Requests cache names.
func (d *domainClient) RequestCacheNames(ctx context.Context, args *RequestCacheNamesArgs) (reply *RequestCacheNamesReply, err error) {
	reply = new(RequestCacheNamesReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CacheStorage.requestCacheNames", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CacheStorage.requestCacheNames", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CacheStorage", Op: "RequestCacheNames", Err: err}
	}
	return
}

// RequestEntries invokes the CacheStorage method. Requests data from cache.
func (d *domainClient) RequestEntries(ctx context.Context, args *RequestEntriesArgs) (reply *RequestEntriesReply, err error) {
	reply = new(RequestEntriesReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "CacheStorage.requestEntries", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CacheStorage.requestEntries", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CacheStorage", Op: "RequestEntries", Err: err}
	}
	return
}

// DeleteCache invokes the CacheStorage method. Deletes a cache.
func (d *domainClient) DeleteCache(ctx context.Context, args *DeleteCacheArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "CacheStorage.deleteCache", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CacheStorage.deleteCache", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CacheStorage", Op: "DeleteCache", Err: err}
	}
	return
}

// DeleteEntry invokes the CacheStorage method. Deletes a cache entry.
func (d *domainClient) DeleteEntry(ctx context.Context, args *DeleteEntryArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "CacheStorage.deleteEntry", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "CacheStorage.deleteEntry", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "CacheStorage", Op: "DeleteEntry", Err: err}
	}
	return
}
