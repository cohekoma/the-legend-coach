import { loadData } from '../dataLoader';

export default class SquadScene extends Phaser.Scene {
  constructor() { super('squad'); }
  async create() {
    const { players, clubs } = await loadData();
    const clubId = 1; // pick a club for now
    const club = clubs.find(c => c.id === clubId);
    this.add.text(20, 20, `${club?.name ?? 'Club'} Squad`, { fontSize: '28px', color: '#fff' });

    const squad = players.filter(p => p.clubId === clubId);
    squad.forEach((p, i) => {
      this.add.text(40, 70 + i * 28, `${p.name} — ${p.position} (${p.overall})`, { fontSize: '20px', color: '#fff' });
    });
  }
}