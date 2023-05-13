package grpc

//func TestServer_ConsumeMessage_OK(t *testing.T) {
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	bytes := []byte(`{"msg":"hello world!"}`)
//
//	grpcServer := grpc.NewServer()
//
//	consumerUsecase := mockConsumer.NewMockUsecase(ctl)
//
//	consumerService := NewConsumerServiceGRPCServer(grpcServer, consumerUsecase)
//
//	consumerUsecase.EXPECT().ConsumeMessage(context.TODO()).Return(bytes).Times(1)
//
//	protoBytes, err := consumerService.ConsumeMessage(context.TODO(), nil)
//
//	require.NoError(t, err)
//	require.Equal(t, bytes, model_conversion.FromProtoBytesToBytes(protoBytes))
//}

//func TestServer_StartConsumeMessages_OK(t *testing.T) {
//	ctl := gomock.NewController(t)
//	defer ctl.Finish()
//
//	grpcServer := grpc.NewServer()
//
//	consumerUsecase := mockConsumer.NewMockUsecase(ctl)
//
//	consumerService := NewConsumerServiceGRPCServer(grpcServer, consumerUsecase)
//
//	_, err := consumerService.StartConsumeMessages(context.TODO(), nil)
//
//	require.NoError(t, err)
//}
