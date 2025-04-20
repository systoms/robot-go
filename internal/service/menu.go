package service

import (
    "robot-go/internal/model"
    "gorm.io/gorm"
    "log"
)

type MenuService struct {
    db *gorm.DB
}

func NewMenuService(db *gorm.DB) *MenuService {
    return &MenuService{db: db}
}

func (s *MenuService) GetUserMenuTree(userID, companyID uint64) ([]model.MenuTree, error) {
    var menus []model.Menu
    
    query := `
        SELECT DISTINCT m.* FROM menus m
        INNER JOIN permissions p ON p.menu_id = m.id
        INNER JOIN role_permissions rp ON rp.permission_id = p.id
        INNER JOIN user_roles ur ON ur.role_id = rp.role_id
        WHERE ur.user_id = ? AND m.status = 1
        ORDER BY m.sort ASC
    `
    if err := s.db.Raw(query, userID).Scan(&menus).Error; err != nil {
        return nil, err
    }

    return buildMenuTree(menus), nil
}

func buildMenuTree(menus []model.Menu) []model.MenuTree {
    log.Printf("Building menu tree with %d menus", len(menus))
    for i, m := range menus {
        log.Printf("Menu %d: ID=%d, ParentID=%d, Title=%s", i+1, m.ID, m.ParentID, m.TitleZh)
    }

    menuMap := make(map[uint64]*model.MenuTree)
    var roots []model.MenuTree

    // 首先创建所有节点
    for _, m := range menus {
        menuTree := model.MenuTree{
            Menu: m,
            Meta: model.MenuMeta{
                Title: map[string]string{
                    "zh_CN": m.TitleZh,
                    "en_US": m.TitleEn,
                },
                Icon: m.Icon,
            },
            Children: []model.MenuTree{}, // 初始化子节点
        }
        menuMap[m.ID] = &menuTree
    }

    // 构建树结构
    for _, m := range menus {
        if m.ParentID == 0 {
            roots = append(roots, *menuMap[m.ID])
        } else {
            if parent, exists := menuMap[m.ParentID]; exists {
                parent.Children = append(parent.Children, *menuMap[m.ID])
            }
        }
    }

    // 如果没有根菜单，返回所有菜单
    if len(roots) == 0 {
        for _, m := range menus {
            roots = append(roots, *menuMap[m.ID])
        }
    }

    log.Printf("Built %d root menus", len(roots))
    return roots
}