package component

import "fmt"

// DockerBuildInfo containts information about the docker image to to be created
type DockerBuildInfo struct {
	Name    string
	Version string
}

func (d *DockerBuildInfo) FullUrl(registry string) string {
	return fmt.Sprintf("%s/%s:%s", registry, d.Name, d.Version)
}
