INSERT INTO companies (company_code, company_name, status) VALUES
('company_001', '测试公司1', 1),
('company_002', '测试公司2', 1);

-- 使用 bcrypt 加密密码，明文为 'admin123'
INSERT INTO users (company_id, username, password, status) VALUES
(1, 'admin', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 1),
(2, 'admin', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 1);

INSERT INTO roles (company_id, role_name, role_code, status) VALUES
(1, '管理员', 'admin', 1),
(2, '管理员', 'admin', 1);

INSERT INTO user_roles (user_id, role_id) VALUES
(1, 1),
(2, 2);

INSERT INTO menus (parent_id, path, name, component, redirect, title_zh, title_en, icon, sort, status) VALUES
(0, '/system', 'system', 'LAYOUT', '/system/user', '系统管理', 'System', 'setting', 1, 1),
(1, 'user', 'SystemUser', 'LAYOUT', NULL, '用户管理', 'User Management', NULL, 1, 1),
(1, 'role', 'SystemRole', 'LAYOUT', NULL, '角色管理', 'Role Management', NULL, 2, 1);

INSERT INTO permissions (permission_name, permission_code, menu_id, status) VALUES
('用户管理', 'system:user', 2, 1),
('角色管理', 'system:role', 3, 1);

INSERT INTO role_permissions (role_id, permission_id) VALUES
(1, 1),
(1, 2),
(2, 1),
(2, 2);