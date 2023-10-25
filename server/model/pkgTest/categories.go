// 自动生成模板Categories
package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// categories表 结构体  Categories
type Categories struct {
      global.GVA_MODEL
      Description  string `json:"description" form:"description" gorm:"column:description;comment:;size:191;"`  //description字段 
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:191;"`  //name字段 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName categories表 Categories自定义表名 categories
func (Categories) TableName() string {
  return "categories"
}

