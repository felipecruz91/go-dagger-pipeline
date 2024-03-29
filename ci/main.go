// A generated module for GoDagger functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/go-dagger/internal/dagger"
	"fmt"
	"path/filepath"
	"runtime"

	platformFormat "github.com/containerd/containerd/platforms"
)

type GoDagger struct{}

func (m *GoDagger) BuildAll(ctx context.Context,
	// dir is the directory containing the Go source code
	// +required
	dir *Directory,
	// goVersion is the version of Go to use for building the binary
	// +optional
	// +default="1.22.0"
	goVersion string,
	// platforms is the list of platforms to build the binary for
	// +optional
	// +default=["linux/amd64", "linux/arm64"]
	platforms []string) (*Directory, error) {
	var files []*File

	for _, platform := range platforms {
		files = append(files, m.BuildBinary(ctx, dir, goVersion, platform))
	}

	return dag.Directory().WithFiles(".", files), nil
}

func (m *GoDagger) BuildBinary(ctx context.Context,
	// dir is the directory containing the Go source code
	// +required
	dir *Directory,
	// goVersion is the version of Go to use for building the binary
	// +optional
	// +default="1.22.0"
	goVersion string,
	// platform is the platform to build the binary for
	// +optional
	platform string) *File {
	if platform == "" {
		platform = "linux/" + runtime.GOARCH
	}

	return m.buildBinary(goVersion, dir, platform)
}

func (m *GoDagger) BuildImage(ctx context.Context,
	// dir is the directory containing the Go source code
	// +required
	dir *Directory,
	// goVersion is the version of Go to use for building the binary
	// +optional
	// +default="1.22.0"
	goVersion string,
	// platform is the platform to build the binary for
	// +optional
	platform string) *Container {

	if platform == "" {
		platform = "linux/" + runtime.GOARCH
	}

	binary := m.buildBinary(goVersion, dir, platform)

	return dag.Container().From("alpine:latest").
		WithFile("/bin/app", binary).
		WithEntrypoint([]string{"/bin/app"})
}

func (*GoDagger) buildBinary(goVersion string, dir *dagger.Directory, platform string) *dagger.File {
	// TODO: use GOOS and GOARCH in cache value key to avoid cache conflicts
	os := platformFormat.MustParse(string(platform)).OS
	arch := platformFormat.MustParse(string(platform)).Architecture
	binaryName := fmt.Sprintf("app_%s_%s", os, arch)

	binary := dag.Container().
		From("golang:"+goVersion).
		WithDirectory("/src", dir).
		WithWorkdir("/src").
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("go-mod")).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", dag.CacheVolume("go-build")).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithEnvVariable("CGO_ENABLED", "0").
		WithEnvVariable("GOOS", os).
		WithEnvVariable("GOARCH", arch).
		WithExec([]string{"go", "build", "-ldflags", "-s -w", "-o", binaryName, "."}).
		File(filepath.Join("/src", binaryName))
	return binary
}

// Test runs the Go tests
func (*GoDagger) Tests(ctx context.Context,
	// dir is the directory containing the Go source code
	// +required
	dir *Directory,
	// goVersion is the version of Go to use for building the binary
	// +optional
	// +default="1.22.0"
	goVersion string,
) (string, error) {
	return dag.Container().
		From("golang:"+goVersion).
		WithDirectory("/src", dir).
		WithWorkdir("/src").
		WithExec([]string{"go", "test", "-v", "--count=1", "./..."}).
		Stdout(ctx)
}

// Returns a container that echoes whatever string argument is provided
func (m *GoDagger) ContainerEcho(stringArg string) *Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *GoDagger) GrepDir(ctx context.Context, directoryArg *Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}
