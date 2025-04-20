package handler

import (
    "strconv"

    "github.com/gin-gonic/gin"
    "robot-go/internal/model"
    "robot-go/internal/service"
)

type RoleHandler struct {
    roleService *service.RoleService
}

func NewRoleHandler(roleService *service.RoleService) *RoleHandler {
    return &RoleHandler{roleService: roleService}
}

func (h *RoleHandler) GetRoles(c *gin.Context) {
    companyID := c.GetUint64("company_id")
    roles, err := h.roleService.GetRoles(companyID)
    if err != nil {
        Error(c, 500, err.Error())
        return
    }
    Success(c, roles)
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
    var role model.Role
    if err := c.ShouldBindJSON(&role); err != nil {
        Error(c, 400, err.Error())
        return
    }
    role.CompanyID = c.GetUint64("company_id")
    
    if err := h.roleService.CreateRole(&role); err != nil {
        Error(c, 500, err.Error())
        return
    }
    Success(c, role)
}

func (h *RoleHandler) AssignPermissions(c *gin.Context) {
    roleID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
    var permissionIDs []uint64
    if err := c.ShouldBindJSON(&permissionIDs); err != nil {
        Error(c, 400, err.Error())
        return
    }

    if err := h.roleService.AssignPermissions(roleID, permissionIDs); err != nil {
        Error(c, 500, err.Error())
        return
    }
    Success(c, "权限分配成功")
}