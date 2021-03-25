package schema

import (
	"errors"
	"github.com/Fiber-Man/albumEditorPlugin/model"
	"github.com/Fiber-Man/funplugin"
	"github.com/Fiber-Man/funplugin/plugin"
	"github.com/graphql-go/graphql"
)

var templateSchema *funplugin.ObjectSchema
var materialSchema *funplugin.ObjectSchema
var albumorderSchema *funplugin.ObjectSchema
var bannerSchema *funplugin.ObjectSchema
var proJStoreSchema *funplugin.ObjectSchema
var proJSchema *funplugin.ObjectSchema

//var orderInfoSchema *funplugin.ObjectSchema

//var pageSchema *funplugin.ObjectSchema

var load = false

func Init() {
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

	if field, err := plugin.AutoField("ProId:proj"); err != nil {
		panic(errors.New("not have object type"))
	} else {
		templateSchema.GraphQLType.AddFieldConfig("proJ", field)
		albumorderSchema.GraphQLType.AddFieldConfig("proJ", field)
	}

	if field, err := plugin.AutoField("UserId:employee"); err != nil {
		panic(errors.New("not have object type"))
	} else {
		proJStoreSchema.GraphQLType.AddFieldConfig("user", field)
	}
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

		albumorderSchema, _ = pls.NewSchemaBuilder(model.Albumorder{})
		marge(albumorderSchema)

		bannerSchema, _ = pls.NewSchemaBuilder(model.Banner{})
		marge(bannerSchema)

		proJStoreSchema, _ = pls.NewSchemaBuilder(model.ProJStore{})
		marge(proJStoreSchema)

		proJSchema, _ = pls.NewSchemaBuilder(model.ProJ{})
		marge(proJSchema)

		//orderInfoSchema, _ = pls.NewSchemaBuilder(model.OrderInfo{})
		//marge(orderInfoSchema)

		//pageSchema, _ = pls.NewSchemaBuilder(model.Page{})
		//marge(pageSchema)
		load = true

		load = true
	}

	return funplugin.Schema{
		Object: map[string]*graphql.Object{
			"template":      templateSchema.GraphQLType,
			"material":      materialSchema.GraphQLType,
			"albumorder":    albumorderSchema.GraphQLType,
			"banner":        bannerSchema.GraphQLType,
			"templateStore": proJStoreSchema.GraphQLType,
			"proJ":          proJSchema.GraphQLType,
			//"orderInfo":     orderInfoSchema.GraphQLType,

			//"page":          pageSchema.GraphQLType,
		},
		Query:    queryFields,
		Mutation: mutationFields,
	}
}
