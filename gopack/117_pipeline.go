package gopack

func Make117Pipeline() {
	pipeline := CreatePipeline("to 1.17 (remove secrets, flattening)", TransPath("work/117/"), "1.17.zip")

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(7))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())
	pipeline.SaveUntouched()
	pipeline.SkipGeneralProcessing = true
	CopyItems(pipeline, TransPath("work/115"), true)

	AddPipeline(pipeline)
}
