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

type ProJStore struct {
	ID        uint                 `gorm:"primary_key" gqlschema:"update!;delete!;query!;querys" description:"ID"`
	UserId    uint                 `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;querys" description:"用户id" funservice:"employee"`
	ProJId    uint                 `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys" description:"设计器id" funservice:"proJ"`
	Status    CommonStatusEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"update;querys" description:"状态"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	v2        int `gorm:"-" exclude:"true"`
}

type ProJStores struct {
	TotalCount int
	Edges      []ProJStore
}

func (o ProJStore) Query(params graphql.ResolveParams) (ProJStore, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o ProJStore) Querys(params graphql.ResolveParams) (ProJStores, error) {
	var result ProJStores

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o ProJStore) Create(params graphql.ResolveParams) (ProJStore, error) {
	p := params.Args
	if p["userId"] != nil {
		o.UserId = uint(p["userId"].(int))
	}
	if p["proJId"] != nil {
		o.ProJId = uint(p["proJId"].(int))
	}
	err := db.Create(&o).Error
	return o, err
}

func (o ProJStore) Update(params graphql.ResolveParams) (ProJStore, error) {
	v, ok := params.Source.(ProJStore)
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

func (o ProJStore) Delete(params graphql.ResolveParams) (ProJStore, error) {
	v, ok := params.Source.(ProJStore)
	if !ok {
		return o, errors.New("delete param")
	}
	err := db.Delete(&v).Error
	return v, err
}
