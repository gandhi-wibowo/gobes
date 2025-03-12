package config

import "path/filepath"

func Convert(fullPath string) ConfigFile {
	basedir := filepath.Dir(fullPath)
	filename := filepath.Base(fullPath)
	extname := filepath.Ext(filename)
	if len(extname) > 0 {
		filename = filename[:len(filename)-len(extname)]
	}

	return ConfigFile{
		Name: filename,
		Path: basedir,
	}
}
