import random


class GuessGame:
    def __init__(self, min_number=1, max_number=100, max_attempts=10):
        self.min_number = min_number
        self.max_number = max_number
        self.max_attempts = max_attempts
        self.secret_number = random.randint(min_number, max_number)
        self.attempts_left = max_attempts

    def guess(self, number: int) -> str:
        self.attempts_left -= 1

        if number < self.secret_number:
            return "Menor"
        elif number > self.secret_number:
            return "Maior"
        else:
            return "Correto"

    def has_attempts(self) -> bool:
        return self.attempts_left > 0
