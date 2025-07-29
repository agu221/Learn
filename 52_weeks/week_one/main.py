import curses
from draw.draw_components import Screen


def main():
    new_screen = curses.wrapper(Screen)
    new_screen.CreateTitleScreen()


if __name__ == "__main__":
    main()
