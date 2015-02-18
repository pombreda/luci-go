// Copyright (c) 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

func mainImpl() error {
	fmt.Printf("swarming communicates with the Swarming server.\n")
	return nil
}

func main() {
	if err := mainImpl(); err != nil {
		fmt.Fprintf(os.Stderr, "failure: %s\n", err)
		os.Exit(1)
	}
}
