// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                       = new(Query)
	Article                 *article
	Brand                   *brand
	CasbinRule              *casbinRule
	Category                *category
	Demo                    *demo
	Express                 *express
	ShippingTemplates       *shippingTemplates
	ShippingTemplatesFree   *shippingTemplatesFree
	ShippingTemplatesRegion *shippingTemplatesRegion
	Sku                     *sku
	SkuAttribute            *skuAttribute
	SmsRecord               *smsRecord
	SmsTemplate             *smsTemplate
	Spu                     *spu
	SpuLog                  *spuLog
	SpuLogistics            *spuLogistics
	SpuPicture              *spuPicture
	SpuService              *spuService
	SpuStockStream          *spuStockStream
	StoreBargain            *storeBargain
	StoreBargainUser        *storeBargainUser
	StoreBargainUserHelp    *storeBargainUserHelp
	StoreCart               *storeCart
	StoreCombination        *storeCombination
	StoreCoupon             *storeCoupon
	StoreCouponUser         *storeCouponUser
	StoreOrder              *storeOrder
	StoreOrderInfo          *storeOrderInfo
	StoreOrderStatus        *storeOrderStatus
	StorePink               *storePink
	StoreProduct            *storeProduct
	StoreProductAttr        *storeProductAttr
	StoreProductAttrResult  *storeProductAttrResult
	StoreProductAttrValue   *storeProductAttrValue
	StoreProductCate        *storeProductCate
	StoreProductCoupon      *storeProductCoupon
	StoreProductDescription *storeProductDescription
	StoreProductLog         *storeProductLog
	StoreProductRelation    *storeProductRelation
	StoreProductReply       *storeProductReply
	StoreProductRule        *storeProductRule
	StoreSeckill            *storeSeckill
	StoreSeckillManger      *storeSeckillManger
	SystemAdmin             *systemAdmin
	SystemAttachment        *systemAttachment
	SystemCity              *systemCity
	SystemConfig            *systemConfig
	SystemFormTemp          *systemFormTemp
	SystemGroup             *systemGroup
	SystemGroupData         *systemGroupData
	SystemMenu              *systemMenu
	SystemNotification      *systemNotification
	SystemRole              *systemRole
	SystemRoleMenu          *systemRoleMenu
	SystemStore             *systemStore
	SystemStoreStaff        *systemStoreStaff
	SystemUserLevel         *systemUserLevel
	TemplateMessage         *templateMessage
	User                    *user
	UserAddress             *userAddress
	UserBill                *userBill
	UserBrokerageRecord     *userBrokerageRecord
	UserExperienceRecord    *userExperienceRecord
	UserExtract             *userExtract
	UserGroup               *userGroup
	UserIntegralRecord      *userIntegralRecord
	UserLevel               *userLevel
	UserRecharge            *userRecharge
	UserSign                *userSign
	UserTag                 *userTag
	UserToken               *userToken
	UserVisitRecord         *userVisitRecord
	WechatCallback          *wechatCallback
	WechatExceptions        *wechatExceptions
	WechatPayInfo           *wechatPayInfo
	WechatReply             *wechatReply
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Article = &Q.Article
	Brand = &Q.Brand
	CasbinRule = &Q.CasbinRule
	Category = &Q.Category
	Demo = &Q.Demo
	Express = &Q.Express
	ShippingTemplates = &Q.ShippingTemplates
	ShippingTemplatesFree = &Q.ShippingTemplatesFree
	ShippingTemplatesRegion = &Q.ShippingTemplatesRegion
	Sku = &Q.Sku
	SkuAttribute = &Q.SkuAttribute
	SmsRecord = &Q.SmsRecord
	SmsTemplate = &Q.SmsTemplate
	Spu = &Q.Spu
	SpuLog = &Q.SpuLog
	SpuLogistics = &Q.SpuLogistics
	SpuPicture = &Q.SpuPicture
	SpuService = &Q.SpuService
	SpuStockStream = &Q.SpuStockStream
	StoreBargain = &Q.StoreBargain
	StoreBargainUser = &Q.StoreBargainUser
	StoreBargainUserHelp = &Q.StoreBargainUserHelp
	StoreCart = &Q.StoreCart
	StoreCombination = &Q.StoreCombination
	StoreCoupon = &Q.StoreCoupon
	StoreCouponUser = &Q.StoreCouponUser
	StoreOrder = &Q.StoreOrder
	StoreOrderInfo = &Q.StoreOrderInfo
	StoreOrderStatus = &Q.StoreOrderStatus
	StorePink = &Q.StorePink
	StoreProduct = &Q.StoreProduct
	StoreProductAttr = &Q.StoreProductAttr
	StoreProductAttrResult = &Q.StoreProductAttrResult
	StoreProductAttrValue = &Q.StoreProductAttrValue
	StoreProductCate = &Q.StoreProductCate
	StoreProductCoupon = &Q.StoreProductCoupon
	StoreProductDescription = &Q.StoreProductDescription
	StoreProductLog = &Q.StoreProductLog
	StoreProductRelation = &Q.StoreProductRelation
	StoreProductReply = &Q.StoreProductReply
	StoreProductRule = &Q.StoreProductRule
	StoreSeckill = &Q.StoreSeckill
	StoreSeckillManger = &Q.StoreSeckillManger
	SystemAdmin = &Q.SystemAdmin
	SystemAttachment = &Q.SystemAttachment
	SystemCity = &Q.SystemCity
	SystemConfig = &Q.SystemConfig
	SystemFormTemp = &Q.SystemFormTemp
	SystemGroup = &Q.SystemGroup
	SystemGroupData = &Q.SystemGroupData
	SystemMenu = &Q.SystemMenu
	SystemNotification = &Q.SystemNotification
	SystemRole = &Q.SystemRole
	SystemRoleMenu = &Q.SystemRoleMenu
	SystemStore = &Q.SystemStore
	SystemStoreStaff = &Q.SystemStoreStaff
	SystemUserLevel = &Q.SystemUserLevel
	TemplateMessage = &Q.TemplateMessage
	User = &Q.User
	UserAddress = &Q.UserAddress
	UserBill = &Q.UserBill
	UserBrokerageRecord = &Q.UserBrokerageRecord
	UserExperienceRecord = &Q.UserExperienceRecord
	UserExtract = &Q.UserExtract
	UserGroup = &Q.UserGroup
	UserIntegralRecord = &Q.UserIntegralRecord
	UserLevel = &Q.UserLevel
	UserRecharge = &Q.UserRecharge
	UserSign = &Q.UserSign
	UserTag = &Q.UserTag
	UserToken = &Q.UserToken
	UserVisitRecord = &Q.UserVisitRecord
	WechatCallback = &Q.WechatCallback
	WechatExceptions = &Q.WechatExceptions
	WechatPayInfo = &Q.WechatPayInfo
	WechatReply = &Q.WechatReply
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                      db,
		Article:                 newArticle(db, opts...),
		Brand:                   newBrand(db, opts...),
		CasbinRule:              newCasbinRule(db, opts...),
		Category:                newCategory(db, opts...),
		Demo:                    newDemo(db, opts...),
		Express:                 newExpress(db, opts...),
		ShippingTemplates:       newShippingTemplates(db, opts...),
		ShippingTemplatesFree:   newShippingTemplatesFree(db, opts...),
		ShippingTemplatesRegion: newShippingTemplatesRegion(db, opts...),
		Sku:                     newSku(db, opts...),
		SkuAttribute:            newSkuAttribute(db, opts...),
		SmsRecord:               newSmsRecord(db, opts...),
		SmsTemplate:             newSmsTemplate(db, opts...),
		Spu:                     newSpu(db, opts...),
		SpuLog:                  newSpuLog(db, opts...),
		SpuLogistics:            newSpuLogistics(db, opts...),
		SpuPicture:              newSpuPicture(db, opts...),
		SpuService:              newSpuService(db, opts...),
		SpuStockStream:          newSpuStockStream(db, opts...),
		StoreBargain:            newStoreBargain(db, opts...),
		StoreBargainUser:        newStoreBargainUser(db, opts...),
		StoreBargainUserHelp:    newStoreBargainUserHelp(db, opts...),
		StoreCart:               newStoreCart(db, opts...),
		StoreCombination:        newStoreCombination(db, opts...),
		StoreCoupon:             newStoreCoupon(db, opts...),
		StoreCouponUser:         newStoreCouponUser(db, opts...),
		StoreOrder:              newStoreOrder(db, opts...),
		StoreOrderInfo:          newStoreOrderInfo(db, opts...),
		StoreOrderStatus:        newStoreOrderStatus(db, opts...),
		StorePink:               newStorePink(db, opts...),
		StoreProduct:            newStoreProduct(db, opts...),
		StoreProductAttr:        newStoreProductAttr(db, opts...),
		StoreProductAttrResult:  newStoreProductAttrResult(db, opts...),
		StoreProductAttrValue:   newStoreProductAttrValue(db, opts...),
		StoreProductCate:        newStoreProductCate(db, opts...),
		StoreProductCoupon:      newStoreProductCoupon(db, opts...),
		StoreProductDescription: newStoreProductDescription(db, opts...),
		StoreProductLog:         newStoreProductLog(db, opts...),
		StoreProductRelation:    newStoreProductRelation(db, opts...),
		StoreProductReply:       newStoreProductReply(db, opts...),
		StoreProductRule:        newStoreProductRule(db, opts...),
		StoreSeckill:            newStoreSeckill(db, opts...),
		StoreSeckillManger:      newStoreSeckillManger(db, opts...),
		SystemAdmin:             newSystemAdmin(db, opts...),
		SystemAttachment:        newSystemAttachment(db, opts...),
		SystemCity:              newSystemCity(db, opts...),
		SystemConfig:            newSystemConfig(db, opts...),
		SystemFormTemp:          newSystemFormTemp(db, opts...),
		SystemGroup:             newSystemGroup(db, opts...),
		SystemGroupData:         newSystemGroupData(db, opts...),
		SystemMenu:              newSystemMenu(db, opts...),
		SystemNotification:      newSystemNotification(db, opts...),
		SystemRole:              newSystemRole(db, opts...),
		SystemRoleMenu:          newSystemRoleMenu(db, opts...),
		SystemStore:             newSystemStore(db, opts...),
		SystemStoreStaff:        newSystemStoreStaff(db, opts...),
		SystemUserLevel:         newSystemUserLevel(db, opts...),
		TemplateMessage:         newTemplateMessage(db, opts...),
		User:                    newUser(db, opts...),
		UserAddress:             newUserAddress(db, opts...),
		UserBill:                newUserBill(db, opts...),
		UserBrokerageRecord:     newUserBrokerageRecord(db, opts...),
		UserExperienceRecord:    newUserExperienceRecord(db, opts...),
		UserExtract:             newUserExtract(db, opts...),
		UserGroup:               newUserGroup(db, opts...),
		UserIntegralRecord:      newUserIntegralRecord(db, opts...),
		UserLevel:               newUserLevel(db, opts...),
		UserRecharge:            newUserRecharge(db, opts...),
		UserSign:                newUserSign(db, opts...),
		UserTag:                 newUserTag(db, opts...),
		UserToken:               newUserToken(db, opts...),
		UserVisitRecord:         newUserVisitRecord(db, opts...),
		WechatCallback:          newWechatCallback(db, opts...),
		WechatExceptions:        newWechatExceptions(db, opts...),
		WechatPayInfo:           newWechatPayInfo(db, opts...),
		WechatReply:             newWechatReply(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Article                 article
	Brand                   brand
	CasbinRule              casbinRule
	Category                category
	Demo                    demo
	Express                 express
	ShippingTemplates       shippingTemplates
	ShippingTemplatesFree   shippingTemplatesFree
	ShippingTemplatesRegion shippingTemplatesRegion
	Sku                     sku
	SkuAttribute            skuAttribute
	SmsRecord               smsRecord
	SmsTemplate             smsTemplate
	Spu                     spu
	SpuLog                  spuLog
	SpuLogistics            spuLogistics
	SpuPicture              spuPicture
	SpuService              spuService
	SpuStockStream          spuStockStream
	StoreBargain            storeBargain
	StoreBargainUser        storeBargainUser
	StoreBargainUserHelp    storeBargainUserHelp
	StoreCart               storeCart
	StoreCombination        storeCombination
	StoreCoupon             storeCoupon
	StoreCouponUser         storeCouponUser
	StoreOrder              storeOrder
	StoreOrderInfo          storeOrderInfo
	StoreOrderStatus        storeOrderStatus
	StorePink               storePink
	StoreProduct            storeProduct
	StoreProductAttr        storeProductAttr
	StoreProductAttrResult  storeProductAttrResult
	StoreProductAttrValue   storeProductAttrValue
	StoreProductCate        storeProductCate
	StoreProductCoupon      storeProductCoupon
	StoreProductDescription storeProductDescription
	StoreProductLog         storeProductLog
	StoreProductRelation    storeProductRelation
	StoreProductReply       storeProductReply
	StoreProductRule        storeProductRule
	StoreSeckill            storeSeckill
	StoreSeckillManger      storeSeckillManger
	SystemAdmin             systemAdmin
	SystemAttachment        systemAttachment
	SystemCity              systemCity
	SystemConfig            systemConfig
	SystemFormTemp          systemFormTemp
	SystemGroup             systemGroup
	SystemGroupData         systemGroupData
	SystemMenu              systemMenu
	SystemNotification      systemNotification
	SystemRole              systemRole
	SystemRoleMenu          systemRoleMenu
	SystemStore             systemStore
	SystemStoreStaff        systemStoreStaff
	SystemUserLevel         systemUserLevel
	TemplateMessage         templateMessage
	User                    user
	UserAddress             userAddress
	UserBill                userBill
	UserBrokerageRecord     userBrokerageRecord
	UserExperienceRecord    userExperienceRecord
	UserExtract             userExtract
	UserGroup               userGroup
	UserIntegralRecord      userIntegralRecord
	UserLevel               userLevel
	UserRecharge            userRecharge
	UserSign                userSign
	UserTag                 userTag
	UserToken               userToken
	UserVisitRecord         userVisitRecord
	WechatCallback          wechatCallback
	WechatExceptions        wechatExceptions
	WechatPayInfo           wechatPayInfo
	WechatReply             wechatReply
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                      db,
		Article:                 q.Article.clone(db),
		Brand:                   q.Brand.clone(db),
		CasbinRule:              q.CasbinRule.clone(db),
		Category:                q.Category.clone(db),
		Demo:                    q.Demo.clone(db),
		Express:                 q.Express.clone(db),
		ShippingTemplates:       q.ShippingTemplates.clone(db),
		ShippingTemplatesFree:   q.ShippingTemplatesFree.clone(db),
		ShippingTemplatesRegion: q.ShippingTemplatesRegion.clone(db),
		Sku:                     q.Sku.clone(db),
		SkuAttribute:            q.SkuAttribute.clone(db),
		SmsRecord:               q.SmsRecord.clone(db),
		SmsTemplate:             q.SmsTemplate.clone(db),
		Spu:                     q.Spu.clone(db),
		SpuLog:                  q.SpuLog.clone(db),
		SpuLogistics:            q.SpuLogistics.clone(db),
		SpuPicture:              q.SpuPicture.clone(db),
		SpuService:              q.SpuService.clone(db),
		SpuStockStream:          q.SpuStockStream.clone(db),
		StoreBargain:            q.StoreBargain.clone(db),
		StoreBargainUser:        q.StoreBargainUser.clone(db),
		StoreBargainUserHelp:    q.StoreBargainUserHelp.clone(db),
		StoreCart:               q.StoreCart.clone(db),
		StoreCombination:        q.StoreCombination.clone(db),
		StoreCoupon:             q.StoreCoupon.clone(db),
		StoreCouponUser:         q.StoreCouponUser.clone(db),
		StoreOrder:              q.StoreOrder.clone(db),
		StoreOrderInfo:          q.StoreOrderInfo.clone(db),
		StoreOrderStatus:        q.StoreOrderStatus.clone(db),
		StorePink:               q.StorePink.clone(db),
		StoreProduct:            q.StoreProduct.clone(db),
		StoreProductAttr:        q.StoreProductAttr.clone(db),
		StoreProductAttrResult:  q.StoreProductAttrResult.clone(db),
		StoreProductAttrValue:   q.StoreProductAttrValue.clone(db),
		StoreProductCate:        q.StoreProductCate.clone(db),
		StoreProductCoupon:      q.StoreProductCoupon.clone(db),
		StoreProductDescription: q.StoreProductDescription.clone(db),
		StoreProductLog:         q.StoreProductLog.clone(db),
		StoreProductRelation:    q.StoreProductRelation.clone(db),
		StoreProductReply:       q.StoreProductReply.clone(db),
		StoreProductRule:        q.StoreProductRule.clone(db),
		StoreSeckill:            q.StoreSeckill.clone(db),
		StoreSeckillManger:      q.StoreSeckillManger.clone(db),
		SystemAdmin:             q.SystemAdmin.clone(db),
		SystemAttachment:        q.SystemAttachment.clone(db),
		SystemCity:              q.SystemCity.clone(db),
		SystemConfig:            q.SystemConfig.clone(db),
		SystemFormTemp:          q.SystemFormTemp.clone(db),
		SystemGroup:             q.SystemGroup.clone(db),
		SystemGroupData:         q.SystemGroupData.clone(db),
		SystemMenu:              q.SystemMenu.clone(db),
		SystemNotification:      q.SystemNotification.clone(db),
		SystemRole:              q.SystemRole.clone(db),
		SystemRoleMenu:          q.SystemRoleMenu.clone(db),
		SystemStore:             q.SystemStore.clone(db),
		SystemStoreStaff:        q.SystemStoreStaff.clone(db),
		SystemUserLevel:         q.SystemUserLevel.clone(db),
		TemplateMessage:         q.TemplateMessage.clone(db),
		User:                    q.User.clone(db),
		UserAddress:             q.UserAddress.clone(db),
		UserBill:                q.UserBill.clone(db),
		UserBrokerageRecord:     q.UserBrokerageRecord.clone(db),
		UserExperienceRecord:    q.UserExperienceRecord.clone(db),
		UserExtract:             q.UserExtract.clone(db),
		UserGroup:               q.UserGroup.clone(db),
		UserIntegralRecord:      q.UserIntegralRecord.clone(db),
		UserLevel:               q.UserLevel.clone(db),
		UserRecharge:            q.UserRecharge.clone(db),
		UserSign:                q.UserSign.clone(db),
		UserTag:                 q.UserTag.clone(db),
		UserToken:               q.UserToken.clone(db),
		UserVisitRecord:         q.UserVisitRecord.clone(db),
		WechatCallback:          q.WechatCallback.clone(db),
		WechatExceptions:        q.WechatExceptions.clone(db),
		WechatPayInfo:           q.WechatPayInfo.clone(db),
		WechatReply:             q.WechatReply.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                      db,
		Article:                 q.Article.replaceDB(db),
		Brand:                   q.Brand.replaceDB(db),
		CasbinRule:              q.CasbinRule.replaceDB(db),
		Category:                q.Category.replaceDB(db),
		Demo:                    q.Demo.replaceDB(db),
		Express:                 q.Express.replaceDB(db),
		ShippingTemplates:       q.ShippingTemplates.replaceDB(db),
		ShippingTemplatesFree:   q.ShippingTemplatesFree.replaceDB(db),
		ShippingTemplatesRegion: q.ShippingTemplatesRegion.replaceDB(db),
		Sku:                     q.Sku.replaceDB(db),
		SkuAttribute:            q.SkuAttribute.replaceDB(db),
		SmsRecord:               q.SmsRecord.replaceDB(db),
		SmsTemplate:             q.SmsTemplate.replaceDB(db),
		Spu:                     q.Spu.replaceDB(db),
		SpuLog:                  q.SpuLog.replaceDB(db),
		SpuLogistics:            q.SpuLogistics.replaceDB(db),
		SpuPicture:              q.SpuPicture.replaceDB(db),
		SpuService:              q.SpuService.replaceDB(db),
		SpuStockStream:          q.SpuStockStream.replaceDB(db),
		StoreBargain:            q.StoreBargain.replaceDB(db),
		StoreBargainUser:        q.StoreBargainUser.replaceDB(db),
		StoreBargainUserHelp:    q.StoreBargainUserHelp.replaceDB(db),
		StoreCart:               q.StoreCart.replaceDB(db),
		StoreCombination:        q.StoreCombination.replaceDB(db),
		StoreCoupon:             q.StoreCoupon.replaceDB(db),
		StoreCouponUser:         q.StoreCouponUser.replaceDB(db),
		StoreOrder:              q.StoreOrder.replaceDB(db),
		StoreOrderInfo:          q.StoreOrderInfo.replaceDB(db),
		StoreOrderStatus:        q.StoreOrderStatus.replaceDB(db),
		StorePink:               q.StorePink.replaceDB(db),
		StoreProduct:            q.StoreProduct.replaceDB(db),
		StoreProductAttr:        q.StoreProductAttr.replaceDB(db),
		StoreProductAttrResult:  q.StoreProductAttrResult.replaceDB(db),
		StoreProductAttrValue:   q.StoreProductAttrValue.replaceDB(db),
		StoreProductCate:        q.StoreProductCate.replaceDB(db),
		StoreProductCoupon:      q.StoreProductCoupon.replaceDB(db),
		StoreProductDescription: q.StoreProductDescription.replaceDB(db),
		StoreProductLog:         q.StoreProductLog.replaceDB(db),
		StoreProductRelation:    q.StoreProductRelation.replaceDB(db),
		StoreProductReply:       q.StoreProductReply.replaceDB(db),
		StoreProductRule:        q.StoreProductRule.replaceDB(db),
		StoreSeckill:            q.StoreSeckill.replaceDB(db),
		StoreSeckillManger:      q.StoreSeckillManger.replaceDB(db),
		SystemAdmin:             q.SystemAdmin.replaceDB(db),
		SystemAttachment:        q.SystemAttachment.replaceDB(db),
		SystemCity:              q.SystemCity.replaceDB(db),
		SystemConfig:            q.SystemConfig.replaceDB(db),
		SystemFormTemp:          q.SystemFormTemp.replaceDB(db),
		SystemGroup:             q.SystemGroup.replaceDB(db),
		SystemGroupData:         q.SystemGroupData.replaceDB(db),
		SystemMenu:              q.SystemMenu.replaceDB(db),
		SystemNotification:      q.SystemNotification.replaceDB(db),
		SystemRole:              q.SystemRole.replaceDB(db),
		SystemRoleMenu:          q.SystemRoleMenu.replaceDB(db),
		SystemStore:             q.SystemStore.replaceDB(db),
		SystemStoreStaff:        q.SystemStoreStaff.replaceDB(db),
		SystemUserLevel:         q.SystemUserLevel.replaceDB(db),
		TemplateMessage:         q.TemplateMessage.replaceDB(db),
		User:                    q.User.replaceDB(db),
		UserAddress:             q.UserAddress.replaceDB(db),
		UserBill:                q.UserBill.replaceDB(db),
		UserBrokerageRecord:     q.UserBrokerageRecord.replaceDB(db),
		UserExperienceRecord:    q.UserExperienceRecord.replaceDB(db),
		UserExtract:             q.UserExtract.replaceDB(db),
		UserGroup:               q.UserGroup.replaceDB(db),
		UserIntegralRecord:      q.UserIntegralRecord.replaceDB(db),
		UserLevel:               q.UserLevel.replaceDB(db),
		UserRecharge:            q.UserRecharge.replaceDB(db),
		UserSign:                q.UserSign.replaceDB(db),
		UserTag:                 q.UserTag.replaceDB(db),
		UserToken:               q.UserToken.replaceDB(db),
		UserVisitRecord:         q.UserVisitRecord.replaceDB(db),
		WechatCallback:          q.WechatCallback.replaceDB(db),
		WechatExceptions:        q.WechatExceptions.replaceDB(db),
		WechatPayInfo:           q.WechatPayInfo.replaceDB(db),
		WechatReply:             q.WechatReply.replaceDB(db),
	}
}

