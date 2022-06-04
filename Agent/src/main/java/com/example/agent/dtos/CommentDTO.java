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
    private String username;

    public CommentDTO() {
    }

    public CommentDTO(String content, String userSignature, Long companyId, Long userId, String username) {
        this.content = content;
        this.userSignature = userSignature;
        this.companyId = companyId;
        this.userId = userId;
        this.username = username;
    }
}
