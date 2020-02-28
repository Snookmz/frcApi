package core

type Ini struct {
	Api Api `json:"api"`
	Database Database `json:"database"`
	Nsc Nsc `json:"nsc"`
	Nss Nss `json:"nss"`
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
	TimeOut int `json:"timeDuration"`
}

type Nss struct {
	Port int `json:"port"`
	Ip string `json:"ip"`
	TimeOut int `json:"timeOut"`
}

type Log struct {
	File string `json:"file"`
}