package com.example.agent.service;

import com.example.agent.controller.AgentController;
import com.example.agent.domain.*;
import com.example.agent.dtos.*;
import com.example.agent.repository.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.stereotype.Service;

import javax.transaction.Transactional;
import java.util.ArrayList;
import java.util.List;
import java.util.Set;

@Service
public class AgentServiceImpl implements AgentService, UserDetailsService {

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

    @Autowired
    private RoleRepository roleRepository;

    Logger logger = LoggerFactory.getLogger(AgentController.class);

    @Override
    public void saveUser(UserRegistrationDTO userRegistrationDTO) {
        agentUserRepository.save(new AgentUser(userRegistrationDTO.getUsername(),
                userRegistrationDTO.getPassword(),
                userRegistrationDTO.getFirstName(),
                userRegistrationDTO.getLastName(),
                userRegistrationDTO.getDateOfBirth(),
                roleRepository.findById(1L).orElseThrow()));

        logger.info("User registrated");
    }

    @Override
    public void saveCompanyRegistrationRequest(CompanyRegistrationRequestDTO dto) {
        companyRegistrationRequestRepository.save(new CompanyRegistrationRequest(dto.getCompanyOwnerUsername(),
                dto.getCompanyOwnerName(),
                dto.getCompanyName(),
                dto.getCompanyContactInfo(),
                dto.getCompanyDescription(), dto.getUsername()));

        logger.info("Company registration request saved");
    }

    @Override
    @Transactional
    public void saveCompany(CompanyRegistrationRequestDTO dto) {
        AgentUser commonUser = agentUserRepository.findAgentUserByUsername(dto.getCompanyOwnerUsername());
        Company newCompany = new Company(dto.getCompanyName(), dto.getCompanyContactInfo(), dto.getCompanyDescription(), dto.getUsername());

        commonUser.setCompany(newCompany);
        commonUser.setRole(roleRepository.findById(2L).orElseThrow());

        companyRepository.save(newCompany);
        agentUserRepository.save(commonUser);

        companyRegistrationRequestRepository.removeAllByCompanyOwnerUsername(dto.getCompanyOwnerUsername());

        logger.info("Company registrated");
    }

    @Override
    public void editCompanyInfo(CompanyInfoDTO dto) {
        Company company = companyRepository.findById(dto.getId()).orElseGet(null);
        company.setContactInfo(dto.getContactInfo());
        company.setDescription(dto.getDescription());
        companyRepository.save(company);

        logger.info("Company (id: " + company.getId() + " )  edited");
    }

    @Override
    public void addOpenPosition(Long companyId, String positionName, String description, String criteria) {
        OpenPosition newOpenPosition = new OpenPosition(positionName, description, criteria);
        openPositionRepository.save(newOpenPosition);

        Company company = companyRepository.findById(companyId).orElseGet(null);
        company.getOpenPositions().add(newOpenPosition);

        companyRepository.save(company);

        logger.info("Open position in company: " + company.getName() + " created");
    }

    @Override
    public void saveComment(CommentDTO dto) {
        if(userIsNotCommon(dto.getUserId()))
            return;

        CommentOnCompany newComment = new CommentOnCompany(dto.getContent(), dto.getUserSignature(), dto.getUsername());
        commentOnCompanyRepository.save(newComment);

        Company company = companyRepository.findById(dto.getCompanyId()).orElseGet(null);
        company.getComments().add(newComment);

        companyRepository.save(company);

        logger.info("Comment in company: " + company.getName() + " saved");
    }

    @Override
    public void addSalary(SalaryDTO dto) {
        if(userIsNotCommon(dto.getUserId()))
            return;

        Salary newSalary = new Salary(dto.getSalary(), dto.getUserId());
        sallaryRepository.save(newSalary);

        OpenPosition openPosition = openPositionRepository.findById(dto.getPositionId()).orElseGet(null);
        openPosition.getSalaries().add(newSalary);

        openPositionRepository.save(openPosition);

        logger.info("Sallary on open position: " + openPosition.getPositionName() + " added");
    }

    @Override
    public void addInterviewProcess(InterviewProcessDTO dto) {
        if(userIsNotCommon(dto.getUserId()))
            return;

        InterviewProcess newInterviewProcess = new InterviewProcess(dto.getInterviewDescription(), dto.getUserSignature(), dto.getUsername());
        interviewProcessRepository.save(newInterviewProcess);

        Company company = companyRepository.findById(dto.getCompanyId()).orElseGet(null);
        company.getInterviewProcesses().add(newInterviewProcess);

        companyRepository.save(company);

        logger.info("Interview process in company: " + company.getName() + " added");
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

        logger.info("Mark on company: " + company.getName() + " added");
    }

    @Override
    public List<Company> findAllCompanies() {
        return companyRepository.findAll();
    }

    @Override
    public Company findOneCompanyById(Long companyId) {
        return companyRepository.findById(companyId).orElseGet(null);
    }

    @Override
    public AgentUser findUser(String username, String password) {
        return agentUserRepository.findAgentUserByUsernameAndPassword(username, password);
    }

    @Override
    public List<CompanyRegistrationRequestDTO> findAllCompanyRegistrationRequests() {
        List<CompanyRegistrationRequestDTO> requests = new ArrayList<>();
        for(CompanyRegistrationRequest request : companyRegistrationRequestRepository.findAll()) {
            requests.add(new CompanyRegistrationRequestDTO(request.getCompanyName(), request.getCompanyOwnerUsername(),
                    request.getCompanyOwnerName(),
                    request.getCompanyContactInfo(),
                    request.getCompanyDescription(),
                    request.getUsername()));
        }
        return requests;
    }

    private boolean userIsNotCommon(Long userId){
        AgentUser user = agentUserRepository.findById(userId).orElseGet(null);
        return !user.getRole().getName().equals("COMMON");
    }

    public Set<CommentOnCompany> findAllCommentsByCompanyId(Long companyId){
        Company company = companyRepository.findById(companyId).orElseGet(null);
        return company.getComments();
    }

    public Set<InterviewProcess> findAllInterviewsByCompanyId(Long companyId){
        Company company = companyRepository.findById(companyId).orElseGet(null);
        return company.getInterviewProcesses();
    }

    public Set<OpenPosition> findAllPositionsByCompanyId(Long companyId){
        Company company = companyRepository.findById(companyId).orElseGet(null);
        return company.getOpenPositions();
    }

    public AgentUser saveToken(Long userId, String token){
        AgentUser agentUser = agentUserRepository.findById(userId).orElseGet(null);
        agentUser.setApiToken(token);
        agentUserRepository.save(agentUser);

        logger.info("ApiToken saved");

        return agentUser;
    }

    public AgentUser loadUserByUsername(String username) {
        return agentUserRepository.findAgentUserByUsername(username);
    }
}
