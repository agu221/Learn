import random

import pygame

from circleshape import CircleShape
from constants import ASTEROID_MIN_RADIUS


class Asteroid(CircleShape):

    def __init__(self, *args):
        super().__init__(*args)

    def draw(self, screen):
        pygame.draw.circle(screen, "white", self.position, self.radius)

    def update(self, dt):
        self.position += self.velocity * dt

    def split(self):
        self.kill()
        if self.radius <= ASTEROID_MIN_RADIUS:
            return
        angle = random.uniform(20, 50)
        vector_one = self.velocity.rotate(angle)
        vector_two = self.velocity.rotate(-angle)
        new_radius = self.radius - ASTEROID_MIN_RADIUS
        asteroid_one = Asteroid(self.position.x, self.position.y, new_radius)
        asteroid_two = Asteroid(self.position.x, self.position.y, new_radius)
        asteroid_one.velocity = vector_one * 1.2
        asteroid_two.velocity = vector_two * 1.2
