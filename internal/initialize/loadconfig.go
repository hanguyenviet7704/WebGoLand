package initialize

import (
	"awesomeProject5/global"
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() {
	v := viper.New()
	v.AddConfigPath("./config/") // Đường dẫn tới thư mục chứa config
	v.SetConfigName("local")     // Tên file config
	v.SetConfigType("yaml")

	// Đọc file cấu hình
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read configuration: %w", err))
	}

	// Debug check
	fmt.Println("Server Port: ", v.GetInt("server.port"))
	fmt.Println("MySQL User: ", v.GetString("mysql.username"))

	// Unmarshal vào global.Config
	if err := v.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("unable to decode config into struct: %w", err))
	}
}
