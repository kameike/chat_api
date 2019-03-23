package repository

import (
	"github.com/kameike/chat_api/datasource"
)

type chatRepository struct {
	ds datasource.DataSourceDescriptor
}
