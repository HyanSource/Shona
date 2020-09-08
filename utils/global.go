package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/HyanSource/Shona/Slot"
)

//儲存全域使用的變數
var GlobalObject *GlobalObj

/*初始化*/
func init() {
	GlobalObject = &GlobalObj{
		FilePath: "shonaconfig/shona.json",
		Symbols:  make([]Slot.ISymbol, 0),
		Reels:    make([][]string, 0, 5),
	}

	GlobalObject.Load()
}

type GlobalObj struct {
	FilePath string         //載入路徑
	Symbols  []Slot.ISymbol //儲存賠率
	Reels    [][]string     //儲存所有滾輪
}

//判斷文章是否存在
func (t *GlobalObj) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	return !os.IsNotExist(err), err
}

//載入文件
func (t *GlobalObj) Load() {
	if Exists, err := t.PathExists(t.FilePath); !Exists {
		panic(err)
	}
	//load file
	data, err := ioutil.ReadFile(t.FilePath)

	if err != nil {
		panic(err)
	}
	//unmarshal
	if err = json.Unmarshal(data, t); err != nil {
		panic(err)
	}

	fmt.Println("global load ok")
	fmt.Println(GlobalObject)
}
