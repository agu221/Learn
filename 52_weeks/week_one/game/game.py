import curses
import time

from game_objects import blocks, player


class GameSession:

    def __init__(self, scr):
        self.scr = scr
        self.current_blocks = []
        self.player = player.Player()
        self.scr.keypad(True)
        curses.curs_set(0)
        curses.noecho()
        curses.cbreak()
        self.scr.nodelay(True)
        self.running = True

    def draw(self, unit, position):
        for i, piece in enumerate(unit):
            self.scr.addstr(position[0], position[1], unit)

    def update(self):
        self.scr.refresh()

    def add_block(self, block: blocks.Block):
        self.current_blocks.append(block)

    def wait_for_input(self):
        next_input = self.scr.getch()

    def run_game_loop(self):
        self.scr.clear()
        last_time = time.time()
        fps = 60
        delay = 1.0 / fps

        while self.running:
            key = self.scr.getch()
            if key != -1:
                if key in (ord("q"), 27):
                    self.running = False
                else:
                    self.player.hanlde_input(key)

            dt = time.time() - last_time
            last_time = time.time()

            self.player.update(dt)

            if self.player.isDead:
                self.running = False

            time.sleep(max(0, delay - time.time() - last_time))
