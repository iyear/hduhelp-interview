package conf

type AppConf struct {
	Port  int    `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
	Url   string `mapstructure:"url"`
	Limit int    `mapstructure:"limit"`
	Photo *Photo `mapstructure:"photo"`
	Path  *Path  `mapstructure:"path"`
	Auth  *Auth  `mapstructure:"auth"`
	Mysql *Mysql `mapstructure:"mysql"`
}

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
type Auth struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectUrl  string `mapstructure:"redirect_url"`
}
type Path struct {
	Photo string `mapstructure:"photo"`
	Temp  string `mapstructure:"temp"`
}
type Photo struct {
	Min  int64    `mapstructure:"min"`
	Max  int64    `mapstructure:"max"`
	MIME []string `mapstructure:"mime"`
}
