package benchmark

import (
	"database/sql"
	"fmt"
	"testing"

	"gopkg.in/TeaMeow/Reiner.v1"
	// The MySQL driver.
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/gocraft/dbr"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)

var (
	dsn = "root:root@/test?charset=utf8"
	i   = 0
)

var (
	reinerDB *reiner.Wrapper
	sqlDB    *sql.DB
	dbrDB    *dbr.Session
	sqlxDB   *sqlx.DB
	gormDB   *gorm.DB
	xormDB   *xorm.Engine
)

func init() {
	var err error
	reinerDB, err = reiner.New(dsn)
	if err != nil {
		panic(err)
	}
	migration := reinerDB.Migration()
	err = migration.Drop("BrenchmarkTests")
	if err != nil {
		panic(err)
	}
	err = migration.
		Table("BrenchmarkTests").
		Column("ID").Int(10).Unsigned().AutoIncrement().Primary().
		Column("Username").Varchar(255).
		Column("Password").Varchar(255).
		Create()
	if err != nil {
		panic(err)
	}
}

func BenchmarkReinerInsert(b *testing.B) {
	for x := 0; x < b.N; x++ {
		i++
		err := reinerDB.Table("BrenchmarkTests").Insert(map[string]interface{}{
			"ID":       i,
			"Username": fmt.Sprintf("ABCDEFG%d", i),
			"Password": "HELLO, WORLD!",
		})
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSQLInsert(b *testing.B) {
	var err error
	sqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	for x := 0; x < b.N; x++ {
		i++
		stmt, err := sqlDB.Prepare("INSERT INTO BrenchmarkTests (ID, Username, Password) VALUES (?, ?, ?)")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(i, fmt.Sprintf("ABCDEFG%d", i), "HELLO, WORLD!")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkDbrInsert(b *testing.B) {
	conn, err := dbr.Open("mysql", dsn, nil)
	dbrDB := conn.NewSession(nil)
	if err != nil {
		panic(err)
	}
	for x := 0; x < b.N; x++ {
		i++
		_, err := dbrDB.InsertInto("BrenchmarkTests").
			Columns("ID", "Username", "Password").
			Values(i, fmt.Sprintf("ABCDEFG%d", i), "HELLO, WORLD!").
			Exec()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkSqlxInsert(b *testing.B) {
	var err error
	sqlxDB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	for x := 0; x < b.N; x++ {
		i++
		stmt, err := sqlxDB.Prepare("INSERT INTO BrenchmarkTests (ID, Username, Password) VALUES (?, ?, ?)")
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec(i, fmt.Sprintf("ABCDEFG%d", i), "HELLO, WORLD!")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGormInsert(b *testing.B) {
	var err error
	gormDB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	var u struct {
		ID       int
		Username string
		Password string
	}
	for x := 0; x < b.N; x++ {
		i++
		u.ID = i
		u.Username = fmt.Sprintf("ABCDEFG%d", i)
		u.Password = "HELLO, WORLD!"
		gormDB.Table("BrenchmarkTests").Create(u)
		if gormDB.Error != nil {
			panic(err)
		}
	}
}

func BenchmarkXormInsert(b *testing.B) {
	var err error
	xormDB, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}
	var u struct {
		ID       int    `xorm:"ID"`
		Username string `xorm:"Username"`
		Password string `xorm:"Password"`
	}
	for x := 0; x < b.N; x++ {
		i++
		u.ID = i
		u.Username = fmt.Sprintf("ABCDEFG%d", i)
		u.Password = "HELLO, WORLD!"
		_, err := xormDB.Table("BrenchmarkTests").Insert(u)
		if err != nil {
			panic(err)
		}
	}
}
