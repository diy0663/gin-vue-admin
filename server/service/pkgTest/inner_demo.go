package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
    "gorm.io/gorm"
)

type InnerDemoService struct {
}

// CreateInnerDemo 创建内部测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (demoService *InnerDemoService) CreateInnerDemo(demo *pkgTest.InnerDemo) (err error) {
	err = global.GVA_DB.Create(demo).Error
	return err
}

// DeleteInnerDemo 删除内部测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (demoService *InnerDemoService)DeleteInnerDemo(demo pkgTest.InnerDemo) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&pkgTest.InnerDemo{}).Where("id = ?", demo.ID).Update("deleted_by", demo.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&demo).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteInnerDemoByIds 批量删除内部测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (demoService *InnerDemoService)DeleteInnerDemoByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&pkgTest.InnerDemo{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&pkgTest.InnerDemo{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateInnerDemo 更新内部测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (demoService *InnerDemoService)UpdateInnerDemo(demo pkgTest.InnerDemo) (err error) {
	err = global.GVA_DB.Save(&demo).Error
	return err
}

// GetInnerDemo 根据id获取内部测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (demoService *InnerDemoService)GetInnerDemo(id uint) (demo pkgTest.InnerDemo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&demo).Error
	return
}

// GetInnerDemoInfoList 分页获取内部测试记录
// Author [piexlmax](https://github.com/piexlmax)
func (demoService *InnerDemoService)GetInnerDemoInfoList(info pkgTestReq.InnerDemoSearch) (list []pkgTest.InnerDemo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgTest.InnerDemo{})
    var demos []pkgTest.InnerDemo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&demos).Error
	return  demos, total, err
}
