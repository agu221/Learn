class Block():
    def __init__(self, position:list[int] = [6,6], length: int = 5):
        self.art = '0' * length 
        self.x_pos = position[0]
        self.y_pos = position[1]
