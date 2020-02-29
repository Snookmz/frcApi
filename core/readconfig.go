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

	i.Dir = cfg.Section("files").Key("dir").String()

	i.Log.File = cfg.Section("log").Key("file").String()

	i.DropBox.AccessToken = cfg.Section("dropBox").Key("accessToken").String()
	i.DropBox.Path = cfg.Section("dropBox").Key("path").String()
	i.DropBox.Url = cfg.Section("dropBox").Key("url").String()

	return

}
