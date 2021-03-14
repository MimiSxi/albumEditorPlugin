/**
 * @Author zhangfan
 * @create 2021/2/25 上午9:50
 * Description:
 */

package model

import (
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

type AlbumOrder struct {
	ID          uint                        `gorm:"primary_key" gqlschema:"delete!;query!;querys" description:"订单id"`
	UserId      uint                        `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys" description:"创建用户id" funservice:"employee"`
	SinglePrice uint                        `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;querys" description:"单价"`
	Amount      uint                        `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;querys" description:"数量"`
	TotalPrice  uint                        `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;querys" description:"总价"`
	Specs       AlbumOrderSpecsEnumType     `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"create;querys" description:"纸张规格"`
	Material    AlbumOrderMaterialEnumType  `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"create;querys" description:"材质"`
	Template    string                      `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create;querys" description:"使用的相册模板"`
	UsageType   AlbumOrderUsageTypeEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"create;querys"  description:"使用类型枚举"`
	CreatedAt   time.Time                   `description:"创建时间" gqlschema:"querys"`
	UpdatedAt   time.Time                   `description:"更新时间" gqlschema:"querys"`
	DeletedAt   *time.Time
	v2          int `gorm:"-" exclude:"true"`
}

type AlbumOrders struct {
	TotalCount int
	Edges      []AlbumOrder
}

func (o AlbumOrder) Query(params graphql.ResolveParams) (AlbumOrder, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o AlbumOrder) Querys(params graphql.ResolveParams) (AlbumOrders, error) {
	var result AlbumOrders

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o AlbumOrder) Create(params graphql.ResolveParams) (AlbumOrder, error) {
	p := params.Args
	o.UserId = uint(p["userId"].(int))
	// template
	if p["singlePrice"] != nil {
		o.SinglePrice = uint(p["singlePrice"].(int))
	}
	if p["amount"] != nil {
		o.Amount = uint(p["amount"].(int))
	}
	if p["totalPrice"] != nil {
		o.TotalPrice = uint(p["totalPrice"].(int))
	}
	if p["specs"] != nil {
		o.Specs = p["specs"].(AlbumOrderSpecsEnumType)
	}
	if p["material"] != nil {
		o.Material = p["material"].(AlbumOrderMaterialEnumType)
	}
	if p["template"] != nil {
		o.Template = p["template"].(string)
	}
	if p["usageType"] != nil {
		o.UsageType = p["usageType"].(AlbumOrderUsageTypeEnumType)
	}
	if p["freightPrice"] != nil {
		o.TotalPrice = uint(p["freightPrice"].(int))
	}
	err := db.Create(&o).Error
	return o, err
}

func (o AlbumOrder) Delete(params graphql.ResolveParams) (AlbumOrder, error) {
	v, ok := params.Source.(AlbumOrder)
	if !ok {
		return o, errors.New("delete param")
	}
	err := db.Delete(&v).Error
	return v, err
}
