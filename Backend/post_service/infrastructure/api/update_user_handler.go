package api

import (
	saga "dislinkt/common/saga/messaging"
	events "dislinkt/common/saga/update_user"
	"dislinkt/post_service/application"
	"dislinkt/post_service/domain"
)

type UpdateUserCommandHandler struct {
	postService       *application.PostService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateUserCommandHandler(postService *application.PostService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateUserCommandHandler, error) {
	o := &UpdateUserCommandHandler{
		postService:       postService,
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
	case events.UpdateUser:
		if command.User.FirstName == command.OldFirstName && command.User.LastName == command.OldLastName && command.User.Username == command.OldUsername {
			return
		}
		newUser := &domain.User{
			Id:        command.User.Id,
			Username:  command.User.Username,
			FirstName: command.User.FirstName,
			LastName:  command.User.LastName,
		}
		err, _ := handler.postService.UpdateUser(newUser.Username, mapUserToPb(newUser))
		if err != nil {
			return
		}
		reply.Type = events.UserUpdated
		break
	case events.RollbackUpdatedUser:
		oldUser := &domain.User{
			Id:        command.User.Id,
			Username:  command.OldUsername,
			FirstName: command.OldFirstName,
			LastName:  command.OldLastName,
		}
		err, _ := handler.postService.UpdateUser(oldUser.Username, mapUserToPb(oldUser))
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
