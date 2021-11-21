package templating

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"

	"github.com/Masterminds/sprig"
	"github.com/rs/zerolog"
)

type Templating struct {
	logger *zerolog.Logger
}

func NewTemplating(logger *zerolog.Logger) *Templating {
	return &Templating{
		logger: logger,
	}
}

func (t *Templating) Render(sourceFolder, destFolder string, values map[string]interface{}) error {
	err := filepath.Walk(sourceFolder,
		func(walkPath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Template if not a folder
			if !info.IsDir() {
				template, err := template.New(path.Base(walkPath)).Funcs(sprig.FuncMap()).ParseFiles(walkPath)
				if err != nil {
					return err
				}

				// Create the destination file
				valuesFile, err := os.Create(fmt.Sprintf("%s/%s", destFolder, info.Name()))
				if err != nil {
					return err
				}

				// foundation_azure_environment_name
				// Render the file
				return template.Execute(valuesFile, values)
			}

			return nil
		})

	return err
}
