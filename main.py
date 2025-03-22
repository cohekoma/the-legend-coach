from model.Player import Player
class Main:
  def __init__(self):
    firstPlayer = Player('Chan', 25, 98, 12, {"Striker": 100, "Central Back": 10})
    self.printPlayer(firstPlayer)
  
  def printPlayer(self, player: Player):
    print(f"Player name: {player.name}, age: {player.age}, can play: {player.position}, ovr: {player.overrallRating}")
if __name__ == '__main__':
  Main()