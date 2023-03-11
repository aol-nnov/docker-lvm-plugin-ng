package plugin

import (
	"github.com/docker/go-plugins-helpers/volume"
	"sync"
)

type vol struct {
	Name       string `json:"name"`
	MountPoint string `json:"mountpoint"`
	Vg         string `json:"vg"`
	Thinpool   string `json:"thinpool"`
	//Type       string `json:"type,omitempty"`
	//Source     string `json:"source,omitempty"`
	SnapSource string `json:"snap_source,omitempty"`
	Count      int    `json:"count"`
}

type localLvmStoragePlugin struct {
	mutex   sync.RWMutex
	volumes map[string]*vol
}

func New() (*localLvmStoragePlugin, error) {
	storagePlugin := &localLvmStoragePlugin{
		volumes: make(map[string]*vol),
	}

	err := loadFromDisk(storagePlugin)

	return storagePlugin, err
}

// check if localLvmStoragePlugin fully implements volume.Driver
// https://stackoverflow.com/a/27804417
var _ volume.Driver = (*localLvmStoragePlugin)(nil)

func (l *localLvmStoragePlugin) Capabilities() *volume.CapabilitiesResponse {
	return &volume.CapabilitiesResponse{
		Capabilities: volume.Capability{
			Scope: "local",
		},
	}
}
