// const firstTeamOvr = 90;
// const secondTeamOvr = 80;

const startGameBTN = document.getElementById("start-btn");

startGameBTN.onclick = () => {
  console.log( "Win Possibility: " + calculateMatchOutcome(20, 80) );
}

function calculateMatchOutcome(firstTeamOvr, secondTeamOvr) {
  const baseWinRate = 50;
  const strengthDifferent = ( firstTeamOvr > secondTeamOvr ) ? firstTeamOvr - secondTeamOvr : secondTeamOvr - firstTeamOvr;
  const winPossibility = strengthDifferent * 1;

  return winPossibility
}