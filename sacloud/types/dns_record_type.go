package types

// EDNSRecordType DNSレコード種別
type EDNSRecordType string

// DNSRecordTypes DNSレコード種別
var DNSRecordTypes = struct {
	// Unknown 不明
	Unknown EDNSRecordType
	// A Aレコード
	A EDNSRecordType
	// AAAA AAAAレコード
	AAAA EDNSRecordType
	// ALIAS ALIASレコード
	ALIAS EDNSRecordType
	// CNAME CNAMEレコード
	CNAME EDNSRecordType
	// NS NSレコード
	NS EDNSRecordType
	// MX MXレコード
	MX EDNSRecordType
	// TXT TXTレコード
	TXT EDNSRecordType
	// SRV SRVレコード
	SRV EDNSRecordType
	// CAA CAAレコード
	CAA EDNSRecordType
}{
	Unknown: EDNSRecordType(""),
	A:       EDNSRecordType("A"),
	AAAA:    EDNSRecordType("AAAA"),
	ALIAS:   EDNSRecordType("ALIAS"),
	CNAME:   EDNSRecordType("CNAME"),
	NS:      EDNSRecordType("NS"),
	MX:      EDNSRecordType("MX"),
	TXT:     EDNSRecordType("TXT"),
	SRV:     EDNSRecordType("SRV"),
	CAA:     EDNSRecordType("CAA"),
}

// String EDNSRecordTypeの文字列表現
func (t EDNSRecordType) String() string {
	return string(t)
}

// DNSRecordTypesStrings 有効なDNSレコードタイプを示す文字列のリスト
//
// Unknown(空文字)は含まない
func DNSRecordTypesStrings() []string {
	return []string{
		DNSRecordTypes.A.String(),
		DNSRecordTypes.AAAA.String(),
		DNSRecordTypes.ALIAS.String(),
		DNSRecordTypes.CNAME.String(),
		DNSRecordTypes.NS.String(),
		DNSRecordTypes.MX.String(),
		DNSRecordTypes.TXT.String(),
		DNSRecordTypes.SRV.String(),
		DNSRecordTypes.CAA.String(),
	}
}
