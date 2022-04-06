package gopack

func Make118Pipeline() {
	pipeline := CreatePipeline("to 1.18 (remove secrets, flattening)", TransPath("work/118/"), "1.18.zip")

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(8))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	MakeEverythingLowercase(pipeline)
	RenamePackFolders(pipeline)
	ConvertItems(pipeline, Get14Blocks())
	ConvertItems(pipeline, Get14Items())
	ConvertItems(pipeline, GetBlocks())
	ConvertItems(pipeline, GetItems())
	ForceContent(pipeline, GetForcedContent())
	MigrateLanguage(pipeline, GetLang())

	//CompressResources(pipeline)
	//CopyItems(pipeline, TransPath("work/115"))

	pipeline.SaveUntouched()
	AddPipeline(pipeline)
}
