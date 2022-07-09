package update_user

import "dislinkt/common/domain"

type UpdateUserCommandType int8

const (
	UpdateUser UpdateUserCommandType = iota
	RollbackUpdatedUser
	UnknownCommand
)

type UpdateUserCommand struct {
	User         domain.User
	OldUsername  string
	OldFirstName string
	OldLastName  string
	Type         UpdateUserCommandType
}

type UpdateUserReplyType int8

const (
	UserUpdated UpdateUserReplyType = iota
	UserNotUpdated
	RollbackUser
	UnknownReply
)

type UpdateUserReply struct {
	User         domain.User
	OldUsername  string
	OldFirstName string
	OldLastName  string
	Type         UpdateUserReplyType
}
