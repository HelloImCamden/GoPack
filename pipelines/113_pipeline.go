package pipelines

import (
	"github.com/Mindgamesnl/GoPack/loader"
	"github.com/Mindgamesnl/GoPack/pipelines/common_actions"
)

func Make113Pipeline() {
	pipeline := loader.CreatePipeline("to 1.12.3", "work/113/")

	// update format
	pipeline.AddForFileName("pack.mcmeta", common_actions.SetMetaValue(4))

	pipeline.UnhandledFileHandler(func(originalPack loader.ResourcePack, resource loader.Resource, pipeline *loader.Pipeline) {
		// remove comments
		pipeline.SaveBytes(resource, resource.ContentAsBytes())
	})

	loader.AddPipeline(&pipeline)
}
