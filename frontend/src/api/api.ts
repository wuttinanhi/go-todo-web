export class API {
  /**
   * get api url
   */
  public static getAPIHost() {
    if (process.env.NODE_ENV === "development") {
      return "http://localhost:3000";
    }
    return "";
  }
}
