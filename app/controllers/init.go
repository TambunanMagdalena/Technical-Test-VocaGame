package controllers

import (
	"template-go/app/usecases"
	"template-go/pkg/config"
)

type Main struct {
	Item ItemInterface
}

type controller struct {
	Options Options
}

type Options struct {
	Config   *config.Config
	UseCases *usecases.Main
}

func Init(opts Options) *Main {
	ctrl := &controller{opts}

	m := &Main{
		Item: (*itemController)(ctrl),
	}

	return m
}
