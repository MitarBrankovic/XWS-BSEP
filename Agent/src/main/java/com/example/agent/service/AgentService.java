package com.example.agent.service;

import com.example.agent.domain.AgentUser;
import com.example.agent.domain.Company;
import com.example.agent.dtos.*;

import java.util.List;
import java.util.Set;

public interface AgentService {
    void saveUser(UserRegistrationDTO userRegistrationDTO);

    void saveCompanyRegistrationRequest(CompanyRegistrationRequestDTO dto);

    void saveCompany(CompanyRegistrationRequestDTO dto);

    void editCompanyInfo(CompanyInfoDTO dto);

    void addOpenPosition(Long companyId, String positionName);

    void saveComment(CommentDTO dto);

    void addSallary(SallaryDTO dto);

    void addInterviewProcess(InterviewProcessDTO dto);

    void addMark(MarkDTO dto);

    List<Company> findAllCompanies();

    Company findOneCompanyById(Long companyId);

    AgentUser findUser(String username, String password);

    List<CompanyRegistrationRequestDTO> findAllCompanyRegistrationRequests();
}
