package database

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"social-backend/internal/config"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func loadAivenTLS() {
	rootCertPool := x509.NewCertPool()

	// Load CA file
	pem, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		log.Fatalf("❌ Unable to load Aiven CA certificate: %v", err)
	}

	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("❌ Failed to append Aiven PEM certificate.")
	}

	// Register TLS config
	err = mysqlDriver.RegisterTLSConfig("aiven", &tls.Config{
		RootCAs:            rootCertPool,
		InsecureSkipVerify: false,
	})
	if err != nil {
		log.Fatalf("❌ TLS Registration failed: %v", err)
	}
}

func InitMySQL() {
	loadAivenTLS()

	log.Println("🔐 Using DSN:", config.Cfg.MysqlDSN)

	db, err := gorm.Open(mysql.Open(config.Cfg.MysqlDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ MySQL connection failed: %v", err)
	}

	DB = db
	log.Println("✅ MySQL connected (Aiven Cloud)")
}
