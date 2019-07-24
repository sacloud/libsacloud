package types

const (
	// ZoneTk1aID 東京第1ゾーン
	ZoneTk1aID = ID(21001)
	// ZoneIs1aID 石狩第1ゾーン
	ZoneIs1aID = ID(31001)
	// ZoneIs1bID 石狩第1ゾーン
	ZoneIs1bID = ID(31002)
	// ZoneTk1vID サンドボックスゾーン
	ZoneTk1vID = ID(29001)
)

// ZoneNames 利用できるゾーンの一覧
var ZoneNames = []string{"tk1a", "is1a", "is1b", "tk1v"}

// ZoneIDs ゾーンIDと名称のマップ
var ZoneIDs = map[string]ID{
	"tk1a": ZoneTk1aID,
	"is1a": ZoneIs1aID,
	"is1b": ZoneIs1bID,
	"tk1v": ZoneTk1vID,
}
