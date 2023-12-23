# Gateway

**⚠️ 注意：该仓库已停止维护，请前往新仓库进行后续开发。**

**🆕 新仓库地址：[AirAccountGateway](https://github.com/AAStarCommunity/AirAccountGateway)**

The Gateway of all instructions, transfers into on-chain commands to execute.



```
go run main.go
// after downloading the package, it should show this:
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (1 handlers)
[GIN-debug] GET    /api/instructions/balance --> AirAccountGateway/internal/controllers/actions.Balance (1 handlers)
[GIN-debug] POST   /api/instructions/bind    --> AirAccountGateway/internal/controllers/actions.Bind (1 handlers)
[GIN-debug] GET    /api/instructions/transfer/check --> AirAccountGateway/internal/controllers/actions.TransferCheck (1 handlers)
[GIN-debug] POST   /api/instructions/transfer --> AirAccountGateway/internal/controllers/actions.Transfer (1 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000
```

Access localhost:8000/swagger/index.html to see the API documentation and debug.
