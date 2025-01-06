class Customer:
    def __init__(self, name, email, phone):
        self.name = name
        self.email = email
        self.phone = phone

    def display_details(self):
        return f"Name: {self.name}, Email: {self.email}, Phone: {self.phone}"


class CustomerManager:
    def __init__(self,customerlst=[]):
        self.customer_list = customerlst

    def add_customer(self,name, email, phone):
        customer = Customer(name, email, phone)
        self.customer_list.append(customer)
        print("Customer added successfully!")

    def view_customers(self):
        if len(self.customer_list):
            print("No customers available.")
        else:
            for customer in self.customer_list:
                print(customer.display_details())

    
    def search_customer(self,search_term):
        for customer in self.customer_list:
            if search_term==customer.phone:
                return customer
        return None
     
    
    