from book_management import Book,BookManager
from customer_management import Customer,CustomerManager
from sales_management import Transaction,SalesManager

def main():
    # Create a new book manager
    book_manager=BookManager()
    sales_manager=SalesManager()
    customer_manager=CustomerManager()
    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ")

        if choice == "1":
            print("\n1. Add Book")
            print("2. View Books")
            print("3. Search Book")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                title = input("Title: ")
                author = input("Author: ")
                price = input("Price: ")
                quantity = input("Quantity: ")
                book_manager.add_book(title, author, price, quantity)
            elif sub_choice == "2":
                book_manager.view_books()
            elif sub_choice == "3":
                search_term = input("Enter title or author to search: ")
                book_manager.search_book(search_term)

        elif choice == "2":
            print("\n1. Add Customer")
            print("2. View Customers")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                name = input("Name: ")
                email = input("Email: ")
                phone = input("Phone: ")
                customer_manager.add_customer(name, email, phone)
            elif sub_choice == "2":
                customer_manager.view_customers()

        elif choice == "3":
            print("\n1. Sell Book")
            print("2. View Sales")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                phone = input("Enter Customer Phone: ")
                customer_result=customer_manager.search_customer(phone)
                if customer_result is None:
                    print("No customer found")
                    break
                print("searched customer:")
                customer_result.display_details()
                book_title = input("Enter Book Title: ")
                book_result=book_manager.search_book(book_title)
                if book_result is None:
                    print("No book found")
                    break
                print("searched book:")
                book_result.display_details()   
                quantity = input("Quantity: ")
                sales_manager.sell_book(customer_result,book_result,quantity)
            elif sub_choice == "2":
                sales_manager.view_sales()

        elif choice == "4":
            print("Thank you for using BookMart!")
            break

        else:
            print("Invalid choice. Please try again!")

if __name__ == "__main__":
    main()