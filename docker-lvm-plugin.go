package main

import (
	"docker-lvm-plugin-ng/plugin"
	"github.com/docker/go-plugins-helpers/volume"
)

func main() {
	lvmStorageDriver, err := plugin.New()

	if err != nil {
		panic(err)
	}

	pluginHandler := volume.NewHandler(lvmStorageDriver)
	if err := pluginHandler.ServeUnix("lvm", 0); err != nil {
		panic(err)
	}
}
