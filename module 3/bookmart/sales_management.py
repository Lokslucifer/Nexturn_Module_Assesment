from customer_management import Customer

class Transaction(Customer):
    def __init__(self, name, email, phone, book_title, quantity_sold):
        super().__init__(name, email, phone)
        self.book_title = book_title
        self.quantity_sold = quantity_sold

sales_records = []
class SalesManager:
    def __init__(self,sales_records=[]):
        self.sales_records = sales_records

    def sell_book(self,customer,book,quantity):
        try:
            quantity = int(quantity)
            if book.quantity>=quantity:
                book.quantity -= quantity
                transaction = Transaction(customer.name,customer.email,customer.phone, book.title, quantity)
                self.sales_records.append(transaction)
                print(f"Sale successful! Remaining quantity: {book.quantity}")
            else:
                print("no stock")
        except ValueError:
            print("Invalid quantity!")

def view_sales(self):
    if len(self.sales_records):
        print("No sales records found.")
    else:
        for sale in self.sales_records:
            print(f"Customer: {sale.name}, Book: {sale.book_title}, Quantity Sold: {sale.quantity_sold}")
