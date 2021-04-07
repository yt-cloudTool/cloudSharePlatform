const baseUrl = "http://localhost:8000"

export default {
	"baseUrl": 			baseUrl,
	"login": 			baseUrl + "/api/login", 			// POST 登录
	"register":			baseUrl + "/api/register",		// POST 注册
	"serverinfo": 		baseUrl + "/api/serverinfo", 		// GET 服务器信息
	"articleList": 		baseUrl + "/api/articlelist", 	// GET 桌面图标(文章列表)
	"articleInsert": 	baseUrl + "/api/articleinsert", 	// PSOT 提交文章
}