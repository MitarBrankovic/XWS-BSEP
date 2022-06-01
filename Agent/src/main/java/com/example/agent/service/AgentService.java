package com.example.agent.service;

import com.example.agent.dtos.*;

public interface AgentService {
    void saveUser(UserRegistrationDTO userRegistrationDTO);

    void saveCompanyRegistrationRequest(CompanyRegistrationRequestDTO dto);

    void saveCompany(CompanyRegistrationRequestDTO dto);

    void editCompanyInfo(CompanyInfoDTO dto);

    void addOpenPosition(Long companyId, String positionName);

    void saveComment(CommentDTO dto);

    void addSallary(SallaryDTO dto);
}
