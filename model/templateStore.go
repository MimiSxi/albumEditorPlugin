/**
 * @Author zhangfan
 * @create 2021/2/25 上午9:51
 * Description: 草稿箱/我的相册
 */

package model

import (
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

type TemplateStore struct {
	ID        uint                 `gorm:"primary_key" gqlschema:"update!;delete!;query!;querys" description:"ID"`
	UserId    uint                 `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;querys" description:"用户id" funservice:"employee"`
	WorkId    uint                 `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;querys" description:"workId"`
	Work      string               `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create;querys" description:"work"`
	Status    CommonStatusEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"update;querys" description:"状态"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	v2        int `gorm:"-" exclude:"true"`
}

type TemplateStores struct {
	TotalCount int
	Edges      []TemplateStore
}

func (o TemplateStore) Query(params graphql.ResolveParams) (TemplateStore, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o TemplateStore) Querys(params graphql.ResolveParams) (TemplateStores, error) {
	var result TemplateStores

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o TemplateStore) Create(params graphql.ResolveParams) (TemplateStore, error) {
	// todo
	p := params.Args
	if p["userId"] != nil {
		o.UserId = p["userId"].(uint)
	}
	if p["workId"] != nil {
		o.WorkId = p["workId"].(uint)
	}
	if p["work"] != nil {
		o.Work = p["work"].(string)
	}
	err := db.Create(&o).Error
	return o, err
}

func (o TemplateStore) Update(params graphql.ResolveParams) (TemplateStore, error) {
	v, ok := params.Source.(TemplateStore)
	if !ok {
		return o, errors.New("update param")
	}
	p := params.Args
	if p["status"] != nil {
		v.Status = p["status"].(CommonStatusEnumType)
	}
	err := db.Save(&v).Error
	return v, err
}

func (o TemplateStore) Delete(params graphql.ResolveParams) (TemplateStore, error) {
	v, ok := params.Source.(TemplateStore)
	if !ok {
		return o, errors.New("delete param")
	}
	err := db.Delete(&v).Error
	return v, err
}
