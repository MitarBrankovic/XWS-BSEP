package com.example.agent.controller;

import com.example.agent.domain.*;
import com.example.agent.dtos.*;
import com.example.agent.service.AgentService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import javax.websocket.server.PathParam;
import java.util.List;
import java.util.Set;

@Controller
@CrossOrigin(origins = "http://localhost:4201")
@RequestMapping("/api/agent")
public class AgentController {

    @Autowired
    private AgentService agentService;

    @RequestMapping("/saveUser")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<UserRegistrationDTO> saveUser(@RequestBody UserRegistrationDTO dto){
        agentService.saveUser(dto);
        return new ResponseEntity<>(dto, HttpStatus.CREATED);
    }

    @RequestMapping("/saveCompanyRegistrationRequest")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<CompanyRegistrationRequestDTO> saveCompanyRegistrationRequest(@RequestBody CompanyRegistrationRequestDTO dto){
        agentService.saveCompanyRegistrationRequest(dto);
        return new ResponseEntity<>(dto, HttpStatus.CREATED);
    }

    @RequestMapping("/saveCompany")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<CompanyRegistrationRequestDTO> saveCompany(@RequestBody CompanyRegistrationRequestDTO dto){
        agentService.saveCompany(dto);
        return new ResponseEntity<>(dto, HttpStatus.CREATED);
    }

    @RequestMapping("/editCompanyInfo")
    @PutMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity editCompanyInfo(@RequestBody CompanyInfoDTO dto){
        agentService.editCompanyInfo(dto);
        return new ResponseEntity(HttpStatus.ACCEPTED);
    }

    @RequestMapping("/addOpenPosition/{companyId}/{positionName}/{description}/{criteria}")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity addOpenPosition(@PathVariable("companyId") Long companyId, @PathVariable("positionName") String positionName, @PathVariable("description") String description, @PathVariable("criteria") String criteria){
        agentService.addOpenPosition(companyId, positionName, description, criteria);
        return new ResponseEntity(HttpStatus.CREATED);
    }

    @RequestMapping("/saveComment")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<CommentOnCompany> saveComment(@RequestBody CommentDTO dto){
        agentService.saveComment(dto);
        return new ResponseEntity<>(HttpStatus.CREATED);
    }

    @RequestMapping("/addSallary")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity addSallary(@RequestBody SalaryDTO dto){
        agentService.addSalary(dto);
        return new ResponseEntity(HttpStatus.CREATED);
    }

    @RequestMapping("/addInterviewProcess")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity addInterviewProcess(@RequestBody InterviewProcessDTO dto){
        agentService.addInterviewProcess(dto);
        return new ResponseEntity(HttpStatus.CREATED);
    }

    @RequestMapping("/addMark")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity addMark(@RequestBody MarkDTO dto){
        agentService.addMark(dto);
        return new ResponseEntity(HttpStatus.CREATED);
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
        return new ResponseEntity<>(agentService.saveToken(userId, token) ,HttpStatus.CREATED);
    }
}
