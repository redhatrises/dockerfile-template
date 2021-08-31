package api

import "fmt"

type Container struct {
	Name         string
	Image        string
	Tag          string
	MajorVersion uint64
	MinorVersion uint64
}

func Filename(c Container) string {
	filename := "Dockerfile"
	if c.Name != "" {
		filename = fmt.Sprintf("%s.%s", filename, c.Name)
	}
	if c.Tag != "" {
		filename = fmt.Sprintf("%s%d", filename, c.MajorVersion)
	}
	return filename
}
