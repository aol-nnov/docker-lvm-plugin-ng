package plugin

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

const (
	//vgConfigPath         = "/etc/docker/docker-lvm-plugin"
	mountRoot            = "/var/lib/docker-lvm-plugin"
	lvmVolumesConfigPath = "/var/lib/docker-lvm-plugin/lvmVolumesConfig.json"
)

func getVgName(name string) (string, error) {
	vgName := name

	if vgName == "" {
		if val, hasEnv := os.LookupEnv("DEFAULT_VOLUME_GROUP"); hasEnv {
			vgName = val
		}
	}

	if vgName == "" {
		return "", fmt.Errorf("VG name is not present neither in request body nor in DEFAULT_VOLUME_GROUP env var")
	}

	return vgName, nil
}

func provisionMoutpoint(volName string) (string, error) {
	mountPoint := getMountpoint(volName)

	err := os.MkdirAll(mountPoint, 0700)
	if err != nil {
		os.RemoveAll(mountPoint)
		return "", err
	}

	return mountPoint, nil
}

func getMountpoint(volName string) string {
	return path.Join(mountRoot, volName)
}

func saveToDisk(volumes map[string]*vol) error {
	// Save volume store metadata.
	fhVolumes, err := os.Create(lvmVolumesConfigPath)
	if err != nil {
		return err
	}
	defer fhVolumes.Close()

	return json.NewEncoder(fhVolumes).Encode(&volumes)
}

func loadFromDisk(l *localLvmStoragePlugin) error {
	if _, err := os.Stat(lvmVolumesConfigPath); err == nil {
		// Load volume store metadata
		jsonVolumes, err := os.Open(lvmVolumesConfigPath)
		if err != nil {
			return err
		}
		defer jsonVolumes.Close()

		return json.NewDecoder(jsonVolumes).Decode(&l.volumes)
	}

	return nil
}
