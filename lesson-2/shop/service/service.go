package service

import (
	"errors"
	"log"

	"GB/lesson-2/shop/models"
	"GB/lesson-2/shop/pkg/notification"
	rep "GB/lesson-2/shop/repository"
)

type Service interface {
	CreateOrder(order *models.Order) (*models.Order, error)
}

type service struct {
	notif notification.Notification
	rep   rep.Repository
}

var (
	ErrItemNotExists = errors.New("item not exists")
)

func (s *service) CreateOrder(order *models.Order) (*models.Order, error) {
	for _, itemID := range order.ItemIDs {
		_, err := s.rep.GetItem(int32(itemID))
		if err != nil {
			return nil, ErrItemNotExists
		}
	}

	order, err := s.rep.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	if err := s.notif.SendOrderCreated(order); err != nil {
		log.Println(err)
	}

	return order, err
}

func (s *service) CreateItem(item *models.Item) (*models.Item, error) {
	if item.Name == "" {
		return nil, errors.New("item name is empty")
	}
	if item.Price <= 0 {
		return nil, errors.New("item price should be positive")
	}

	return s.rep.CreateItem(item)
}

func NewService(rep rep.Repository, notif notification.Notification) Service {
	return &service{
		rep: rep,
	}
}
