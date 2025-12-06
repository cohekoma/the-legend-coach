import Phaser from 'phaser';

class MenuScene extends Phaser.Scene {

  btn!: Phaser.GameObjects.Text;
  bg!: Phaser.GameObjects.Image;

  constructor() { super('menu'); }

  preload() { this.load.image('background', '/assets/images/main_menu_bg.png'); }
  create() {
    this.bg = this.add.image(0, 0, 'background').setOrigin(0);
    this.btn = this.add.text(0, 0, 'Start Game', { fontSize: '32px', color: '#fff' });

    this.resize(this.scale.gameSize.width, this.scale.gameSize.height);
    this.scale.on('resize', (size: Phaser.Structs.Size) => this.resize(size.width, size.height));
    this.btn.setInteractive().on('pointerup', () => this.scene.start('game'));
  }

  resize(w: number, h: number) {
    this.bg.setDisplaySize(w, h);
    this.btn.setPosition(w / 2, h / 2).setOrigin(0.5);
  }
}

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
  scene: [MenuScene, GameScene],
  physics: { default: 'arcade' },
  scale: {
    mode: Phaser.Scale.RESIZE,
    autoCenter: Phaser.Scale.CENTER_BOTH
  }
});

window.addEventListener('resize', () => {
  game.scale.resize(window.innerWidth, window.innerHeight);
});