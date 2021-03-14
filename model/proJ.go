/**
 * @Author zhangfan
 * @create 2021/2/27 下午2:58
 * Description: 设计器项目
 */

package model

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"time"
)

type ProJ struct {
	ID         uint                     `gorm:"primary_key" gqlschema:"update!;delete!;query!;querys" description:"设计器项目id"`
	UserId     uint                     `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys" description:"创建用户id" funservice:"employee"`
	Status     ProJCommonStatusEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"update;querys" description:"状态"`
	Name       string                   `gorm:"Type:varchar(64);DEFAULT:'';NOT NULL;" gqlschema:"create!;update;querys" description:"项目名称"`
	Cover      string                   `gorm:"Type:longText;" gqlschema:"create!;update;querys" description:"封面"`
	Pages      string                   `gorm:"Type:longText;" gqlschema:"create;update" description:"画布"`
	ImgUpload  string                   `gorm:"Type:text;" gqlschema:"create;update" description:"图片json"`
	TempUsedId uint                     `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;update;querys" description:"使用的模版id" funservice:"template"`
	IsCopy     uint                     `gorm:"DEFAULT:1;NOT NULL;" exclude:"true"` // 1。代表原生 2拷贝
	CreatedAt  time.Time                `description:"创建时间" gqlschema:"querys"`
	UpdatedAt  time.Time                `description:"更新时间" gqlschema:"querys"`
	DeletedAt  *time.Time
	v2         int `gorm:"-" exclude:"true"`
}

// 设计器项目集合
type ProJs struct {
	TotalCount int
	Edges      []ProJ
}

func (o ProJ) Query(params graphql.ResolveParams) (ProJ, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o ProJ) Querys(params graphql.ResolveParams) (ProJs, error) {
	var result ProJs
	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)
	err := dbselect.Where("is_copy = 1").Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o ProJ) Create(params graphql.ResolveParams) (ProJ, error) {
	p := params.Args
	if p["name"] != nil {
		o.Name = p["name"].(string)
	}
	if p["imgUpload"] != nil {
		o.ImgUpload = p["imgUpload"].(string)
	}
	if p["tempUsedId"] != nil {
		o.TempUsedId = uint(p["tempUsedId"].(int))
	}
	if p["pages"] != nil {
		o.Pages = p["pages"].(string)
	}
	//var pages []Page
	//if p["pages"] != nil {
	//	pageJson := p["pages"].(string)
	//	err := json.Unmarshal([]byte(pageJson), &pages)
	//	if err != nil {
	//		return o, err
	//	}
	//}
	err := db.Transaction(func(tx *gorm.DB) error {
		o.UserId = uint(p["userId"].(int))
		o.Cover = p["cover"].(string)
		err := tx.Create(&o).Error
		if err != nil {
			return err
		}
		// 创建page
		//for _, v := range pages {
		//	if err := tx.Create(
		//		&Page{
		//			ProJId:     o.ID,
		//			RenderRes:  v.RenderRes,
		//			Status:     v.Status,
		//			Direction:  v.Direction,
		//			PType:      v.PType,
		//			CanvasJson: v.CanvasJson,
		//			Font:       v.Font,
		//		}).Error; err != nil {
		//		return err
		//	}
		//}
		return nil
	})
	return o, err
}

func (o ProJ) Update(params graphql.ResolveParams) (ProJ, error) {
	v, ok := params.Source.(ProJ)
	if !ok {
		return o, errors.New("update param")
	}
	p := params.Args
	//var pages []Page
	//if p["pages"] != nil {
	//	pageJson := p["pages"].(string)
	//	err := json.Unmarshal([]byte(pageJson), &pages)
	//	if err != nil {
	//		return o, err
	//	}
	//}
	//err := db.Where("pro_j_id = ?", v.ID).Delete(&Page{}).Error
	//err = db.Transaction(func(tx *gorm.DB) error {
	//	// 创建page
	//	for _, v := range pages {
	//		if err := tx.Create(
	//			&Page{
	//				ProJId:     o.ID,
	//				RenderRes:  v.RenderRes,
	//				Status:     v.Status,
	//				Direction:  v.Direction,
	//				PType:      v.PType,
	//				CanvasJson: v.CanvasJson,
	//				Font:       v.Font,
	//			}).Error; err != nil {
	//			return err
	//		}
	//	}
	//	return nil
	//})
	if p["pages"] != nil {
		v.Pages = p["pages"].(string)
	}
	if p["tempUsedId"] != nil {
		v.TempUsedId = uint(p["tempUsedId"].(int))
	}
	if p["status"] != nil {
		v.Status = p["status"].(ProJCommonStatusEnumType)
	}
	if p["name"] != nil {
		v.Name = p["name"].(string)
	}
	if p["cover"] != nil {
		v.Cover = p["cover"].(string)
	}
	if p["imgUpload"] != nil {
		v.ImgUpload = p["imgUpload"].(string)
	}
	err := db.Save(&v).Error
	return v, err
}

func (o ProJ) Delete(params graphql.ResolveParams) (ProJ, error) {
	v, ok := params.Source.(ProJ)
	if !ok {
		return o, errors.New("delete param")
	}
	err := db.Delete(&v).Error
	err = db.Where("pro_j_id = ?", v.ID).Delete(&Page{}).Error
	return v, err
}
