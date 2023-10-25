// 自动生成模板InnerDemo
package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 内部测试 结构体  InnerDemo
type InnerDemo struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`  //名称 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 内部测试 InnerDemo自定义表名 inner_demo
func (InnerDemo) TableName() string {
  return "inner_demo"
}

