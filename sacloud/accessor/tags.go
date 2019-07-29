package accessor

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// Tags Tagsを持つリソース向けのインターフェース
type Tags interface {
	GetTags() types.Tags
	SetTags(v types.Tags)
}

// HasTag 指定のタグが存在する場合trueを返す
func HasTag(target Tags, tag string) bool {
	tags := target.GetTags()
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

// AppendTag 指定のタグを追加
func AppendTag(target Tags, tag string) {
	if HasTag(target, tag) {
		return
	}
	tags := target.GetTags()
	target.SetTags(append(tags, tag))
}

// RemoveTag 指定のタグを削除
func RemoveTag(target Tags, tag string) {
	if !HasTag(target, tag) {
		return
	}

	tags := target.GetTags()
	nt := types.Tags{}
	for _, t := range tags {
		if t != tag {
			nt = append(nt, t)
		}
	}
	target.SetTags(nt)
}

// ClearTags 全タグをクリア
func ClearTags(target Tags) {
	target.SetTags(types.Tags{})
}
