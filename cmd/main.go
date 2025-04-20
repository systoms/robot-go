package main

import (
    "fmt"
    "log"

    "robot-go/internal/config"
    "robot-go/internal/handler"
    "robot-go/internal/router"
    "robot-go/internal/service"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    // 加载配置
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // 连接数据库
    db, err := gorm.Open(mysql.Open(cfg.Database.DSN()), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect database: %v", err)
    }

    // 初始化服务
    authService := service.NewAuthService(cfg.JWT.Secret, db)
    roleService := service.NewRoleService(db)
    menuService := service.NewMenuService(db)

    // 初始化处理器
    authHandler := handler.NewAuthHandler(db, authService)
    roleHandler := handler.NewRoleHandler(roleService)

    // 设置路由
    r := router.SetupRouter(authHandler, roleHandler, menuService, authService)

    // 启动服务器
    addr := fmt.Sprintf(":%d", cfg.Server.Port)
    log.Printf("Server is running at %s", addr)
    if err := r.Run(addr); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}