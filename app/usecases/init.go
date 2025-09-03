package usecases

import (
	"template-go/app/repositories"
	"template-go/pkg/config"
)

type Main struct {
	Item ItemInterface
}

type usecase struct {
	Options Options
}

type Options struct {
	Repository *repositories.Main
	Config     *config.Config
}

func Init(opts Options) *Main {
	ucs := &usecase{
		Options: opts,
	}

	m := &Main{
		Item: (*itemUsecase)(ucs),
	}

	return m
}
