package mapper

import (
	"tanyaustadz/domain/entity"
	"tanyaustadz/internal/repository/psql/models"
	"tanyaustadz/pkg/common"
)

func ToDomainDataObject(m *models.DataObject) *entity.DataObject {
	id, _ := common.StringToID(m.Id)
	dataObject := &entity.DataObject{
		Id:        id,
		Path:      m.Path,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}

	return dataObject
}

func ToModelDataObject(d *entity.DataObject) *models.DataObject {
	dataObject := &models.DataObject{
		Id:        d.Id.String(),
		Path:      d.Path,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}

	return dataObject
}
