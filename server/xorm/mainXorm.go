/**
 * 應用程序
 * 同目錄下多文件引用的問題解決方法:
 * https://blog.csdn.net/pringD/article/details/79143235
 * 方法1 1 go build ./ 2 運行編譯後的文件
 * 方法2 go run *.go
 * xorm 工具使用 :
 * 安裝 : go get github.com/go-xorm/cmd/xorm
 * WIN下使用路徑 : C:\Users\rorast\go\src\github.com\go-xorm\cmd\xorm
 * 使用 : xorm reverse mysql "root:123456@tcp(127.0.0.1:3306)/superstar?charset=utf8" templates/goxorm/
 */
package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

const DriverName = "mysql"
const MasterDataSourceName = "root:123456@tcp(127.0.0.1:3306)/superstar?charset=utf8"

/**
CREATE TABLE `user_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主鍵ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '中文名',
  `sys_created` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '創建時間',
  `sys_updated` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最後修改時間',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
type UserInfo struct {
	Id         int `xorm:"not null pk autoincr"`
	Name       string
	SysCreated int
	SysUpdated int
}

var engine *xorm.Engine

func main() {
	engine = newEngin()

	//execute()
	//ormInsert()
	query()
	ormGet()
	ormGetCols()
	ormCount()
	//ormFindRows()
	//ormUpdate()
	//ormOmitUpdate()
	//ormMustColsUpdate()

}

// 連接到資料庫
func newEngin() *xorm.Engine {
	engine, err := xorm.NewEngine(DriverName, MasterDataSourceName)
	if err != nil {
		log.Fatal(newEngin, err)
		return nil
	}
	// Debug 模式，打印全部的SQL語句，幫助對比，看ORM與SQL執行的對照關係
	engine.ShowSQL(true)
	return engine
}

// 通過 query 方法查詢
func query() {
	sql := "SELECT * FROM user_info"
	// results, err := engine.Query(sql)
	// results, err := engine.QueryInterface(sql)
	results, err := engine.QueryString(sql)
	if err != nil {
		log.Fatal("query", sql, err)
		return
	}
	total := len(results)
	if total == 0 {
		fmt.Println("沒有任何數據", sql)
	} else {
		for i, data := range results {
			fmt.Printf("%d = %v\n", i, data)
		}
	}

}

// 通過 execute 方法執行更新
func execute() {
	sql := `INSERT INTO user_info values(NULL, 'name', 0, 0)`
	affected, err := engine.Exec(sql)
	if err != nil {
		log.Fatal("execute error", err)
	} else {
		id, _ := affected.LastInsertId()
		rows, _ := affected.RowsAffected()
		fmt.Println("execute id=", id, ", rows=", rows)
	}
}

// 根據 models 的結構映射數據表
func ormInsert() {
	UserInfo := &UserInfo{
		Id:         0,
		Name:       "梅西",
		SysCreated: 0,
		SysUpdated: 0,
	}
	id, err := engine.Insert(UserInfo)
	if err != nil {
		log.Fatal("ormInsert error", err)
	} else {
		fmt.Println("ormInsert id=", id)
		fmt.Printf("%v\n", *UserInfo)
	}
}

// 根據 models 的結構讀取數據
func ormGet() {
	UserInfo := &UserInfo{Id: 2}
	ok, err := engine.Get(UserInfo)
	if ok {
		fmt.Printf("%v\n", *UserInfo)
	} else if err != nil {
		log.Fatal("ormGet error", err)
	} else {
		fmt.Println("orgGet empty id=", UserInfo.Id)
	}
}

// 獲取指定的字段
func ormGetCols() {
	UserInfo := &UserInfo{Id: 2}
	ok, err := engine.Cols("name").Get(UserInfo)
	if ok {
		fmt.Printf("%v\n", *UserInfo)
	} else if err != nil {
		log.Fatal("ormGetCols error", err)
	} else {
		fmt.Println("orgGetCols empty id=", UserInfo.Id)
	}
}

// 統計
func ormCount() {
	//count, err := engine.Count(&UserInfo{})
	//count, err := engine.Where("name_zh=?", "梅西").Count(&UserInfo{})
	count, err := engine.Count(&UserInfo{Name: "梅西"})
	if err == nil {
		fmt.Printf("count=%v\n", count)
	} else {
		log.Fatal("ormCount error", err)
	}
}

// 查找多行數據
func ormFindRows() {
	list := make([]UserInfo, 0)
	//list := make(map[int]UserInfo
	//err := engine.Find(&list)
	//err := engine.Where("id>?", 1).Limit(100, 0).Find(&list)
	err := engine.Cols("id", "name").Where("id>?", 0).Limit(10).Asc("id", "sys_created").Find(&list)

	//list := make([]map[string]string, 0)
	//err :=engine.Table("star_info").Cols("id", "name_zh", "name_en").
	// Where("id>?", 1).Find(&list)

	if err == nil {
		fmt.Printf("%v\n", list)
	} else {
		log.Fatal("ormFindRows error", err)
	}
}

// 更新一個數據
func ormUpdate() {
	// 全部更新
	//UserInfo := &UserInfo{NameZh:"測試名"}
	//ok, err := engine.Update(UserInfo)
	// 指定 ID 更新
	UserInfo := &UserInfo{Name: "梅西"}
	ok, err := engine.ID(2).Update(UserInfo)
	fmt.Println(ok, err)
}

// 排除某字段
func ormOmitUpdate() {
	info := &UserInfo{Id: 1}
	ok, _ := engine.Get(info)
	if ok {
		if info.SysCreated > 0 {
			ok, _ := engine.ID(info.Id).Omit("sys_created").Update(&UserInfo{SysCreated: 0, SysUpdated: int(time.Now().Unix())})
			fmt.Printf("ormOmitUpdate, rows=%d, "+"sys_created=%d\n", ok, 0)
		} else {
			ok, _ := engine.ID(info.Id).Omit("sys_created").Update(&UserInfo{SysCreated: 1, SysUpdated: int(time.Now().Unix())})
			fmt.Printf("ormOmitUpdate, rows=%d, "+"sys_created=%d\n", ok, 0)
		}
	}
}

// 字段為空也可以更新 (0, 空字符串等)
func ormMustColsUpdate() {
	info := &UserInfo{Id: 1}
	ok, _ := engine.Get(info)
	if ok {
		if info.SysCreated > 0 {
			ok, _ := engine.ID(info.Id).MustCols("sys_created").Update(&UserInfo{SysCreated: 0, SysUpdated: int(time.Now().Unix())})
			fmt.Printf("ormMustColsUpdate, rows=%d, "+"sys_created=%d\n", ok, 0)
		} else {
			ok, _ := engine.ID(info.Id).MustCols("sys_created").Update(&UserInfo{SysCreated: 1, SysUpdated: int(time.Now().Unix())})
			fmt.Printf("ormMustColsUpdate, rows=%d, "+"sys_created=%d\n", ok, 0)
		}
	}
}
