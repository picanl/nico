// Copyright (c) 2020-present Cloud <cloud@txthinking.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of version 3 of the GNU General Public
// License as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"net/http"
	"os"
)

type WebRoot struct {
	Dir http.Dir
}

func NewWebRoot(dir string) *WebRoot {
	return &WebRoot{
		Dir: http.Dir(dir),
	}
}

func (w *WebRoot) Open(name string) (http.File, error) {
	f, err := w.Dir.Open(name)
	if err == nil {
		return f, nil
	}
	if os.IsNotExist(err) {
		return w.Dir.Open("/index.html")
	}
	return nil, err
}
