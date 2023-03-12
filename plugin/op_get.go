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

	vi, err := lvm.GetVolumeInfo(v.Vg, req.Name)
	if err != nil {
		return nil, err
	}

	statusMap := vi[0].Map()
	statusMap["RefCount"] = v.RefCount

	var res volume.GetResponse
	res.Volume = &volume.Volume{
		Name:       req.Name,
		Mountpoint: getMountpoint(req.Name),
		CreatedAt:  vi[0].Created.Format(time.RFC3339),
		Status:     statusMap,
	}
	return &res, nil
}
