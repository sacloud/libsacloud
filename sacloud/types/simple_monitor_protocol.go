package types

// ESimpleMonitorProtocol シンプル監視 プロトコル
type ESimpleMonitorProtocol string

// SimpleMonitorProtocols シンプル監視 プロトコル
var SimpleMonitorProtocols = struct {
	// Unknown 不明
	Unknown ESimpleMonitorProtocol
	// HTTP http
	HTTP ESimpleMonitorProtocol
	// HTTPS https
	HTTPS ESimpleMonitorProtocol
	// Ping ping
	Ping ESimpleMonitorProtocol
	// TCP tcp
	TCP ESimpleMonitorProtocol
	// DNS dns
	DNS ESimpleMonitorProtocol
	// SSH ssh
	SSH ESimpleMonitorProtocol
	// SMTP smtp
	SMTP ESimpleMonitorProtocol
	// POP3 pop3
	POP3 ESimpleMonitorProtocol
	// SNMP snmp
	SNMP ESimpleMonitorProtocol
	// SSLCertificate sslcertificate
	SSLCertificate ESimpleMonitorProtocol
}{
	Unknown:        ESimpleMonitorProtocol(""),
	HTTP:           ESimpleMonitorProtocol("http"),
	HTTPS:          ESimpleMonitorProtocol("https"),
	Ping:           ESimpleMonitorProtocol("ping"),
	TCP:            ESimpleMonitorProtocol("tcp"),
	DNS:            ESimpleMonitorProtocol("dns"),
	SSH:            ESimpleMonitorProtocol("ssh"),
	SMTP:           ESimpleMonitorProtocol("smtp"),
	POP3:           ESimpleMonitorProtocol("pop3"),
	SNMP:           ESimpleMonitorProtocol("snmp"),
	SSLCertificate: ESimpleMonitorProtocol("sslcertificate"),
}
