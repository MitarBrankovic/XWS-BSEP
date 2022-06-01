package com.example.agent.dtos;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class MarkDTO {
    private Long companyId;
    private Long userId;
    private Integer mark;

    public MarkDTO() {
    }
}
