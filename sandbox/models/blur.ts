import Filter from "./filter";

export default class Blur extends Filter {
  public sigma: number;
  public minAmpl: number;

  constructor(sigma: number, minAmpl: number) {
    super();
    this.sigma = sigma;
    this.minAmpl = minAmpl;
  }
}
