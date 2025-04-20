package handler

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "robot-go/internal/model"
    "robot-go/internal/service"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type AuthHandler struct {
    db          *gorm.DB
    authService *service.AuthService
}

func NewAuthHandler(db *gorm.DB, authService *service.AuthService) *AuthHandler {
    return &AuthHandler{
        db:          db,
        authService: authService,
    }
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req struct {
        Username    string `json:"username" binding:"required"`
        Password    string `json:"password" binding:"required"`
        CompanyCode string `json:"company_code" binding:"required"`
        Phone       string `json:"phone"`
        VerifyCode  string `json:"verifyCode"`
        Checked     bool   `json:"checked"`
    }

    // 直接解析请求体
    if err := c.ShouldBindJSON(&req); err != nil {
        log.Printf("Bind error: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    400,
            "message": "请求参数错误",
            "details": err.Error(),
        })
        return
    }

    log.Printf("Parsed request: %+v", req)

    // 检查必要字段
    if req.CompanyCode == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "code":    400,
            "message": "公司代码不能为空",
        })
        return
    }

    // 查找公司
    var company model.Company
    if err := h.db.Where("company_code = ? AND status = 1", req.CompanyCode).First(&company).Error; err != nil {
        Error(c, 401, "公司不存在或已禁用")
        return
    }

    // 查找用户
    var user model.User
    if err := h.db.Where("username = ? AND company_id = ? AND status = 1", req.Username, company.ID).First(&user).Error; err != nil {
        Error(c, 401, "用户名或密码错误")
        return
    }

    // 验证密码
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        Error(c, 401, "用户名或密码错误")
        return
    }

    // 生成token
    token, err := h.authService.GenerateToken(&user)
    if err != nil {
        Error(c, 500, "生成token失败")
        return
    }

    Success(c, gin.H{
        "token": token,
        "user": gin.H{
            "id":       user.ID,
            "username": user.Username,
        },
    })
}