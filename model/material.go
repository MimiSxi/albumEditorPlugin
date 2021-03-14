/**
 * @Author zhangfan
 * @create 2021/2/25 上午9:50
 * Description: 素材
 */

package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/graphql-go/graphql"
	"time"
)

type Material struct {
	ID        uint                 `gorm:"primary_key" gqlschema:"update!;delete!;query!;querys" description:"素材id"`
	Status    CommonStatusEnumType `gorm:"DEFAULT:1;NOT NULL;" gqlschema:"update;querys" description:"状态"`
	Kind1     string               `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"一级分类"`
	Kind2     string               `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"二级分类"`
	Kind3     string               `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"三级分类"`
	Name      string               `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"名称"`
	Hash      string               `gorm:"Type:varchar(255);DEFAULT:'';NOT NULL;" gqlschema:"create;update;querys" description:"图片类型hash地址"`
	Json      AnnexJSON            `gorm:"Type:text;" gqlschema:"create;update" description:"非图片类型json数据"`
	CreatedAt time.Time            `description:"创建时间" gqlschema:"querys"`
	UpdatedAt time.Time            `description:"更新时间" gqlschema:"querys"`
	DeletedAt *time.Time
	v2        int `gorm:"-" exclude:"true"`
}

type AnnexJSON map[string]interface{}

func (c AnnexJSON) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *AnnexJSON) Scan(input interface{}) error {
	v, ok := input.([]byte)
	if !ok {
		v = []byte(input.(string))
	}
	err := json.Unmarshal(v, c)
	return err
}

// 素材集合
type Materials struct {
	TotalCount int
	Edges      []Material
	Groups     MaterialsGroup
}

type MaterialsGroup struct {
	Kind1 []MaterialsGroupType
	Kind2 []MaterialsGroupType
	Kind3 []MaterialsGroupType
}

type MaterialsGroupType struct {
	Name  string
	Count int
}

func (o Material) Query(params graphql.ResolveParams) (Material, error) {
	p := params.Args
	err := db.Where(p).First(&o).Error
	return o, err
}

func (o Material) Querys(params graphql.ResolveParams) (Materials, error) {
	var result Materials

	dbselect := GenSelet(db, params)
	dbcount := GenWhere(db.Model(o), params)

	err := dbselect.Find(&result.Edges).Error
	if err != nil {
		return result, err
	}
	err = dbcount.Count(&result.TotalCount).Error
	err = db.Model(&Material{}).Select("kind1 as name, COUNT(id) as count").Group("kind1").Scan(&result.Groups.Kind1).Error
	if err != nil {
		return result, err
	}

	err = db.Model(&Material{}).Select("kind2 as name, COUNT(id) as count").Group("kind2").Scan(&result.Groups.Kind2).Error
	if err != nil {
		return result, err
	}

	err = db.Model(&Material{}).Select("kind3 as name, COUNT(id) as count").Group("kind3").Scan(&result.Groups.Kind3).Error
	if err != nil {
		return result, err
	}
	return result, err
}

func (o Material) Create(params graphql.ResolveParams) (Material, error) {
	p := params.Args
	if p["name"] != nil {
		o.Name = p["name"].(string)
	}
	if p["hash"] != nil {
		o.Hash = p["hash"].(string)
	}
	if p["kind1"] != nil {
		o.Kind1 = p["kind1"].(string)
	}
	if p["kind2"] != nil {
		o.Kind2 = p["kind2"].(string)
	}
	if p["kind3"] != nil {
		o.Kind3 = p["kind3"].(string)
	}
	if p["json"] != nil {
		config := p["json"].(string)
		if err := json.Unmarshal([]byte(config), &o.Json); err != nil {
			return o, err
		}
	}
	err := db.Create(&o).Error
	return o, err
}

func (o Material) Update(params graphql.ResolveParams) (Material, error) {
	v, ok := params.Source.(Material)
	if !ok {
		return o, errors.New("update param")
	}
	p := params.Args
	if p["status"] != nil {
		v.Status = p["status"].(CommonStatusEnumType)
	}
	if p["name"] != nil {
		v.Name = p["name"].(string)
	}
	if p["hash"] != nil {
		v.Hash = p["hash"].(string)
	}
	if p["kind1"] != nil {
		v.Kind1 = p["kind1"].(string)
	}
	if p["kind2"] != nil {
		v.Kind2 = p["kind2"].(string)
	}
	if p["kind3"] != nil {
		v.Kind3 = p["kind3"].(string)
	}
	if p["json"] != nil {
		config := p["json"].(string)
		if err := json.Unmarshal([]byte(config), &v.Json); err != nil {
			return o, err
		}
	}
	err := db.Save(&v).Error
	return v, err
}

func (o Material) Delete(params graphql.ResolveParams) (Material, error) {
	v, ok := params.Source.(Material)
	if !ok {
		return o, errors.New("delete param")
	}
	err := db.Delete(&v).Error
	return v, err
}
