import Filter from "./filter";

export default class Background extends Filter {
  public color: string;

  constructor(color: string) {
    super();
    this.color = color;
  }
}
