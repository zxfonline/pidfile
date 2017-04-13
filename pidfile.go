// Copyright 2016 zxfonline@sina.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pidfile

import (
	"os"
	"strconv"

	"github.com/zxfonline/fileutil"
	"github.com/zxfonline/golog"
)

type Pidfile struct {
	Pathfile string
}

var _localpid *Pidfile

func Init(pathfile string) {
	_localpid = New(pathfile)
}

// New creates a new Pidfile and writes the current PID
func New(pathfile string) *Pidfile {
	file, err := fileutil.OpenFile(pathfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fileutil.DefaultFileMode)
	if err != nil {
		golog.Warnf("pidfile: failed to open %s (%s)", pathfile, err)
		return nil
	}
	defer file.Close()
	pid := strconv.Itoa(os.Getpid())
	golog.Infof("start process pid:%v --> %s", pid, pathfile)
	file.WriteString(pid)
	return &Pidfile{pathfile}
}

func (pf *Pidfile) Remove() {
	err := os.Remove(pf.Pathfile)
	if err != nil {
		golog.Warnf("pidfile: failed to remove %s (%s)", pf.Pathfile, err)
	}
}
func Remove() {
	defer func() { recover() }()
	if _localpid != nil {
		_localpid.Remove()
	}
}
