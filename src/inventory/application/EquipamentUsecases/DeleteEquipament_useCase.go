package equipamentusecases

import "gym-system/src/inventory/domain/repositories"

type DeleteEquipment struct {
	db repositories.IEquipamentRepository
}

func NewDeleteEquipment (db repositories.IEquipamentRepository) *DeleteEquipment {
	return &DeleteEquipment{db: db}
}

func (deleteEquipment *DeleteEquipment) Execute(id int){
	deleteEquipment.db.Delete(id)
}