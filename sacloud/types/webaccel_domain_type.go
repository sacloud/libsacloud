package types

// EWebAccelDomainType ウェブアクセラレータ ドメイン種別
type EWebAccelDomainType string

// WebAccelDomainTypes ウェブアクセラレータ ドメイン種別
var WebAccelDomainTypes = struct {
	// Own 独自ドメイン
	Own EWebAccelDomainType
	// SubDomain サブドメイン
	SubDomain EWebAccelDomainType
}{
	Own:       EWebAccelDomainType("own_domain"),
	SubDomain: EWebAccelDomainType("subdomain"),
}
