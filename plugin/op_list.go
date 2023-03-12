package plugin

import "github.com/docker/go-plugins-helpers/volume"

func (l *localLvmStoragePlugin) List() (*volume.ListResponse, error) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	var ls []*volume.Volume
	for volName, _ := range l.volumes {
		v := &volume.Volume{
			Name:       volName,
			Mountpoint: getMountpoint(volName),
			// TO-DO: Find the significance of status field, and add that to volume.Volume
		}
		ls = append(ls, v)
	}
	return &volume.ListResponse{Volumes: ls}, nil
}
