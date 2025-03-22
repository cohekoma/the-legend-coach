class Player:
  def __init__(self, name:str, age:int, att:int, dfn:int, position:dict[str:int]):
    self.name = name
    self.age = age
    self.position = position
    self.att = att
    self.dfn = dfn
    self.overrallRating = self.generateOverrall()

  def generateOverrall(self):
    return (self.att + self.dfn) // 2
  