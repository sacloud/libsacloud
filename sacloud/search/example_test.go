package search_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
)

func Example() {

	// API Keys
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")
	zone := os.Getenv("SAKURACLOUD_ZONE")

	// API Client
	caller := sacloud.NewClient(token, secret)
	serverOp := sacloud.NewArchiveOp(caller)

	// ******************************************
	// Find
	// ******************************************

	// 名称に"Example"を含むサーバを検索
	condition := &sacloud.FindCondition{
		Filter: search.Filter{
			search.Criterion{Key: search.Key("Name"), Value: search.PartialMatch("Example")},
		},
	}
	searched, err := serverOp.Find(context.Background(), zone, condition)
	if err != nil {
		panic(err)
	}

	fmt.Printf("searched: %#v", searched)

	// 以下の条件で検索
	//   - 名称に"test"と"example"を含む
	//   - ゾーンが"is1a"または"is1b"
	//   - 作成日時が1週間以上前
	condition = &sacloud.FindCondition{
		Filter: search.Filter{
			search.Criterion{Key: search.Key("Name"), Value: search.AndEqual("test", "example")},
			search.Criterion{Key: search.Key("Zone.Name"), Value: search.OrEqual("is1a", "is1b")},
			search.Criterion{
				Key:   search.KeyWithOp("CreatedAt", search.OpLessThan),
				Value: time.Now().Add(-7 * 24 * time.Hour),
			},
		},
	}
	searched, err = serverOp.Find(context.Background(), zone, condition)
	if err != nil {
		panic(err)
	}

	fmt.Printf("searched: %#v", searched)
}
