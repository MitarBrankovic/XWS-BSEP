export class Company {
    id: number;
    name: string;
    contactInfo: string;
    description: string;
    username: string;
    interviewProcesses: [];
    comments: [];
    openPositions = [];
    marks = [];

    constructor() {
        this.id = 0;
        this.name = '';
        this.contactInfo = "";
        this.description = "";
        this.username = "";
        this.interviewProcesses = [];
        this.comments = [];
        this.openPositions = [];
        this.marks = []
    }
}