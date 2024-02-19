package learnviper

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	var vp *viper.Viper = viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	var err error = vp.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "Habi-Viper", vp.GetString("app.name"))
	assert.Equal(t, "Ahmad Rizky Nusantara Habibi", vp.GetString("app.author"))
	assert.Equal(t, "localhost", vp.GetString("database.host"))
	assert.Equal(t, 3306, vp.GetInt("database.port"))
	assert.Equal(t, true, vp.GetBool("database.show_sql"))

	fmt.Println("App name:", vp.GetString("app.name"))
	fmt.Println("App:", vp.Get("app"))
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var vp *viper.Viper = viper.New()
		vp.SetConfigName("config")
		vp.SetConfigType("json")
		vp.AddConfigPath(".")

		_ = vp.ReadInConfig()
	}
}

func TestYAML(t *testing.T) {
	var vp *viper.Viper = viper.New()
	//config.SetConfigName("config")
	//config.SetConfigType("yaml")
	vp.SetConfigFile("config.yml")
	vp.AddConfigPath(".")

	// read config
	err := vp.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "Habi-Viper", vp.GetString("app.name"))
	assert.Equal(t, "Ahmad Rizky Nusantara Habibi", vp.GetString("app.author"))
	assert.Equal(t, "localhost", vp.GetString("database.host"))
	assert.Equal(t, 3306, vp.GetInt("database.port"))
	assert.Equal(t, true, vp.GetBool("database.show_sql"))

	fmt.Println("App name:", vp.GetString("app.name"))
	fmt.Println("App:", vp.Get("app"))
}

func BenchmarkYAML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var vp *viper.Viper = viper.New()
		vp.SetConfigName("config")
		vp.SetConfigType("yml")
		vp.AddConfigPath(".")

		_ = vp.ReadInConfig()
	}
}

func TestENVFile(t *testing.T) {
	var vp *viper.Viper = viper.New()
	vp.SetConfigFile("vp.env")
	vp.AddConfigPath(".")

	// read config
	err := vp.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "Habi-Viper", vp.GetString("APP_NAME"))
	assert.Equal(t, "Ahmad Rizky Nusantara Habibi", vp.GetString("APP_AUTHOR"))
	assert.Equal(t, "localhost", vp.GetString("DATABASE_HOST"))
	assert.Equal(t, 3306, vp.GetInt("DATABASE_PORT"))
	assert.Equal(t, true, vp.GetBool("DATABASE_SHOW_SQL"))
}

type ConfigENV struct {
	AppName    string `mapstructure:"APP_NAME"`
	AppAuthor  string `mapstructure:"APP_AUTHOR"`
	AppVersion string `mapstructure:"APP_VERSION"`

	DatabaseHost string `mapstructure:"DATABASE_HOST"`
	DatabasePort int    `mapstructure:"DATABASE_PORT"`
}

func TestENV(t *testing.T) {
	vp := viper.New()
	vp.SetConfigFile("config.env")
	// vp.SetConfigType("env")
	vp.AddConfigPath(".")
	vp.AutomaticEnv()
	// vp.BindEnv("APP_NAME", "APP_AUTHOR", "APP_VERSION", "DATABASE_HOST", "DATABASE_PORT")
	vp.BindEnv("AppName")

	if err := viper.ReadInConfig(); err != nil {
		t.Error(err)
	}

	var config ConfigENV
	if err := viper.Unmarshal(&config); err != nil {
		t.Error(err)
	}

	fmt.Println("App Name:", config.AppName)
	fmt.Println("Database Host:", config.DatabaseHost)

	fmt.Println("Config:", config)
}
