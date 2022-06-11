package com.example.agent.service;

import com.example.agent.domain.*;
import com.example.agent.dtos.*;
import org.springframework.security.core.userdetails.UserDetailsService;

import java.util.List;
import java.util.Set;

public interface AgentService{
    void saveUser(UserRegistrationDTO userRegistrationDTO);

    void saveCompanyRegistrationRequest(CompanyRegistrationRequestDTO dto);

    void saveCompany(CompanyRegistrationRequestDTO dto);

    void editCompanyInfo(CompanyInfoDTO dto);

    void addOpenPosition(Long companyId, String positionName, String description, String criteria);

    void saveComment(CommentDTO dto);

    void addSalary(SalaryDTO dto);

    void addInterviewProcess(InterviewProcessDTO dto);

    void addMark(MarkDTO dto);

    List<Company> findAllCompanies();

    Company findOneCompanyById(Long companyId);

    AgentUser findUser(String username, String password);

    List<CompanyRegistrationRequestDTO> findAllCompanyRegistrationRequests();

    Set<CommentOnCompany> findAllCommentsByCompanyId(Long companyId);

    Set<InterviewProcess> findAllInterviewsByCompanyId(Long companyId);

    Set<OpenPosition> findAllPositionsByCompanyId(Long companyId);

    AgentUser saveToken(Long userId, String token);
}
