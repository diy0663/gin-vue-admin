// 自动生成模板Links
package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// links表 结构体  Links
type Links struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:255;"`  //name字段 
      Url  string `json:"url" form:"url" gorm:"column:url;comment:;size:600;"`  //url字段 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName links表 Links自定义表名 links
func (Links) TableName() string {
  return "links"
}

