package api

import (
	saga "dislinkt/common/saga/messaging"
	events "dislinkt/common/saga/update_user"
	"dislinkt/connection_service/application"
	"dislinkt/connection_service/domain"
)

type UpdateUserCommandHandler struct {
	connectionService *application.ConnectionService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateUserCommandHandler(connectionService *application.ConnectionService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateUserCommandHandler, error) {
	o := &UpdateUserCommandHandler{
		connectionService: connectionService,
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
		User:        command.User,
		Type:        events.UnknownReply,
		OldUsername: command.OldUsername,
		OldPrivate:  command.OldPrivate,
	}
	switch command.Type {
	case events.UpdateUser:
		if command.User.Username == command.OldUsername && command.User.Private == command.OldPrivate {
			return
		}
		newUser := &domain.User{
			Username: command.User.Username,
			Private:  command.User.Private,
		}
		err, _ := handler.connectionService.UpdateUser(newUser.Username, mapSagaUserToPb(newUser))
		if err != nil {
			return
		}
		reply.Type = events.UserUpdated
		break
	case events.RollbackUpdatedUser:
		oldUser := &domain.User{
			Username: command.OldUsername,
			Private:  command.OldPrivate,
		}
		err, _ := handler.connectionService.UpdateUser(oldUser.Username, mapSagaUserToPb(oldUser))
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
