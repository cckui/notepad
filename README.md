"# notepad" 

# Gin Download
```go
    func FileDownload(c *gin.Context){
    	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
        c.Writer.Header().Add("Content-Type", "application/octet-stream")
    	c.File("./file/a.txt")
    }
```

