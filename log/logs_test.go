// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2018, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package log

import (
	"os"
	"testing"
)

// Logger.
var logger = NewLogger(os.Stdout)

func TestSetLevel(t *testing.T) {
	SetLevel("trace")
}

func TestTrace(t *testing.T) {
	logger.SetLevel("trace")
	logger.Trace("trace")
	logger.SetLevel("off")
	logger.Trace("trace")
}

func TestTracef(t *testing.T) {
	logger.SetLevel("trace")
	logger.Tracef("tracef")
	logger.SetLevel("off")
	logger.Tracef("tracef")
}

func TestDebug(t *testing.T) {
	logger.SetLevel("debug")
	logger.Debug("debug")
	logger.SetLevel("off")
	logger.Debug("debug")
}

func TestDebugf(t *testing.T) {
	logger.SetLevel("debug")
	logger.Debugf("debugf")
	logger.SetLevel("off")
	logger.Debug("debug")
}

func TestInfo(t *testing.T) {
	logger.SetLevel("info")
	logger.Info("info")
	logger.SetLevel("off")
	logger.Info("info")
}

func TestInfof(t *testing.T) {
	logger.SetLevel("info")
	logger.Infof("infof")
	logger.SetLevel("off")
	logger.Infof("infof")
}

func TestWarn(t *testing.T) {
	logger.SetLevel("warn")
	logger.Warn("warn")
	logger.SetLevel("off")
	logger.Warn("warn")
}

func TestWarnf(t *testing.T) {
	logger.SetLevel("warn")
	logger.Warnf("warnf")
	logger.SetLevel("off")
	logger.Warnf("warnf")
}

func TestError(t *testing.T) {
	logger.SetLevel("error")
	logger.Error("error")
	logger.SetLevel("off")
	logger.Error("error")
}

func TestErrorf(t *testing.T) {
	logger.SetLevel("error")
	logger.Errorf("errorf")
	logger.SetLevel("off")
	logger.Errorf("errorf")
}

func TestGetLevel(t *testing.T) {
	if getLevel("trace") != Trace {
		t.FailNow()

		return
	}

	if getLevel("debug") != Debug {
		t.FailNow()

		return
	}

	if getLevel("info") != Info {
		t.FailNow()

		return
	}

	if getLevel("warn") != Warn {
		t.FailNow()

		return
	}

	if getLevel("error") != Error {
		t.FailNow()

		return
	}

	if getLevel("fatal") != Fatal {
		t.FailNow()

		return
	}
}

func TestLoggerSetLevel(t *testing.T) {
	logger.SetLevel("trace")

	if logger.level != Trace {
		t.FailNow()

		return
	}
}

func TestIsTraceEnabled(t *testing.T) {
	logger.SetLevel("trace")

	if !logger.IsTraceEnabled() {
		t.FailNow()

		return
	}
}

func TestIsDebugEnabled(t *testing.T) {
	logger.SetLevel("debug")

	if !logger.IsDebugEnabled() {
		t.FailNow()

		return
	}
}

func TestIsWarnEnabled(t *testing.T) {
	logger.SetLevel("warn")

	if !logger.IsWarnEnabled() {
		t.FailNow()

		return
	}
}
