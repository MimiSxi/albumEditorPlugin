package model

import "github.com/Fiber-Man/funplugin"

// 通用状态枚举类型
type CommonStatusEnumType uint

const (
	C_ENABLE  CommonStatusEnumType = 1 // 可用
	C_DISABLE CommonStatusEnumType = 2 // 不可用
	C_DELETE  CommonStatusEnumType = 3 // 删除
)

func (s CommonStatusEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"C_ENABLE": funplugin.EnumValue{
			Value:       C_ENABLE,
			Description: "可用",
		},
		"C_DISABLE": funplugin.EnumValue{
			Value:       C_DISABLE,
			Description: "不可用",
		},
		"C_DELETE": funplugin.EnumValue{
			Value:       C_DELETE,
			Description: "删除",
		},
	}
}

// 分类级别枚举类型
type MaterialKindLevelEnumType uint

const (
	ONE   MaterialKindLevelEnumType = 1 // 一级分类
	TWO   MaterialKindLevelEnumType = 2 // 二级分类
	THREE MaterialKindLevelEnumType = 2 // 三级分类
)

func (s MaterialKindLevelEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"ONE": funplugin.EnumValue{
			Value:       ONE,
			Description: "一级分类",
		},
		"TWO": funplugin.EnumValue{
			Value:       TWO,
			Description: "二级分类",
		},
		"THREE": funplugin.EnumValue{
			Value:       THREE,
			Description: "三级分类",
		},
	}
}

// 纸张规格枚举类型
type AlbumOrderSpecsEnumType uint

const (
	K16_A3 AlbumOrderSpecsEnumType = 1 // 16开A3纸
	K32_A4 AlbumOrderSpecsEnumType = 2 // 32开A4纸
)

func (s AlbumOrderSpecsEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"K16_A3": funplugin.EnumValue{
			Value:       K16_A3,
			Description: "16开A3纸",
		},
		"K32_A4": funplugin.EnumValue{
			Value:       K32_A4,
			Description: "32开A4纸",
		},
	}
}

// 纸张材质枚举类型
type AlbumOrderMaterialEnumType uint

const (
	COPPERPLATE_200G AlbumOrderMaterialEnumType = 1 // 200g铜版纸
)

func (s AlbumOrderMaterialEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"COPPERPLATE_200G": funplugin.EnumValue{
			Value:       COPPERPLATE_200G,
			Description: "200g铜版纸",
		},
	}
}

// 使用类型枚举类型
type AlbumOrderUsageTypeEnumType uint

const (
	PERSONAL_OR_CHARITY AlbumOrderUsageTypeEnumType = 1 //个人/公益使用
)

func (s AlbumOrderUsageTypeEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"PERSONAL_OR_CHARITY": funplugin.EnumValue{
			Value:       PERSONAL_OR_CHARITY,
			Description: "个人/公益使用",
		},
	}
}

// 订单支付方式枚举类型
type AlbumOrderPayWayEnumType uint

const (
	WECHAT AlbumOrderPayWayEnumType = 1 //微信
	ALIPAY AlbumOrderPayWayEnumType = 2 //支付宝
)

func (s AlbumOrderPayWayEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"WECHAT": funplugin.EnumValue{
			Value:       WECHAT,
			Description: "微信",
		},
		"ALIPAY": funplugin.EnumValue{
			Value:       ALIPAY,
			Description: "支付宝",
		},
	}
}

// 订单状态枚举类型
type AlbumOrderStatusEnumType uint

const (
	TO_BE_PAID    AlbumOrderStatusEnumType = 1 //待支付
	TO_BE_DELIVER AlbumOrderStatusEnumType = 2 //待发货
	HAS_DELIVER   AlbumOrderStatusEnumType = 3 //已发货
	HAS_RECEIVED  AlbumOrderStatusEnumType = 4 //已收货
	CANCELED      AlbumOrderStatusEnumType = 5 //订单取消
)

func (s AlbumOrderStatusEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"TO_BE_PAID": funplugin.EnumValue{
			Value:       TO_BE_PAID,
			Description: "待支付",
		},
		"TO_BE_DELIVER": funplugin.EnumValue{
			Value:       TO_BE_DELIVER,
			Description: "待发货",
		},
		"HAS_DELIVER": funplugin.EnumValue{
			Value:       HAS_DELIVER,
			Description: "已发货",
		},
		"HAS_RECEIVED": funplugin.EnumValue{
			Value:       HAS_RECEIVED,
			Description: "已收货",
		},
		"CANCELED": funplugin.EnumValue{
			Value:       CANCELED,
			Description: "订单取消",
		},
	}
}

// 设计器项目通用状态枚举类型
type ProJCommonStatusEnumType uint

const (
	P_ENABLE  ProJCommonStatusEnumType = 1 // 正常
	P_DISABLE ProJCommonStatusEnumType = 2 // 停用
	P_DELETE  ProJCommonStatusEnumType = 2 // 删除
)

func (s ProJCommonStatusEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"P_ENABLE": funplugin.EnumValue{
			Value:       P_ENABLE,
			Description: "正常",
		},
		"P_DISABLE": funplugin.EnumValue{
			Value:       P_DISABLE,
			Description: "停用",
		},
		"P_DELETE": funplugin.EnumValue{
			Value:       P_DELETE,
			Description: "删除",
		},
	}
}

// 画布页面种类枚举类型
type PageTypeEnumType uint

const (
	COVER       PageTypeEnumType = 1 // 封面
	BACK_COVER  PageTypeEnumType = 2 // 封底
	CERTIFICATE PageTypeEnumType = 3 // 证书
	NORMAL      PageTypeEnumType = 4 // 普通
	TITLE_PAGE  PageTypeEnumType = 5 // 扉页
)

func (s PageTypeEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"COVER": funplugin.EnumValue{
			Value:       COVER,
			Description: "封面",
		},
		"BACK_COVER": funplugin.EnumValue{
			Value:       BACK_COVER,
			Description: "封底",
		},
		"CERTIFICATE": funplugin.EnumValue{
			Value:       CERTIFICATE,
			Description: "证书",
		},
		"NORMAL": funplugin.EnumValue{
			Value:       NORMAL,
			Description: "普通",
		},
		"TITLE_PAGE": funplugin.EnumValue{
			Value:       TITLE_PAGE,
			Description: "扉页",
		},
	}
}

// 画布页面方向枚举类型
type PageDirectionEnumType uint

const (
	V PageDirectionEnumType = 1 // 垂直方向
	H PageDirectionEnumType = 2 // 水平方向
)

func (s PageDirectionEnumType) Enum() map[string]funplugin.EnumValue {
	return map[string]funplugin.EnumValue{
		"V": funplugin.EnumValue{
			Value:       V,
			Description: "垂直方向",
		},
		"H": funplugin.EnumValue{
			Value:       H,
			Description: "水平方向",
		},
	}
}
