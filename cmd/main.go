package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/antoha2/urlCoder/config"
	Repository "github.com/antoha2/urlCoder/repository"
	Service "github.com/antoha2/urlCoder/service"
	webHTTP "github.com/antoha2/urlCoder/transport"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	Run()
}

func Run() {

	cfg := config.GetConfig()
	gormDB, err := initDb(cfg)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	Rep := Repository.NewRepository(gormDB)
	Ser := Service.NewService(Rep)

	Tran := webHTTP.NewHTTP(Ser)

	go Tran.StartHTTP()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	Tran.Stop()
}

func initDb(cfg *config.Config) (*gorm.DB, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Dbname,
		cfg.DB.Sslmode,
	)

	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("1 failed to parse config: %v", err)
	}

	// Make connections
	dbx, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return nil, fmt.Errorf("2 failed to create connection db: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbx,
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("3 gorm.Open(): %v", err)
	}

	err = dbx.Ping()
	if err != nil {
		return nil, fmt.Errorf("4 error to ping connection pool: %v", err)
	}
	log.Printf("Подключение к базе данных на http://127.0.0.1:%d\n", cfg.DB.Port)
	return gormDB, nil
}
