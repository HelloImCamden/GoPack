package gopack

func Make119Pipeline() {
	pipeline := CreatePipeline("to 1.19 (remove secrets, flattening)", TransPath("work/119/"), "1.19.zip")

	// update format
	pipeline.AddForFileName("pack.mcmeta", SetMetaRevision(8))
	// remove comments
	pipeline.AddForFileType("json", RemoveComment())

	//MakeEverythingLowercase(pipeline)
	//RenamePackFolders(pipeline)
	//ConvertItems(pipeline, Get14Blocks())
	//ConvertItems(pipeline, Get14Items())
	//ConvertItems(pipeline, GetBlocks())
	//ConvertItems(pipeline, GetItems())
	//ForceContent(pipeline, GetForcedContent())
	//MigrateLanguage(pipeline, GetLang())

	pipeline.SaveUntouched()
	pipeline.SkipGeneralProcessing = true

	//CompressResources(pipeline)
	CopyItems(pipeline, TransPath("work/115"), true)


	AddPipeline(pipeline)
}
