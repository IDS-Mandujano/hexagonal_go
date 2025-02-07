package machineusecases

import "gym-system/src/inventory/domain/repositories/Machine"

type GetStatusMachine struct {
	db repositories.IMachineRepository
}

func NewMachineStatus(db repositories.IMachineRepository) *GetStatusMachine {
	return &GetStatusMachine{db: db}
}

func (getMStatus *GetStatusMachine) Execute(id int) (string, error) {
	return getMStatus.db.GetStatus(id)
}