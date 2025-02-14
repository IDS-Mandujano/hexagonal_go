package equipamentusecases

import "gym-system/src/inventory/domain/repositories/Equipment"

type GetEquipmentByCondition struct {
	db repositories.IEquipamentRepository
}

func NewEquipmentByCondition (db repositories.IEquipamentRepository) *GetEquipmentByCondition {
	return &GetEquipmentByCondition{db: db}
}

func (getEquipmentByCondition *GetEquipmentByCondition) Execute (condition string) ([]map[string]interface{}, error){
	return getEquipmentByCondition.db.GetCondition(condition)
}