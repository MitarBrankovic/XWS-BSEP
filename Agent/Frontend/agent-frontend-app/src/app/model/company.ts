export class Company {
    id: number;
    contactInfo: string;
    description: string;
    interviewProcesses: [];
    comments: [];
    openPositions = [];

    constructor() {
        this.id = 0;
        this.contactInfo = "";
        this.description = "";
        this.interviewProcesses = [];
        this.comments = [];
        this.openPositions = [];
    }
}