import Filter from "./filter";

export default class Crop extends Filter {
  public width: number;
  public height: number;
  public cropType: boolean;

  constructor(width: number, height: number, cropType: boolean) {
    super();
    this.width = width;
    this.height = height;
    this.cropType = cropType;
  }
}
