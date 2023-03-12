PLUGIN_NAME = i13n/docker-lvm-plugin
PLUGIN_TAG ?= overhaul
DISTRO ?= buster
ALLSRC = $(wildcard *.go) $(wildcard plugin/*.go) $(wildcard lvm/*.go)

all: create

dist/rootfs/docker-lvm-plugin: $(ALLSRC)
	docker build -t ${PLUGIN_NAME}:rootfs --build-arg DISTRO=${DISTRO} .

	@rm -rf ./dist/rootfs
	@mkdir -p ./dist/rootfs

	docker create --name tmp ${PLUGIN_NAME}:rootfs
	docker export tmp | tar -x -C ./dist/rootfs

	docker rm -vf tmp
	docker image rm ${PLUGIN_NAME}:rootfs

create: dist/rootfs/docker-lvm-plugin
	docker plugin rm -f ${PLUGIN_NAME}:${PLUGIN_TAG}-${DISTRO} || true

	docker plugin create ${PLUGIN_NAME}:${PLUGIN_TAG}-${DISTRO} ./dist

push: create
	@echo "### push plugin ${PLUGIN_NAME}:${PLUGIN_TAG}-${DISTRO}"
	@docker plugin push ${PLUGIN_NAME}:${PLUGIN_TAG}-${DISTRO}

clean:
	rm -rf ./dist/rootfs ./docker-lvm-plugin-ng
