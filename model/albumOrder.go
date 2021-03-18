/**
 * @Author zhangfan
 * @create 2021/2/25 上午9:50
 * Description:
 */

package model

import (
	"errors"
	"github.com/Fiber-Man/funplugin/plugin"
	"github.com/graphql-go/graphql"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Albumorder struct {
	ID           uint                        `gorm:"primary_key" gqlschema:"delete!;query!;querys" description:"订单id"`
	UserId       uint                        `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create!;querys" description:"创建用户id" funservice:"employee"`
	SinglePrice  uint                        `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;querys" description:"单价"`
	Amount       uint                        `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;querys" description:"数量"`
	TotalPrice   uint                        `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"create;querys" description:"总价"`
	Specs        AlbumOrderSpecsEnumType     `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"create;querys" description:"纸张规格"`
	Material     AlbumOrderMaterialEnumType  `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"create;querys" description:"材质"`
	Template     string                      `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create;querys" description:"使用的相册模板"`
	UsageType    AlbumOrderUsageTypeEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"create;querys"  description:"使用类型枚举"`
	OrderInfoId  uint                        `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"querys" description:"父订单id"`
	CreatedAt    time.Time                   `description:"创建时间" gqlschema:"querys"`
	UpdatedAt    time.Time                   `description:"更新时间" gqlschema:"querys"`
	DeletedAt    *time.Time
	v2           int    `gorm:"-" exclude:"true"`
	Remark       string `gorm:"-" exclude:"true" gqlschema:"create" description:"订单备注"`
	Address      string `gorm:"-" exclude:"true" gqlschema:"create" description:"收货地址"`
	FreightPrice uint   `gorm:"-" exclude:"true" gqlschema:"create" description:"运费"`
}

type Albumorders struct {
	TotalCount int
	Edges      []Albumorder
}

type OrderPluginData struct {
	Order OrderPlugin
}

type OrderPlugin struct {
	Orderinfos OrderInfoData
}

type OrderInfoData struct {
	Action OrderAction
}

type OrderAction struct {
	Create OrderActionProp
}

type OrderActionProp struct {
	Id string
}

// 跨接口创建订单 childrenType = "albumorder"
func createOrder(userId uint, childrenId uint, childrenType string, remark string, address string, freightPrice uint, goodsPrice uint) (result interface{}, err error) {
	mutation := `mutation ($address: String, $childrenId: Int!, $childrenType: String!, $freightPrice: Int, $goodsPrice: Int, $remark: String, $status: OrderStatusEnumType!, $userId: Int!) {
				  order {
					orderinfos {
					  action {
						create(childrenType: $childrenType, userId: $userId, address: $address, childrenId: $childrenId, freightPrice: $freightPrice, goodsPrice: $goodsPrice, remark: $remark, status: $status) {
						  id
						}
					  }
					}
				  }
				}`
	params := map[string]interface{}{
		"address":      address,
		"childrenId":   childrenId,
		"childrenType": childrenType,
		"freightPrice": freightPrice,
		"goodsPrice":   goodsPrice,
		"remark":       remark,
		"status":       "TO_BE_PAID",
		"userId":       userId,
	}
	e := OrderPluginData{}
	_, err = plugin.Go(mutation, params, &e)
	if err != nil {
		return nil, err
	}
	result = e.Order.Orderinfos.Action.Create.Id
	return
}

//ID2id is string to uint
func ID2id(ID interface{}) (uint, error) {
	if ID == nil || reflect.TypeOf(ID).String() != "string" {
		return 0, errors.New("ID Type error")
	}
	ID2 := ID.(string)
	if p1 := strings.Index(ID2, "-"); p1 > -1 {
		ID2 = ID2[p1+1:]
	}
	id, err := strconv.ParseUint(ID2, 10, 64)
	return uint(id), err
}

func (o *Albumorder) QueryByID(id uint) (err error) {
	return db.Where("id = ?", id).First(&o).Error
}

func (o Albumorder) Query(params graphql.ResolveParams) (Albumorder, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o Albumorder) Querys(params graphql.ResolveParams) (Albumorders, error) {
	var result Albumorders

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	return result, err
}

func (o Albumorder) Create(params graphql.ResolveParams) (Albumorder, error) {
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
	if p["remark"] != nil {
		o.Remark = p["remark"].(string)
	}
	if p["address"] != nil {
		o.Address = p["address"].(string)
	}
	if p["freightPrice"] != nil {
		o.TotalPrice = uint(p["freightPrice"].(int))
	}
	// TODO 事务
	OrderId, err := createOrder(o.UserId, o.ID, "albumorder", o.Remark, o.Address, o.FreightPrice, o.TotalPrice)
	if err != nil {
		return o, err
	}
	o.OrderInfoId, err = ID2id(OrderId)
	err = db.Create(&o).Error
	return o, err
}

func (o Albumorder) Delete(params graphql.ResolveParams) (Albumorder, error) {
	v, ok := params.Source.(Albumorder)
	if !ok {
		return o, errors.New("delete param")
	}
	err := db.Delete(&v).Error
	return v, err
}
