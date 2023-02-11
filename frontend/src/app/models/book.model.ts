export class Book {
    private title: string;
    private author: string;
    private description: string;
    private cover: string;

    constructor(title: string, author: string, description: string, cover?: string) {
        this.title = title;
        this.author = author;
        this.description = description;
        this.cover = cover ? cover : 'https://via.placeholder.com/75';
    }

    public getTitle(): string {
        return this.title;
    }

    public getAuthor(): string {
        return this.author;
    }

    public getDescription(): string {
        return this.description;
    }

    public getCover(): string {
        return this.cover;
    }

    public setTitle(title: string): void {
        this.title = title;
    }

    public setAuthor(author: string): void {
        this.author = author;
    }

    public setDescription(description: string): void {
        this.description = description;
    }

    public setCover(cover: string): void {
        this.cover = cover;
    }
    
}

