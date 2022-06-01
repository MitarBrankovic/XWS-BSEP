package com.example.agent.service;

import com.example.agent.dtos.CommentDTO;
import com.example.agent.dtos.CompanyInfoDTO;
import com.example.agent.dtos.CompanyRegistrationRequestDTO;
import com.example.agent.dtos.UserRegistrationDTO;

public interface AgentService {
    void saveUser(UserRegistrationDTO userRegistrationDTO);

    void saveCompanyRegistrationRequest(CompanyRegistrationRequestDTO dto);

    void saveCompany(CompanyRegistrationRequestDTO dto);

    void editCompanyInfo(CompanyInfoDTO dto);

    void addOpenPosition(Long companyId, String positionName);

    void saveComment(CommentDTO dto);
}
