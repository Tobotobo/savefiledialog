# savefiledialog
golang savefiledialog

一番シンプルな使い方
``` go
package main

import (
	"fmt"

	"github.com/Tobotobo/savefiledialog")


func main() {
	if filePath, ok := savefiledialog.Show(); ok {
		fmt.Println(filePath)
	}
}
```
タイトル、フィルター指定
``` go
savefiledialog.Title("Excelを保存").Filter("Excel(*.xlsx)|*.xlsx").Show()
```
