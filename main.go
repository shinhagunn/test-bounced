package main

import (
	"fmt"
	"time"

	"github.com/zsmartex/pkg/v2/config"
	"github.com/zsmartex/pkg/v2/infrastructure/gorm_fx"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// func main() {
// 	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=6000 sslmode=disable TimeZone=Asia/Shanghai"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(db)
// }

type ix struct {
	fx.Out

	MaxIdleConns    int           `name:"max_idle_conns"`
	MaxOpenConns    int           `name:"max_open_conns"`
	ConnMaxLifetime time.Duration `name:"conn_max_lifetime"`
}

func main() {
	count := 100

	for i := 0; i < count; i++ {
		go func() {
			app := fx.New(
				fx.Supply(config.Postgres{
					Host: "localhost",
					Port: 6000,
					User: "postgres",
					Pass: "123456",
					Name: "postgres",
				}),
				fx.Supply(ix{}),
				gorm_fx.Module,
				fx.Invoke(func(db *gorm.DB) {
					fmt.Println(db)
				}),
			)

			app.Run()
		}()
	}

	for {
		time.Sleep(time.Second)
	}
}
