package internal

type Config struct {
	RoleID         string
	SecretID       string
	VaultAddress   string
	AppRolePath    string
	CronExpression string
	ServerName     string
	CertPath       string
	KeyPath        string
	CaChainPath    string
	PkiPath        string
	PkiRole        string
	PkiIssuer      string
	IpAddresses    []string
	CertTtl        string
}
