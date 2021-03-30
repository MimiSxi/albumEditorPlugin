package model

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
	"time"
)

type Share struct {
	ID        uint      `gorm:"primary_key" gqlschema:"delete!;query!;querys" description:"分享id"`
	ProId     uint      `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys" funservice:"proJ" description:"ProId"`
	CreatedAt time.Time `description:"创建时间" gqlschema:"querys"`
	UpdatedAt time.Time `description:"更新时间" gqlschema:"querys"`
	DeletedAt *time.Time
	v2        int `gorm:"-" exclude:"true"`
}

type Shares struct {
	TotalCount int
	Edges      []Share
}

func (o Share) Query(params graphql.ResolveParams) (Share, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o Share) Querys(params graphql.ResolveParams) (Shares, error) {
	var result Shares

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o Share) Create(params graphql.ResolveParams) (Share, error) {
	p := params.Args
	if p["proId"] != nil {
		fmt.Println(reflect.TypeOf(p["proId"]))
		OldProJId := p["proId"].(int)
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
	err := db.Create(&o).Error
	return o, err
}

func (o Share) Delete(params graphql.ResolveParams) (Share, error) {
	v, ok := params.Source.(Share)
	if !ok {
		return o, errors.New("delete param")
	}
	err := db.Delete(&v).Error
	return v, err
}