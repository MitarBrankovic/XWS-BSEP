import { Content } from "./content";
import { PostUser } from "./post-user";

export class Post{
    createdAt: Date;
    content: Content;
    user: PostUser;
    reactions: any;
    comments: any;

    constructor() {
        this.content = new Content();
        this.user = new PostUser();
        this.createdAt = new Date();
        this.reactions = [];
        this.comments = [];
    }  
}
