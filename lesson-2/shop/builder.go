package main

type Builder struct {
	service service.Service
	rep repository.Repository
	notif notification.Notification
	server     *server
	serviceFund func(rep repository.Repository, notif notification.Notification) service.Service
	rep        func() repository.Repository
	notifFunc      func NewTelegramBot(token string, chatID int64) (Notification, errror)
}

func (b *Builder) Build(token string, chatID int64) (err error) {
b.notif, err := b.notifFunc(tokenStr, 38266)
	if err != nil {
		return err
	}
	b.rep = b.repFunc()
	b.service = b.serviceFund(b.rep, b.notif)
	b.server := &server{

		service: b.service,
		rep:     b.rep,
	}
return nil
}
