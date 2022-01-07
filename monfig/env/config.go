package env

import (
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// Config ...
type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	GetStringMapString(key string) map[string]string
	GetStringMap(key string) map[string]interface{}
	Init()
}

type viperConfig struct {
}

func (v *viperConfig) Init() {
	viper.SetEnvPrefix(`go-clean`)
	viper.AutomaticEnv()

	basePath := GetBasePath()

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigName("config")
	viper.AddConfigPath(basePath)
	//added current folder path too
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (v *viperConfig) GetStringMapString(key string) map[string]string {

	return viper.GetStringMapString(key)
}

// GetStringMap ....
func (v *viperConfig) GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

// NewViperConfig ...
func NewViperConfig() Config {
	v := &viperConfig{}
	v.Init()
	return v
}

//CustomClaims ...
type CustomClaims struct {
	CustomerCode     string `json:"customerCode"`
	VerificationCode string `json:"verificationCode"`
	jwt.StandardClaims
}
