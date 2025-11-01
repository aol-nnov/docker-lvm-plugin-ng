# Docker lvm volume driver

Highly inspired by @nickbreen's [docker-lvm-plugin](https://github.com/nickbreen/docker-lvm-plugin)

Though, it is used in production for 6+ years by now, use it with agrain of salt, you've been warned ))

For plugin to work correctly, it is recommended to add `--timeout` parameter to the corresponding
`docker plugin enable` command, i.e.:

```
docker plugin enable --timeout 300 <your_plugin_name>
```

## Building

```
make
```

## Usage

(somewhere in docker-compose.yml)

```yml
volumes:
  grafana-config:
    driver: "<your_plugin_name>"
    driver_opts:
      thinpool: data
      size: 100M
```
