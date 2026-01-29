import pygame
import random
import sys

# -----------------------
# Configurações
# -----------------------
WIDTH = 400
HEIGHT = 600
FPS = 60

PLAYER_WIDTH = 50
PLAYER_HEIGHT = 80

ENEMY_WIDTH = 50
ENEMY_HEIGHT = 80

INITIAL_SPEED = 5
SPEED_INCREMENT = 0.5

# -----------------------
# Inicialização
# -----------------------
pygame.init()
screen = pygame.display.set_mode((WIDTH, HEIGHT))
pygame.display.set_caption("Jogo do Carrinho")

clock = pygame.time.Clock()
font = pygame.font.SysFont(None, 36)

# -----------------------
# Jogador
# -----------------------
player_x = WIDTH // 2 - PLAYER_WIDTH // 2
player_y = HEIGHT - PLAYER_HEIGHT - 20
player_speed = 6

# -----------------------
# Inimigos
# -----------------------
enemies = []

def create_enemy():
    x = random.randint(50, WIDTH - ENEMY_WIDTH - 50)
    y = -ENEMY_HEIGHT
    return pygame.Rect(x, y, ENEMY_WIDTH, ENEMY_HEIGHT)

for _ in range(3):
    enemies.append(create_enemy())

# -----------------------
# Estado do jogo
# -----------------------
speed = INITIAL_SPEED
score = 0
game_over = False
running = True

# -----------------------
# Loop principal
# -----------------------
while running:
    clock.tick(FPS)

    # Eventos
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

    keys = pygame.key.get_pressed()

    # Movimento do jogador
    if not game_over:
        if keys[pygame.K_LEFT] and player_x > 50:
            player_x -= player_speed
        if keys[pygame.K_RIGHT] and player_x < WIDTH - PLAYER_WIDTH - 50:
            player_x += player_speed

    player_rect = pygame.Rect(player_x, player_y, PLAYER_WIDTH, PLAYER_HEIGHT)

    # Movimento inimigos
    if not game_over:
        for enemy in enemies:
            enemy.y += speed

            if enemy.y > HEIGHT:
                enemy.y = -ENEMY_HEIGHT
                enemy.x = random.randint(50, WIDTH - ENEMY_WIDTH - 50)
                score += 1

                if score % 5 == 0:
                    speed += SPEED_INCREMENT

            if enemy.colliderect(player_rect):
                game_over = True

    # -----------------------
    # DESENHO DA TELA
    # -----------------------
    screen.fill((30, 30, 30))

    # Estrada
    pygame.draw.rect(screen, (80, 80, 80), (50, 0, 300, HEIGHT))

    # Faixas da pista
    for y in range(0, HEIGHT, 40):
        pygame.draw.rect(screen, (255, 255, 255), (195, y, 10, 20))

    # Jogador
    pygame.draw.rect(screen, (0, 200, 255), player_rect)

    # Inimigos
    for enemy in enemies:
        pygame.draw.rect(screen, (255, 0, 0), enemy)

    # HUD
    score_text = font.render(f"Score: {score}", True, (255, 255, 255))
    speed_text = font.render(f"Speed: {round(speed,1)}", True, (255, 255, 255))
    screen.blit(score_text, (10, 10))
    screen.blit(speed_text, (10, 40))

    # Game Over
    if game_over:
        over_text = font.render("GAME OVER", True, (255, 255, 0))
        restart_text = font.render("Pressione R para reiniciar", True, (255, 255, 255))
        screen.blit(over_text, (WIDTH//2 - 90, HEIGHT//2 - 40))
        screen.blit(restart_text, (WIDTH//2 - 160, HEIGHT//2))

        if keys[pygame.K_r]:
            speed = INITIAL_SPEED
            score = 0
            enemies.clear()
            for _ in range(3):
                enemies.append(create_enemy())
            game_over = False

    pygame.display.flip()

pygame.quit()
sys.exit()
