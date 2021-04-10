/**
 * @Author zhangfan
 * @create 2021/4/7 下午8:59
 * Description:统计
 */

package model

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"time"
)

// 相册统计
type Statistic struct {
	ID         uint                 `gorm:"primary_key" gqlschema:"query!;querys;" description:"id"`
	BeginTime  time.Time            `gorm:"-" gqlschema:"countprojs!;countorder!;countordermoney!;countloginrecord!;" description:"开始时间"`    // 统计开始时间
	EndTime    time.Time            `gorm:"-" gqlschema:"countprojs!;countorder!;countordermoney!;countloginrecord!;" description:"结束时间"`    // 统计结束时间
	ReturnKind StatisticsReturnKind `gorm:"-" gqlschema:"countprojs!;countorder!;countordermoney!;countloginrecord!;" description:"按什么类型返回"` // 数据返回样式
	v2         int                  `gorm:"-" exclude:"true"`
}

type Statistics struct {
	TotalCount int
	Edges      []Statistic
}

// 返回总户数
type TotalOverview struct {
	Abscissa int64 // 横坐标
	Quantity int64 // 数量
}

type TotalOverviews struct {
	TotalOverviews []TotalOverview
}

type TotalMoney struct {
	Abscissa int64   // 横坐标
	Quantity float64 // 数量
}

type TotalMoneys struct {
	TotalOverviews []TotalMoney
}

// 相册使用数量
func (o Statistic) Countprojs(params graphql.ResolveParams) (TotalOverviews, error) {
	var result TotalOverviews
	p := params.Args
	beginTime, ok := p["beginTime"].(time.Time)
	if !ok {
		return result, errors.New("beginTime type error")
	}
	endTime, ok := p["endTime"].(time.Time)
	if !ok {
		return result, errors.New("endTime type error")
	}
	returnKind := p["returnKind"].(StatisticsReturnKind)
	fmt.Println(beginTime, endTime, returnKind)
	dbx := db
	dbx = dbx.Table("pro_j")
	if returnKind == RETURN_DAY {
		dbx = dbx.Select("count(*) as quantity,day(created_at) as abscissa")
	}
	if returnKind == RETURN_MONTH {
		dbx = dbx.Select("count(*) as quantity,month(created_at) as abscissa")
	}
	dbx = dbx.Where("created_at > ? and created_at < ?", beginTime, endTime)
	dbx = dbx.Group("abscissa")
	err := dbx.Find(&result.TotalOverviews).Error
	return result, err
}

// 订单数量
func (o Statistic) Countorder(params graphql.ResolveParams) (TotalOverviews, error) {
	var result TotalOverviews
	p := params.Args
	beginTime, ok := p["beginTime"].(time.Time)
	if !ok {
		return result, errors.New("beginTime type error")
	}
	endTime, ok := p["endTime"].(time.Time)
	if !ok {
		return result, errors.New("endTime type error")
	}
	returnKind := p["returnKind"].(StatisticsReturnKind)
	fmt.Println(beginTime, endTime, returnKind)
	dbx := db
	dbx = dbx.Table("albumorder")
	if returnKind == RETURN_DAY {
		dbx = dbx.Select("count(*) as quantity,day(created_at) as abscissa")
	}
	if returnKind == RETURN_MONTH {
		dbx = dbx.Select("count(*) as quantity,month(created_at) as abscissa")
	}
	dbx = dbx.Where("created_at > ? and created_at < ?", beginTime, endTime)
	dbx = dbx.Group("abscissa")
	err := dbx.Find(&result.TotalOverviews).Error
	return result, err
}

// 统计订单金额
func (o Statistic) Countordermoney(params graphql.ResolveParams) (TotalMoneys, error) {
	var result TotalMoneys
	p := params.Args
	beginTime, ok := p["beginTime"].(time.Time)
	if !ok {
		return result, errors.New("beginTime type error")
	}
	endTime, ok := p["endTime"].(time.Time)
	if !ok {
		return result, errors.New("endTime type error")
	}
	returnKind := p["returnKind"].(StatisticsReturnKind)
	fmt.Println(beginTime, endTime, returnKind)
	dbx := db
	dbx = dbx.Table("albumorder")
	if returnKind == RETURN_DAY {
		dbx = dbx.Select("sum(total_price) as quantity,day(created_at) as abscissa")
	}
	if returnKind == RETURN_MONTH {
		dbx = dbx.Select("sum(total_price) as quantity,month(created_at) as abscissa")
	}
	dbx = dbx.Where("created_at > ? and created_at < ?", beginTime, endTime)
	dbx = dbx.Group("abscissa")
	err := dbx.Find(&result.TotalOverviews).Error
	return result, err
}

// 统计
func (o Statistic) Countloginrecord(params graphql.ResolveParams) (TotalOverviews, error) {
	var result TotalOverviews
	p := params.Args
	beginTime, ok := p["beginTime"].(time.Time)
	if !ok {
		return result, errors.New("beginTime type error")
	}
	endTime, ok := p["endTime"].(time.Time)
	if !ok {
		return result, errors.New("endTime type error")
	}
	returnKind := p["returnKind"].(StatisticsReturnKind)
	fmt.Println(beginTime, endTime, returnKind)
	dbx := db
	dbx = dbx.Table("login_record")
	if returnKind == RETURN_DAY {
		dbx = dbx.Select("count(*) as quantity,day(created_at) as abscissa")
	}
	if returnKind == RETURN_MONTH {
		dbx = dbx.Select("count(*) as quantity,month(created_at) as abscissa")
	}
	dbx = dbx.Where("created_at > ? and created_at < ?", beginTime, endTime)
	dbx = dbx.Group("abscissa")
	err := dbx.Find(&result.TotalOverviews).Error
	return result, err
}

// 不写代码跑不了 ????????????????
func (o Statistic) Query(params graphql.ResolveParams) (Statistic, error) {
	return o, nil
}

func (o Statistic) Querys(params graphql.ResolveParams) (Statistics, error) {
	var result Statistics
	return result, nil
}
