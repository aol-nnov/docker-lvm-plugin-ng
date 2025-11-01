module docker-lvm-plugin-ng

// fixme https://github.com/docker/go-plugins-helpers/issues/134
// https://github.com/coreos/go-systemd/issues/321#issuecomment-552745196
replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.5.0

go 1.21

toolchain go1.22.2

require (
	github.com/docker/go-plugins-helpers v0.0.0-20211224144127-6eecb7beb651
	github.com/moby/sys/mountinfo v0.7.1
	golang.org/x/tools v0.13.0
)

require (
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
)
