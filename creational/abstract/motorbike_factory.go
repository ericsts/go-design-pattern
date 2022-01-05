package abstract_factory

import (
	"errors"
	"fmt"
)

type MotorbikeFactory struct{}

const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}
