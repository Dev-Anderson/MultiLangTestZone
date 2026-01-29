from game import GuessGame


def main():
    print("🎯 Jogo da Adivinhação")

    game = GuessGame()

    while game.has_attempts():
        try:
            guess = int(input("Digite um número entre 1 e 100: "))
        except ValueError:
            print("Digite apenas números!")
            continue

        result = game.guess(guess)

        if result == "Correto":
            print("🎉 Parabéns! Você acertou!")
            break
        else:
            print(f"Seu palpite é {result}")
            print(f"Tentativas restantes: {game.attempts_left}")

    else:
        print("😢 Suas tentativas acabaram!")
        print(f"O número era: {game.secret_number}")

    play_again = input("Jogar novamente? (s/n): ").lower()
    if play_again == "s":
        main()


if __name__ == "__main__":
    main()
