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
	CNAME:   EDNSRecordType("CNAME"),
	NS:      EDNSRecordType("NS"),
	MX:      EDNSRecordType("MX"),
	TXT:     EDNSRecordType("TXT"),
	SRV:     EDNSRecordType("SRV"),
	CAA:     EDNSRecordType("CAA"),
}