type queryCtx struct {
	Article                 IArticleDo
	Brand                   IBrandDo
	CasbinRule              ICasbinRuleDo
	Category                ICategoryDo
	Demo                    IDemoDo
	Express                 IExpressDo
	ShippingTemplates       IShippingTemplatesDo
	ShippingTemplatesFree   IShippingTemplatesFreeDo
	ShippingTemplatesRegion IShippingTemplatesRegionDo
	Sku                     ISkuDo
	SkuAttribute            ISkuAttributeDo
	SmsRecord               ISmsRecordDo
	SmsTemplate             ISmsTemplateDo
	Spu                     ISpuDo
	SpuLog                  ISpuLogDo
	SpuLogistics            ISpuLogisticsDo
	SpuPicture              ISpuPictureDo
	SpuService              ISpuServiceDo
	SpuStockStream          ISpuStockStreamDo
	StoreBargain            IStoreBargainDo
	StoreBargainUser        IStoreBargainUserDo
	StoreBargainUserHelp    IStoreBargainUserHelpDo
	StoreCart               IStoreCartDo
	StoreCombination        IStoreCombinationDo
	StoreCoupon             IStoreCouponDo
	StoreCouponUser         IStoreCouponUserDo
	StoreOrder              IStoreOrderDo
	StoreOrderInfo          IStoreOrderInfoDo
	StoreOrderStatus        IStoreOrderStatusDo
	StorePink               IStorePinkDo
	StoreProduct            IStoreProductDo
	StoreProductAttr        IStoreProductAttrDo
	StoreProductAttrResult  IStoreProductAttrResultDo
	StoreProductAttrValue   IStoreProductAttrValueDo
	StoreProductCate        IStoreProductCateDo
	StoreProductCoupon      IStoreProductCouponDo
	StoreProductDescription IStoreProductDescriptionDo
	StoreProductLog         IStoreProductLogDo
	StoreProductRelation    IStoreProductRelationDo
	StoreProductReply       IStoreProductReplyDo
	StoreProductRule        IStoreProductRuleDo
	StoreSeckill            IStoreSeckillDo
	StoreSeckillManger      IStoreSeckillMangerDo
	SystemAdmin             ISystemAdminDo
	SystemAttachment        ISystemAttachmentDo
	SystemCity              ISystemCityDo
	SystemConfig            ISystemConfigDo
	SystemFormTemp          ISystemFormTempDo
	SystemGroup             ISystemGroupDo
	SystemGroupData         ISystemGroupDataDo
	SystemMenu              ISystemMenuDo
	SystemNotification      ISystemNotificationDo
	SystemRole              ISystemRoleDo
	SystemRoleMenu          ISystemRoleMenuDo
	SystemStore             ISystemStoreDo
	SystemStoreStaff        ISystemStoreStaffDo
	SystemUserLevel         ISystemUserLevelDo
	TemplateMessage         ITemplateMessageDo
	User                    IUserDo
	UserAddress             IUserAddressDo
	UserBill                IUserBillDo
	UserBrokerageRecord     IUserBrokerageRecordDo
	UserExperienceRecord    IUserExperienceRecordDo
	UserExtract             IUserExtractDo
	UserGroup               IUserGroupDo
	UserIntegralRecord      IUserIntegralRecordDo
	UserLevel               IUserLevelDo
	UserRecharge            IUserRechargeDo
	UserSign                IUserSignDo
	UserTag                 IUserTagDo
	UserToken               IUserTokenDo
	UserVisitRecord         IUserVisitRecordDo
	WechatCallback          IWechatCallbackDo
	WechatExceptions        IWechatExceptionsDo
	WechatPayInfo           IWechatPayInfoDo
	WechatReply             IWechatReplyDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Article:                 q.Article.WithContext(ctx),
		Brand:                   q.Brand.WithContext(ctx),
		CasbinRule:              q.CasbinRule.WithContext(ctx),
		Category:                q.Category.WithContext(ctx),
		Demo:                    q.Demo.WithContext(ctx),
		Express:                 q.Express.WithContext(ctx),
		ShippingTemplates:       q.ShippingTemplates.WithContext(ctx),
		ShippingTemplatesFree:   q.ShippingTemplatesFree.WithContext(ctx),
		ShippingTemplatesRegion: q.ShippingTemplatesRegion.WithContext(ctx),
		Sku:                     q.Sku.WithContext(ctx),
		SkuAttribute:            q.SkuAttribute.WithContext(ctx),
		SmsRecord:               q.SmsRecord.WithContext(ctx),
		SmsTemplate:             q.SmsTemplate.WithContext(ctx),
		Spu:                     q.Spu.WithContext(ctx),
		SpuLog:                  q.SpuLog.WithContext(ctx),
		SpuLogistics:            q.SpuLogistics.WithContext(ctx),
		SpuPicture:              q.SpuPicture.WithContext(ctx),
		SpuService:              q.SpuService.WithContext(ctx),
		SpuStockStream:          q.SpuStockStream.WithContext(ctx),
		StoreBargain:            q.StoreBargain.WithContext(ctx),
		StoreBargainUser:        q.StoreBargainUser.WithContext(ctx),
		StoreBargainUserHelp:    q.StoreBargainUserHelp.WithContext(ctx),
		StoreCart:               q.StoreCart.WithContext(ctx),
		StoreCombination:        q.StoreCombination.WithContext(ctx),
		StoreCoupon:             q.StoreCoupon.WithContext(ctx),
		StoreCouponUser:         q.StoreCouponUser.WithContext(ctx),
		StoreOrder:              q.StoreOrder.WithContext(ctx),
		StoreOrderInfo:          q.StoreOrderInfo.WithContext(ctx),
		StoreOrderStatus:        q.StoreOrderStatus.WithContext(ctx),
		StorePink:               q.StorePink.WithContext(ctx),
		StoreProduct:            q.StoreProduct.WithContext(ctx),
		StoreProductAttr:        q.StoreProductAttr.WithContext(ctx),
		StoreProductAttrResult:  q.StoreProductAttrResult.WithContext(ctx),
		StoreProductAttrValue:   q.StoreProductAttrValue.WithContext(ctx),
		StoreProductCate:        q.StoreProductCate.WithContext(ctx),
		StoreProductCoupon:      q.StoreProductCoupon.WithContext(ctx),
		StoreProductDescription: q.StoreProductDescription.WithContext(ctx),
		StoreProductLog:         q.StoreProductLog.WithContext(ctx),
		StoreProductRelation:    q.StoreProductRelation.WithContext(ctx),
		StoreProductReply:       q.StoreProductReply.WithContext(ctx),
		StoreProductRule:        q.StoreProductRule.WithContext(ctx),
		StoreSeckill:            q.StoreSeckill.WithContext(ctx),
		StoreSeckillManger:      q.StoreSeckillManger.WithContext(ctx),
		SystemAdmin:             q.SystemAdmin.WithContext(ctx),
		SystemAttachment:        q.SystemAttachment.WithContext(ctx),
		SystemCity:              q.SystemCity.WithContext(ctx),
		SystemConfig:            q.SystemConfig.WithContext(ctx),
		SystemFormTemp:          q.SystemFormTemp.WithContext(ctx),
		SystemGroup:             q.SystemGroup.WithContext(ctx),
		SystemGroupData:         q.SystemGroupData.WithContext(ctx),
		SystemMenu:              q.SystemMenu.WithContext(ctx),
		SystemNotification:      q.SystemNotification.WithContext(ctx),
		SystemRole:              q.SystemRole.WithContext(ctx),
		SystemRoleMenu:          q.SystemRoleMenu.WithContext(ctx),
		SystemStore:             q.SystemStore.WithContext(ctx),
		SystemStoreStaff:        q.SystemStoreStaff.WithContext(ctx),
		SystemUserLevel:         q.SystemUserLevel.WithContext(ctx),
		TemplateMessage:         q.TemplateMessage.WithContext(ctx),
		User:                    q.User.WithContext(ctx),
		UserAddress:             q.UserAddress.WithContext(ctx),
		UserBill:                q.UserBill.WithContext(ctx),
		UserBrokerageRecord:     q.UserBrokerageRecord.WithContext(ctx),
		UserExperienceRecord:    q.UserExperienceRecord.WithContext(ctx),
		UserExtract:             q.UserExtract.WithContext(ctx),
		UserGroup:               q.UserGroup.WithContext(ctx),
		UserIntegralRecord:      q.UserIntegralRecord.WithContext(ctx),
		UserLevel:               q.UserLevel.WithContext(ctx),
		UserRecharge:            q.UserRecharge.WithContext(ctx),
		UserSign:                q.UserSign.WithContext(ctx),
		UserTag:                 q.UserTag.WithContext(ctx),
		UserToken:               q.UserToken.WithContext(ctx),
		UserVisitRecord:         q.UserVisitRecord.WithContext(ctx),
		WechatCallback:          q.WechatCallback.WithContext(ctx),
		WechatExceptions:        q.WechatExceptions.WithContext(ctx),
		WechatPayInfo:           q.WechatPayInfo.WithContext(ctx),
		WechatReply:             q.WechatReply.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
