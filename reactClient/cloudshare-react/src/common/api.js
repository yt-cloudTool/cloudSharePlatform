const baseUrl = "http://localhost:8000"

export default {
	"baseUrl": 			baseUrl,
	"login": 			baseUrl + "/api/user/login", 			// POST 登录
	"register":			baseUrl + "/api/user/register",			// POST 注册
	"checkLogin":		baseUrl + "/api/user/checkLogin",		// 检查登录
	"serverinfo": 		baseUrl + "/api/server/info", 		// GET 服务器信息
	"articleList": 		baseUrl + "/api/article/list", 		// GET 桌面图标(文章列表)
	"articleInsert": 	baseUrl + "/api/article/insert", 	// PSOT 提交文章
	"fileUpload":		baseUrl + "/api/file/upload",		// POST 上传文件
	"createFileBox":	baseUrl + "/api/filebox/create",		// POST 建立文件box
	"fileBoxInsert": 	baseUrl + "/api/filebox/insertInto"	// POST 文件box追加文件
}