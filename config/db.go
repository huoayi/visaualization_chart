package config

type Db struct {
	Mysql Mysql `yaml:"mysql"`
}
type Mysql struct {
	Address 	string 		`yaml:"address" json:"address"`
	Port		string		`yaml:"port" json:"port"`
	UserName 	string		`yaml:"user_name" json:"user_name"`
	Password 	string		`yaml:"password" json:"password"`
	Database 	string 		`yaml:"database" json:"database"`
	DbTable		string		`yaml:"db_table" json:"db_table"`
}

