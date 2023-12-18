package Controller

import (
	"WebBack/Util"
	"WebBack/dao"
	"baliance.com/gooxml/document"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

type EnterMessage struct {
	Id       int32  `json:"Id"`
	Password string `json:"Password"`
}
type NewCountMessage struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
}
type Answer struct {
	Files []string `json:"Files"`
}
type DownMessage struct {
	TargetFolderName string `json:"targetFolderName"`
	Index            int    `json:"index"`
}
type KeySearchFile struct {
	TargetFolderName string `json:"targetfoldername"`
	KeyWord          string `json:"keyword"`
	Index            int    `json:"index"`
}

type FileArray struct {
	Files []int    `json:"Files"`
	Names []string `json:"Names"`
}

func UserLogin(c *gin.Context) {
	var body EnterMessage
	if err := c.BindJSON(&body); err != nil {
		fmt.Println("Error:", err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ans := dao.CheckExist(body.Id, body.Password)
	if ans {
		c.JSON(200, gin.H{"message": "Login successful"})
	} else {
		c.JSON(200, gin.H{"message": "password or id is wrong"})
	}
}
func NewCount(c *gin.Context) {
	var body NewCountMessage
	if err := c.BindJSON(&body); err != nil {
		fmt.Println("Error:", err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	Count := dao.Regist(body.Password, body.Name)
	c.JSON(200, gin.H{"message": Count})
}
func GetClient(c *gin.Context) {
	fmt.Println("Clinet")
	filePath := "F:/nginx-1.22.0-tlias.zip"
	dao.DownLoadCount()
	c.File(filePath)
}
func TestJson(c *gin.Context) {
	c.JSON(200, "OK")
}
func GetAllFilesName(c *gin.Context) {
	var TargetFoldName KeySearchFile
	c.BindJSON(&TargetFoldName)
	TargetName := TargetFoldName.TargetFolderName
	files := Util.GetFiles(TargetName)
	var FileNames []string
	for i := 0; i < len(files); i++ {
		FileNames = append(FileNames, files[i].Name())
		fmt.Println(files[i].Name())
	}
	var FileAnswer Answer
	FileAnswer.Files = FileNames
	c.JSON(200, FileAnswer)
}
func DownLoadFile(c *gin.Context) {
	var UserDownLoadIndex DownMessage
	c.BindJSON(&UserDownLoadIndex)
	Index := UserDownLoadIndex.Index

	files := Util.GetFiles(UserDownLoadIndex.TargetFolderName)
	if files == nil {
		c.JSON(404, "Can not Find")
	}
	if Index >= len(files) || Index < 0 {
		c.JSON(404, "Out of Index")
	}
	c.File(UserDownLoadIndex.TargetFolderName + "/" + string(files[UserDownLoadIndex.Index].Name()))
}
func Search(c *gin.Context) {
	var Keyword KeySearchFile
	c.ShouldBindJSON(&Keyword)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	files := Util.GetFiles(currentDir + "/" + Keyword.TargetFolderName)
	answer := []int{}
	filenames := []string{}
	for i := 0; i < len(files); i++ { //有txt
		fmt.Println(files[i].Name())
		last := ".docx"
		if !strings.Contains(files[i].Name(), last) {
			continue
		}
		doc, _ := document.Open(currentDir + "/" + Keyword.TargetFolderName + "/" + files[i].Name())
		var Temp string
		for _, para := range doc.Paragraphs() {
			for _, run := range para.Runs() {
				Temp += string(run.Text())
			}
		}
		if strings.Contains(Temp, Keyword.KeyWord) {
			answer = append(answer, i)
			filenames = append(filenames, files[i].Name())
		}
	}
	var fileArray FileArray
	fileArray.Files = answer
	fileArray.Names = filenames
	c.JSON(200, fileArray)
}

func DownLoadByIndex(c *gin.Context) {
	var Keyword KeySearchFile
	c.BindJSON(&Keyword)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	files := Util.GetFiles(currentDir + Keyword.TargetFolderName)

	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}

	filePath := currentDir + Keyword.TargetFolderName + "/" + files[Keyword.Index].Name()
	c.File(filePath)
}

func ReceiveFile(c *gin.Context) {

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	file, err := c.FormFile("file")
	uploadInfoStr := c.PostForm("uploadInfo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存文件到服务器
	fmt.Println(file.Filename)
	fmt.Println(currentDir + "/" + uploadInfoStr)
	dst := fmt.Sprintf("%s/%s", uploadInfoStr, file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("Failed")
		return
	}
	fmt.Println("Success")
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
