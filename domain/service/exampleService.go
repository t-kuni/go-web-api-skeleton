package service

import (
	"github.com/t-kuni/go-cli-app-skeleton/domain/infrastructure/api"
)

type ExampleService struct {
	BinanceApi api.BinanceApiInterface
}

func ProvideExampleService(binanceApi api.BinanceApiInterface) ExampleService {
	return ExampleService{binanceApi}
}

func (s ExampleService) Exec(baseAsset string) (string, error) {
	info, err := s.BinanceApi.GetExchangeInfo(baseAsset)
	if err != nil {
		return "", err
	}
	return info.Symbols[0].Status, nil
}
