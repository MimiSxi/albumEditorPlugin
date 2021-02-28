package schema

import (
	"github.com/Fiber-Man/albumEditorPlugin/model"
	"github.com/Fiber-Man/funplugin"
	"github.com/graphql-go/graphql"
)

var templateSchema *funplugin.ObjectSchema
var materialSchema *funplugin.ObjectSchema
var albumOrderSchema *funplugin.ObjectSchema
var bannerSchema *funplugin.ObjectSchema
var templateStoreSchema *funplugin.ObjectSchema

var load = false

func Init() {
	// InitAccount()
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

		templateStoreSchema, _ = pls.NewSchemaBuilder(model.TemplateStores{})
		marge(templateStoreSchema)
		load = true
	}

	// roleSchema, _ := pls.NewSchemaBuilder(model.Role{})
	// marge(roleSchema)

	// roleAccountSchema, _ := pls.NewSchemaBuilder(model.RoleAccount{})
	// marge(roleAccountSchema)

	return funplugin.Schema{
		Object: map[string]*graphql.Object{
			// "account": accountType,
			"template":      templateSchema.GraphQLType,
			"material":      materialSchema.GraphQLType,
			"albumOrder":    albumOrderSchema.GraphQLType,
			"banner":        bannerSchema.GraphQLType,
			"templateStore": templateStoreSchema.GraphQLType,
			// "role":        roleSchema.GraphQLType,
			// "roleaccount": roleAccountSchema.GraphQLType,
		},
		Query:    queryFields,
		Mutation: mutationFields,
	}
}
