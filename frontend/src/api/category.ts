import { API } from "./api";

export interface ICategory {
  ID: number;
  name: string;
  CreatedAt?: Date;
  UpdatedAt?: Date;
  DeletedAt?: any;
}

export class CategoryAPI extends API {
  /**
   * getCategory
   * @returns {Promise<ICategory[]>}
   */
  public static async getCategory(): Promise<ICategory[]> {
    const request = await fetch(this.getAPIHost() + "/api/category");
    return request.json();
  }

  /**
   * addCategory
   * @param categoryName
   * @returns {Promise<void>}
   */
  public static async addCategory(categoryName: string): Promise<void> {
    await fetch(this.getAPIHost() + "/api/category/create", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name: categoryName }),
    });
  }

  /**
   * updateCategory
   * @param categoryId
   * @param categoryName
   * @returns {Promise<void>}
   */
  public static async updateCategory(
    categoryId: number,
    categoryName: string
  ): Promise<void> {
    await fetch(this.getAPIHost() + "/api/category/update", {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        id: categoryId,
        name: categoryName,
      }),
    });
  }

  /**
   * deleteCategory
   * @param categoryId
   * @returns {Promise<void>}
   */
  public static async deleteCategory(categoryId: number): Promise<void> {
    await fetch(this.getAPIHost() + "/api/category/delete", {
      method: "DELETE",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ id: categoryId }),
    });
  }
}
