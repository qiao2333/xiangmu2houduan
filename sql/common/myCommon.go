package common

import (
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

type DBEngine struct {
	*xorm.Engine
}

var engine *DBEngine

var once sync.Once

func GetDBEngine() *DBEngine {
	once.Do(func() {
		var tmp, err = xorm.NewEngine("mysql", "work:123456@tcp(localhost:3306)/light?charset=utf8&parseTime=True")
		if err != nil {
			panic(err)
		}
		engine = &DBEngine{
			Engine: tmp,
		}

		engine.DatabaseTZ = time.Local
		engine.TZLocation = time.Local
		engine.ShowSQL(true)
		engine.ShowExecTime(true)
		engine.SetMaxOpenConns(1500)
		engine.SetMaxIdleConns(1200)
		fmt.Println("**************************************************************")
	})
	return engine
}
