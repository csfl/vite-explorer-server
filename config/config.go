package config

import "github.com/gin-gonic/gin"

var baseConfig *gin.H

func mergeConfig (config1, config2 gin.H) (gin.H) {
	//var newConfig gin.H
	//for key, value := range config1 {
	//
	//}
	return nil
}

func InitConfig (env string) {
	//baseConfig = &{
	//	"env": "dev",
	//}


}

func Get (key string) (interface{}, bool){
	//configItem, ok := _config[key]
	//if !ok {
	//	return nil, ok
	//}
	//
	//value.
	return nil, nil
}
