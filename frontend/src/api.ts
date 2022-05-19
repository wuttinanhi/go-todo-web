export interface ICategory {
  ID: number;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt?: any;
  name: string;
}

export interface ITodo {
  id: number;
  category_id: number;
  name: string;
}

// function fetch from api
export async function fetchTodo(): Promise<ITodo[]> {
  const request = await fetch("http://localhost:3000/api/todo");
  return request.json();
}

// function add todo
export async function addTodo(todoName: string, categoryId: number) {
  const request = await fetch("http://localhost:3000/api/todo/create", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      name: todoName,
      category_id: categoryId,
    }),
  });
  return request.json();
}

// function update todo
export async function updateTodo(
  todoId: number,
  todoName: string,
  categoryId: number
) {
  const request = await fetch("http://localhost:3000/api/todo/update", {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      id: todoId,
      name: todoName,
      category_id: categoryId,
    }),
  });
  return request.json();
}

// function delete todo
export async function deleteTodo(todoId: number) {
  const request = await fetch("http://localhost:3000/api/todo/delete", {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      id: todoId,
    }),
  });
  return request.json();
}

// function fetch category
export async function fetchCategory(): Promise<ICategory[]> {
  const request = await fetch("http://localhost:3000/api/category");
  return request.json();
}
