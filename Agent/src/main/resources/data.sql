insert into sallary (id, sallary_value) values (1, '500')

insert into open_position (id, position_name) values (1, 'React developer')

insert into open_position_sallarys (open_position_id, sallarys_id) values (1, 1)

insert into interview_process (id, description) values (1, 'Interview process')

insert into comment_on_company (id, comment, user_signature) values (1, 'Comment on company', 'Software developer (Medior)')

insert into company (id, contact_info, description) values (1, 'Company contact info', 'Company description')

insert into company_comments (company_id, comments_id) values (1, 1)

insert into company_interview_processes (company_id, interview_processes_id) values (1, 1)

insert into company_open_positions (company_id, open_positions_id) values (1, 1)

insert into agent_user (id, first_name, last_name, username, password, date_of_birth, role, company_id) values (1, 'Marko', 'Markovic', 'markoagent', 'markoagent', '2000-01-01', 0, null )
insert into agent_user (id, first_name, last_name, username, password, date_of_birth, role, company_id) values (2, 'Petar', 'Petrovic', 'petaragent', 'petaragent', '2001-01-01', 1, null )
insert into agent_user (id, first_name, last_name, username, password, date_of_birth, role, company_id) values (3, 'Stevan', 'Stevanovic', 'stevanagent', 'stevanagent', '2002-01-01', 2, 1)


insert into company_registration_request (id, company_owner, company_contact_info, company_description) values (1, 'Mirko Vojinovic', 'GracanITa, Obrovac 21423, Vojvodjanska 28', 'IT company')