import curses

from draw import statics
from game_objects import blocks


class Screen:
    def __init__(self, stdscr):
        self.stdscr = stdscr
        self.screen_height, self.screen_width = stdscr.getmaxyx()

    def WriteTitleToScreen(self):
        title = statics.TITLE_NAME
        lines = title.split("//")
        start_y = (self.screen_height - len(lines)) // 2
        for i, line in enumerate(lines):
            start_x = (self.screen_width - len(line)) // 2
            self.stdscr.addstr(start_y + i, start_x, line)

    def DrawBorders(self):
        for row in range(1, self.screen_height - 1):
            self.stdscr.addstr(row, 1, "|")
            self.stdscr.addstr(row, self.screen_width - 2, "|")

        for column in range(1, self.screen_width - 1):
            self.stdscr.addstr(1, column, "-")
            self.stdscr.addstr(self.screen_height - 2, column, "-")

    def CreateTitleScreen(self):

        curses.curs_set(0)
        self.stdscr.clear()

        self.WriteTitleToScreen()
        self.DrawBorders()
        self.stdscr.addstr(
            self.screen_height - self.screen_height // 3,
            self.screen_width // 2 - len(statics.WELCOME_TEXT) // 2,
            statics.WELCOME_TEXT,
            curses.A_BOLD,
        )
        self.stdscr.refresh()

        self.wait_for_input()

    def wait_for_input(self):
        self.stdscr.nodelay(True)
        while True:
            key = self.stdscr.getch()
            if key in [10, 13, ord(" "), curses.KEY_ENTER]:
                self.stdscr.clear()
                return
            elif key == ord("q"):
                break
