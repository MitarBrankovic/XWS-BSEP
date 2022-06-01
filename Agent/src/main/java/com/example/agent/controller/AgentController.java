package com.example.agent.controller;

import com.example.agent.domain.InterviewProcess;
import com.example.agent.dtos.*;
import com.example.agent.service.AgentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import javax.websocket.server.PathParam;

@Controller
@CrossOrigin(origins = "http://localhost:4200")
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

    @RequestMapping("/addOpenPosition")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity addOpenPosition(@PathParam("companyId") Long companyId, @PathParam("positionName") String positionName){
        agentService.addOpenPosition(companyId, positionName);
        return new ResponseEntity(HttpStatus.CREATED);
    }

    @RequestMapping("/saveComment")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity saveComment(@RequestBody CommentDTO dto){
        agentService.saveComment(dto);
        return new ResponseEntity(HttpStatus.CREATED);
    }

    @RequestMapping("/addSallary")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity addSallary(@RequestBody SallaryDTO dto){
        agentService.addSallary(dto);
        return new ResponseEntity(HttpStatus.CREATED);
    }

    @RequestMapping("/addInterviewProcess")
    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity addInterviewProcess(@RequestBody InterviewProcessDTO dto){
        agentService.addInterviewProcess(dto);
        return new ResponseEntity(HttpStatus.CREATED);
    }
}
