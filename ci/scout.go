package main

import "context"

const (
	dockerScoutImage = "index.docker.io/docker/scout-cli:latest"
)

// DockerScoutCves displays CVEs identified in a container image
func (m *GoDagger) DockerScoutCves(
	ctx context.Context,
	// dockerScoutHubUser is the username for Docker Scout Hub
	// +required
	dockerScoutHubUser string,
	// dockerScoutHubPassword is the password for Docker Scout Hub
	// +required
	dockerScoutHubPassword *Secret,
	// imageRef is the reference to the image to analyze
	// +required
	imageRef string) (*Container, error) {

	cli := dag.Pipeline("docker-scout-cves")

	return cli.Container().From(dockerScoutImage).
		WithEnvVariable("DOCKER_SCOUT_HUB_USER", dockerScoutHubUser).
		WithSecretVariable("DOCKER_SCOUT_HUB_PASSWORD", dockerScoutHubPassword).
		WithExec([]string{"cves", imageRef}).
		Sync(ctx)
}
