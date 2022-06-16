package com.example.agent.validator;

import com.example.agent.dtos.*;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;

import java.time.LocalDateTime;

public class Validators {
    private static PasswordConstraintValidator passwordConstraintValidator = new PasswordConstraintValidator();

    public static boolean isValidUserDto(UserRegistrationDTO dto) {
        return  dto.getUsername() != ""
                && passwordConstraintValidator.isValid(dto.getPassword(), null)
                && dto.getFirstName() != ""
                && dto.getLastName() != ""
                && dto.getDateOfBirth().isBefore(LocalDateTime.now());
    }

    public static boolean isValidCompanyRegistrationRequestDto(CompanyRegistrationRequestDTO dto) {
        return dto.getCompanyName() != ""
                && dto.getCompanyOwnerUsername() != ""
                && dto.getCompanyOwnerName() != ""
                && dto.getCompanyContactInfo() != ""
                && dto.getCompanyDescription() != ""
                && dto.getUsername() != "";
    }

    public static boolean isValidCompanyInfoDTO(CompanyInfoDTO dto) {
        return dto.getContactInfo() != ""
                && dto.getDescription() != ""
                && dto.getId() >= 0;
    }

    public static boolean isValidOpenPositionDto(Long companyId, String positionName, String description, String criteria) {
        return companyId >= 0
                && positionName != ""
                && description != ""
                && criteria != "";
    }

    public static boolean isValidCommentDto(CommentDTO dto) {
        return  dto.getUserSignature() != ""
                && dto.getContent() != ""
                && dto.getCompanyId() >= 0
                && dto.getUsername() != ""
                && dto.getUserSignature() != "";
    }

    public static boolean isValidSallaryDto(SalaryDTO dto) {
        return dto.getPositionId() >= 0
                && dto.getUserId() >= 0
                && dto.getSalary() >= 400 && dto.getSalary() <= 30000;
    }

    public static boolean isValidInterviewProcessDto(InterviewProcessDTO dto) {
        return dto.getUsername() != ""
                && dto.getUserSignature() != ""
                && dto.getCompanyId() >= 0
                && dto.getInterviewDescription() != ""
                && dto.getUserId() >= 0;
    }

    public static boolean isValidMarkDto(MarkDTO dto) {
        return dto.getMark() >= 1 && dto.getMark() <= 5
                && dto.getUserId() >= 0
                && dto.getCompanyId() >= 0;
    }

    public static boolean isValidToken(Long userId, String token) {
        return userId >= 0 && token.length() == 32;
    }

}
