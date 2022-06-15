package gopack

func Make119Pipeline() {
	pipeline := CreatePipeline("to 1.19 (remove secrets, flattening)", TransPath("work/119/"), "1.19.zip")
	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(9))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())
	pipeline.SaveUntouched()
	pipeline.SkipGeneralProcessing = true

	CopyItems(pipeline, TransPath("work/115"), true)
	AddPipeline(pipeline)
}
