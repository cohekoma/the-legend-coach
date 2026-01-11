from random import randint

class Team:
    def __init__(self, name, ovr):
        self.name = name
        self.ovr = ovr

def calculate_performance(ovr_point) -> int:
    return ovr_point + randint(-20, 20)

def test_match(team_a : Team, team_b : Team):
    team_a_perf = calculate_performance(team_a.ovr)
    team_b_perf = calculate_performance(team_b.ovr)
    print(team_a_perf)
    print(team_b_perf)

def main() -> None:
    team_a = Team("Real Madrid", 80)
    team_b = Team("Wrexham", 50)
    test_match(team_a, team_b)

if __name__ == "__main__":
    main()
