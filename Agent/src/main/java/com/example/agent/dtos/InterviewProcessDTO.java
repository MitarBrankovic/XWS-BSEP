package com.example.agent.dtos;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class InterviewProcessDTO {
    private Long userId;
    private Long companyId;
    private String interviewDescription;
    private String userSignature;
    private String username;

    public InterviewProcessDTO(){}

    public InterviewProcessDTO(Long userId, Long companyId, String interviewDescription, String userSignature, String username) {
        this.userId = userId;
        this.companyId = companyId;
        this.interviewDescription = interviewDescription;
        this.userSignature = userSignature;
        this.username = username;
    }
}
