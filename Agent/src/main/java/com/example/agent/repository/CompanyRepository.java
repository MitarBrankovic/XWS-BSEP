package com.example.agent.repository;

import com.example.agent.domain.CommentOnCompany;
import com.example.agent.domain.Company;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

public interface CompanyRepository extends JpaRepository<Company, Long> {

}
