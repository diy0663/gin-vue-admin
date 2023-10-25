package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
    "gorm.io/gorm"
)

type LinksService struct {
}

// CreateLinks 创建links表记录
// Author [piexlmax](https://github.com/piexlmax)
func (linksService *LinksService) CreateLinks(links *pkgTest.Links) (err error) {
	err = global.GVA_DB.Create(links).Error
	return err
}

// DeleteLinks 删除links表记录
// Author [piexlmax](https://github.com/piexlmax)
func (linksService *LinksService)DeleteLinks(links pkgTest.Links) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&pkgTest.Links{}).Where("id = ?", links.ID).Update("deleted_by", links.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&links).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteLinksByIds 批量删除links表记录
// Author [piexlmax](https://github.com/piexlmax)
func (linksService *LinksService)DeleteLinksByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&pkgTest.Links{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&pkgTest.Links{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateLinks 更新links表记录
// Author [piexlmax](https://github.com/piexlmax)
func (linksService *LinksService)UpdateLinks(links pkgTest.Links) (err error) {
	err = global.GVA_DB.Save(&links).Error
	return err
}

// GetLinks 根据id获取links表记录
// Author [piexlmax](https://github.com/piexlmax)
func (linksService *LinksService)GetLinks(id uint) (links pkgTest.Links, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&links).Error
	return
}

// GetLinksInfoList 分页获取links表记录
// Author [piexlmax](https://github.com/piexlmax)
func (linksService *LinksService)GetLinksInfoList(info pkgTestReq.LinksSearch) (list []pkgTest.Links, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgTest.Links{})
    var linkss []pkgTest.Links
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
         	orderMap["url"] = true
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
	
	err = db.Find(&linkss).Error
	return  linkss, total, err
}
