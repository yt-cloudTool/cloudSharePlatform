const baseUrl = "http://localhost:8000"

export default {
	"baseUrl": 			baseUrl,
	"login": 			baseUrl + "/api/login", 			//  POST 登录
	"serverinfo": 		baseUrl + "/api/serverinfo", 		// 服务器信息
	"articleList": 		baseUrl + "/api/articlelist", 		// 桌面图标(文章列表)
	"articleInsert": 	baseUrl + "/api/articleinsert", 	// 提交文章
}