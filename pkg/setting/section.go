package setting

type Config struct {
	Mysql  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
}
type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}
type LoggerSetting struct {
	Log_Level     string `mapstructure:"log_level"`     // Mức độ log
	File_log_name string `mapstructure:"file_log_name"` // đường dẫn và tên file log
	Max_backups   int    `mapstructure:"max_backups"`   // số lượng file log cũ tối đa được giữ lại
	Max_age       int    `mapstructure:"max_age"`       // số ngày tối đa lưu trữ file log
	Max_size      int    `mapstructure:"max_size"`      // kích thước tối đa của một file log
	Compress      bool   `mapstructure:"compress"`      // có nên nén file log cũ hay không nếu true sẽ tiết kiệm ổ đĩa
}
