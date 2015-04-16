// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package isolateserver

import (
	"crypto/sha1"
	"fmt"
	"hash"

	"github.com/luci/luci-go/client/internal/common"
)

// IsolateServer is the client interface to interact with an Isolate server.
type IsolateServer interface {
	ServerCapabilities() (*ServerCapabilities, error)
}

// ServerCapabilities is the server details as exposed by the server.
type ServerCapabilities struct {
	ServerVersion string `json:"server_version"`
}

// Namespace is the bucket in which content is saved into.
type Namespace struct {
	Namespace   string `json:"namespace"`
	DigestAlgo  string `json:"digest_hash"`
	Compression string `json:"compression"`
}

// Returns the valid hash.Hash instance for this namespace.
func (n *Namespace) GetHashAlgo() (hash.Hash, error) {
	switch n.DigestAlgo {
	case "sha-1":
		return sha1.New(), nil
	default:
		return nil, fmt.Errorf("unknown hash algo \"%s\"", n.DigestAlgo)
	}
}

// New returns a new IsolateServer client.
func New(url, namespace, digestAlgo, compression string) IsolateServer {
	return &isolateServer{
		url: url,
		namespace: Namespace{
			Namespace:   namespace,
			DigestAlgo:  digestAlgo,
			Compression: compression,
		},
	}
}

// Private details.

type isolateServer struct {
	url       string
	namespace Namespace
}

func (i *isolateServer) ServerCapabilities() (*ServerCapabilities, error) {
	url := i.url + "/_ah/api/isolateservice/v1/server_details"
	out := &ServerCapabilities{}
	if _, err := common.PostJSON(nil, url, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}
