package gopack

func Make118Pipeline() {
	pipeline := CreatePipeline("to 1.18 (remove secrets, flattening)", TransPath("work/118/"), "1.18.zip")
	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(8))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())
	pipeline.SaveUntouched()
	pipeline.SkipGeneralProcessing = true
	CopyItems(pipeline, TransPath("work/115"), true)

	AddPipeline(pipeline)
}
