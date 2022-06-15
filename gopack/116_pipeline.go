package gopack

func Make116Pipeline() {
	pipeline := CreatePipeline("to 1.16 (remove secrets, flattening)", TransPath("work/116/"), "1.16.zip")
	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(6))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())
	pipeline.SaveUntouched()
	pipeline.SkipGeneralProcessing = true
	CopyItems(pipeline, TransPath("work/115"), true)

	AddPipeline(pipeline)
}
