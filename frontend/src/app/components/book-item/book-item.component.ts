import { Component } from '@angular/core';
import { Book } from '../../models/book.model';
import { Input } from '@angular/core';

@Component({
  selector: 'app-book-item',
  templateUrl: './book-item.component.html',
  styleUrls: ['./book-item.component.css']
})
export class BookItemComponent {
  @Input()
  public book: Book;

  constructor() { }

}
