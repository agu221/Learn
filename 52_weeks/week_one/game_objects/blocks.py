import random

from game_objects.object import Object


class Block(Object):
    def __init__(self, position: list[int] = [0, 0], length: int = 5, y_limit=0):
        self._icon = "0" * length
        self._x_pos = position[0]
        self._y_pos = position[1]
        self.isOffScreen = False
        self.y_limit = y_limit

    def update(self):
        self.checkIsOffScreen()
        if not self.isOffScreen:
            self._y_pos += 1

    def checkIsOffScreen(self):
        self.isOffScreen = self._y_pos >= self.y_limit

    @property
    def x_pos(self) -> int:
        return self._x_pos

    @property
    def y_pos(self) -> int:
        return self._y_pos

    @property
    def icon(self) -> str:
        return self._icon


class BlockManager:
    def __init__(self, screen_height, screen_width):
        self.blocks_queue = []
        self.x_bounds = [screen_width // 3, screen_width * 2 // 3]
        self.y_bounds = [3, screen_height - 3]
        self.max_blocks = self.y_bounds[1] // 2

    def update(self):
        for i, block in enumerate(self.blocks_queue):
            block.update()
        if self.blocks_queue and self.blocks_queue[0].isOffScreen:
            self.blocks_queue.pop(0)

        if len(self.blocks_queue) < self.max_blocks:
            self.add_block_to_queue()

    def add_block_to_queue(self):
        x_pos = random.randint(self.x_bounds[0] + 6, self.x_bounds[1] - 6)
        y_pos = self.y_bounds[0] - 1
        new_block = Block(position=[x_pos, y_pos], y_limit=self.y_bounds[1])
        self.blocks_queue.append(new_block)

    def start_screen(self):
        for n in range(1, self.max_blocks):
            x_pos = random.randint(self.x_bounds[0] + 6, self.x_bounds[1] - 6)
            y_pos = self.y_bounds[1] - 2 * n
            new_block = Block(position=[x_pos, y_pos], y_limit=self.y_bounds[1])
            self.blocks_queue.append(new_block)
