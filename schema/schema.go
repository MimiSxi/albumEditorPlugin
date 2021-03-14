package schema

import (
	"errors"
	"github.com/Fiber-Man/albumEditorPlugin/model"
	"github.com/Fiber-Man/funplugin"
	"github.com/Fiber-Man/funplugin/plugin"
	"github.com/graphql-go/graphql"
	"reflect"
)

var templateSchema *funplugin.ObjectSchema
var materialSchema *funplugin.ObjectSchema
var albumOrderSchema *funplugin.ObjectSchema
var bannerSchema *funplugin.ObjectSchema
var proJStoreSchema *funplugin.ObjectSchema
var proJSchema *funplugin.ObjectSchema

//var pageSchema *funplugin.ObjectSchema

var load = false

func Init() {
	InitAlbumEditor()
	//proJSchema.GraphQLType.AddFieldConfig("pages", pageSchema.Query["pages"])

	if field, err := plugin.AutoField("TempUsedId:template"); err != nil {
		panic(errors.New("not have object type"))
	} else {
		proJSchema.GraphQLType.AddFieldConfig("tempUsed", field)
	}

	if field, err := plugin.AutoField("ProJId:proj"); err != nil {
		panic(errors.New("not have object type"))
	} else {
		proJStoreSchema.GraphQLType.AddFieldConfig("proJ", field)
	}

	if field, err := plugin.AutoField("ProJId:proj"); err != nil {
		panic(errors.New("not have object type"))
	} else {
		templateSchema.GraphQLType.AddFieldConfig("proJ", field)
	}

	if field, err := plugin.AutoField("UserId:employee"); err != nil {
		panic(errors.New("not have object type"))
	} else {
		proJStoreSchema.GraphQLType.AddFieldConfig("user", field)
	}
}

func InitAlbumEditor() {
	obj, ok := plugin.GetObject("orderInfo")
	if !ok {
		panic(errors.New("not have object type"))
	}

	obj.AddFieldConfig("albumOrder", &graphql.Field{
		Type:        albumOrderSchema.GraphQLType,
		Description: "albumOrder type",
		Args:        graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (result interface{}, err error) {
			v := reflect.ValueOf(p.Source)
			var id uint
			{
				struct_model := v.FieldByName("Model")
				if !struct_model.IsValid() {
					panic("bad field Model")
				}
				model := struct_model.Interface()
				struct_gorm_model := reflect.ValueOf(model)
				idx := struct_gorm_model.FieldByName("ID")
				if !idx.IsValid() {
					panic("bad field in gorm.Model id")
				}
				id = uint(idx.Uint())
			}

			var typestr string
			{
				referx := v.FieldByName("ChildrenType")
				if !referx.IsValid() {
					panic("bad field ReferType")
				}
				typestr = referx.String()
			}

			if typestr != "albumOrder" {
				return nil, nil
			}

			obj := &model.AlbumOrder{}
			err = obj.QueryByID(id)
			if err != nil {
				return nil, err
			}
			return *obj, nil
		},
	})
}


func marge(oc *funplugin.ObjectSchema) {
	for k, v := range oc.Query {
		queryFields[k] = v
	}
	for k, v := range oc.Mutation {
		mutationFields[k] = v
	}
}

var queryFields = graphql.Fields{
	// "account":  &queryAccount,
	// "accounts": &queryAccountList,
	// "authority":  &queryAuthority,
	// "authoritys": &queryAuthorityList,
}

var mutationFields = graphql.Fields{
	// "createAccount": &createAccount,
	// "updateAccount": &updateAccount,
}

// NewSchema 用于插件主程序调用
func NewPlugSchema(pls funplugin.PluginManger) funplugin.Schema {
	if load != true {
		templateSchema, _ = pls.NewSchemaBuilder(model.Template{})
		marge(templateSchema)

		materialSchema, _ = pls.NewSchemaBuilder(model.Material{})
		marge(materialSchema)

		albumOrderSchema, _ = pls.NewSchemaBuilder(model.AlbumOrder{})
		marge(albumOrderSchema)

		bannerSchema, _ = pls.NewSchemaBuilder(model.Banner{})
		marge(bannerSchema)

		proJStoreSchema, _ = pls.NewSchemaBuilder(model.ProJStore{})
		marge(proJStoreSchema)

		proJSchema, _ = pls.NewSchemaBuilder(model.ProJ{})
		marge(proJSchema)

		//pageSchema, _ = pls.NewSchemaBuilder(model.Page{})
		//marge(pageSchema)
		load = true

		load = true
	}

	return funplugin.Schema{
		Object: map[string]*graphql.Object{
			"template":      templateSchema.GraphQLType,
			"material":      materialSchema.GraphQLType,
			"albumOrder":    albumOrderSchema.GraphQLType,
			"banner":        bannerSchema.GraphQLType,
			"templateStore": proJStoreSchema.GraphQLType,
			"proJ":          proJSchema.GraphQLType,

			//"page":          pageSchema.GraphQLType,
		},
		Query:    queryFields,
		Mutation: mutationFields,
	}
}
