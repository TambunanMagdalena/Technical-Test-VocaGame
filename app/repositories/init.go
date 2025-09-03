package repositories

import (
	
	"template-go/pkg/config"

	"gorm.io/gorm"
)

type Main struct {
	Item        ItemInterface
}


type repository struct {
	Options Options
}

type Options struct {
	Postgres *gorm.DB
	Config   *config.Config
}

func Init(opts Options) *Main {
	repo := &repository{opts}

	m := &Main{
		Item:        (*itemRepository)(repo),
		
	}

	return m
}
