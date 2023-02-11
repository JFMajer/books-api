import { Component } from '@angular/core';
import { Book } from '../../models/book.model';

@Component({
  selector: 'app-list-books',
  templateUrl: './list-books.component.html',
  styleUrls: ['./list-books.component.css']
})
export class ListBooksComponent {
  public books: Book[];

  constructor() {
    this.books = [
      new Book('The Hobbit', 'J.R.R. Tolkien', 'The Hobbit is a children\'s fantasy novel by English author J. R. R. Tolkien. It was published on 21 September 1937 to wide critical acclaim, being nominated for the Carnegie Medal and awarded a prize from the New York Herald Tribune for best juvenile fiction. The book remains popular and is recognized as a classic in children\'s literature.'),
      new Book('The Lord of the Rings', 'J.R.R. Tolkien', 'The Lord of the Rings is an epic high fantasy novel written by English author and scholar J. R. R. Tolkien. The story began as a sequel to Tolkien\'s 1937 fantasy novel The Hobbit, but eventually developed into a much larger work. Written in stages between 1937 and 1949, The Lord of the Rings is one of the best-selling novels ever written, with over 150 million copies sold.'),
      new Book('The Silmarillion', 'J.R.R. Tolkien', 'The Silmarillion is an unfinished mythopoeic history of Middle-earth, written by J. R. R. Tolkien. It was published posthumously in 1977, with revisions and additions by Christopher Tolkien, and is one of the few works of Tolkien\'s to be published during his lifetime. The Silmarillion is set in the First Age of Tolkien\'s legendarium, and tells of the Elder Days, the First Age, and the events before the War of the Ring.'),
      new Book('The Chronicles of Narnia', 'C.S. Lewis', 'The Chronicles of Narnia is a series of seven fantasy novels by C. S. Lewis. It is considered a classic of children\'s literature and is the author\'s best-known work, having sold over 100 million copies in 47 languages. The series was written by Lewis between 1949 and 1954, during the years when he was a professor of medieval and renaissance literature at Magdalen College, Oxford, and a fellow and tutor of Magdalen College. The first published book in the series, The Lion, the Witch and the Wardrobe, has been adapted several times for radio, television, the stage, and film.')
    ];
  }

  // print all books
  public printBooks(): void {
    console.log(this.books);
  }


}
