from draw import statics


class Player:

    def __init__(self, curr_position: list[int] = [0, 0]):
        self.isDead = False
        self.icon = statics.PLAYER_ICON
        self.curr_x = curr_position[0]
        self.curr_y = curr_position[1]
        self.isOnGround = True
        self.fallSpeed = -3

    def jump(self):
        if not self.isFalling():
            self.curr_y += 3
        else:
            pass

    def move(self, input):
        if input == "L":
            self.curr_x -= 1
        else:
            self.curr_x += 1

    def isFalling(self):
        pass

    def update(self):
        if not self.isOnGround:
            self.curr_y -= self.fallSpeed
        if self.isDead:
            self.icon = statics.DEAD_PLAYER_ICON
