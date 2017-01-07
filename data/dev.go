// +build dev

// +-------------------=M=a=r=c=h=-=E=n=g=i=n=e=---------------------+
// | Copyright (C) 2016-2017 Andreas T Jonsson. All rights reserved. |
// | Contact <mail@andreasjonsson.se>                                |
// +-----------------------------------------------------------------+

package data

import (
	"net/http"
	"path"
)

type fsWrapper struct {
	http.FileSystem
}

func (fs fsWrapper) Open(name string) (http.File, error) {
	return fs.FileSystem.Open(path.Join("data", "src", name))
}

var FS = func() http.FileSystem {
	return &fsWrapper{http.Dir("")}
}()
