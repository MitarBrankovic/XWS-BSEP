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

    public InterviewProcessDTO(){}
}
