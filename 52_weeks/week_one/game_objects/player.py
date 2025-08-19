import math

from draw import statics
from game_objects.object import Object


class Player(Object):

    def __init__(self, curr_position: list[int] = [0, 0]):
        self.isDead = False
        self._icon = statics.PLAYER_ICON
        self._x_pos = curr_position[0]
        self._y_pos = curr_position[1]
        self.isOnGround = True
        self.fallSpeed = -1

    @property
    def x_pos(self) -> int:
        return self._x_pos

    @property
    def y_pos(self) -> int:
        return self._y_pos

    @property
    def icon(self) -> str:
        return self._icon

    def jump(self):
        if not self.isFalling():
            self._y_pos += 3
        else:
            pass

    def move(self, input):
        if input == "L":
            self._x_pos -= 1
        else:
            self._x_pos += 1

    def isFalling(self):
        self.isOnGround = False

    def update(self, dt):
        if not self.isOnGround:
            self._y_pos -= math.floor(self.fallSpeed * dt)
        if self.isDead:
            self.icon = statics.DEAD_PLAYER_ICON

    def handle_input(self, input):
        if input == ord("a"):
            self.move("L")
        elif input == ord("d"):
            self.move("R")
        elif input == ord(" "):
            self.jump()
