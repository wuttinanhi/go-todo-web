import { API } from "./api";

export interface ITodo {
  id: number;
  category_id: number;
  name: string;
}

export class TodoAPI extends API {
  /**
   * getTodo
   * @returns {Promise<ITodo[]>}
   */
  public static async getTodo(): Promise<ITodo[]> {
    const request = await fetch(this.getAPIHost() + "/api/todo");
    return request.json();
  }

  /**
   * addTodo
   * @param todoName
   * @param categoryId
   * @returns {Promise<void>}
   */
  public static async addTodo(
    todoName: string,
    categoryId: number
  ): Promise<void> {
    await fetch(this.getAPIHost() + "/api/todo/create", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name: todoName, category_id: categoryId }),
    });
  }

  /**
   * updateTodo
   * @param todoId
   * @param todoName
   * @param categoryId
   * @returns {Promise<void>}
   */
  public static async updateTodo(
    todoId: number,
    todoName: string,
    categoryId: number
  ): Promise<void> {
    await fetch(this.getAPIHost() + "/api/todo/update", {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        id: todoId,
        name: todoName,
        category_id: categoryId,
      }),
    });
  }

  /**
   * deleteTodo
   * @param todoId
   * @returns {Promise<void>}
   */
  public static async deleteTodo(todoId: number): Promise<void> {
    await fetch(this.getAPIHost() + "/api/todo/delete", {
      method: "DELETE",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        id: todoId,
      }),
    });
  }
}
