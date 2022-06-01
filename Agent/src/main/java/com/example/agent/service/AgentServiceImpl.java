package com.example.agent.service;

import com.example.agent.domain.*;
import com.example.agent.dtos.CommentDTO;
import com.example.agent.dtos.CompanyInfoDTO;
import com.example.agent.dtos.CompanyRegistrationRequestDTO;
import com.example.agent.dtos.UserRegistrationDTO;
import com.example.agent.repository.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class AgentServiceImpl implements AgentService {

    @Autowired
    private AgentUserRepository agentUserRepository;

    @Autowired
    private CompanyRegistrationRequestRepository companyRegistrationRequestRepository;

    @Autowired
    private CompanyRepository companyRepository;

    @Autowired
    private OpenPositionRepository openPositionRepository;

    @Autowired
    private CommentOnCompanyRepository commentOnCompanyRepository;

    @Override
    public void saveUser(UserRegistrationDTO userRegistrationDTO) {
        agentUserRepository.save(new AgentUser(userRegistrationDTO.getUsername(),
                userRegistrationDTO.getPassword(),
                userRegistrationDTO.getFirstName(),
                userRegistrationDTO.getLastName(),
                userRegistrationDTO.getDateOfBirth(),
                UserRole.Common));
    }

    @Override
    public void saveCompanyRegistrationRequest(CompanyRegistrationRequestDTO dto) {
        companyRegistrationRequestRepository.save(new CompanyRegistrationRequest(dto.getCompanyOwnerUsername(),
                dto.getCompanyOwnerName(),
                dto.getCompanyContactInfo(),
                dto.getCompanyDescription()));
    }

    @Override
    public void saveCompany(CompanyRegistrationRequestDTO dto) {
        AgentUser commonUser = agentUserRepository.findAgentUserByUsername(dto.getCompanyOwnerUsername());
        Company newCompany = new Company(dto.getCompanyContactInfo(), dto.getCompanyDescription());

        commonUser.setCompany(newCompany);
        commonUser.setRole(UserRole.CompanyOwner);

        companyRepository.save(newCompany);
        agentUserRepository.save(commonUser);
    }

    @Override
    public void editCompanyInfo(CompanyInfoDTO dto) {
        Company company = companyRepository.findById(dto.getId()).orElseGet(null);
        company.setContactInfo(dto.getContactInfo());
        company.setDescription(dto.getDescription());
        companyRepository.save(company);
    }

    @Override
    public void addOpenPosition(Long companyId, String positionName) {
        OpenPosition newOpenPosition = new OpenPosition(positionName);
        openPositionRepository.save(newOpenPosition);

        Company company = companyRepository.findById(companyId).orElseGet(null);
        company.getOpenPositions().add(newOpenPosition);

        companyRepository.save(company);
    }

    @Override
    public void saveComment(CommentDTO dto) {
        AgentUser user = agentUserRepository.findById(dto.getUserId()).orElseGet(null);
        if(user.getRole().equals(UserRole.CompanyOwner) || user.getRole().equals(UserRole.Admin))
            return;

        CommentOnCompany newComment = new CommentOnCompany(dto.getContent(), dto.getUserSignature());
        commentOnCompanyRepository.save(newComment);

        Company company = companyRepository.findById(dto.getCompanyId()).orElseGet(null);
        company.getComments().add(newComment);

        companyRepository.save(company);
    }
}
