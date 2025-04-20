package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "robot-go/internal/service"
)

func JWT(authService *service.AuthService) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权"})
            c.Abort()
            return
        }

        // 移除Bearer前缀
        token = strings.Replace(token, "Bearer ", "", 1)
        claims, err := authService.ParseToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "无效的token"})
            c.Abort()
            return
        }

        // 将用户信息存储到上下文
        c.Set("user_id", claims.UserID)
        c.Set("company_id", claims.CompanyID)
        c.Set("username", claims.Username)
        c.Next()
    }
}