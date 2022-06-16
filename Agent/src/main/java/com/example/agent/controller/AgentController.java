package com.example.agent.controller;

import com.example.agent.domain.*;
import com.example.agent.dtos.*;
import com.example.agent.service.AgentService;
import com.example.agent.validator.Validators;
import org.apache.logging.log4j.LogManager;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import javax.websocket.server.PathParam;
import java.util.List;
import java.util.Set;

import org.slf4j.Logger;

@Controller
@CrossOrigin(origins = "http://localhost:4201")
@RequestMapping("/api/agent")
public class AgentController {

    @Autowired
    private AgentService agentService;

    private static final Logger logSuccess = LoggerFactory.getLogger("success-log");
    private static final Logger logError = LoggerFactory.getLogger("error-log");

    @RequestMapping("/saveUser")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity saveUser(@RequestBody UserRegistrationDTO dto){
        if(Validators.isValidUserDto(dto)){
            agentService.saveUser(dto);
            logSuccess.info("User saved successfully");
            return new ResponseEntity(HttpStatus.CREATED);
        }
        logError.error("Invalid UserRegistrationDTO");
        return new ResponseEntity(HttpStatus.BAD_REQUEST);
    }

    @RequestMapping("/saveCompanyRegistrationRequest")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity saveCompanyRegistrationRequest(@RequestBody CompanyRegistrationRequestDTO dto){
        if(Validators.isValidCompanyRegistrationRequestDto(dto)){
            agentService.saveCompanyRegistrationRequest(dto);
            logSuccess.info("Company registration request saved successfully");
            return new ResponseEntity(HttpStatus.CREATED);
        }
        logError.error("Invalid CompanyRegistrationRequestDTO");
        return new ResponseEntity(HttpStatus.BAD_REQUEST);
    }

    @RequestMapping("/saveCompany")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity saveCompany(@RequestBody CompanyRegistrationRequestDTO dto){
        if(Validators.isValidCompanyRegistrationRequestDto(dto))
        {
            agentService.saveCompany(dto);
            logSuccess.info("Company saved successfully");
            return new ResponseEntity(HttpStatus.CREATED);
        }
        logError.error("Invalid CompanyRegistrationRequestDTO");
        return new ResponseEntity(HttpStatus.BAD_REQUEST);
    }

    @RequestMapping("/editCompanyInfo")
    @PutMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity editCompanyInfo(@RequestBody CompanyInfoDTO dto){
        if(Validators.isValidCompanyInfoDTO(dto))
        {
            agentService.editCompanyInfo(dto);
            logSuccess.info("Company info edited successfully");
            return new ResponseEntity(HttpStatus.OK);
        }
        logError.error("Invalid CompanyInfoDTO");
        return new ResponseEntity(HttpStatus.BAD_REQUEST);
    }

    @RequestMapping("/addOpenPosition/{companyId}/{positionName}/{description}/{criteria}")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity addOpenPosition(@PathVariable("companyId") Long companyId, @PathVariable("positionName") String positionName, @PathVariable("description") String description, @PathVariable("criteria") String criteria){
        if(Validators.isValidOpenPositionDto(companyId, positionName, description, criteria))
        {
            agentService.addOpenPosition(companyId, positionName, description, criteria);
            logSuccess.info("Open position added successfully");
            return new ResponseEntity(HttpStatus.OK);
        }
        logError.error("Invalid OpenPositionDTO");
        return new ResponseEntity(HttpStatus.BAD_REQUEST);
    }

    @RequestMapping("/saveComment")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity saveComment(@RequestBody CommentDTO dto){
        if(Validators.isValidCommentDto(dto))
        {
            agentService.saveComment(dto);
            logSuccess.info("Comment saved successfully");
            return new ResponseEntity<>(HttpStatus.CREATED);
        }
        logError.error("Invalid CommentDTO");
        return new ResponseEntity(HttpStatus.BAD_REQUEST);
    }

    @RequestMapping("/addSallary")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity addSallary(@RequestBody SalaryDTO dto){
        if(Validators.isValidSallaryDto(dto))
        {
            agentService.addSalary(dto);
            logSuccess.info("Sallary added successfully");
            return new ResponseEntity(HttpStatus.CREATED);
        }
        logError.error("Invalid SalaryDTO");
        return new ResponseEntity(HttpStatus.BAD_REQUEST);
    }

