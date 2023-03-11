module docker-lvm-plugin-ng

// fixme https://github.com/docker/go-plugins-helpers/issues/134
// https://github.com/coreos/go-systemd/issues/321#issuecomment-552745196
replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.5.0

go 1.19

require (
	github.com/docker/go-plugins-helpers v0.0.0-20211224144127-6eecb7beb651
	github.com/moby/sys/mountinfo v0.6.2
)

require (
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/tools v0.1.12 // indirect
)
