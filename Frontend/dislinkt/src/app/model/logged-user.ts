export class LoggedUser {
    username: string;
    role: string;
    exp: number;

    constructor() {
        this.username = '';
        this.role = 'guest';
        this.exp = 0;
    }
}
