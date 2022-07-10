package api

import (
	commonDomain "dislinkt/common/domain"
	pb "dislinkt/common/proto/user_service"
	"dislinkt/user_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapUserToPb(user *domain.User) *pb.User {
	pbUser := &pb.User{
		Id:               user.Id.Hex(),
		Username:         user.Username,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		DateOfBirth:      timestamppb.New(user.DateOfBirth),
		Email:            user.Email,
		Education:        make([]*pb.Education, 0),
		WorkExperience:   make([]*pb.WorkExperience, 0),
		Skills:           make([]string, 0),
		Interests:        make([]string, 0),
		Activated:        user.Activated,
		Private:          user.Private,
		TwoFactorEnabled: user.TwoFactorEnabled,
	}

	for _, education := range user.Education {
		educationPb := &pb.Education{
			School:       education.School,
			Degree:       education.Degree,
			FieldOfStudy: education.FieldOfStudy,
			StartDate:    timestamppb.New(education.StartDate),
			EndDate:      timestamppb.New(education.EndDate),
		}
		pbUser.Education = append(pbUser.Education, educationPb)
	}

	for _, workExperience := range user.WorkExperience {
		workExperiencePb := &pb.WorkExperience{
			Title:          workExperience.Title,
			Company:        workExperience.Company,
			EmploymentType: 0,
			Location:       workExperience.Location,
			StartDate:      timestamppb.New(workExperience.StartDate),
			EndDate:        timestamppb.New(workExperience.StartDate),
		}
		pbUser.WorkExperience = append(pbUser.WorkExperience, workExperiencePb)
	}

	for _, skill := range user.Skills {
		pbUser.Skills = append(pbUser.Skills, skill)
	}

	for _, interest := range user.Interests {
		pbUser.Interests = append(pbUser.Interests, interest)
	}

	return pbUser
}

func mapPbToUser(pbUser *pb.User) *domain.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pbUser.Password), bcrypt.DefaultCost)
	user := &domain.User{
		Id:               getObjectId(pbUser.Id),
		Username:         pbUser.Username,
		HashedPassword:   string(hashedPassword),
		FirstName:        pbUser.FirstName,
		LastName:         pbUser.LastName,
		FullName:         pbUser.FirstName + " " + pbUser.LastName,
		DateOfBirth:      pbUser.DateOfBirth.AsTime(),
		Email:            pbUser.Email,
		Role:             "unverified",
		Private:          pbUser.Private,
		Activated:        pbUser.Activated,
		TwoFactorEnabled: pbUser.TwoFactorEnabled,
		Education:        make([]domain.Education, 0),
		WorkExperience:   make([]domain.WorkExperience, 0),
		Skills:           make([]string, 0),
		Interests:        make([]string, 0),
	}
	for _, education := range pbUser.Education {
		education := &domain.Education{
			School:       education.School,
			Degree:       education.Degree,
			FieldOfStudy: education.FieldOfStudy,
			StartDate:    education.StartDate.AsTime(),
			EndDate:      education.EndDate.AsTime(),
		}
		user.Education = append(user.Education, *education)
	}

	for _, workExperience := range pbUser.WorkExperience {
		workExperience := &domain.WorkExperience{
			Title:          workExperience.Title,
			Company:        workExperience.Company,
			EmploymentType: workExperience.EmploymentType.String(),
			Location:       workExperience.Location,
			StartDate:      workExperience.StartDate.AsTime(),
			EndDate:        workExperience.StartDate.AsTime(),
		}
		user.WorkExperience = append(user.WorkExperience, *workExperience)
	}

	for _, skill := range pbUser.Skills {
		user.Skills = append(user.Skills, skill)
	}

	for _, interest := range pbUser.Interests {
		user.Interests = append(user.Interests, interest)
	}

	return user
}

func mapCommonUserToUser(commonUser *commonDomain.User) *domain.User {
	user := &domain.User{
		Id:                 commonUser.Id,
		Username:           commonUser.Username,
		FirstName:          commonUser.FirstName,
		LastName:           commonUser.LastName,
		FullName:           commonUser.FirstName + " " + commonUser.LastName,
		Email:              commonUser.Email,
		Role:               commonUser.Role,
		Private:            commonUser.Private,
		Activated:          commonUser.Activated,
		TwoFactorEnabled:   commonUser.TwoFactorEnabled,
		TwoFactorToken:     commonUser.TwoFactorToken,
		TwoFactorTokenDate: commonUser.TwoFactorTokenDate,
		ApiToken:           commonUser.ApiToken,
		HashedPassword:     commonUser.HashedPassword,
		DateOfBirth:        commonUser.DateOfBirth,
	}
	for _, education := range user.Education {
		education := &domain.Education{
			School:       education.School,
			Degree:       education.Degree,
			FieldOfStudy: education.FieldOfStudy,
			StartDate:    education.StartDate,
			EndDate:      education.EndDate,
		}
		user.Education = append(user.Education, *education)
	}

	for _, workExperience := range user.WorkExperience {
		workExperience := &domain.WorkExperience{
			Title:          workExperience.Title,
			Company:        workExperience.Company,
			EmploymentType: workExperience.EmploymentType,
			Location:       workExperience.Location,
			StartDate:      workExperience.StartDate,
			EndDate:        workExperience.StartDate,
		}
		user.WorkExperience = append(user.WorkExperience, *workExperience)
	}

	for _, skill := range user.Skills {
		user.Skills = append(user.Skills, skill)
	}

	for _, interest := range user.Interests {
		user.Interests = append(user.Interests, interest)
	}
	return user
}

func mapPbBlockToBlock(pbBlock *pb.Block) *domain.Block {
	block := &domain.Block{
		IssuerUsername:  pbBlock.IssuerUsername,
		SubjectUsername: pbBlock.SubjectUsername,
	}
	return block
}

func mapBlockToPbBlock(block *domain.Block) *pb.Block {
	pbBlock := &pb.Block{
		IssuerUsername:  block.IssuerUsername,
		SubjectUsername: block.SubjectUsername,
	}
	return pbBlock
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
