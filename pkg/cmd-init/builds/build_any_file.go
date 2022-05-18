package builds

type BuildAnyFile struct {
	BaseBuild
	outFile string
	values  map[string]interface{}
}

func NewBuildAnyFile(base BaseBuild, values map[string]interface{}, tmplFile, outFile string) *BuildAnyFile {
	res := &BuildAnyFile{
		BaseBuild: base,
		values:    values,
		outFile:   outFile,
	}
	res.TmplFile = tmplFile
	res.OutFile = outFile
	res.ValuesFunc = res.Values
	return res
}
