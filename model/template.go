/**¡
 * @Author zhangfan
 * @create 2021/2/25 上午9:48
 * Description: 相册模板
 */

package model

import (
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

type Template struct {
	ID                  uint                 `gorm:"primary_key" gqlschema:"update!;delete!;query!;querys" description:"ID"`
	Name                string               `gorm:"Type:varchar(1000);DEFAULT:'';NOT NULL;" gqlschema:"create!;update;querys" description:"模板名称"`
	ProId               uint                 `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;update;querys" funservice:"proJ" description:"ProJId"`
	Kind                string               `gorm:"Type:text" gqlschema:"create;update;querys" description:"模板分类"`
	Theme               string               `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"模板主题"`
	Usage               string               `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"模板用途"`
	UseCounts           uint                 `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"querys;update;" description:"使用次数"`
	Status              CommonStatusEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"create;update;querys" description:"模板状态"`
	MedalId             string               `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;update;querys" description:"勋章id"`
	BasicPage           uint                 `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;update;querys" description:"基础页数"`
	BasicPrice16K       uint                 `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;update;querys" description:"16k纸张基础价格"`
	BasicPrice32K       uint                 `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;update;querys" description:"32k纸张基础价格"`
	OneMorePagePrice16K uint                 `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;update;querys" description:"16k纸张每页增加价格"`
	OneMorePagePrice32K uint                 `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;update;querys" description:"32k纸张每页增加价格"`
	MusicLink           string               `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"相册音乐链接"`
	CreatedAt           time.Time            `description:"创建时间" gqlschema:"querys"`
	UpdatedAt           time.Time            `description:"更新时间" gqlschema:"querys"`
	DeletedAt           *time.Time           `description:"删除时间" gqlschema:"querys"`
	v2                  int                  `gorm:"-" exclude:"true"`
}

// 模版集合
type Templates struct {
	TotalCount int
	Edges      []Template
	Groups     TemplateGroup
}

type TemplateGroup struct {
	Kind  []TemplateGroupType
	Usage []TemplateGroupType
	Theme []TemplateGroupType
}

type TemplateGroupType struct {
	Name  string
	Count int
}

func (o Template) Query(params graphql.ResolveParams) (Template, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o Template) Querys(params graphql.ResolveParams) (Templates, error) {
	var result Templates

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	if err != nil {
		return result, err
	}

	err = db.Model(&Template{}).Select("kind as name, COUNT(id) as count").Group("kind").Scan(&result.Groups.Kind).Error
	if err != nil {
		return result, err
	}

	err = db.Model(&Template{}).Select("'usage' as name, COUNT(id) as count").Group("'usage'").Scan(&result.Groups.Usage).Error
	if err != nil {
		return result, err
	}

	err = db.Model(&Template{}).Select("theme as name, COUNT(id) as count").Group("theme").Scan(&result.Groups.Theme).Error
	if err != nil {
		return result, err
	}
	return result, err
}

func (o Template) Create(params graphql.ResolveParams) (Template, error) {
	p := params.Args
	o.Name = p["name"].(string)
	if p["kind"] != nil {
		o.Kind = p["kind"].(string)
	}
	if p["proJ"] != nil {
		OldProJId := uint(p["proJ"].(int))
		p := &ProJ{}
		err := db.Where("id = ?", OldProJId).First(p).Error
		if err != nil {
			return o, err
		}
		n := &ProJ{}
		n.UserId = p.UserId
		n.Status = p.Status
		n.Name = p.Name
		n.Cover = p.Cover
		n.Pages = p.Pages
		n.ImgUpload = p.ImgUpload
		n.TempUsedId = p.TempUsedId
		n.IsCopy = 2
		err = db.Create(n).Error
		if err != nil {
			return o, err
		}
		o.ProId = n.ID
	}
	if p["status"] != nil {
		o.Status = p["status"].(CommonStatusEnumType)
	}
	if p["theme"] != nil {
		o.Theme = p["theme"].(string)
	}
	if p["usage"] != nil {
		o.Usage = p["usage"].(string)
	}
	if p["medalId"] != nil {
		o.MedalId = p["medalId"].(string)
	}
	if p["basicPage"] != nil {
		o.BasicPage = uint(p["basicPage"].(int))
	}
	if p["basicPrice16K"] != nil {
		o.BasicPrice16K = uint(p["basicPrice16K"].(int))
	}
	if p["basicPrice32K"] != nil {
		o.BasicPrice32K = uint(p["basicPrice32K"].(int))
	}
	if p["oneMorePagePrice16K"] != nil {
		o.OneMorePagePrice16K = uint(p["oneMorePagePrice16K"].(int))
	}
	if p["oneMorePagePrice32K"] != nil {
		o.OneMorePagePrice32K = uint(p["oneMorePagePrice32K"].(int))
	}
	if p["musicLink"] != nil {
		o.MusicLink = p["musicLink"].(string)
	}
	err := db.Create(&o).Error
	return o, err
}

func (o Template) Update(params graphql.ResolveParams) (Template, error) {
	v, ok := params.Source.(Template)
	if !ok {
		return o, errors.New("update param")
	}
	p := params.Args
	if p["name"] != nil {
		v.Name = p["name"].(string)
	}
	if p["proJ"] != nil {
		OldProJId := uint(p["proJ"].(int))
		p := &ProJ{}
		err := db.Where("id = ?", OldProJId).First(p).Error
		if err != nil {
			return o, err
		}
		n := &ProJ{}
		n.UserId = p.UserId
		n.Status = p.Status
		n.Name = p.Name
		n.Cover = p.Cover
		n.Pages = p.Pages
		n.ImgUpload = p.ImgUpload
		n.TempUsedId = p.TempUsedId
		n.IsCopy = 2
		err = db.Create(n).Error
		if err != nil {
			return o, err
		}
		v.ProId = n.ID
	}
	if p["useCounts"] != nil {
		v.UseCounts = uint(p["useCounts"].(int))
	}
	if p["kind"] != nil {
		v.Kind = p["kind"].(string)
	}
	if p["status"] != nil {
		v.Status = p["status"].(CommonStatusEnumType)
	}
	if p["theme"] != nil {
		v.Theme = p["theme"].(string)
	}
	if p["usage"] != nil {
		v.Usage = p["usage"].(string)
	}
	if p["medalId"] != nil {
		v.MedalId = p["medalId"].(string)
	}
	if p["basicPage"] != nil {
		v.BasicPage = uint(p["basicPage"].(int))
	}
	if p["basicPrice16K"] != nil {
		v.BasicPrice16K = uint(p["basicPrice16K"].(int))
	}
	if p["basicPrice32K"] != nil {
		v.BasicPrice32K = uint(p["basicPrice32K"].(int))
	}
	if p["oneMorePagePrice16K"] != nil {
		v.OneMorePagePrice16K = uint(p["oneMorePagePrice16K"].(int))
	}
	if p["oneMorePagePrice32K"] != nil {
		v.OneMorePagePrice32K = uint(p["oneMorePagePrice32K"].(int))
	}
	if p["musicLink"] != nil {
		v.MusicLink = p["musicLink"].(string)
	}
	err := db.Save(&v).Error
	return v, err
}

func (o Template) Delete(params graphql.ResolveParams) (Template, error) {
	v, ok := params.Source.(Template)
	if !ok {
		return o, errors.New("delete param")
	}
	err := db.Delete(&v).Error
	return v, err
}
