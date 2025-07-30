import curses

from draw.draw_components import Screen
from game import game


def main():
    new_screen = curses.wrapper(Screen)
    new_screen.CreateTitleScreen()
    new_game = curses.wrapper(game.GameSession)
    new_game.run_game_loop()


if __name__ == "__main__":
    main()
