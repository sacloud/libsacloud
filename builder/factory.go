package builder

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud/ostype"
)

func (b *serverBuilder) ServerPublicArchiveUnix(os ostype.ArchiveOSTypes, password string) {
	if !os.IsSupportDiskEdit() {
		b.errors = append(b.errors, fmt.Errorf("%q is not support EditDisk", os))
	}

	archive, err := b.client.Archive.FindByOSType(os)
	if err != nil {
		b.errors = append(b.errors, err)
	}

	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = archive.ID
	b.disk.password = password

}

func (b *serverBuilder) ServerPublicArchiveFixedUnix(os ostype.ArchiveOSTypes) {
	archive, err := b.client.Archive.FindByOSType(os)
	if err != nil {
		b.errors = append(b.errors, err)
	}

	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = archive.ID
}

func (b *serverBuilder) ServerPublicArchiveWindows(os ostype.ArchiveOSTypes) {
	if !os.IsWindows() {
		b.errors = append(b.errors, fmt.Errorf("%q is not windows", os))
	}

	archive, err := b.client.Archive.FindByOSType(os)
	if err != nil {
		b.errors = append(b.errors, err)
	}

	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = archive.ID
	b.disk.sourceDiskID = 0
	b.disk.forceEditDisk = true
}

func (b *serverBuilder) ServerFromDisk(sourceDiskID int64) {
	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = 0
	b.disk.sourceDiskID = sourceDiskID
}

func (b *serverBuilder) ServerFromArchive(sourceArchiveID int64) {

	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = sourceArchiveID
	b.disk.sourceDiskID = 0
}

func (b *serverBuilder) ServerFromBlank() {
	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = 0
	b.disk.sourceDiskID = 0
}
