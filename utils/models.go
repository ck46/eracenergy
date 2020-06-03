package utils

type Config struct {
	Host  string `json:"host"`
	Port  string `json:"port"`
	Dbcon struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DbName   string `json:"dbname"`
	} `json:"database"`
}

type Context struct {
	Title    string
	IsLogin  bool
	Fname    string
	Userid   uint
	Redirect string
	Content  interface{}
}
