package config

type Profile struct {
	Server Server `json:"server" yaml:"server"`
	Db     Db     `json:"db" yaml:"db"`
	Jwt    Jwt    `json:"jwt" yaml:"jwt"`
	Log    Log    `json:"log" yaml:"log"`
	Cors   Cors   `json:"cors" yaml:"cors"`
}

type Server struct {
	Port      string `yaml:"port" json:"port"`
	Env       string `yaml:"env" json:"env"`
	SignToken string `yaml:"signToken" json:"signToken"`
}

type Db struct {
	Url string `yaml:"url" json:"url"`
}

type Jwt struct {
	Key string `yaml:"key" json:"key"`
}

type Log struct {
	Level      string `yaml:"level" json:"level"`
	Filename   string `yaml:"filename" json:"filename"`
	MaxSize    int    `yaml:"maxSize" json:"maxSize"`
	MaxBackups int    `yaml:"maxBackups" json:"maxBackups"`
	MaxAge     int    `yaml:"maxAge" json:"maxAge"`
	Compress   bool   `yaml:"compress" json:"compress"`
	LocalTime  bool   `yaml:"localTime" json:"localTime"`
}

type Cors struct {
	AllowedOriginPatterns []string `yaml:"allowedOriginPatterns" json:"allowedOriginPatterns"`
	AllowedMethods        string   `yaml:"allowedMethods" json:"allowedMethods"`
	AllowedHeaders        string   `yaml:"allowedHeaders" json:"allowedHeaders"`
	ExposeHeaders         string   `yaml:"exposeHeaders" json:"exposeHeaders"`
	MaxAge                int64    `yaml:"maxAge" json:"maxAge"`
	AllowCredentials      bool     `yaml:"allowCredentials" json:"allowCredentials"`
}
