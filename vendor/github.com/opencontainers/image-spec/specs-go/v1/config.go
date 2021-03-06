// Copyright 2016 The Linux Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

// ImageConfig defines the execution parameters which should be used as a base when running a container using an image.
type ImageConfig struct {
	// User defines the username or UID which the process in the container should run as.
	User string `json:"User"`

	// Memory defines the memory limit.
	Memory int64 `json:"Memory"`

	// MemorySwap defines the total memory usage limit (memory + swap).
	MemorySwap int64 `json:"MemorySwap"`

	// CPUShares is the CPU shares (relative weight vs. other containers).
	CPUShares int64 `json:"CpuShares"`

	// ExposedPorts a set of ports to expose from a container running this image.
	ExposedPorts map[string]struct{} `json:"ExposedPorts"`

	// Env is a list of environment variables to be used in a container.
	Env []string `json:"Env"`

	// Entrypoint defines a list of arguments to use as the command to execute when the container starts.
	Entrypoint []string `json:"Entrypoint"`

	// Cmd defines the default arguments to the entrypoint of the container.
	Cmd []string `json:"Cmd"`

	// Volumes is a set of directories which should be created as data volumes in a container running this image.
	Volumes map[string]struct{} `json:"Volumes"`

	// WorkingDir sets the current working directory of the entrypoint process in the container.
	WorkingDir string `json:"WorkingDir"`
}

// RootFS describes a layer content addresses
type RootFS struct {
	// Type is the type of the rootfs.
	Type string `json:"type"`

	// DiffIDs is an array of layer content hashes (DiffIDs), in order from bottom-most to top-most.
	DiffIDs []string `json:"diff_ids"`
}

// History describes the history of a layer.
type History struct {
	// Created is the creation time.
	Created string `json:"created"`

	// CreatedBy is the command which created the layer.
	CreatedBy string `json:"created_by"`

	// Author is the author of the build point.
	Author string `json:"author"`

	// Comment is a custom message set when creating the layer.
	Comment string `json:"comment"`

	// EmptyLayer is used to mark if the history item created a filesystem diff.
	EmptyLayer bool `json:"empty_layer"`
}

// Image is the JSON structure which describes some basic information about the image.
type Image struct {
	// Created defines an ISO-8601 formatted combined date and time at which the image was created.
	Created string `json:"created"`

	// Author defines the name and/or email address of the person or entity which created and is responsible for maintaining the image.
	Author string `json:"author"`

	// Architecture is the CPU architecture which the binaries in this image are built to run on.
	Architecture string `json:"architecture"`

	// OS is the name of the operating system which the image is built to run on.
	OS string `json:"os"`

	// Config defines the execution parameters which should be used as a base when running a container using the image.
	Config ImageConfig `json:"config"`

	// RootFS references the layer content addresses used by the image.
	RootFS RootFS `json:"rootfs"`

	// History describes the history of each layer.
	History []History `json:"history"`
}
