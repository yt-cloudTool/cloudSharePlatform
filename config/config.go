package config

var fileStorePath string = "fileStore/" // 存储文件路径(以可执行文件位置为相对路径)

// -----------------------------------------------------------------------------
// 存储文件路径
func SetFileStorePath(path string) {
	fileStorePath = path
}

func GetFileStorePath() string {
	return fileStorePath
}

// -----------------------------------------------------------------------------
