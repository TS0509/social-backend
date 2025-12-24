package database

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	gomysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"social-backend/internal/config"
	"social-backend/internal/model"
)

var DB *gorm.DB

// ===== TLS for Aiven =====
func loadAivenTLS() {
	rootCertPool := x509.NewCertPool()

	pem, err := os.ReadFile("ca.pem")
	if err != nil {
		log.Fatalf("❌ 无法读取 CA 文件: %v", err)
	}

	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("❌ 无法加载 CA PEM")
	}

	err = gomysql.RegisterTLSConfig("aiven", &tls.Config{
		RootCAs:    rootCertPool,
		MinVersion: tls.VersionTLS12,
	})
	if err != nil {
		log.Fatalf("❌ 注册 TLS 配置失败: %v", err)
	}

	log.Println("🔐 TLS 配置 'aiven' 已成功注册")
}

func InitMySQL() {
	if config.Cfg == nil {
		log.Fatal("❌ config 未加载")
	}

	loadAivenTLS()

	dsn := config.Cfg.MysqlDSN

	// 自动加入 tls=aiven
	if !strings.Contains(dsn, "tls=") {
		if strings.Contains(dsn, "?") {
			dsn += "&tls=aiven"
		} else {
			dsn += "?tls=aiven"
		}
	}

	log.Printf("🔍 连接 MySQL: %s\n", maskPassword(dsn))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ MySQL 连接失败: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// ⭐ 自动创建 users 表
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("❌ AutoMigrate 失败: %v", err)
	}

	DB = db
	log.Println("✅ MySQL 连接成功，User 表已同步")
}

// 隐藏密码
func maskPassword(dsn string) string {
	parts := strings.Split(dsn, "@")
	if len(parts) != 2 {
		return dsn
	}

	cred := strings.Split(parts[0], ":")
	if len(cred) < 2 {
		return dsn
	}

	return fmt.Sprintf("%s:****@%s", cred[0], parts[1])
}
