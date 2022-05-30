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
    private: boolean;
    activated: boolean;

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
        this.private = false;
        this.activated = false;
    }
}
