package core

type Ini struct {
	Api Api `json:"api"`
	Database Database `json:"database"`
	Nsc Nsc `json:"nsc"`
	Log Log `json:"log"`
}

type Api struct {
	Port int `json:"port"`
}

type Database struct {
	User string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Host string `json:"host"`
}

type Nsc struct {
	Port int `json:"port"`
	Ip string `json:"ip"`
	GoApiPort int `json:"goApiPort"`
}

type Log struct {
	File string `json:"file"`
}