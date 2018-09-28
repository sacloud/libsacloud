package api_test

import (
	"fmt"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"

	"os"
	"time"
)

func Example_basic() {
	token := "PUT YOUR TOKEN"
	secret := "PUT YOUR SECRET"
	zone := "tk1a"

	// クライアントの作成
	client := api.NewClient(token, secret, zone)

	// 以降はこのクライアントを通じて操作を行う
	authStatus, err := client.AuthStatus.Read()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Auth status : [%#v]", authStatus)
}

func Example_find() {
	token := "PUT YOUR TOKEN"
	secret := "PUT YOUR SECRET"
	zone := "tk1a"

	// クライアントの作成
	client := api.NewClient(token, secret, zone)

	// サーバーの検索
	res, err := client.Server.
		WithNameLike("server name"). // サーバー名に"server name"が含まれる
		Offset(0).                   // 検索結果の位置0(先頭)から取得
		Limit(5).                    // 5件取得
		Include("Name").             // 結果にName列を含める
		Include("Description").      // 結果にDescription列を含める
		Find()                       // 検索実施

	if err != nil {
		panic(err)
	}

	fmt.Printf("response: %#v", res.Servers)
}

func Example_create() {
	token := "PUT YOUR TOKEN"
	secret := "PUT YOUR SECRET"
	zone := "tk1a"

	// クライアントの作成
	client := api.NewClient(token, secret, zone)

	// スイッチの作成
	param := client.Switch.New()           // 新規作成用パラメーターの作成
	param.Name = "example"                 // 値の設定(名前)
	param.Description = "example"          // 値の設定(説明)
	sw, err := client.Switch.Create(param) // 作成

	if err != nil {
		panic(err)
	}

	fmt.Printf("Created switch: %#v", sw)

}

func Example_wait() {
	token := "PUT YOUR TOKEN"
	secret := "PUT YOUR SECRET"
	zone := "tk1a"

	// クライアントの作成
	client := api.NewClient(token, secret, zone)

	// パブリックアーカイブからディスク作成
	archive, _ := client.Archive.FindLatestStableCentOS()
	// ディスクの作成
	param := client.Disk.New()
	param.Name = "example"                 // 値の設定(名前)
	param.SetSourceArchive(archive.ID)     // コピー元にCentOSパブリックアーカイブを指定
	disk, err := client.Disk.Create(param) // 作成

	if err != nil {
		panic(err)
	}

	// 作成完了まで待機
	err = client.Disk.SleepWhileCopying(disk.ID, client.DefaultTimeoutDuration)

	// タイムアウト発生の場合errが返る
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created disk: %#v", disk)

}

func Example_update() {
	token := "PUT YOUR TOKEN"
	secret := "PUT YOUR SECRET"
	zone := "tk1a"

	// クライアントの作成
	client := api.NewClient(token, secret, zone)

	req, err := client.Server.Find()
	if err != nil {
		panic(err)
	}
	server := &req.Servers[0]

	// 更新
	server.Name = "update"                                        // サーバー名を変更
	server.AppendTag("example-tag")                               // タグを追加
	updatedServer, err := client.Server.Update(server.ID, server) //更新実行
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated server: %#v", updatedServer)

}

func Example_delete() {
	token := "PUT YOUR TOKEN"
	secret := "PUT YOUR SECRET"
	zone := "tk1a"

	// クライアントの作成
	client := api.NewClient(token, secret, zone)

	req, err := client.Switch.Find()
	if err != nil {
		panic(err)
	}
	sw := req.Switches[0]

	// 削除
	deletedSwitch, err := client.Switch.Delete(sw.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted switch: %#v", deletedSwitch)
}

func Example_power() {
	token := "PUT YOUR TOKEN"
	secret := "PUT YOUR SECRET"
	zone := "tk1a"

	// クライアントの作成
	client := api.NewClient(token, secret, zone)

	req, err := client.Server.Find()
	if err != nil {
		panic(err)
	}
	server := req.Servers[0]

	// 起動
	_, err = client.Server.Boot(server.ID)
	if err != nil {
		panic(err)
	}

	// 起動完了まで待機
	err = client.Server.SleepUntilUp(server.ID, client.DefaultTimeoutDuration)

	// シャットダウン
	_, err = client.Server.Shutdown(server.ID) // gracefulシャットダウン
	//_, err = client.Server.Stop(server.ID)   // forceシャットダウン

	if err != nil {
		panic(err)
	}

	// ダウンまで待機
	err = client.Server.SleepUntilDown(server.ID, client.DefaultTimeoutDuration)
	if err != nil {
		panic(err)
	}
}

func Example() {

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
	client := api.NewClient(token, secret, zone)

	//search archives
	fmt.Println("searching archives")
	archive, _ := client.Archive.FindLatestStableCentOS()

	// search scripts
	fmt.Println("searching scripts")
	res, _ := client.Note.
		WithNameLike("WordPress").
		WithSharedScope().
		Limit(1).
		Find()
	script := res.Notes[0]

	// create a disk
	fmt.Println("creating a disk")
	disk := client.Disk.New()
	disk.Name = name
	disk.Description = description
	disk.Tags = []string{tag}
	disk.SetDiskPlanToSSD()
	disk.SetSourceArchive(archive.ID)

	disk, _ = client.Disk.Create(disk)

	// create a server
	fmt.Println("creating a server")
	server := client.Server.New()
	server.Name = name
	server.Description = description
	server.Tags = []string{tag}

	// (set ServerPlan)
	plan, _ := client.Product.Server.GetBySpec(cpu, mem, sacloud.PlanDefault)
	server.SetServerPlanByID(plan.GetStrID())

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
	diskConf := client.Disk.NewCondig()
	diskConf.SetHostName(hostName)
	diskConf.SetPassword(password)
	diskConf.AddSSHKeyByString(sshPublicKey)
	diskConf.AddNote(script.GetStrID())
	client.Disk.Config(disk.ID, diskConf)

	// connect to server
	client.Disk.ConnectToServer(disk.ID, server.ID)

	// boot
	fmt.Println("booting the server")
	client.Server.Boot(server.ID)

	// stop
	time.Sleep(3 * time.Second)
	fmt.Println("stopping the server")
	client.Server.Stop(server.ID)

	// wait for server to down
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
