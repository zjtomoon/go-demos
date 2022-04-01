package config

type searchReq struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
}

type searchRes struct {
	FilePath string
	FileSize int64
}
