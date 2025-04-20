package service

import (
    "robot-go/internal/model"
    "gorm.io/gorm"
)

type PermissionService struct {
    db *gorm.DB
}

func NewPermissionService(db *gorm.DB) *PermissionService {
    return &PermissionService{db: db}
}

func (s *PermissionService) GetPermissions() ([]model.Permission, error) {
    var permissions []model.Permission
    err := s.db.Where("status = 1").Find(&permissions).Error
    return permissions, err
}

func (s *PermissionService) GetRolePermissions(roleID uint64) ([]uint64, error) {
    var permissionIDs []uint64
    err := s.db.Model(&model.RolePermission{}).
        Where("role_id = ?", roleID).
        Pluck("permission_id", &permissionIDs).Error
    return permissionIDs, err
}

func (s *PermissionService) CreatePermission(permission *model.Permission) error {
    return s.db.Create(permission).Error
}

func (s *PermissionService) UpdatePermission(permission *model.Permission) error {
    return s.db.Model(permission).Updates(permission).Error
}

func (s *PermissionService) DeletePermission(id uint64) error {
    return s.db.Model(&model.Permission{}).Where("id = ?", id).Update("status", 0).Error
}