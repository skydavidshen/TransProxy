package translator

import (
	"TransProxy/model/business"
	"TransProxy/model/request"
)

type Handler interface {
	InsertItem(item request.Item) error
	Translate(item request.Item) (business.TranslateItem, error)
}