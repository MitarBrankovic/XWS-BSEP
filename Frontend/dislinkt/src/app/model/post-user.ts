import { StringLiteralLike } from "typescript";

export class PostUser {
    username: string;
    firstName: string;
    lastName: string;

    constructor(){
        this.username = '';
        this.firstName = '';
        this.lastName = '';
    }
}
