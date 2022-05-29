export class Connection {

    issuerUsername: string;
    subjectUsername: string;

    constructor(issuerUsername: string, subjectUsername: string) {
        this.issuerUsername = issuerUsername;
        this.subjectUsername = subjectUsername;
    }
}
