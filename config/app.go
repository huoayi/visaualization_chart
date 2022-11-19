package config

type App struct {
	Name 	string 	`yaml:"name" json:"name"`
	Host 	string	`yaml:"host" json:"host"`
	Port	int		`yaml:"port" json:"port"`

}