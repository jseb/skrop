import Filter from "./filter";

export default class Resize extends Filter {
  public width: number;
  public height: number;
  public keepAspectRatio: boolean;

  constructor(width: number, height: number, keepAspectRatio: boolean) {
    super();
    this.width = width;
    this.height = height;
    this.keepAspectRatio = keepAspectRatio;
  }
}
