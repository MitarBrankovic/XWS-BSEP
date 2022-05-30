package com.example.agent.service;


import com.example.agent.domain.TestClass;
import com.example.agent.repository.TestRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class TestServiceImpl implements TestService{

    @Autowired
    private TestRepository testRepository;

    @Override
    public List<TestClass> findAll() {
        return testRepository.findAll();
    }
}
