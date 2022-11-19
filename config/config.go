package config

type Server struct {
	Db Db	`yaml:"DB" json:"DB"`
	App App	`yaml:"App" json:"App"`
}
