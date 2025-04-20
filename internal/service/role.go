package service

import (
    "robot-go/internal/model"
    "gorm.io/gorm"
)

type RoleService struct {
    db *gorm.DB
}

func NewRoleService(db *gorm.DB) *RoleService {
    return &RoleService{db: db}
}

func (s *RoleService) GetRoles(companyID uint64) ([]model.Role, error) {
    var roles []model.Role
    err := s.db.Where("company_id = ? AND status = 1", companyID).Find(&roles).Error
    return roles, err
}

func (s *RoleService) CreateRole(role *model.Role) error {
    return s.db.Create(role).Error
}

func (s *RoleService) UpdateRole(role *model.Role) error {
    return s.db.Model(role).Updates(role).Error
}

func (s *RoleService) DeleteRole(id uint64) error {
    return s.db.Model(&model.Role{}).Where("id = ?", id).Update("status", 0).Error
}

func (s *RoleService) AssignPermissions(roleID uint64, permissionIDs []uint64) error {
    return s.db.Transaction(func(tx *gorm.DB) error {
        // 删除原有权限
        if err := tx.Where("role_id = ?", roleID).Delete(&model.RolePermission{}).Error; err != nil {
            return err
        }

        // 添加新权限
        for _, pid := range permissionIDs {
            rp := model.RolePermission{
                RoleID:       roleID,
                PermissionID: pid,
            }
            if err := tx.Create(&rp).Error; err != nil {
                return err
            }
        }
        return nil
    })
}