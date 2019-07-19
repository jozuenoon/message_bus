package query

import "github.com/jozuenoon/message_bus/pkg/types"

type Repository interface {
	GetDetectors(latitude, longitude types.DecimalDegrees, radius int64)
}
