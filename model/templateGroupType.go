/**
 * @Author zhangfan
 * @create 2021/2/28 下午5:21
 * Description:
 */

package model

import "time"

type templateGroupType struct {
	ID        uint       `gorm:"primary_key" gqlschema:"update!;delete!;query!;querys" description:"ID"`
	Name      string     `gorm:"Type:varchar(1000);DEFAULT:'';NOT NULL;" gqlschema:"create!;update!;querys"`
	Count     uint       `gorm:"DEFAULT:0;NOT NULL;" gqlschema:"querys" `
	CreatedAt time.Time  `description:"创建时间" gqlschema:"querys"`
	UpdatedAt time.Time  `description:"更新时间" gqlschema:"querys"`
	DeletedAt *time.Time `description:"删除时间" gqlschema:"querys"`
	v2        int        `gorm:"-" exclude:"true"`
	//projId:String 模板使用的设计器项目id
	//proj:PhotoShop.Proj  模板使用的设计器项目
}
