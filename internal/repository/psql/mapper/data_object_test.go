package mapper_test

import (
	"tanyaustadz/internal/repository/psql/mapper"
	"tanyaustadz/test/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataObjectMapper(t *testing.T) {
	dataObjectModel := testdata.NewDataObjectModel()
	dataObjectDomain := testdata.NewDataObject()

	res := mapper.ToDomainDataObject(dataObjectModel)

	assert.Equal(t, res, dataObjectDomain)
}
