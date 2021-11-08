package generate

import (
	"path/filepath"
	"strings"

	"github.com/blang/semver/v4"
	"github.com/pkg/errors"
	"github.com/redhatrises/dockerfile-template/cmd/dft/common"
	"github.com/redhatrises/dockerfile-template/pkg/dft/api"
	"github.com/redhatrises/dockerfile-template/pkg/dft/engine"
	"github.com/spf13/cobra"
)

type generateOptions struct {
	File                string
	Image               string
	Name                string
	Print               bool
	Tag                 string
	Template            string
	OverrideFilename    string
	OverrideFilenameTag bool
}

var (
	opts generateOptions

	// Command: dft _generate_
	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate Dockerfiles from templates",
		Long:  "Generate Dockerfiles from templates",
		RunE: func(cmd *cobra.Command, args []string) error {
			return generateTemplates(cmd, args, opts)
		},
	}
)

func init() {
	common.Commands = common.AllCommands(common.CliCommand{
		Command: generateCmd,
	})

	generateFlags(generateCmd)
}

func generateFlags(cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.StringVarP(&opts.Tag, "tag", "t", "", "specify an image tag")
	flags.BoolVarP(&opts.Print, "print", "p", false, "output generated template to stdout")
	flags.StringVar(&opts.Template, "template", "", "`pathname or URL` of a Dockerfile.template")
	flags.StringVar(&opts.OverrideFilename, "override-filename", "", "override the auto-generated filename")
	flags.BoolVar(&opts.OverrideFilenameTag, "override-filename-tag", false, "drop the auto-generated filename tag")
}

func generateTemplates(cmd *cobra.Command, args []string, opts generateOptions) error {
	if len(args) < 1 {
		return errors.Errorf("unknown image: missing '<image>[:<tag>]' or '<image>'")
	}
	if len(opts.Template) == 0 {
		return errors.Errorf("missing argument '--template'\nTry '%[1]s --help' for more information.", cmd.CommandPath())
	}

	image := args[0]
	opts.Image = image

	if opts.Tag == "" {
		opts.Image = strings.Split(image, ":")[0]
		if strings.Contains(image, ":") {
			opts.Tag = strings.Split(image, ":")[1]
		}
	}

	opts.Name = filepath.Base(opts.Image)

	containerInfo := api.Container{
		Image: opts.Image,
		Tag:   opts.Tag,
		Name:  opts.Name,
	}

	if opts.Tag != "" {
		version, err := semver.ParseTolerant(opts.Tag)
		if err != nil {
			return errors.Errorf("Unable to parse semver: %v", err)
		}
		containerInfo.MajorVersion = version.Major
		containerInfo.MinorVersion = version.Minor
	}

	filename := api.Filename(containerInfo, opts.OverrideFilenameTag)
	if opts.OverrideFilename != "" {
		filename = opts.OverrideFilename
	}

	engine.FileTemplate(containerInfo, opts.Template, filename, opts.Print)

	return nil
}