    @RequestMapping("/addInterviewProcess")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity addInterviewProcess(@RequestBody InterviewProcessDTO dto){
        if(Validators.isValidInterviewProcessDto(dto))
        {
            agentService.addInterviewProcess(dto);
            logSuccess.info("Interview process added successfully");
            return new ResponseEntity(HttpStatus.CREATED);
        }
        logError.error("Invalid InterviewProcessDTO");
        return new ResponseEntity(HttpStatus.BAD_REQUEST);
    }

    @RequestMapping("/addMark")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity addMark(@RequestBody MarkDTO dto){
        if(Validators.isValidMarkDto(dto))
        {
            agentService.addMark(dto);
            logSuccess.info("Mark added successfully");
            return new ResponseEntity(HttpStatus.CREATED);
        }
        logError.error("Invalid MarkDTO");
        return new ResponseEntity(HttpStatus.BAD_REQUEST);
    }

    @RequestMapping("/findAllCompanies")
    @GetMapping(produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<List<Company>> findAllCompanies(){
        List<Company> companies = agentService.findAllCompanies();
        return new ResponseEntity<>(companies, HttpStatus.OK);
    }

    @RequestMapping("/findAllCompanyRegistrationRequests")
    @GetMapping(produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<List<CompanyRegistrationRequestDTO>> findAllCompanyRegistrationRequests(){
        List<CompanyRegistrationRequestDTO> request = agentService.findAllCompanyRegistrationRequests();
        return new ResponseEntity<>(request, HttpStatus.OK);
    }

    @RequestMapping("/findOneCompanyById")
    @GetMapping(produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<Company> findOneCompanyById(@PathParam("companyId") Long companyId){
        Company company = agentService.findOneCompanyById(companyId);
        return new ResponseEntity<Company>(company, HttpStatus.OK);
    }

    @RequestMapping("/findUser")
    @GetMapping(produces = MediaType.APPLICATION_JSON_VALUE, consumes = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<AgentUser> findUser(@PathParam("username") String username, @PathParam("password") String password){
        AgentUser agentUser = agentService.findUser(username, password);
        return new ResponseEntity<>(agentUser, HttpStatus.OK);
    }

    @RequestMapping("/findAllCommentsByCompanyId/{companyId}")
    @GetMapping(produces = MediaType.APPLICATION_JSON_VALUE, consumes = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<Set<CommentOnCompany>> findAllCommentsByCompanyId(@PathVariable("companyId") Long companyId){
        Set<CommentOnCompany> comments = agentService.findAllCommentsByCompanyId(companyId);
        return new ResponseEntity<>(comments, HttpStatus.OK);
    }

    @RequestMapping("/findAllInterviewsByCompanyId/{companyId}")
    @GetMapping(produces = MediaType.APPLICATION_JSON_VALUE, consumes = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<Set<InterviewProcess>> findAllInterviewsByCompanyId(@PathVariable("companyId") Long companyId){
        Set<InterviewProcess> interviews = agentService.findAllInterviewsByCompanyId(companyId);
        return new ResponseEntity<>(interviews, HttpStatus.OK);
    }

    @RequestMapping("/findAllPositionsByCompanyId/{companyId}")
    @GetMapping(produces = MediaType.APPLICATION_JSON_VALUE, consumes = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<Set<OpenPosition>> findAllPositionsByCompanyId(@PathVariable("companyId") Long companyId){
        Set<OpenPosition> positions = agentService.findAllPositionsByCompanyId(companyId);
        return new ResponseEntity<>(positions, HttpStatus.OK);
    }

    @RequestMapping("/saveToken/{userId}/{token}")
    @PutMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<AgentUser> saveToken(@PathVariable("userId") Long userId, @PathVariable("token") String token){
        if(Validators.isValidToken(userId, token))
        {
            AgentUser agentUser = agentService.saveToken(userId, token);
            logSuccess.info("Token saved successfully");
            return new ResponseEntity<>(agentUser, HttpStatus.OK);
        }
        logError.error("Invalid token");
        return new ResponseEntity(HttpStatus.BAD_REQUEST);
    }
}
