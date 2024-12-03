package config

type Config struct {
	Server     Server      `mapstructure:"server"`
	Log        Log         `mapstructure:"logging"`
	PostgreSQL *PostgreSQL `mapstructure:"postgresql"`
	Ink        Ink         `mapstructure:"ink"`
}

type PostgreSQL struct {
	Dsn string `mapstructure:"dsn"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Log struct {
	FileName string `mapstructure:"filename"`
	Levels   Levels `mapstructure:"level"`
	MaxSize  int    `mapstructure:"maxsize"`
	MaxAge   int    `mapstructure:"maxage"`
	Compress bool   `mapstructure:"compress"`
}

type Levels struct {
	App  string `mapstructure:"app"`
	Gorm string `mapstructure:"gorm"`
}

type LogMode string

const (
	Console LogMode = "console"
	File    LogMode = "file"
)

type Ink struct {
	Mode         string  `mapstructure:"mode"`
	LogMode      LogMode `mapstructure:"log_mode"`
	WorkDir      string  `mapstructure:"work_dir"`
	UploadDir    string
	LogDir       string `mapstructure:"log_dir"`
	AdminURLPath string `mapstructure:"admin_url_path"`
}
