class Book:
    def __init__(self, title, author, price, quantity):
        self.title = title
        self.author = author
        self.price = price
        self.quantity = quantity

    def display_details(self):
        return f"Title: {self.title}, Author: {self.author}, Price: {self.price}, Quantity: {self.quantity}"

class BookManager:
    def __init__(self, book_inventory=[]):
        self.book_inventory = book_inventory


    def add_book(self,title, author, price, quantity):
        try:
            price = float(price)
            quantity = int(quantity)
            if price < 0 or quantity < 0:
                raise ValueError("Price and quantity must be positive numbers.")
            book = Book(title, author, price, quantity)
            self.book_inventory.append(book)
            print("Book added successfully!")
        except ValueError as e:
            print(f"Error: {e}")

    def view_books(self):
        if len(self.book_inventory)==0:
            print("No books available.")
        else:
            for book in self.book_inventory:
                print(book.display_details())

    def search_book(self,search_term):
        found_books = [book for book in self.book_inventory if search_term.lower() in book.title.lower() or search_term.lower() in book.author.lower()]
        if len(found_books)==0:
            print("No matching books found.")
        else:
            for book in found_books:
                print(book.display_details())

    def search_specific_book(self,search_term):
        found_books = [book for book in self.book_inventory if search_term.lower() in book.title.lower() or search_term.lower() in book.author.lower()]
        if len(found_books)==0:
            print("No matching books found.")
        else:
            return found_books[0]
 
