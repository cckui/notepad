package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type fileInfo struct {
	FileName string      `json:"FileName"`
	Size     int64       `json:"Size"`
	IsDir    bool        `json:"IsDir"`
	Mode     fs.FileMode `json:"Mode"`
	ModTime  time.Time   `json:"ModTime"`
}

func GetAllFileDetial(pathname string, s []fileInfo) ([]fileInfo, error) {

	// fileDetial := []fileInfo{}

	fromSlash := filepath.FromSlash(pathname)
	//fmt.Println(fromSlash)
	rd, err := ioutil.ReadDir(fromSlash)
	if err != nil {
		//log.LOGGER("Error").Error("read dir fail %vn", err)
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := filepath.Join(fromSlash, fi.Name())
			temp := fileInfo{}

			fi, err := os.Stat(fullDir)
			if err == nil {
				temp.FileName = fi.Name()
				temp.Size = fi.Size()
				temp.IsDir = fi.IsDir()
				temp.Mode = fi.Mode()
				temp.ModTime = fi.ModTime()
			}

			s = append(s, temp)

			s, err = GetAllFileDetial(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				//log.LOGGER("Error").Error("read dir fail: %vn", err)
				return s, err
			}
		} else {
			fullName := filepath.Join(fromSlash, fi.Name())
			temp := fileInfo{}

			fi, err := os.Stat(fullName)
			if err == nil {
				temp.FileName = fi.Name()
				temp.Size = fi.Size()
				temp.IsDir = fi.IsDir()
				temp.Mode = fi.Mode()
				temp.ModTime = fi.ModTime()
			}

			s = append(s, temp)
		}
	}
	return s, nil
}

func GetAllFile(pathname string, s []string) ([]string, error) {
	fromSlash := filepath.FromSlash(pathname)
	//fmt.Println(fromSlash)
	rd, err := ioutil.ReadDir(fromSlash)
	if err != nil {
		//log.LOGGER("Error").Error("read dir fail %vn", err)
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := filepath.Join(fromSlash, fi.Name())
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				//log.LOGGER("Error").Error("read dir fail: %vn", err)
				return s, err
			}
		} else {
			fullName := filepath.Join(fromSlash, fi.Name())
			s = append(s, fullName)
		}
	}
	return s, nil
}

func GetALLFIles_walk(pathname string) []string {
	StartTime := time.Now()
	dst_target := []string{}
	err := filepath.Walk(pathname, func(src string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		} else { //進行檔案的複製
			dst_target = append(dst_target, src)

			//return s
		}
		//println(path)
		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %vn", err)
		return nil
		//log.LOGGER("Error").Error("filepath.Walk() returned %vn", err)
	}
	fmt.Println("Cost Time:", time.Since(StartTime))
	return dst_target
}

func main() {
	s := []string{}
	n := []string{}
	test := []fileInfo{}

	pathname := "D:\\_ProtableTools\\_SyncBackPro64_Setup_NI"

	fmt.Printf("filepath walk cost time returned \n")
	n = GetALLFIles_walk(pathname)
	fmt.Println("the number of file is %v,content is:%v", len(n), n)

	fmt.Printf("io util cost time returned \n")
	StartTime := time.Now()
	s, _ = GetAllFile(pathname, s)
	fmt.Println("Cost Time:", time.Since(StartTime))
	fmt.Println("the number of file is %v,content is:%v", len(s), s)

	fmt.Printf("io util detial cost time returned \n")
	StartTime = time.Now()
	test, _ = GetAllFileDetial(pathname, test)
	fmt.Println("Cost Time:", time.Since(StartTime))
	j, _ := json.Marshal(test)
	fmt.Println(string(j))
	// fmt.Println("the number of file is %v,content is:%v", len(s), s)

}

func sizeCalc(bytes int64) {
	var kilobytes int64
	kilobytes = (bytes / 1024)

	var megabytes float64
	megabytes = (float64)(kilobytes / 1024) // cast to type float64

	var gigabytes float64
	gigabytes = (megabytes / 1024)

	var terabytes float64
	terabytes = (gigabytes / 1024)

	var petabytes float64
	petabytes = (terabytes / 1024)

	var exabytes float64
	exabytes = (petabytes / 1024)

	var zettabytes float64
	zettabytes = (exabytes / 1024)

	fmt.Println(zettabytes)
}
