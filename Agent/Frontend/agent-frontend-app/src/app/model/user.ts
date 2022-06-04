export class User {
    username: string;
    password: string;
    firstName: string;
    lastName: string;
    dateOfBirth: string;
    role: string;

    constructor(){
        this.username = "";
        this.password = ""; 
        this.firstName = "";
        this.lastName = "";
        this.dateOfBirth = "";
        this.role = "";
    }
}