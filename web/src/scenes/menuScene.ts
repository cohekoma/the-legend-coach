export default class MenuScene extends Phaser.Scene {

  btn!: Phaser.GameObjects.Text;
  bg!: Phaser.GameObjects.Image;

  constructor() { super('menu'); }

  preload() { this.load.image('background', '/assets/images/main_menu_bg.png'); }
  create() {
    this.bg = this.add.image(0, 0, 'background').setOrigin(0);
    this.btn = this.add.text(0, 0, 'Start Game', { fontSize: '32px', color: '#fff' });

    this.resize(this.scale.gameSize.width, this.scale.gameSize.height);
    this.scale.on('resize', (size: Phaser.Structs.Size) => this.resize(size.width, size.height));
    this.btn.setInteractive().on('pointerup', () => this.scene.start('squad'));
  }

  resize(w: number, h: number) {
    this.bg.setDisplaySize(w, h);
    this.btn.setPosition(w / 2, h / 2).setOrigin(0.5);
  }
}