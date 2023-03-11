package plugin

import (
	"docker-lvm-plugin-ng/lvm"
	"fmt"
	"github.com/docker/go-plugins-helpers/volume"
	"log"
	"time"
)

func (l *localLvmStoragePlugin) Get(req *volume.GetRequest) (*volume.GetResponse, error) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	log.Printf("lvm driver get request %v", req)

	v, exists := l.volumes[req.Name]
	log.Printf("volume in registry %v exists %v", v, exists)

	if !exists {
		return &volume.GetResponse{}, fmt.Errorf("No such volume '%s'", req.Name)
	}

	createdAt, err := lvm.GetVolumeCreationDateTime(v.Vg, v.Name)
	if err != nil {
		return nil, err
	}

	var res volume.GetResponse
	res.Volume = &volume.Volume{
		Name:       v.Name,
		Mountpoint: v.MountPoint,
		CreatedAt:  fmt.Sprintf(createdAt.Format(time.RFC3339)),
	}
	return &res, nil
}
