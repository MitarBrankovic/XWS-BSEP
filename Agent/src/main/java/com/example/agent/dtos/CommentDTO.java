package com.example.agent.dtos;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class CommentDTO {
    private String content;
    private String userSignature;
    private Long companyId;
    private Long userId;

    public CommentDTO() {
    }
}
