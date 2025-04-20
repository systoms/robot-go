package router

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "net/http"
    "robot-go/internal/handler"
    "robot-go/internal/middleware"
    "robot-go/internal/service"
    "time"
    "bytes"
    "log"
    "io/ioutil"
    "fmt"
)

func SetupRouter(
    authHandler *handler.AuthHandler,
    roleHandler *handler.RoleHandler,
    menuService *service.MenuService,
    authService *service.AuthService,
) *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger(), gin.Recovery())

    // 自定义 CORS 中间件
    r.Use(func(c *gin.Context) {
        origin := c.GetHeader("Origin")
        if origin != "" {
            cors.New(cors.Config{
                AllowOrigins:     []string{origin},
                AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
                AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
                ExposeHeaders:    []string{"Content-Length"},
                AllowCredentials: true,
                MaxAge:           12 * time.Hour,
            })(c)
        }
        c.Next()
    })

    // 公开路由
    // 在路由配置中添加中间件
    public := r.Group("/api")
    {
        public.POST("/login", func(c *gin.Context) {
            // 确保请求体不被提前消耗
            body, err := c.GetRawData()
            if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{
                    "code":    400,
                    "message": "请求体读取失败",
                })
                return
            }

            // 将请求体重新设置回上下文
            c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
            authHandler.Login(c)
        })
    }

    // 需要认证的路由
    authorized := r.Group("/api")
    authorized.Use(middleware.JWT(authService))
    {
        // 用户信息路由
        authorized.GET("/user/info", func(c *gin.Context) {
            userID := c.GetUint64("user_id")
            companyID := c.GetUint64("company_id")
            
            // 调用服务层获取用户信息
            userInfo, err := authService.GetUserInfo(userID, companyID)
            if err != nil {
                handler.Error(c, 500, err.Error())
                return
            }
            handler.Success(c, userInfo)
        })

        // 菜单路由
        authorized.GET("/get-menu-list", func(c *gin.Context) {
            userID := c.GetUint64("user_id")
            companyID := c.GetUint64("company_id")
            
            // 打印调试信息
            log.Printf("Getting menu list for userID: %d, companyID: %d", userID, companyID)
            
            menus, err := menuService.GetUserMenuTree(userID, companyID)
            fmt.Println("menus")
            fmt.Println(menus)
            if err != nil {
                log.Printf("Error getting menu list: %v", err)
                handler.Error(c, 500, err.Error())
                return
            }
            
            // 打印查询结果
            log.Printf("Menu list result: %+v", menus)
            
            // 返回包含 list 字段的数据
            handler.Success(c, gin.H{
                "list": menus,
            })
        })

        // 角色路由
        roles := authorized.Group("/roles")
        {
            roles.GET("", roleHandler.GetRoles)
            roles.POST("", roleHandler.CreateRole)
            roles.POST("/:id/permissions", roleHandler.AssignPermissions)
        }
    }

    r.OPTIONS("/*path", func(c *gin.Context) {
        c.Status(http.StatusNoContent)
        return
    })

    return r
}