package zoom

import (
	config "cloudSharePlatform/config"
	db "cloudSharePlatform/db"
	utils "cloudSharePlatform/utils"

	// json "encoding/json"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	// "time"

	// "strconv"

	gin "github.com/gin-gonic/gin"
	// bson "go.mongodb.org/mongo-driver/bson"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// 存储成功项类型定义
type fileSuccStore struct {
	FileName      string `json:"filename"`
	StoreFileName string `json:"store_filename"`
	FileSize      int64  `json:"size"`
	FileId        string `json:"fileid"`
}

/*
   header传token
   params:
        count: int64 文件数量(可选)
        is_tmp: int 是否是临时文件 0:否 1:是
        file: fileObj[]

*/
func FileUpload(c *gin.Context) {

	user_id, isExist := c.Get("user_id")

	/*
	   如果存在user_id则检索此用户的文章
	   如果不存在user_id则返回空
	*/
	if isExist == false {
		c.JSON(200, gin.H{"status": 401, "message": "no user_id", "data": ""})
		return
	}

	// -------------------------------------------------------------------------
	// 参数
	param_count := c.PostForm("count")
	// isTmp参数必须
	param_isTmp := c.PostForm("is_tmp")
	if param_isTmp == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough: is_tmp", "data": ""})
		return
	}
	param_isTmp_int, err := strconv.Atoi(param_isTmp)
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "params is_tmp wrong format", "data": ""})
		return
	}
	// isPub参数必须
	param_isPub := c.PostForm("is_pub")
	if param_isPub == "" {
		c.JSON(500, gin.H{"status": 0, "message": "params not enough: is_pub", "data": ""})
		return
	}
	param_isPub_int, err := strconv.Atoi(param_isPub)
	if err != nil {
		c.JSON(500, gin.H{"status": 0, "message": "params is_pub wrong format", "data": ""})
		return
	}
	// -------------------------------------------------------------------------

	// Multipart form
	multiForm, err := c.MultipartForm()
	if err != nil {
		c.JSON(500, gin.H{"status": -1, "message": "get form err", "data": ""})
		return
	}

	// file arr
	fileArr := multiForm.File["file"]

	// 存储成功的文件名/id数组
	var succStoreFileArr []fileSuccStore
	// 存储出错的文件名数组
	var errStoreFileArr []string = []string{}

	// 锁
	var lock sync.Mutex
	// 同步等待组
	var wg sync.WaitGroup

	// 遍历文件info
	for _, file := range fileArr {
		wg.Add(1)

		go func(file *(multipart.FileHeader)) {

			fileName := filepath.Base(file.Filename)
			fileSize := file.Size

			// 存储用的文件名
			storeFileName := utils.SnowflakeGenerate() + "_" + fileName

			// 创建目录
			storeFilePath := config.GetFileStorePath() + user_id.(primitive.ObjectID).Hex() + "/"
			os.MkdirAll(storeFilePath, 0766)

			// 存储文件
			err := c.SaveUploadedFile(file, storeFilePath+storeFileName)
			if err != nil {
				lock.Lock()
				defer lock.Unlock()
				// 存储失败的数组追加
				errStoreFileArr = append(errStoreFileArr, fileName)
				wg.Done()
				return
			}
			fmt.Println("fileName ================>", param_count+fileName)

			// 记录存储到数据库
			MongoFile := db.MongoFile{
				Id_:           primitive.NewObjectID(),
				User_id:       user_id.(primitive.ObjectID),
				IsTmp:         param_isTmp_int,
				IsPub:         param_isPub_int,
				FileName:      fileName,
				StoreFileName: storeFileName,
				Size:          fileSize,
			}

			dbResult, err := db.MongoInsertOne("cloudshareplatform", "file", MongoFile)
			if err != nil {
				// 数据库存储失败也要文件存储失败 失败处理
				lock.Lock()
				defer lock.Unlock()
				errStoreFileArr = append(errStoreFileArr, fileName)
				wg.Done()
				return
			}
			fmt.Println("dbResult ===========>", dbResult)

			// 存储成功的数组追加
			succIte := fileSuccStore{
				FileName:      fileName,
				StoreFileName: storeFileName,
				FileSize:      fileSize,
				FileId:        dbResult.InsertedID.(primitive.ObjectID).Hex(),
			}
			lock.Lock()
			defer lock.Unlock()
			succStoreFileArr = append(succStoreFileArr, succIte)
			wg.Done()
		}(file)
	}

	// 等待执行完成
	wg.Wait()

	if len(errStoreFileArr) != 0 {
		c.JSON(500, gin.H{"status": -2, "message": "store file err", "data": errStoreFileArr})
	} else {
		c.JSON(200, gin.H{"status": 1, "message": "ok", "data": succStoreFileArr})
	}

}
