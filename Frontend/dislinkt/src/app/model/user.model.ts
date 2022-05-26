export class User {
    username: string;
    password: string;
    firstName: string;
    lastName: string;
    dateOfBirth: string;
    email: string;
    skills: any;
    interests: any;
    education: any;
    workExperience: any;

    constructor(){
        this.username = "";
        this.password = ""; 
        this.firstName = "";
        this.lastName = "";
        this.dateOfBirth = "";
        this.email = "";
        this.skills = [];
        this.interests = [];
        this.education = [];
        this.workExperience = [];
    }
}
