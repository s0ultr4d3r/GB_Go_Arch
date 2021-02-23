package main

import{
	"GB/shop/Lesson-1/notification"
	"GB/shop/Lesson-1/repository"
	"GB/shop/Lesson-1/service"
}

constructors []func(interface)

type Builder struct {
	service service.Service
	rep repository.Repository
	notif notification.Notification
	server     *server
	ServiceFund func(rep repository.Repository, notif notification.Notification) service.Service
	Rep        func() repository.Repository
	NotifFunc      func NewTelegramBot(token string, chatID int64) (Notification, errror)
}

func (b *Builder) Build(token string, chatID int64) (err error) {
b.notif, err := b.NotifFunc(tokenStr, 38266)
	if err != nil {
		return err
	}
	b.rep = b.RepFunc()
	b.service = b.ServiceFund(b.rep, b.notif)
	b.server := &server{

		service: b.service,
		rep:     b.rep,
	}
return nil
}
