# SOLID Design Principles — Examples

SOLID is five object-oriented design principles that help keep code maintainable, flexible, and easier to test.

---

## S — Single Responsibility Principle (SRP)

**A class should have only one reason to change.**

**Bad:** One class does too much.

```python
class User:
    def save_to_database(self): ...
    def send_welcome_email(self): ...
    def generate_report(self): ...
```

**Good:** Each class has one job.

```python
class UserRepository:
    def save(self, user): ...

class EmailService:
    def send_welcome(self, user): ...

class UserReportGenerator:
    def generate(self, user): ...
```

If email templates change, you only touch `EmailService` — not database or reporting logic.

---

## O — Open/Closed Principle (OCP)

**Open for extension, closed for modification.**

**Bad:** Adding a new payment type means editing existing code.

```python
def process_payment(payment_type, amount):
    if payment_type == "credit_card":
        ...
    elif payment_type == "paypal":
        ...
    elif payment_type == "crypto":  # new type = modify this function
        ...
```

**Good:** Extend via new classes, not by changing core logic.

```python
class PaymentProcessor(ABC):
    @abstractmethod
    def pay(self, amount): ...

class CreditCardPayment(PaymentProcessor):
    def pay(self, amount): ...

class PayPalPayment(PaymentProcessor):
    def pay(self, amount): ...

class CryptoPayment(PaymentProcessor):  # new type = new class only
    def pay(self, amount): ...
```

---

## L — Liskov Substitution Principle (LSP)

**Subtypes must be substitutable for their base types without breaking behavior.**

**Bad:** A subclass violates the parent's contract.

```python
class Bird:
    def fly(self): ...

class Penguin(Bird):
    def fly(self):
        raise Exception("Penguins can't fly!")  # breaks callers expecting Bird.fly()
```

**Good:** Subclasses honor the same expectations.

```python
class Bird: ...

class FlyingBird(Bird):
    def fly(self): ...

class Penguin(Bird):
    def swim(self): ...  # Penguin is a Bird, but not a FlyingBird
```

Callers that need flying birds use `FlyingBird`, not a broken `Bird` subtype.

---

## I — Interface Segregation Principle (ISP)

**Clients should not depend on interfaces they don't use.**

**Bad:** One fat interface forces unused methods.

```python
class Worker(ABC):
    @abstractmethod
    def work(self): ...
    @abstractmethod
    def eat(self): ...

class Robot(Worker):
    def work(self): ...
    def eat(self):
        pass  # Robots don't eat — forced to implement anyway
```

**Good:** Smaller, focused interfaces.

```python
class Workable(ABC):
    @abstractmethod
    def work(self): ...

class Eatable(ABC):
    @abstractmethod
    def eat(self): ...

class Human(Workable, Eatable): ...
class Robot(Workable): ...
```

---

## D — Dependency Inversion Principle (DIP)

**Depend on abstractions, not concrete implementations.**

**Bad:** High-level code is tightly coupled to low-level details.

```python
class OrderService:
    def __init__(self):
        self.db = MySQLDatabase()  # hard-coded dependency

    def place_order(self, order):
        self.db.save(order)
```

**Good:** Inject an abstraction so you can swap implementations.

```python
class Database(ABC):
    @abstractmethod
    def save(self, data): ...

class MySQLDatabase(Database): ...
class PostgresDatabase(Database): ...

class OrderService:
    def __init__(self, db: Database):  # depends on abstraction
        self.db = db

    def place_order(self, order):
        self.db.save(order)
```

Now you can use MySQL, Postgres, or a mock in tests without changing `OrderService`.

---

## Quick Reference

| Principle | One-liner |
|-----------|-----------|
| **S** | One class → one job |
| **O** | Extend with new code, don't rewrite old code |
| **L** | Subclasses must behave like their parent |
| **I** | Many small interfaces > one big interface |
| **D** | Code against interfaces, inject implementations |
