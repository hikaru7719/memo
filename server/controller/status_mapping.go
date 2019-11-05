package controller

import "github.com/hikaru7719/memo/server/domain"

func convertDomainStatus(status int32) domain.Status {
	switch status {
	case 0:
		return domain.Unknown
	case 1:
		return domain.Todo
	case 2:
		return domain.Progress
	case 3:
		return domain.Done
	default:
		return domain.Unknown
	}
}

func convertProtoStatus(status domain.Status) int32 {
	switch status {
	case domain.Unknown:
		return 0
	case domain.Todo:
		return 1
	case domain.Progress:
		return 2
	case domain.Done:
		return 3
	default:
		return 0
	}
}
