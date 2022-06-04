package com.example.agent.dtos;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class SalaryDTO {
    private Long userId;
    private Long positionId;
    private Double salary;

    public SalaryDTO() {
    }
}
