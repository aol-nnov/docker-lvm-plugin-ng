//go:generate go run golang.org/x/tools/cmd/stringer -linecomment -type=LvVolumeType,LvPermissions,LvAllocPolicy,LvState,LvDeviceStatus,LvTargetType,LvHealth

package lvm

import "fmt"

type LvVolumeType int

const (
	Cache                 LvVolumeType = iota // cache
	Mirrored                                  // mirrored
	MirroredNoInitialSync                     // mirrored without initial sync
	Origin                                    // origin
	OriginMerging                             // origin with merging snapshot
	Raid                                      // raid
	RaidNoInitialSync                         // raid without initial sync
	Snapshot                                  // snapshot
	SnapshotMerging                           // merging snapshot
	PvMove                                    // pvmove
	Virtual                                   // virtual
	Image                                     // mirror or raid image
	ImageOutOfSync                            // mirror or raid image out-of-sync
	Log                                       // log device
	UnderConversion                           // under conversion
	ThinVolume                                // thin volume
	ThinPool                                  // thin pool
	ThinPoolData                              // thin pool data
	VdoPool                                   // VDO pool
	VdoPoolData                               // VDO pool data
	PoolMetaSpare                             // raid or pool metadata or pool metadata spare
)

type LvPermissions int

const (
	Writeable        LvPermissions = iota // writeable
	ReadOnly                              // read-only
	RoActivationOfRw                      // read-only activation of non-read-only volume
)

type LvAllocPolicy int

const (
	Anywhere         LvAllocPolicy = iota // anywhere
	Contiguous                            // contiguous
	Inherited                             // inherited
	Cling                                 // cling
	Normal                                // normal
	AnywhereLocked                        // anywhere (currently locked)
	ContiguousLocked                      // contiguous (currently locked)
	InheritedLocked                       // inherited (currently locked)
	ClingLocked                           // cling (currently locked)
	NormalLocked                          // normal (currently locked)
)

type LvState int

const (
	Inactive                   LvState = iota // inactive
	Active                                    // active
	Historical                                // historical
	Suspended                                 // suspended
	InvalidSnap                               // invalid snapshot
	InvalidSuspendedSnap                      // invalid suspended snapshot
	SnapMergeFailed                           // snapshot merge failed
	SuspendedSnapMergeFailed                  // suspended snapshot merge failed
	MappedDevWithoutTables                    // mapped device present without tables
	MappedDevWithInactiveTable                // mapped device present with inactive table
	CheckThinPool                             // thin-pool check needed
	CheckSuspendedThinPool                    // suspended thin-pool check needed
	StateUnknown                              // unknown

)

type LvDeviceStatus int

const (
	Mounted         LvDeviceStatus = iota // mounted (open)
	NotMounted                            // not mounted
	DevStateUnknown                       // unknown
)

type LvTargetType int

const (
	TargetTypeCache    LvTargetType = iota // cache
	TargetTypeMirror                       // mirror
	TargetTypeRaid                         // raid
	TargetTypeSnapshot                     // snapshot
	TargetTypeThin                         // thin
	TargetTypeUnknown                      // unknown
	TargetTypeVirtual                      // virtual

)

type LvHealth int

const (
	HealthPartial LvHealth = iota // partial
	HealthUnknown                 //unknown

	RaidRefreshNeeded      // refresh needed
	RaidMismatches         // mismatches exist
	RaidWritemostly        // writemostly
	RaidReshaping          // reshaping
	RaidRemoveAfterReshape // remove after reshape

	ThinFailed     // failed
	ThinOutOfData  // out of data space
	ThinMetaFailed // metadata read only

	WritecacheError // error
)

type LvAttributes struct {
	VolumeType           LvVolumeType
	Permissions          LvPermissions
	AllocPolicy          LvAllocPolicy
	FixedMinor           bool
	State                LvState
	Device               LvDeviceStatus
	TargetType           LvTargetType
	ZeroNewBlocksOnAlloc bool
	Health               LvHealth
	SkipActivation       bool
	Raw                  string
}

//var fields = []string{
//	"Volume type",
//	"Permissions",
//	"Allocation policy",
//	"Fixed minor",
//	"State",
//	"Device",
//	"Target type",
//	"Newly-allocated data blocks are zeroed before use",
//	"Health",
//	"Skip activation",
//}

func (attrs *LvAttributes) String() string {
	return fmt.Sprintf(
		"\tVolume type: %s\n"+
			"\tPermissions: %s\n"+
			"\tAllocation policy: %s\n"+
			"\tFixed minor: %t\n"+
			"\tState: %s\n"+
			"\tDevice: %s\n"+
			"\tTarget type: %s\n"+
			"\tNewly-allocated data blocks are zeroed before use: %t\n"+
			"\tHealth: %s\n"+
			"\tSkip activation: %t\n"+
			"\tRaw value: '%s'",
		attrs.VolumeType.String(),
		attrs.Permissions.String(),
		attrs.AllocPolicy.String(),
		attrs.FixedMinor,
		attrs.State.String(),
		attrs.Device.String(),
		attrs.TargetType.String(),
		attrs.ZeroNewBlocksOnAlloc,
		attrs.Health.String(),
		attrs.SkipActivation,
		attrs.Raw,
	)
}
