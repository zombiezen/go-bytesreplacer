// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bytesreplacer provides a utility for replacing parts of byte slices.
package bytesreplacer // import "zombiezen.com/go/bytesreplacer"

import (
	"io"

	"go4.org/bytereplacer"
)

// Replacer replaces a list of strings with replacements.
// It is safe for concurrent use by multiple goroutines.
type Replacer struct {
	*bytereplacer.Replacer
}

// New returns a new Replacer from a list of old, new string pairs.
// Replacements are performed in order, without overlapping matches.
func New(oldnew ...string) *Replacer {
	return &Replacer{bytereplacer.New(oldnew...)}
}

// Write writes s to w with all replacements performed.
func (r *Replacer) Write(w io.Writer, s []byte) (n int, err error) {
	ss := make([]byte, len(s))
	copy(ss, s)
	ss = r.Replace(ss)
	return w.Write(ss)
}
