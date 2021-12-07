package generate

//func customMiddleware(handler middleware.Handler) middleware.Handler {
//	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
//		if tr, ok := transport.FromServerContext(ctx); ok {
//			fmt.Println("operation:", tr.Operation())
//		}
//		reply, err = handler(ctx, req)
//		return
//	}
//}

//func main() {
//	router := gin.Default()
//	// 使用kratos中间件
//	router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))
//
//	router.GET("/helloworld/:name", func(ctx *gin.Context) {
//		name := ctx.Param("name")
//		if name == "error" {
//			// 返回kratos error
//			kgin.Error(ctx, errors.Unauthorized("auth_error", "no authentication"))
//		} else {
//			ctx.JSON(200, map[string]string{"welcome": name})
//		}
//	})
//
//	httpSrv := http.NewServer(http.Address(":9000"))
//	httpSrv.HandlePrefix("/", router)
//	//grpcSrv := grpc.NewServer(grpc.Address(":8000"))
//
//	app := kratos.New(
//		kratos.Name("gin"),
//		kratos.Server(
//			httpSrv,
//		),
//	)
//	if err := app.Run(); err != nil {
//		log.Fatal(err)
//	}
//}
