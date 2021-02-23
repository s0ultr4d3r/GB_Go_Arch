package service

import (
	"GB/lesson-2/shop/models"
	"reflect"
	"testing"
)

func Test_service_CreateItem(t *testing.T) {
	type args struct {
		item *models.Item
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    *models.Item
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateItem(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateOrder(t *testing.T) {
	type args struct {
		order *models.Order
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    *models.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateOrder(tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
