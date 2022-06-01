package com.example.agent.dtos;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class SallaryDTO {
    private Long userId;
    private Long positionId;
    private Double sallary;

    public SallaryDTO() {
    }
}
