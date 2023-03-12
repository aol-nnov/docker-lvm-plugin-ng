package plugin

import (
	"fmt"
	"github.com/docker/go-plugins-helpers/volume"
	"github.com/moby/sys/mountinfo"
	"os/exec"
)

// https://github.com/moby/moby/blob/master/volume/local/local.go
// https://github.com/moby/moby/blob/master/volume/local/local_unix.go

func (l *localLvmStoragePlugin) Unmount(req *volume.UnmountRequest) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.volumes[req.Name].RefCount == 1 {
		mp := getMountpoint(req.Name)

		isVolMounted, err := mountinfo.Mounted(mp)
		if err != nil {
			return fmt.Errorf("error unmounting volume '%s': %s", req.Name, err)
		}
		if isVolMounted {
			cmd := exec.Command("umount", mp)
			if out, err := cmd.CombinedOutput(); err != nil {
				return fmt.Errorf("error unmounting volume '%s' unmount error: %s output %s", req.Name, err, string(out))
			}
		}
	}
	l.volumes[req.Name].RefCount--
	//if err := saveToDisk(l.volumes); err != nil {
	//	return err
	//}
	return nil
}
