package plugin

import (
	"docker-lvm-plugin-ng/lvm"
	"fmt"
	"github.com/docker/go-plugins-helpers/volume"
	"log"
	"os"
)

func (l *localLvmStoragePlugin) Remove(req *volume.RemoveRequest) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	v, exists := l.volumes[req.Name]
	log.Printf("volume in registry %v exists %v", v, exists)

	if !exists {
		return fmt.Errorf("remove: No such volume '%s'", req.Name)
	}

	isOrigin := func() bool {
		for volName, vol := range l.volumes {
			if volName == req.Name {
				continue
			}
			if vol.Origin == req.Name {
				return true
			}
		}
		return false
	}()

	if isOrigin {
		return fmt.Errorf("Error removing volume, all snapshot destinations must be removed before removing the original volume")
	}

	if err := os.RemoveAll(getMountpoint(req.Name)); err != nil {
		return err
	}

	if out, err := lvm.RemoveLogicalVolume(v.Vg, req.Name); err != nil {
		return fmt.Errorf("Remove: removeLogicalVolume error %s output %s", err, string(out))
	}

	delete(l.volumes, req.Name)
	if err := saveToDisk(l.volumes); err != nil {
		return err
	}
	return nil
}
