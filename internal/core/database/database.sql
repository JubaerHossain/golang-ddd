-- Customers table
CREATE TABLE Customers (
    customer_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100),
    phone VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index on email for quick lookup
CREATE UNIQUE INDEX idx_customers_email ON Customers(email);

-- Tables table
CREATE TABLE Tables (
    table_id SERIAL PRIMARY KEY,
    seating_capacity INT NOT NULL,
    status VARCHAR(20) DEFAULT 'available' CHECK (status IN ('available', 'occupied', 'reserved')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Orders table
CREATE TABLE Orders (
    order_id SERIAL PRIMARY KEY,
    table_id INT REFERENCES Tables(table_id),
    customer_id INT REFERENCES Customers(customer_id),
    order_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) DEFAULT 'open' CHECK (status IN ('open', 'closed')),
    total_amount DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index on order_time for ordering
CREATE INDEX idx_orders_order_time ON Orders(order_time);

-- Menu_Items table
CREATE TABLE Menu_Items (
    item_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    category_id INT REFERENCES Categories(category_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Categories table
CREATE TABLE Categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index on category_id for menu item category lookup
CREATE INDEX idx_menu_items_category_id ON Menu_Items(category_id);

-- Order_Items table
CREATE TABLE Order_Items (
    order_item_id SERIAL PRIMARY KEY,
    order_id INT REFERENCES Orders(order_id),
    item_id INT REFERENCES Menu_Items(item_id),
    quantity INT NOT NULL,
    modifiers TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Ingredients table
CREATE TABLE Ingredients (
    ingredient_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    stock_quantity INT NOT NULL,
    unit_of_measure VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Transactions table
CREATE TABLE Transactions (
    transaction_id SERIAL PRIMARY KEY,
    order_id INT REFERENCES Orders(order_id),
    amount DECIMAL(10, 2) NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    transaction_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Employees table
CREATE TABLE Employees (
    employee_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    position VARCHAR(50),
    hourly_rate DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index on position for employee lookup
CREATE INDEX idx_employees_position ON Employees(position);

-- Shifts table
CREATE TABLE Shifts (
    shift_id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES Employees(employee_id),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    hours_worked DECIMAL(5, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index on start_time for shift lookup
CREATE INDEX idx_shifts_start_time ON Shifts(start_time);

-- Payroll table
CREATE TABLE Payroll (
    payroll_id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES Employees(employee_id),
    pay_period VARCHAR(20),
    total_hours DECIMAL(5, 2),
    total_pay DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Reservations table
CREATE TABLE Reservations (
    reservation_id SERIAL PRIMARY KEY,
    table_id INT REFERENCES Tables(table_id),
    customer_id INT REFERENCES Customers(customer_id),
    reservation_time TIMESTAMP,
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'cancelled')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index on reservation_time for reservation lookup
CREATE INDEX idx_reservations_reservation_time ON Reservations(reservation_time);
