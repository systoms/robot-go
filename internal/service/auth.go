package service

import (
    "errors"
    "time"

    "github.com/dgrijalva/jwt-go"
    "robot-go/internal/model"
    "gorm.io/gorm"
)

type Claims struct {
    UserID    uint64 `json:"user_id"`
    CompanyID uint64 `json:"company_id"`
    Username  string `json:"username"`
    jwt.StandardClaims
}

type AuthService struct {
    jwtSecret []byte
    db        *gorm.DB
}

func NewAuthService(secret string, db *gorm.DB) *AuthService {
    return &AuthService{
        jwtSecret: []byte(secret),
        db:        db,
    }
}

func (s *AuthService) GenerateToken(user *model.User) (string, error) {
    nowTime := time.Now()
    expireTime := nowTime.Add(24 * time.Hour)

    claims := Claims{
        UserID:    user.ID,
        CompanyID: user.CompanyID,
        Username:  user.Username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expireTime.Unix(),
            IssuedAt:  nowTime.Unix(),
        },
    }

    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := tokenClaims.SignedString(s.jwtSecret)

    return token, err
}

func (s *AuthService) ParseToken(token string) (*Claims, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return s.jwtSecret, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
}

func (s *AuthService) GetUserInfo(userID, companyID uint64) (map[string]interface{}, error) {
    var user model.User
    if err := s.db.Where("id = ? AND company_id = ?", userID, companyID).First(&user).Error; err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "id":       user.ID,
        "username": user.Username,
        "companyId": user.CompanyID,
        // 可以根据需要添加更多字段
    }, nil
}