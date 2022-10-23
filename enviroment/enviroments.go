package enviroment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type enviromentVariables struct {
	Port       string `json:"port"`
	DirPath    string `json:"dirpath"`
	JwtSecret  string `json:"jwtsecret"`
	AdminLogin string `json:"adminlogin"`
	AdminPwd   string `json:"adminpwd"`
}

func SetEnviromentVariables() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	data := &enviromentVariables{}
	_ = json.Unmarshal([]byte(file), &data)

	reflectedStruct := reflect.ValueOf(data).Elem()
	for i := 0; i < reflectedStruct.NumField(); i++ {
		varName := reflectedStruct.Type().Field(i).Name
		varValue := reflectedStruct.Field(i).Interface()
		os.Setenv(varName, fmt.Sprintf("%v", varValue))
	}
}
