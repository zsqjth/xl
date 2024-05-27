package controller

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"xl/utils"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 限制为10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.MultipartForm.RemoveAll()

	// 获取文件
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	ext := filepath.Ext(handler.Filename)

	// 生成新的文件名
	newFileName := generateUniqueFileName(ext)

	// 构建保存路径
	savePath := filepath.Join("C:\\Users\\张绍启\\Desktop\\xl\\file", newFileName)

	// 创建目标文件并保存上传的文件
	dest, err := os.Create(savePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	if _, err = io.Copy(dest, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 响应成功信息和文件保存路径
	utils.RespondWithJSON(w, 0, "success", savePath)
}

// 生成新的文件名
func generateUniqueFileName(ext string) string {
	// 使用 UUID 生成唯一的文件名
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "error" + ext
	}
	return uuid.String() + ext
}
