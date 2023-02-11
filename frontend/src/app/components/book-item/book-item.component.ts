import { Component } from '@angular/core';
import { Book } from '../../models/book.model';

@Component({
  selector: 'app-book-item',
  templateUrl: './book-item.component.html',
  styleUrls: ['./book-item.component.css']
})
export class BookItemComponent {

  public book: Book;

  constructor() { }

}
