package plugin

import "github.com/docker/go-plugins-helpers/volume"

func (l *localLvmStoragePlugin) Path(req *volume.PathRequest) (*volume.PathResponse, error) {
	return &volume.PathResponse{Mountpoint: req.Name}, nil
}
