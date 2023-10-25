package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
    "gorm.io/gorm"
)

type CategoriesService struct {
}

// CreateCategories 创建categories表记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoriesService *CategoriesService) CreateCategories(categories *pkgTest.Categories) (err error) {
	err = global.GVA_DB.Create(categories).Error
	return err
}

// DeleteCategories 删除categories表记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoriesService *CategoriesService)DeleteCategories(categories pkgTest.Categories) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&pkgTest.Categories{}).Where("id = ?", categories.ID).Update("deleted_by", categories.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&categories).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCategoriesByIds 批量删除categories表记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoriesService *CategoriesService)DeleteCategoriesByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&pkgTest.Categories{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&pkgTest.Categories{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCategories 更新categories表记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoriesService *CategoriesService)UpdateCategories(categories pkgTest.Categories) (err error) {
	err = global.GVA_DB.Save(&categories).Error
	return err
}

// GetCategories 根据id获取categories表记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoriesService *CategoriesService)GetCategories(id uint) (categories pkgTest.Categories, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&categories).Error
	return
}

// GetCategoriesInfoList 分页获取categories表记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoriesService *CategoriesService)GetCategoriesInfoList(info pkgTestReq.CategoriesSearch) (list []pkgTest.Categories, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgTest.Categories{})
    var categoriess []pkgTest.Categories
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["name"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&categoriess).Error
	return  categoriess, total, err
}
