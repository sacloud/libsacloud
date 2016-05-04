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


```go

package main

import (
	"fmt"
	"github.com/yamamoto-febc/libsacloud/api"
	"github.com/yamamoto-febc/libsacloud/sacloud"
	"os"
	"time"
)

func main() {

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

	client := api.NewClient(token, secret, zone)

	//search archives
	fmt.Println("searching archives")
	res, _ := client.Archive.
		WithNameLike("CentOS 6.7 64bit").
		WithSharedScope().
		Limit(1).
		Find()

	archive := res.Archives[0]

	// search scripts
	fmt.Println("searching scripts")
	res, _ = client.Note.
		WithNameLike("WordPress").
		WithSharedScope().
		Limit(1).
		Find()
	script := res.Notes[0]

	// create a disk
	fmt.Println("creating a disk")
	disk := client.Disk.New()
	disk.Name = name
	disk.Name = name
	disk.Description = description
	disk.Tags = []string{tag}
	disk.Plan = sacloud.DiskPlanSSD
	disk.SetSourceArchive(archive.ID)

	disk, _ = client.Disk.Create(disk)

	// create a server
	fmt.Println("creating a server")
	server := client.Server.New()
	server.Name = name
	server.Description = description
	server.Tags = []string{tag}

	// (set ServerPlan)
	plan, _ := client.Product.Server.GetBySpec(cpu, mem)
	server.SetServerPlanByID(plan.ID.String())

	server, _ = client.Server.Create(server)

	// connect to shared segment

	fmt.Println("connecting the server to shared segment")
	iface, _ := client.Interface.CreateAndConnectToServer(server.ID)
	client.Interface.ConnectToSharedSegment(iface.ID)

	// wait disk copy
	err := client.Disk.SleepWhileCopying(disk.ID, 120*time.Second)
	if err != nil {
		fmt.Println("failed")
		os.Exit(1)
	}

	// config the disk
	diskconf := client.Disk.NewCondig()
	diskconf.HostName = hostName
	diskconf.Password = password
	diskconf.SSHKey.PublicKey = sshPublicKey
	diskconf.AddNote(script.ID)
	client.Disk.Config(disk.ID, diskconf)

	// boot
	fmt.Println("booting the server")
	client.Server.Boot(server.ID)

	// stop
	time.Sleep(3 * time.Second)
	fmt.Println("stopping the server")
	client.Server.Stop(server.ID)

	err = client.Server.SleepUntilDown(server.ID, 120*time.Second)
	if err != nil {
		fmt.Println("failed")
		os.Exit(1)
	}

	// disconnect the disk from the server
	fmt.Println("disconnecting the disk")
	client.Disk.DisconnectFromServer(disk.ID)

	// delete the server
	fmt.Println("deleting the server")
	client.Server.Delete(server.ID)

	// delete the disk
	fmt.Println("deleting the disk")
	client.Disk.Delete(disk.ID)

}

````



# License

This project is published under [Apache 2.0 License](LICENSE).

# Author

* Kazumichi Yamamoto ([@yamamoto-febc](https://github.com/yamamoto-febc))
