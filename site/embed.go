package site

import (
	"embed"
	"io/fs"
)

// DistFS contains the Vite build output under site/dist.
//
//go:embed all:dist
var DistFS embed.FS

// DistDirFS is dist rooted at "/" for static serving.
var DistDirFS fs.FS

func init() {
	sub, err := fs.Sub(DistFS, "dist")
	if err != nil {
		panic(err)
	}
	DistDirFS = sub
}
