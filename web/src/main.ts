import Phaser from 'phaser';
class MenuScene extends Phaser.Scene {
  constructor() { super('menu'); }
  preload() { this.load.image('background', '/assets/images/main_menu_bg.png'); }
  create() {
    this.add.image(640, 360, 'background').setDisplaySize(1280, 720);
    const btn = this.add.text(640, 360, 'Start Game', { fontSize: '32px', color: '#000' }).setOrigin(0.5);
    btn.setInteractive().on('pointerup', () => this.scene.start('game'));
  }
}
class GameScene extends Phaser.Scene { create() { this.add.text(640, 360, 'Game loop'); } }

new Phaser.Game({
  type: Phaser.AUTO,
  width: 1280,
  height: 720,
  parent: 'app',
  scene: [MenuScene, GameScene],
  physics: { default: 'arcade' }
});
