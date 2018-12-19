export default abstract class Filter {
  public getProps() {
    return Object.entries(this).map(([key, value]) => {
      return { prop: key, type: typeof value, value };
    });
  }

  public toJson() {
    return {
      [this.constructor.name.toLowerCase()]: Object.entries(this).map(
        ([key, value]) => value
      )
    };
  }
}
