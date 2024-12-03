package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/project-sistem-voucher/api/handler"
	"github.com/project-sistem-voucher/config"
	"github.com/project-sistem-voucher/manager"
	"github.com/project-sistem-voucher/middleware"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(router *gin.Engine) error {

	router.Use(middleware.LogRequestMiddleware(logrus.New()))

	infraManager := manager.NewInfraManager(config.Cfg)
	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager)

	voucherHandler := handler.NewVoucherHandler(repoManager.VoucherService())
	redeemHandler := handler.NewRedeemHandler(repoManager.RedeemService())
	handlerApplicationVoucher := handler.NewHandlerApplicationVoucher(repoManager.ServiceApplicationVoucher())
	handlerHistoryVoucher := handler.NewHandlerHistoryVoucher(repoManager.ServiceHistoryVoucher())
	v1 := router.Group("/api/v1")
	{
		sistemVoucher := v1.Group("/management-voucher")
		{
			sistemVoucher.POST("/create", voucherHandler.CreateVoucher)
			sistemVoucher.DELETE("/delete/:id", voucherHandler.DeleteVoucher)
			sistemVoucher.PUT("update/:id", voucherHandler.UpdateVoucher)
			sistemVoucher.GET("list", voucherHandler.GetVouchers)
			sistemVoucher.GET("/redeem-list", voucherHandler.GetVouchersForRedeem)
			sistemVoucher.POST("/redeem", redeemHandler.RedeemVoucher)
		}
		applyVoucher := v1.Group("/apply-voucher")
		{
			applyVoucher.POST("/use", handlerApplicationVoucher.CreateUseVoucher)
			applyVoucher.POST("/validate", handlerApplicationVoucher.ValidateVoucher)
			applyVoucher.GET("/:userID/:voucherType", handlerApplicationVoucher.GetMyVoucherByCategory)
		}
		historyVoucher := v1.Group("/history-voucher")
		{
			historyVoucher.GET("/reedem/:userID", handlerHistoryVoucher.GetReedemVoucherByUserId)
			historyVoucher.GET("/use/:userID", handlerHistoryVoucher.GetUseVoucherByUserId)
			historyVoucher.GET("/all/:kode_voucher", handlerHistoryVoucher.GetAllUseByVoucherCode)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router.Run()

}
