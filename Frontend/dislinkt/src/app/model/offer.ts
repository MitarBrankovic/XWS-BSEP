export class Offer {
    
	username : string;   
	company : string;
	description :string;
	position : string;
	criteria : string;
	createdAt : Date;

    constructor(username: string, company: string, description: string, position: string, criteria: string, createdAt: Date) {
        this.username = username;
        this.company = company;
        this.description = description;
        this.position = position;
        this.criteria = criteria;
        this.createdAt = createdAt;
    }
}