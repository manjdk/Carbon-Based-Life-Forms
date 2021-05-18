package repository

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/google/go-cmp/cmp"
	"github.com/manjdk/Carbon-Based-Life-Forms/domain"
	"github.com/stretchr/testify/assert"
)

func TestDynamoDB_Create(t *testing.T) {
	mineral := &domain.Mineral{
		ID:        uuid.Must(uuid.NewV4()).String(),
		ClientID:  "testClientID",
		Name:      "testName",
		State:     "testState",
		Fractures: 56,
	}

	if err := testDynamoDB.Create(mineral); err != nil {
		t.Errorf(err.Error())
		return
	}

	if err := testDynamoDB.Create(mineral); err == nil {
		t.Errorf("should fail because of duplication")
		return
	}

	deleteAfterTest(t, mineral.ID)
}

func TestDynamoDB_Update(t *testing.T) {
	mineral := &domain.Mineral{
		ID:        uuid.Must(uuid.NewV4()).String(),
		ClientID:  "testClientID",
		Name:      "testName",
		State:     "testState",
		Fractures: 56,
	}

	if err := testDynamoDB.Create(mineral); err != nil {
		t.Errorf(err.Error())
		return
	}

	mineral.SetState(domain.Solid)
	if err := testDynamoDB.Update(mineral); err != nil {
		t.Errorf(err.Error())
		return
	}

	mineralAfterUpdate, err := testDynamoDB.Get(mineral.ID)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if diff := cmp.Diff(mineral, mineralAfterUpdate); diff != "" {
		t.Errorf(diff)
	}

	deleteAfterTest(t, mineral.ID)
}

func TestDynamoDB_Get(t *testing.T) {
	mineral := &domain.Mineral{
		ID:        uuid.Must(uuid.NewV4()).String(),
		ClientID:  "testClientID",
		Name:      "testName",
		State:     "testState",
		Fractures: 56,
	}

	if err := testDynamoDB.Create(mineral); err != nil {
		t.Errorf(err.Error())
		return
	}

	gotMineral, err := testDynamoDB.Get(mineral.ID)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if diff := cmp.Diff(gotMineral, mineral); diff != "" {
		t.Errorf(diff)
	}

	deleteAfterTest(t, mineral.ID)
}

func TestDynamoDB_GetAll(t *testing.T) {
	a := assert.New(t)

	mineral := &domain.Mineral{
		ID:        uuid.Must(uuid.NewV4()).String(),
		ClientID:  "testClientID",
		Name:      "testName",
		State:     "testState",
		Fractures: 56,
	}

	mineral2 := &domain.Mineral{
		ID:        uuid.Must(uuid.NewV4()).String(),
		ClientID:  "testClientID",
		Name:      "testName",
		State:     "testState",
		Fractures: 56,
	}

	mineral3 := &domain.Mineral{
		ID:        uuid.Must(uuid.NewV4()).String(),
		ClientID:  "testClientID",
		Name:      "testName",
		State:     "testState",
		Fractures: 56,
	}

	var err error
	err = testDynamoDB.Create(mineral)
	err = testDynamoDB.Create(mineral2)
	err = testDynamoDB.Create(mineral3)

	if err != nil {
		t.Errorf(err.Error())
		deleteAfterTest(t, mineral.ID, mineral2.ID, mineral3.ID)
		return
	}

	minerals, err := testDynamoDB.GetAll()
	if err != nil {
		t.Errorf(err.Error())
		deleteAfterTest(t, mineral.ID, mineral2.ID, mineral3.ID)
		return
	}

	a.Equal(3, len(minerals))
	deleteAfterTest(t, mineral.ID, mineral2.ID, mineral3.ID)
}

func TestDynamoDB_GetByClientID(t *testing.T) {
	a := assert.New(t)

	mineral := &domain.Mineral{
		ID:        uuid.Must(uuid.NewV4()).String(),
		ClientID:  "testClientID",
		Name:      "testName",
		State:     "testState",
		Fractures: 56,
	}

	mineral2 := &domain.Mineral{
		ID:        uuid.Must(uuid.NewV4()).String(),
		ClientID:  "testClientID",
		Name:      "testName",
		State:     "testState",
		Fractures: 56,
	}

	mineral3 := &domain.Mineral{
		ID:        uuid.Must(uuid.NewV4()).String(),
		ClientID:  "testClientID2",
		Name:      "testName",
		State:     "testState",
		Fractures: 56,
	}

	var err error
	err = testDynamoDB.Create(mineral)
	err = testDynamoDB.Create(mineral2)
	err = testDynamoDB.Create(mineral3)

	if err != nil {
		t.Errorf(err.Error())
		deleteAfterTest(t, mineral.ID, mineral2.ID, mineral3.ID)
		return
	}

	minerals, err := testDynamoDB.GetByClientID("testClientID")
	if err != nil {
		t.Errorf(err.Error())
		deleteAfterTest(t, mineral.ID, mineral2.ID, mineral3.ID)
		return
	}
	a.Equal(2, len(minerals))

	minerals, err = testDynamoDB.GetByClientID("testClientID2")
	if err != nil {
		t.Errorf(err.Error())
		deleteAfterTest(t, mineral.ID, mineral2.ID, mineral3.ID)
		return
	}
	a.Equal(1, len(minerals))

	deleteAfterTest(t, mineral.ID, mineral2.ID, mineral3.ID)
}

func deleteAfterTest(t *testing.T, mineralID ...string) {
	for _, id := range mineralID {
		if err := testDynamoDB.Delete(id); err != nil {
			t.Errorf(err.Error())
		}
	}
}
