package gopack

import (
	cp "github.com/otiai10/copy"
	"os"
	"strings"
)

func CopyItems(pipeline *Pipeline, directory string, removeFirst bool) {
	finish := func(pipeline *Pipeline) {
		if removeFirst {
			_ = os.RemoveAll(pipeline.OutFolder + "assets")
		}
		opt := cp.Options{
			Skip: func(src string) (bool, error) {
				return strings.HasPrefix(src, ".git"), nil
			},
			OnDirExists: func(src, dest string) cp.DirExistsAction {
				return cp.Replace
			},
		}
		cp.Copy(directory + "/assets", pipeline.OutFolder + "assets", opt);
	}

	pipeline.AddFinishHandler(finish)
}

