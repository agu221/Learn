from game_objects import blocks, player


class GameSession:

    def __init__(self, scr):
        self.scr = scr
        self.current_blocks = []
        self.player = player.Player()

    def draw(self, unit, position):
        for i, piece in enumerate(unit):
            self.scr.addstr(position[0], position[1], unit)

    def update(self):
        self.scr.refresh()

    def add_block(self, block: blocks.Block):
        self.current_blocks.append(block)

    def wait_for_input(self):
        next_input = self.scr.getch()
