package com.example.agent.dtos;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CompanyInfoDTO {
    private Long id;
    private String contactInfo;
    private String description;

    public CompanyInfoDTO() {
    }
}
