# libsacloud

This project provides various Go packages to perform operations
on [`SAKURA CLOUD APIs`](http://developer.sakura.ad.jp/cloud/api/1.1/).

[![GoDoc](https://godoc.org/github.com/yamamoto-febc/libsacloud?status.svg)](https://godoc.org/github.com/yamamoto-febc/libsacloud)
[![Build Status](https://travis-ci.org/yamamoto-febc/libsacloud.svg?branch=master)](https://travis-ci.org/yamamoto-febc/libsacloud)

See list of implemented API clients [here](https://godoc.org/github.com/yamamoto-febc/libsacloud).

# Installation

    go get -d github.com/yamamoto-febc/libsacloud

# Sample

This sample is a translation of the examples of [saklient](http://sakura-internet.github.io/saklient.doc/) to golang.

Original sample codes is [here](http://sakura-internet.github.io/saklient.doc/).


##  Create a server

```go

package main

import (
	"fmt"
	API "github.com/yamamoto-febc/libsacloud/api"
	"os"
	"time"
)

func main() {

	// settings
	var (
		token        = os.Args[1]
		secret       = os.Args[2]
		zone         = os.Args[3]
		name         = "libsacloud demo"
		description  = "libsacloud demo description"
		tag          = "libsacloud-test"
		cpu          = 1
		mem          = 2
		hostName     = "libsacloud-test"
		password     = "C8#mf92mp!*s"
		sshPublicKey = "ssh-rsa AAAA..."
	)

	// authorize
	api := API.NewClient(token, secret, zone)

	//search archives
	fmt.Println("searching archives")
	res, _ := api.Archive.
		WithNameLike("CentOS 6.7 64bit").
		WithSharedScope().
		Limit(1).
		Find()

	archive := res.Archives[0]

	// search scripts
	fmt.Println("searching scripts")
	res, _ = api.Note.
		WithNameLike("WordPress").
		WithSharedScope().
		Limit(1).
		Find()
	script := res.Notes[0]

	// create a disk
	fmt.Println("creating a disk")
	disk := api.Disk.New()
	disk.Name = name
	disk.Name = name
	disk.Description = description
	disk.Tags = []string{tag}
	disk.SetDiskPlanToSSD()
	disk.SetSourceArchive(archive.ID)

	disk, _ = api.Disk.Create(disk)

	// create a server
	fmt.Println("creating a server")
	server := api.Server.New()
	server.Name = name
	server.Description = description
	server.Tags = []string{tag}

	// (set ServerPlan)
	plan, _ := api.Product.Server.GetBySpec(cpu, mem)
	server.SetServerPlanByID(plan.ID.String())

	server, _ = api.Server.Create(server)

	// connect to shared segment

	fmt.Println("connecting the server to shared segment")
	iface, _ := api.Interface.CreateAndConnectToServer(server.ID)
	api.Interface.ConnectToSharedSegment(iface.ID)

	// wait disk copy
	err := api.Disk.SleepWhileCopying(disk.ID, 120*time.Second)
	if err != nil {
		fmt.Println("failed")
		os.Exit(1)
	}

	// config the disk
	diskconf := api.Disk.NewCondig()
	diskconf.HostName = hostName
	diskconf.Password = password
	diskconf.SSHKey.PublicKey = sshPublicKey
	diskconf.AddNote(script.ID)
	api.Disk.Config(disk.ID, diskconf)

	// boot
	fmt.Println("booting the server")
	api.Server.Boot(server.ID)

	// stop
	time.Sleep(3 * time.Second)
	fmt.Println("stopping the server")
	api.Server.Stop(server.ID)

	err = api.Server.SleepUntilDown(server.ID, 120*time.Second)
	if err != nil {
		fmt.Println("failed")
		os.Exit(1)
	}

	// disconnect the disk from the server
	fmt.Println("disconnecting the disk")
	api.Disk.DisconnectFromServer(disk.ID)

	// delete the server
	fmt.Println("deleting the server")
	api.Server.Delete(server.ID)

	// delete the disk
	fmt.Println("deleting the disk")
	api.Disk.Delete(disk.ID)

}

```

## Download a disk image

**Pre requirements**
  * install ftps libs. please run `go get github.com/webguerilla/ftps`
  * create a disk named "GitLab"

```go

package main

import (
	"fmt"
	"github.com/webguerilla/ftps"
	API "github.com/yamamoto-febc/libsacloud/api"
	"os"
	"time"
)

func main() {

	// settings
	var (
		token   = os.Args[1]
		secret  = os.Args[2]
		zone    = os.Args[3]
		srcName = "GitLab"
	)

	// authorize
	api := API.NewClient(token, secret, zone)

	// search the source disk
	res, _ := api.Disk.
		WithNameLike(srcName).
		Limit(1).
		Find()
	if res.Count == 0 {
		panic("Disk `GitLab` not found")
	}

	disk := res.Disks[0]

	// copy the disk to a new archive
	fmt.Println("copying the disk to a new archive")

	archive := api.Archive.New()
	archive.Name = fmt.Sprintf("Copy:%s", disk.Name)
	archive.SetSourceDisk(disk.ID)
	archive, _ = api.Archive.Create(archive)
	api.Archive.SleepWhileCopying(archive.ID, 180*time.Second)

	// get FTP information
	ftp, _ := api.Archive.OpenFTP(archive.ID, false)
	fmt.Println("FTP information:")
	fmt.Println("  user: " + ftp.User)
	fmt.Println("  pass: " + ftp.Password)
	fmt.Println("  host: " + ftp.HostName)

	// download the archive via FTPS
	ftpsClient := &ftps.FTPS{}
	ftpsClient.TLSConfig.InsecureSkipVerify = true
	ftpsClient.Connect(ftp.HostName, 21)
	ftpsClient.Login(ftp.User, ftp.Password)
	err := ftpsClient.RetrieveFile("archive.img", "archive.img")
	if err != nil {
		panic(err)
	}
	ftpsClient.Quit()

	// delete the archive after download
	fmt.Println("deleting the archive")
	api.Archive.CloseFTP(archive.ID)
	api.Archive.Delete(archive.ID)

}

```

# License

This project is published under [Apache 2.0 License](LICENSE).

# Author

* Kazumichi Yamamoto ([@yamamoto-febc](https://github.com/yamamoto-febc))
