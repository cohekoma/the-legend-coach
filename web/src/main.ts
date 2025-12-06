import Phaser from 'phaser';
import MenuScene from './scenes/menuScene';
import SquadScene from './scenes/squadScene';

class GameScene extends Phaser.Scene { 
  constructor() {
    super('game')
  }

  create() { 
    this.add.text(16, 16, 'Game loop', {fontSize: '50px', color: '#fff'}); 
  } 
}

const game = new Phaser.Game({
  type: Phaser.AUTO,
  parent: 'app',
  scene: [MenuScene, GameScene, SquadScene],
  physics: { default: 'arcade' },
  scale: {
    mode: Phaser.Scale.RESIZE,
    autoCenter: Phaser.Scale.CENTER_BOTH
  }
});

window.addEventListener('resize', () => {
  game.scale.resize(window.innerWidth, window.innerHeight);
});