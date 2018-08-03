package setup

import (
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/sacloud"
)

// MaxRetryCountExceededError リトライ最大数超過エラー
type MaxRetryCountExceededError error

// DefaultMaxRetryCount デフォルトリトライ最大数
const DefaultMaxRetryCount = 3

// DefaultDeleteRetryCount リソースごとの削除API呼び出しのリトライ最大数
const DefaultDeleteRetryCount = 10

// DefaultDeleteWaitInterval リソースごとの削除API呼び出しのリトライ間隔
const DefaultDeleteWaitInterval = 10 * time.Second

// CreateFunc リソース作成関数
type CreateFunc func() (sacloud.ResourceIDHolder, error)

// AsyncWaitForCopyFunc リソース作成時のコピー待ち(非同期)関数
type AsyncWaitForCopyFunc func(id int64) (
	chan interface{}, chan interface{}, chan error,
)

// ProvisionBeforeUpFunc リソース作成後、起動前のプロビジョニング関数
//
// リソース作成後に起動が行われないリソース(VPCルータなど)向け。
// 必要であればこの中でリソース起動処理を行う。
type ProvisionBeforeUpFunc func(target interface{}) error

// DeleteFunc リソース削除関数。
//
// リソース作成時のコピー待ちの間にリソースのAvailabilityがFailedになった場合に利用される。
type DeleteFunc func(id int64) error

// WaitForUpFunc リソース起動待ち関数
type WaitForUpFunc func(id int64) error

// RetryableSetup リソース作成時にコピー待ちや起動待ちが必要なリソースのビルダー。
//
// リソースのビルドの際、必要に応じてリトライ(リソースの削除&再作成)を行う。
type RetryableSetup struct {
	// Create リソース作成用関数
	Create CreateFunc
	// AsyncWaitForCopy コピー待ち用関数
	AsyncWaitForCopy AsyncWaitForCopyFunc
	// ProvisionBeforeUp リソース起動前のプロビジョニング関数
	ProvisionBeforeUp ProvisionBeforeUpFunc
	// Delete リソース削除用関数
	Delete DeleteFunc
	// WaitForUp リソース起動待ち関数
	WaitForUp WaitForUpFunc
	// RetryCount リトライ回数
	RetryCount int
	// DeleteRetryCount 削除リトライ回数
	DeleteRetryCount int
	// DeleteRetryInterval 削除リトライ間隔
	DeleteRetryInterval time.Duration
}

type hasFailed interface {
	IsFailed() bool
}

// Setup リソースのビルドを行う。必要に応じてリトライ(リソースの削除&再作成)を行う。
func (r *RetryableSetup) Setup() (interface{}, error) {

	max := r.RetryCount
	if max == 0 {
		max = DefaultMaxRetryCount
	}
	if r.DeleteRetryCount == 0 {
		r.DeleteRetryCount = DefaultDeleteRetryCount
	}
	if r.DeleteRetryInterval == 0 {
		r.DeleteRetryInterval = DefaultDeleteWaitInterval
	}

	var created interface{}
	for cur := 0; cur < max; cur++ {

		target, err := r.Create()
		if err != nil {
			return nil, err
		}

		id := target.GetID()

		// コピー待ち
		if r.AsyncWaitForCopy != nil {
			//wait
			compChan, progChan, errChan := r.AsyncWaitForCopy(id)
			var state interface{}
			var err error

		loop:
			for {
				select {
				case v := <-compChan:
					state = v
					break loop
				case v := <-progChan:
					state = v
				case e := <-errChan:
					err = e
					break loop
				}
			}

			if state != nil {
				// Availabilityを持ち、Failedになっていた場合はリソースを削除してリトライ
				if f, ok := state.(hasFailed); ok && f.IsFailed() {

					// FailedになったばかりだとDelete APIが失敗する(コピー進行中など)場合があるため、
					// 任意の回数リトライ&待機を行う
					for i := 0; i < r.DeleteRetryCount; i++ {
						if err = r.Delete(id); err == nil {
							break
						}
						time.Sleep(r.DeleteRetryInterval)
					}

					continue
				}

				created = state
			}
			if err != nil {
				return nil, err
			}
		} else {
			created = target
		}

		if created != nil {
			if r.ProvisionBeforeUp != nil {
				if err = r.ProvisionBeforeUp(created); err != nil {
					return nil, err
				}
			}
			if r.WaitForUp != nil {
				if err = r.WaitForUp(id); err != nil {
					return nil, err
				}
			}
			break
		}
	}

	if created == nil {
		return nil, MaxRetryCountExceededError(fmt.Errorf("max retry count exceeded"))
	}
	return created, nil
}
