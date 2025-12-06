import pygame

pygame.init()

SCREEN_WIDTH, SCREEN_HEIGHT = 1280, 720
screen = pygame.display.set_mode((SCREEN_WIDTH, SCREEN_HEIGHT))
font = pygame.font.Font(None, 40)

def draw_menu():
    
    # draw the background img first
    menu_bg : pygame.Surface = pygame.image.load("./assets/images/main_menu_bg.png")
    menu_bg = pygame.transform.scale(menu_bg, (SCREEN_WIDTH, SCREEN_HEIGHT))

    screen.blit(menu_bg)

    # draw the buttons
    btn_rect: pygame.Rect = pygame.Rect(0, 0, 320, 40)
    btn_rect.center = (SCREEN_WIDTH//2, SCREEN_HEIGHT//2)

    btn_text: pygame.Surface = font.render("Start Game", True, "black")
    btn_text_rect = btn_text.get_rect(center=btn_rect.center)

    pygame.draw.rect(screen, "white", btn_rect)
    screen.blit(btn_text, btn_text_rect)

    return

def main() -> None:
    
    pygame.display.set_caption("The Legend Coach")
    clock = pygame.time.Clock()

    game_state = "at_menu"
    state_trigger = {
        "at_menu": draw_menu()
    }

    running = True
    while running:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False
            else:
                state_trigger[game_state]

        pygame.display.flip()
        clock.tick(30)

    pygame.quit()


if __name__ == "__main__":
    main()
