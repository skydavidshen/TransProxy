package translator

import (
	"TransProxy/model/business"
	"TransProxy/model/request"
)

const contentType = "application/json"

type Handler interface {
	InsertItem(item request.Item) error
	Translate(item request.Item) (business.TranslateItem, error)
}