export class LoggedUser {
    username: string;
    role: string;
    exp: Date;

    constructor() {
        this.username = '';
        this.role = 'guest';
        this.exp = new Date();
    }
}
