import type { Player, Club, League } from "./types";

export async function loadData() {
    const [players, clubs, leagues] = await Promise.all(
        [
            fetch('/data/players.json').then(response => response.json() as Promise<Player[]>),
            fetch('/data/clubs.json').then(response => response.json() as Promise<Club[]>),
            fetch('/data/leagues.json').then(response => response.json() as Promise<League[]>)
        ]
    );

    return { players, clubs, leagues }
}