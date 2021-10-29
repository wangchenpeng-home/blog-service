package setting

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  string
	WriteTimeout string
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType        string
	Username      string
	Password      string
	Host          string
	DBName        string
	TablePrefix   string
	Charset       string
	ParseTime     bool
	MaxIdleConnes int
	MaxOpenConnes int
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
