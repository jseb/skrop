import Filter from "./filter";

export default class Overlay extends Filter {
  public file: string;
  public opacity: number;
  public verticalGravity: string;
  public horizontalGravity: string;
  public rightMargin: number;
  public leftMargin: number;
  public topMargin: number;
  public bottomMargin: number;

  constructor(
    file: string,
    opacity: number,
    verticalGravity: string,
    horizontalGravity: string,
    rightMargin: number,
    leftMargin: number,
    topMargin: number,
    bottomMargin: number
  ) {
    super();
    this.file = file;
    this.opacity = opacity;
    this.horizontalGravity = horizontalGravity;
    this.verticalGravity = verticalGravity;
    this.rightMargin = rightMargin;
    this.leftMargin = leftMargin;
    this.topMargin = topMargin;
    this.bottomMargin = bottomMargin;
  }
}
