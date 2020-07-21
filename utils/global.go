package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type GlobalObj struct {
	FilePath string
}

//判斷文章是否存在
func (t *GlobalObj) PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err != nil {
		return os.IsNotExist(err), err
	}
	return true, nil
}

//載入文件
func (t *GlobalObj) Load() {
	if Exists, err := t.PathExists(t.FilePath); !Exists {
		panic(err)
	}

	data, err := ioutil.ReadFile(t.FilePath)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, t)

	if err != nil {
		panic(err)
	}

	fmt.Println("載入成功")
}

//儲存全域使用的變數
var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		FilePath: "slotconfig/shona.json",
	}

	GlobalObject.Load()
}
