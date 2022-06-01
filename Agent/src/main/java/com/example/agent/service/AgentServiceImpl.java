package com.example.agent.service;

import com.example.agent.domain.*;
import com.example.agent.dtos.*;
import com.example.agent.repository.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Set;

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

    @Autowired
    private SallaryRepository sallaryRepository;

    @Autowired
    private InterviewProcessRepository interviewProcessRepository;

    @Autowired
    private MarkRepository markRepository;

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
        if(userIsNotCommon(dto.getUserId()))
            return;

        CommentOnCompany newComment = new CommentOnCompany(dto.getContent(), dto.getUserSignature());
        commentOnCompanyRepository.save(newComment);

        Company company = companyRepository.findById(dto.getCompanyId()).orElseGet(null);
        company.getComments().add(newComment);

        companyRepository.save(company);
    }

    @Override
    public void addSallary(SallaryDTO dto) {
        if(userIsNotCommon(dto.getUserId()))
            return;

        Sallary newSallary = new Sallary(dto.getSallary());
        sallaryRepository.save(newSallary);

        OpenPosition openPosition = openPositionRepository.findById(dto.getPositionId()).orElseGet(null);
        openPosition.getSallarys().add(newSallary);

        openPositionRepository.save(openPosition);
    }

    @Override
    public void addInterviewProcess(InterviewProcessDTO dto) {
        if(userIsNotCommon(dto.getUserId()))
            return;

        InterviewProcess newInterviewProcess = new InterviewProcess(dto.getInterviewDescription(), dto.getUserSignature());
        interviewProcessRepository.save(newInterviewProcess);

        Company company = companyRepository.findById(dto.getCompanyId()).orElseGet(null);
        company.getInterviewProcesses().add(newInterviewProcess);

        companyRepository.save(company);
    }

    @Override
    public void addMark(MarkDTO dto) {
        if(userIsNotCommon(dto.getUserId()))
            return;
        Mark newMark = new Mark(dto.getMark());
        Company company = companyRepository.findById(dto.getCompanyId()).orElseGet(null);

        markRepository.save(newMark);

        company.getMarks().add(newMark);
        companyRepository.save(company);

    }

    @Override
    public List<Company> findAllCompanies() {
        return companyRepository.findAll();
    }

    @Override
    public Company findOneCompany(Long companyId) {
        return companyRepository.findById(companyId).orElseGet(null);
    }

    private boolean userIsNotCommon(Long userId){
        AgentUser user = agentUserRepository.findById(userId).orElseGet(null);
        return !user.getRole().equals(UserRole.Common);
    }
}
