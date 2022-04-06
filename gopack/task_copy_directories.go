package gopack

import (
	cp "github.com/otiai10/copy"
)

func CopyItems(pipeline *Pipeline, directory string) {
	cp.Copy(directory + "/assets", pipeline.OutFolder + "/assets")
}