/**
 * @Author zhangfan
 * @create 2021/2/27 下午2:59
 * Description:
 */

package model

import (
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

type Page struct {
	ID         uint      `gorm:"primary_key" gqlschema:"update!;delete!;query!;querys" description:"单页画布设计id" json:"id"`
	ProJId     uint      `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"querys" description:"项目id" json:"proJId" funservice:"proJId"`
	RenderRes  string    `gorm:"Type:varchar(1000);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"画布的base64格式图片渲染图" json:"renderRes"`
	Status     uint      `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"update;querys" description:"模板页面状态" json:"status"`
	Direction  uint      `gorm:"DEFAULT:2;NOT NULL;" gqlschema:"create;update;querys" description:"画布页面放置方向" json:"direction"`
	PType      uint      `gorm:"DEFAULT:4;NOT NULL;" gqlschema:"create;update;querys" description:"画布页面类型" json:"pType"`
	CanvasJson string    `gorm:"Type:text;" gqlschema:"create;update" description:"画布json数据" json:"canvasJson"`
	Font       string    `gorm:"Type:text;" gqlschema:"create;update" description:"该画布包含的字体json数据" json:"font"`
	CreatedAt  time.Time `description:"创建时间" gqlschema:"querys" json:"createdAt"`
	UpdatedAt  time.Time `description:"更新时间" gqlschema:"querys" json:"updatedAt"`
	DeletedAt  *time.Time
	v2         int `gorm:"-" exclude:"true"`
}

type Pages struct {
	TotalCount int
	Edges      []Page
}

func (o Page) Query(params graphql.ResolveParams) (Page, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o Page) Querys(params graphql.ResolveParams) (Pages, error) {
	var result Pages

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o Page) Create(params graphql.ResolveParams) (Page, error) {
	p := params.Args
	if p["renderRes"] != nil {
		o.RenderRes = p["renderRes"].(string)
	}
	if p["direction"] != nil {
		o.Direction = uint(p["direction"].(int))
	}
	if p["pType"] != nil {
		o.PType = uint(p["pType"].(int))
	}
	if p["canvasJson"] != nil {
		o.CanvasJson = p["canvasJson"].(string)
	}
	if p["font"] != nil {
		o.Font = p["font"].(string)
	}
	err := db.Create(&o).Error
	return o, err
}

func (o Page) Update(params graphql.ResolveParams) (Page, error) {
	v, ok := params.Source.(Page)
	if !ok {
		return o, errors.New("update param")
	}
	p := params.Args
	if p["renderRes"] != nil {
		v.RenderRes = p["renderRes"].(string)
	}
	if p["status"] != nil {
		v.Status = uint(p["status"].(int))
	}
	if p["direction"] != nil {
		v.Direction = uint(p["direction"].(int))
	}
	if p["pType"] != nil {
		v.PType = uint(p["pType"].(int))
	}
	if p["canvasJson"] != nil {
		v.CanvasJson = p["canvasJson"].(string)
	}
	if p["font"] != nil {
		v.Font = p["font"].(string)
	}
	err := db.Save(&v).Error
	return v, err
}

func (o Page) Delete(params graphql.ResolveParams) (Page, error) {
	v, ok := params.Source.(Page)
	if !ok {
		return o, errors.New("delete param")
	}
	err := db.Delete(&v).Error
	return v, err
}
