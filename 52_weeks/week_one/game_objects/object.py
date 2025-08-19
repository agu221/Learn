from abc import ABC, abstractmethod


class Object(ABC):
    @property
    @abstractmethod
    def x_pos(self) -> int:
        pass

    @property
    @abstractmethod
    def y_pos(self) -> int:
        pass

    @property
    @abstractmethod
    def icon(self) -> str:
        pass
