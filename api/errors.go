package api

import (
	"github.com/Conflux-Chain/go-conflux-util/api"
)

const (
	ErrCodeDatabase              = 4
	ErrCodeBlockchainRPC         = 5
	ErrCodeStorageNodeRPC        = 6
	ErrCodeNoMatchingRecordFound = 7
	ErrCodeUnmarshalValue        = 8
)

func ErrDatabase(err error) error {
	return api.NewBusinessError(ErrCodeDatabase, "Database exception", err.Error())
}

func ErrBlockchainRPC(err error) error {
	return api.NewBusinessError(ErrCodeBlockchainRPC, "Blockchain RPC exception", err.Error())
}

func ErrStorageNodeRPC(err error) error {
	return api.NewBusinessError(ErrCodeStorageNodeRPC, "Storage node RPC exception", err.Error())
}

func ErrNoMatchingRecordFound(err error) error {
	return api.NewBusinessError(ErrCodeNoMatchingRecordFound, "No matching record found", err.Error())
}

func ErrUnmarshalValue(err error) error {
	return api.NewBusinessError(ErrCodeUnmarshalValue, "Failed to unmarshal value", err.Error())
}
