/**
 * @Author zhangfan
 * @create 2021/2/27 下午2:59
 * Description: 轮播图
 */

package model

import (
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

type Banner struct {
	ID        uint      `gorm:"primary_key" gqlschema:"update!;delete!;query!;querys" description:"轮播图id"`
	Pic1      string    `gorm:"Type:varchar(1000);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"轮番图1"`
	Pic2      string    `gorm:"Type:varchar(1000);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"轮番图2"`
	Pic3      string    `gorm:"Type:varchar(1000);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"轮番图3"`
	Pic4      string    `gorm:"Type:varchar(1000);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"轮番图4"`
	CreatedAt time.Time `description:"创建时间" gqlschema:"querys"`
	UpdatedAt time.Time `description:"更新时间" gqlschema:"querys"`
	DeletedAt *time.Time
	v2        int `gorm:"-" exclude:"true"`
}

type Banners struct {
	TotalCount int
	Edges      []Banner
}

func (o Banner) Query(params graphql.ResolveParams) (Banner, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o Banner) Querys(params graphql.ResolveParams) (Banners, error) {
	var result Banners

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o Banner) Create(params graphql.ResolveParams) (Banner, error) {
	p := params.Args
	if p["pic1"] != nil {
		o.Pic1 = p["pic1"].(string)
	}
	if p["pic2"] != nil {
		o.Pic2 = p["pic2"].(string)
	}
	if p["pic3"] != nil {
		o.Pic3 = p["pic3"].(string)
	}
	if p["pic4"] != nil {
		o.Pic4 = p["pic4"].(string)
	}
	err := db.Create(&o).Error
	return o, err
}

func (o Banner) Update(params graphql.ResolveParams) (Banner, error) {
	v, ok := params.Source.(Banner)
	if !ok {
		return o, errors.New("update param")
	}
	p := params.Args
	if p["pic1"] != nil {
		v.Pic1 = p["pic1"].(string)
	}
	if p["pic2"] != nil {
		v.Pic2 = p["pic2"].(string)
	}
	if p["pic3"] != nil {
		v.Pic3 = p["pic3"].(string)
	}
	if p["pic4"] != nil {
		v.Pic4 = p["pic4"].(string)
	}
	err := db.Save(&v).Error
	return v, err
}

func (o Banner) Delete(params graphql.ResolveParams) (Banner, error) {
	v, ok := params.Source.(Banner)
	if !ok {
		return o, errors.New("delete param")
	}
	err := db.Delete(&v).Error
	return v, err
}
