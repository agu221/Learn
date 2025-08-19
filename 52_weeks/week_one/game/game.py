import curses
import time

from game_objects import blocks, player
from game_objects.object import Object


class GameSession:

    def __init__(self, scr):
        self.scr = scr
        self.screen_height, self.screen_width = self.scr.getmaxyx()
        self.block_manager = blocks.BlockManager(self.screen_height, self.screen_width)
        self.x_bounds = [self.screen_width // 3, self.screen_width * 2 // 3]
        self.y_bounds = [3, self.screen_height - 3]
        self.starting_pos = [
            (self.x_bounds[1] + self.x_bounds[0]) // 2,
            self.y_bounds[1],
        ]
        self.player = player.Player(curr_position=self.starting_pos)
        self.scr.keypad(True)
        curses.curs_set(0)
        curses.noecho()
        curses.cbreak()
        self.scr.nodelay(True)
        self.running = False

    def draw(self, unit: Object):
        self.scr.addstr(unit.y_pos, unit.x_pos, unit.icon)
        self.scr.refresh()

    def draw_borders(self):
        for x in range(self.x_bounds[0], self.x_bounds[1]):
            self.scr.addstr(self.y_bounds[0], x, "-")
            self.scr.addstr(self.y_bounds[1], x, "-")

        for y in range(self.y_bounds[0], self.y_bounds[1]):
            self.scr.addstr(y, self.x_bounds[0], "|")
            self.scr.addstr(y, self.x_bounds[1], "|")

    def update(self):
        self.scr.refresh()

    def wait_for_player_input(self):
        key = self.scr.getch()
        if key == ord("s"):
            self.player.handle_input(key)
            self.running = True

    def run_game_loop(self):
        self.scr.clear()
        last_time = time.time()
        fps = 5
        delay = 1.0 / fps
        self.block_manager.start_screen()
        self.draw_borders()
        for block in self.block_manager.blocks_queue:
            self.draw(block)

        self.draw(self.player)
        while not self.running:
            self.wait_for_player_input()
        while self.running:
            self.scr.clear()
            self.draw_borders()
            self.draw(self.player)
            for block in self.block_manager.blocks_queue:
                self.draw(block)
            self.scr.refresh()
            key = self.scr.getch()
            if key != -1:
                if key in (ord("q"), 27):
                    self.scr.clear()
                    self.running = False
                else:
                    self.player.handle_input(key)

            dt = time.time() - last_time
            last_time = time.time()

            self.player.update(dt)
            self.block_manager.update()

            if self.player.isDead:
                self.running = False

            time.sleep(delay)
