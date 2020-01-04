package core
import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

var iniBuffer []byte

func TestReadConfig (t *testing.T) {
	iniFile := "../config.ini"

	var err error
	iniBuffer, err = ioutil.ReadFile(iniFile)

	if err != nil {
		t.Errorf("error returned from ioutil.ReadFile: %s", err.Error())
	}

	var i Ini
	i, err = ReadConfig(iniFile)

	if err != nil {
		t.Errorf("error returned from ReadConfig(%s): %s", iniFile, err.Error())
	}

	loopThroughStruct(i, t)
}

func loopThroughStruct(i interface{}, t *testing.T) {
	v := reflect.ValueOf(i)

	for c := 0; c < v.NumField(); c++ {
		switch v.Field(c).Kind() {
		case reflect.String:
			if v.Field(c).String() != "" {
				if !doesStringExistInIniBuffer(v.Type().Field(c).Name, v.Field(c).String()) {
					keySnake := convertCamelToSnakeCase(v.Type().Field(c).Name)
					t.Errorf("expected '%s = %s' to exist in ini buffer", keySnake, v.Field(c).String())
				}
			}
		case reflect.Int, reflect.Int64:
			if v.Field(c).Int() != 0 {
				if !doesStringExistInIniBuffer(v.Type().Field(c).Name, fmt.Sprintf("%v",v.Field(c).Int())) {
					t.Errorf("expected '%s %v' to exist in ini buffer", v.Type().Field(c).Name, v.Field(c).Int())
				}
			}
		case reflect.Float64:
			if v.Field(c).Float() != 0 {
				if !doesStringExistInIniBuffer(v.Type().Field(c).Name, fmt.Sprintf("%v",v.Field(c).Float())) {
					t.Errorf("expected '%s %v' to exist in ini buffer", v.Type().Field(c).Name, v.Field(c).Float())
				}
			}
		case reflect.Struct:
			loopThroughStruct(v.Field(c).Interface(), t)
		default:
			t.Errorf("unmatched default")
		}
	}
}

func doesStringExistInIniBuffer(key string, value string) (result bool) {
	keyValue := convertCamelToSnakeCase(key) + " = " + value
	if strings.Contains(string(iniBuffer), keyValue) {
		return true
	}

	return false
}

func convertCamelToSnakeCase(value string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(value, "${1}_${2}")
	snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)

}
