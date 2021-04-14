package utils

// 生成id
func SnowflakeGenerate() string {
	return snowflakeNodeVar.Generate().String()
}
