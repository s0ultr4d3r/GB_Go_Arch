package repository

import (
	"gb-go-architecture/lesson-1/shop/models"
	"reflect"
	"testing"
	"time"
)

func TestNewMapDB(t *testing.T) {
	mapDB, ok := NewMapDB().(*mapDB)
	if !ok {
		t.Error("Can't open DB")
	}
	if mapDB.itemsTable == nil {
		t.Error("Can't open DB")
	}
}
func TestCreateItem(t *testing.T) {
	db := NewMapDB()

	input := &models.Item{
		Name:  "someName",
		Price: 10,
	}
	expected := &models.Item{
		ID:    1,
		Name:  input.Name,
		Price: input.Price,
	}

	result, err := db.CreateItem(input)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if expected.ID != result.ID {
		t.Errorf("unexpected name: expected %d result: %d", expected.ID, result.ID)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected name: expected %s result: %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected name: expected %d result: %d", expected.Price, result.Price)
	}

	result, err = db.GetItem(expected.ID)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if expected.ID != result.ID {
		t.Errorf("unexpected name: expected %d result: %d", expected.ID, result.ID)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected name: expected %s result: %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected name: expected %d result: %d", expected.Price, result.Price)
	}

	input = &models.Item{
		Name:  "someName2",
		Price: 20,
	}
	expected = &models.Item{
		ID:    2,
		Name:  input.Name,
		Price: input.Price,
	}

	result, err = db.CreateItem(input)
	if err != nil {
		t.Error("unexpected error: ", err)
	}

	if expected.ID != result.ID {
		t.Errorf("unexpected name: expected %d result: %d", expected.ID, result.ID)
	}
	if expected.Name != result.Name {
		t.Errorf("unexpected name: expected %s result: %s", expected.Name, result.Name)
	}
	if expected.Price != result.Price {
		t.Errorf("unexpected name: expected %d result: %d", expected.Price, result.Price)
	}
}

func TestGetItem(t *testing.T) {
	MyDB := &mapDB{
		itemsTable: &itemsTable{
			items: map[int32]*models.Item{
				1: {
					ID:    1,
					Name:  "TestItem_1",
					Price: 10.0,
				},
				2: {
					ID:    2,
					Name:  "TestItem_2",
					Price: 20.0,
				},
				3: {
					ID:    3,
					Name:  "TestItem_3",
					Price: 30.0,
				},
			},
			maxID: 3,
		},
	}

	getTest, err := MyDB.GetItem(3)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(getTest, MyDB.itemsTable.items[3]) {
		t.Fatalf("want %v get %v", MyDB.itemsTable.items[3], getTest)
	}
}

func TestDeleteItem(t *testing.T) {
	MyDB := &mapDB{
		itemsTable: &itemsTable{
			items: map[int32]*models.Item{
				1: {
					ID:    1,
					Name:  "TestItem_1",
					Price: 10.0,
				},
				2: {
					ID:    2,
					Name:  "TestItem_2",
					Price: 20.0,
				},
				3: {
					ID:    3,
					Name:  "TestItem_3",
					Price: 30.0,
				},
			},
			maxID: 3,
		},
	}

	MyDB.DeleteItem(2)

	if len(MyDB.itemsTable.items) != 2 {
		t.Fatalf("Function DeleteItem is not work")
	}
}

func TestUpdateItem(t *testing.T) {
	MyDB := &mapDB{
		itemsTable: &itemsTable{
			items: map[int32]*models.Item{
				1: {
					ID:    1,
					Name:  "TestItem_1",
					Price: 10.0,
				},
				2: {
					ID:    2,
					Name:  "TestItem_2",
					Price: 20.0,
				},
				3: {
					ID:    3,
					Name:  "TestItem_3",
					Price: 30.0,
				},
			},
			maxID: 3,
		},
	}

	want := &models.Item{
		ID:        2,
		Name:      "Update_Test_Ok",
		Price:     20.0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	_, err := MyDB.UpdateItem(want)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, MyDB.itemsTable.items[2]) {
		t.Fatalf("want %v got %v", want, MyDB.itemsTable.items[2])
	}
}
