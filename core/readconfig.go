package core

import (
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
)

func ReadConfig (iniFile string) (i Ini, err error) {
	var cfg *ini.File // a parsed ini file

	cfg, err = ini.Load(iniFile)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from ini.Load(%s): %s", iniFile, err.Error()))
		return
	}

	i.Api.Port = cfg.Section("api").Key("port").MustInt(8002)

	i.Database.User = cfg.Section("database").Key("user").String()
	i.Database.Password = cfg.Section("database").Key("password").String()
	i.Database.Database = cfg.Section("database").Key("database").String()
	i.Database.Host = cfg.Section("database").Key("host").String()

	i.Nsc.Port = cfg.Section("nsc").Key("port").MustInt(4213)
	i.Nsc.Ip = cfg.Section("nsc").Key("ip").String()
	i.Nsc.GoApiPort = cfg.Section("nsc").Key("go_api_port").MustInt(8002)
	i.Nsc.TimeOut = cfg.Section("nsc").Key("timeDuration").MustInt(2000)

	i.Nss.Port = cfg.Section("nss").Key("port").MustInt(9999)
	i.Nss.Ip = cfg.Section("nss").Key("ip").String()
	i.Nss.TimeOut = cfg.Section("nss").Key("timeDuration").MustInt(2000)

	i.Log.File = cfg.Section("log").Key("file").String()



	return

}
