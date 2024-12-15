package handlers

import "github.com/dills122/receipt-rack/store"

var dataStore store.Store

func Init(s store.Store) {
	dataStore = s
}
