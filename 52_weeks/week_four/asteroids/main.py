import pygame

from asteroid import Asteroid
from asteroidfield import AsteroidField
from constants import *
from player import Player
from shot import Shot


def main():
    print("Starting Asteroids!")
    print(f"Screen width: {SCREEN_WIDTH}")
    print(f"Screen height: {SCREEN_HEIGHT}")
    pygame.init()

    screen = pygame.display.set_mode((SCREEN_WIDTH, SCREEN_HEIGHT))
    game_clock = pygame.time.Clock()
    dt = 0
    updatable = pygame.sprite.Group()
    drawable = pygame.sprite.Group()
    asteroids = pygame.sprite.Group()
    shots = pygame.sprite.Group()

    Player.containers = (updatable, drawable)
    Shot.containers = (updatable, drawable, shots)
    Asteroid.containers = (asteroids, updatable, drawable)
    AsteroidField.containers = updatable

    player = Player(x=SCREEN_WIDTH / 2, y=SCREEN_HEIGHT / 2, radius=PLAYER_RADIUS)
    asteroidField = AsteroidField()
    while True:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                return

        for update in updatable:
            update.update(dt)

        for asteroid in asteroids:
            if asteroid.collide(player):
                return
            for shot in shots:
                if asteroid.collide(shot):
                    asteroid.split()

        screen.fill("black")

        for draw in drawable:
            draw.draw(screen)

        pygame.display.flip()

        dt = game_clock.tick(60) / 1000


if __name__ == "__main__":
    main()
    print("GAME OVER")
