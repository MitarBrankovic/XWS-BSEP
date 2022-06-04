package com.example.agent.repository;

import com.example.agent.domain.CommentOnCompany;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.List;

public interface CommentOnCompanyRepository extends JpaRepository<CommentOnCompany, Long> {

}
