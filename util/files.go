// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2019, b3log.org & hacpai.com
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

package util

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/b3log/pipe/log"
)

// Logger
var logger = log.NewLogger(os.Stdout)

type myfile struct{}

// File utilities.
var File = myfile{}

// GetFileSize get the length in bytes of file of the specified path.
func (*myfile) GetFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if nil != err {
		logger.Error(err)

		return -1
	}

	return fi.Size()
}

// IsExist determines whether the file spcified by the given path is exists.
func (*myfile) IsExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

// IsBinary determines whether the specified content is a binary file content.
func (*myfile) IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}

	return false
}

// IsImg determines whether the specified extension is a image.
func (*myfile) IsImg(extension string) bool {
	ext := strings.ToLower(extension)

	switch ext {
	case ".jpg", ".jpeg", ".bmp", ".gif", ".png", ".svg", ".ico":
		return true
	default:
		return false
	}
}

// IsDir determines whether the specified path is a directory.
func (*myfile) IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if nil != err {
		logger.Warnf("Determines whether [%s] is a directory failed: [%v]", path, err)

		return false
	}

	return fio.IsDir()
}

// CopyFile copies the source file to the dest file.
func (*myfile) CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if nil != err {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if nil != err {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		if sourceinfo, e := os.Stat(source); nil != e {
			err = os.Chmod(dest, sourceinfo.Mode())

			return
		}
	}

	return nil
}

// CopyDir copies the source directory to the dest directory.
func (*myfile) CopyDir(source string, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if nil != err {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if nil != err {
		return err
	}

	directory, err := os.Open(source)
	if nil != err {
		return err
	}

	defer directory.Close()

	objects, err := directory.Readdir(-1)
	if nil != err {
		return err
	}

	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// create sub-directories - recursively
			err = File.CopyDir(srcFilePath, destFilePath)
			if nil != err {
				logger.Error(err)
			}
		} else {
			err = File.CopyFile(srcFilePath, destFilePath)
			if nil != err {
				logger.Error(err)
			}
		}
	}

	return nil
}
