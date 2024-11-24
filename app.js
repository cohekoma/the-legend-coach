const startGameBTN = document.getElementById("startGame");
const coachNameTxt = document.querySelector("#coachName");

coachNameTxt.innerText = "Chan"

startGameBTN.onclick = () => {
  console.log( "Win Possibility: " + calculateMatchOutcome(20, 80) );
}

function calculateMatchOutcome(firstTeamOvr, secondTeamOvr) {
  const baseWinRate = 50;
  const strengthDifferent = ( firstTeamOvr > secondTeamOvr ) ? firstTeamOvr - secondTeamOvr : secondTeamOvr - firstTeamOvr;
  const winPossibility = strengthDifferent * 2;

  return winPossibility
}

function manageCoach() {
  console.log("Manage Your Coach");
}

function manageClub() {
  console.log("Manage Club");
}