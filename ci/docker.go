package main

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	platformFormat "github.com/containerd/containerd/platforms"
)

func (m *GoDagger) DockerBuildImage(ctx context.Context,
	// sourcer is the directory containing the Go source code
	// +required
	source *Directory,
	// goVersion is the version of Go to use for building the binary
	// +optional
	// +default="1.22.0"
	goVersion string,
	// platform is the platform to build the binary for
	// +optional
	platform string) (*Container, error) {
	cli := dag.Pipeline("docker-build")

	if platform == "" {
		platform = "linux/" + runtime.GOARCH
	}

	os := platformFormat.MustParse(string(platform)).OS
	arch := platformFormat.MustParse(string(platform)).Architecture
	binaryName := fmt.Sprintf("app_%s_%s", os, arch)

	ctr, err := m.buildBinary(ctx, source, goVersion, platform)
	if err != nil {
		return nil, err
	}
	file := ctr.File(filepath.Join("/src", binaryName))

	return cli.Container().From("alpine:latest").
		WithFile("/bin/app", file).
		WithEntrypoint([]string{"/bin/app"}), nil

}
