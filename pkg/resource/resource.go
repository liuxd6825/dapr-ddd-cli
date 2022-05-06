package resource

import "embed"

var local embed.FS

func Init(aLocal embed.FS) {
	local = aLocal
}

func Local() embed.FS {
	return local
}
