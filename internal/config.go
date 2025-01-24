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
	ChainPath      string
}
