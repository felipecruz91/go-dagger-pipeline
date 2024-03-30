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
)

const (
	defaultGoVersion = "1.22.0"
)

var (
	defaultPlatforms = []string{"linux/amd64", "linux/arm64"}
)

type GoDagger struct{}

// func (m *GoDagger) Ci(ctx context.Context,
// 	// source is the directory containing the Go source code
// 	// +required
// 	source *Directory,
// // dockerScoutHubUser is the username for Docker Scout Hub
// // +required
// dockerScoutHubUser string,
// // dockerScoutHubPassword is the password for Docker Scout Hub
// // +required
// dockerScoutHubPassword *Secret
// ) error {

// 	if _, err := m.Test(ctx, source, defaultGoVersion, false, true, true); err != nil {
// 		return err
// 	}

// 	if _, err := m.BuildBinaries(ctx, source, defaultGoVersion, defaultPlatforms); err != nil {
// 		return err
// 	}

// 	for _, platform := range defaultPlatforms {
// 		_, err := m.DockerBuildImage(ctx, source, defaultGoVersion, platform)
// 		if err != nil {
// 			return err
// 		}
// 		// TODO: not supported to export image because it's running inside a module container?
// 		// if _, err := ctr.Export(ctx, "image"); err != nil {
// 		// 	return err
// 		// }

// 		// _, err = m.Cves(ctx, dockerScoutHubUser, dockerScoutHubPassword, "")
// 		// if err != nil {
// 		// 	return err
// 		// }
// 	}

// 	return nil
// }

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
