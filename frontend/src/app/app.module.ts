import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { BookItemComponent } from './components/book-item/book-item.component';
import { ListBooksComponent } from './components/list-books/list-books.component';

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    BookItemComponent,
    ListBooksComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
