package api

import (
	saga "dislinkt/common/saga/messaging"
	events "dislinkt/common/saga/update_user"
	"dislinkt/user_service/application"
)

type UpdateUserCommandHandler struct {
	userService       *application.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateProfileCommandHandler(userService *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateUserCommandHandler, error) {
	o := &UpdateUserCommandHandler{
		userService:       userService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *UpdateUserCommandHandler) handle(command *events.UpdateUserCommand) {
	reply := &events.UpdateUserReply{
		User:         command.User,
		Type:         events.UnknownReply,
		OldUsername:  command.OldUsername,
		OldFirstName: command.OldFirstName,
		OldLastName:  command.OldLastName,
	}
	switch command.Type {
	case events.RollbackUpdatedUser:
		oldUser := command.User
		oldUser.Username = command.OldUsername
		oldUser.FirstName = command.OldFirstName
		oldUser.LastName = command.OldLastName
		err := handler.userService.RollbackUpdate(mapCommonUserToUser(&oldUser))
		if err != nil {
			return
		}
		reply.Type = events.RollbackUser
	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
