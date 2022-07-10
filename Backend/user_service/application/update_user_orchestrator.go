package application

import (
	"dislinkt/common/domain"
	saga "dislinkt/common/saga/messaging"
	events "dislinkt/common/saga/update_user"
)

type UpdateUserOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewUpdateUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateUserOrchestrator, error) {
	o := &UpdateUserOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *UpdateUserOrchestrator) Start(user *domain.User, oldUsername string,
	oldFirstName string, oldLastName string, oldPrivate bool) error {
	event := &events.UpdateUserCommand{
		User:         *user,
		OldUsername:  oldUsername,
		OldFirstName: oldFirstName,
		OldLastName:  oldLastName,
		OldPrivate:   oldPrivate,
		Type:         events.UpdateUser,
	}
	return o.commandPublisher.Publish(event)
}

func (o *UpdateUserOrchestrator) handle(reply *events.UpdateUserReply) {
	command := events.UpdateUserCommand{
		User:         reply.User,
		OldUsername:  reply.OldUsername,
		OldFirstName: reply.OldFirstName,
		OldLastName:  reply.OldLastName,
		OldPrivate:   reply.OldPrivate,
	}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *UpdateUserOrchestrator) nextCommandType(reply events.UpdateUserReplyType) events.UpdateUserCommandType {
	switch reply {
	case events.UserNotUpdated:
		return events.RollbackUpdatedUser
	default:
		return events.UnknownCommand
	}
}
