package third_party

type SlsService interface {
	CreateLogistic()
	CancelLogistic()
}

type WebapiService interface {
	CreateOrder()
	CancelOrder()
}

type WmsService interface {
	CreateWmsOrder()
	CancelWmsOrder()
}
