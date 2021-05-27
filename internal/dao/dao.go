package dao

import (
	_ "database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/multierr"

	"github.com/jimmykodes/vehicle_maintenance/internal/settings"
)

type DAO struct {
	Service     Service
	ServiceType ServiceType
	User        User
	Vehicle     Vehicle
}

func New(dbSettings settings.DB) (*DAO, error) {
	db, err := sqlx.Open(dbSettings.DriveName, dbSettings.DNS())
	if err != nil {
		return nil, err
	}
	vehicle, err := newVehicle(db, dbSettings.Database)
	if err != nil {
		return nil, err
	}
	service, err := newService(db, dbSettings.Database)
	if err != nil {
		return nil, err
	}
	serviceType, err := newServiceType(db, dbSettings.Database)
	if err != nil {
		return nil, err
	}
	user, err := newUser(db, dbSettings.Database)
	if err != nil {
		return nil, err
	}
	return &DAO{
		Vehicle:     vehicle,
		Service:     service,
		ServiceType: serviceType,
		User:        user,
	}, nil
}

func (d DAO) Close() error {
	return multierr.Combine(
		d.Service.Close(),
		d.ServiceType.Close(),
		d.User.Close(),
		d.Vehicle.Close(),
	)
}
