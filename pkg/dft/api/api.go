package api

import "fmt"

type Container struct {
	Name         string
	Image        string
	Tag          string
	MajorVersion uint64
	MinorVersion uint64
}

func Filename(c Container, removeTag bool) string {
	filename := "Dockerfile"
	if c.Name != "" {
		return fmt.Sprintf("%s.%s", filename, c.Name)
	}
	if c.Tag != "" {
		return fmt.Sprintf("%s%d", filename, c.MajorVersion)
	}
	if removeTag {
		return fmt.Sprintf("%s.%s", filename, c.Name)
	}
	return filename
}
