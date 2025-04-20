package model

type Menu struct {
    ID        uint64 `json:"id" gorm:"primary_key"`
    ParentID  uint64 `json:"parent_id"`
    Path      string `json:"path"`
    Name      string `json:"name"`
    Component string `json:"component"`
    Redirect  string `json:"redirect"`
    TitleZh   string `json:"title_zh"`
    TitleEn   string `json:"title_en"`
    Icon      string `json:"icon"`
    Sort      int    `json:"sort"`
    Status    int8   `json:"status"`
}

type MenuTree struct {
    Menu
    Meta     MenuMeta    `json:"meta"`
    Children []MenuTree  `json:"children,omitempty"`
}

type MenuMeta struct {
    Title map[string]string `json:"title"`
    Icon  string           `json:"icon,omitempty"`
}