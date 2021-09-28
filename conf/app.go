package conf

type AppConf struct {
	Port  int        `mapstructure:"port"`
	Debug bool       `mapstructure:"debug"`
	Url   string     `mapstructure:"url"`
	Photo *PhotoConf `mapstructure:"photo"`
	Path  *PathConf  `mapstructure:"path"`
	Auth  *AuthConf  `mapstructure:"auth"`
	Mysql *MYSQLConf `mapstructure:"mysql"`
}

type MYSQLConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
type AuthConf struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectUrl  string `mapstructure:"redirect_url"`
}
type PathConf struct {
	Photo string `mapstructure:"photo"`
	Temp  string `mapstructure:"temp"`
}
type PhotoConf struct {
	Min  int64    `mapstructure:"min"`
	Max  int64    `mapstructure:"max"`
	MIME []string `mapstructure:"mime"`
}
