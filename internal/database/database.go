package database

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/jtonynet/go-products-api/config"
	"github.com/jtonynet/go-products-api/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/prometheus"
)

func NewDatabase(cfg *config.Database) (*gorm.DB, error) {
	strConn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=true",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)

	var dbErr error
	var db *gorm.DB

	// Retry connecting to the database for RetryMaxElapsedTimeInMs Milliseconds
	retry := backoff.NewExponentialBackOff()
	retry.MaxElapsedTime = time.Duration(cfg.RetryMaxElapsedTimeInMs) * time.Millisecond
	backoff.RetryNotify(func() error {
		db, dbErr = gorm.Open(mysql.Open(strConn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		return dbErr
	}, retry, func(err error, t time.Duration) {
		slog.Info("Retrying connect to Database after error: %v", err)
	})

	if dbErr != nil {
		return nil, dbErr
	}

	if cfg.MetricsEnabled {
		db.Use(prometheus.New(prometheus.Config{
			DBName:          cfg.MetricsDBName,               // `DBName` as metrics label
			RefreshInterval: cfg.MetricsRefreshIntervalInSec, // refresh metrics interval (default 15 seconds)
			StartServer:     cfg.MetricStartServer,           // start http server to expose metrics
			HTTPServerPort:  cfg.MetricServerPort,            // configure http server port, default port 8080 (if you have configured multiple instances, only the first `HTTPServerPort` will be used to start server)
			MetricsCollector: []prometheus.MetricsCollector{
				&prometheus.MySQL{VariableNames: []string{"Threads_running"}},
			},
		}))
	}

	db.AutoMigrate(entity.Product{})

	return db, nil
}
