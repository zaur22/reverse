package config

import (
	"github.com/spf13/viper"
	"path/filepath"
)

func GetString(name string) string {
	return viper.GetString(name)
}

func GetInt32(name string) int32 {
	return viper.GetInt32(name)
}

const (
	//ServerAddr адрес запуска сервиса
	ServerAddr = "server_addr"

	//UtilWorkersMaxCount количество максимальных потоков,
	//которые работают с util
	UtilWorkersMaxCount = "util_worker_max_count"

	//UtilPath путь для запуска сервиса util
	UtilPath = "util_path"
)

func init(){
	viper.AutomaticEnv()
	viper.SetDefault(ServerAddr, "localhost:3000")
	viper.SetDefault(UtilPath, filepath.Join("..", "util", "util"))
	viper.SetDefault(UtilWorkersMaxCount, 30)
}