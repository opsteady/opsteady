package component

import "fmt"

// DockerBuildInfo contains information about the docker image to to be created
type DockerBuildInfo struct {
	Name      string
	Version   string
	BuildArgs map[string]string
}

func (d *DockerBuildInfo) FullURL(registry string) string {
	return fmt.Sprintf("%s/%s:%s", registry, d.Name, d.Version)
}
