<script lang="ts">
  import Todo from "./Todo.svelte";
  import {
    fetchCategory,
    fetchTodo,
    ICategory,
    ITodo,
    addTodo,
    updateTodo,
  } from "./api";

  let categories: ICategory[] = [];
  let todos: ITodo[] = [];

  async function addTodoWrapper(categoryId: number) {
    // prompt user todo name
    let todoName = prompt("Please enter todo:", "");

    // send add todo request
    await addTodo(todoName, categoryId);

    // update todos array
    todos = [
      ...todos,
      {
        id: todos.length + 1,
        name: todoName,
        category_id: categoryId,
      },
    ];
  }

  async function fetcher() {
    categories = await fetchCategory();
    todos = await fetchTodo();
  }

  function onDragStart(event: DragEvent, todoId: number) {
    event.dataTransfer.setData("todoId", todoId.toString());
  }

  function onDrop(event: DragEvent, categoryId: number) {
    const todoId = parseInt(event.dataTransfer.getData("todoId"));
    let arrayIndex = null;

    const todoElement = todos.filter((t, i) => {
      if (todoId === t.id) {
        arrayIndex = i;
        return true;
      }
    })[0];

    todos[arrayIndex].category_id = categoryId;
    updateTodo(todoId, todoElement.name, categoryId);
  }

  fetcher();
</script>

<main>
  <h1 class="text-3xl font-bold my-10 text-center text-orange-500">Todo App</h1>

  <div
    class="w-full h-full grid grid-cols-4 gap-3 content-between justify-between align-middle"
  >
    {#each categories as category}
      <div
        class="grid border-2 border-black w-full h-screen"
        on:dragenter={(event) => event.preventDefault()}
        on:dragover={(event) => event.preventDefault()}
        on:drop={(event) => onDrop(event, category.ID)}
      >
        <div class="flex flex-col px-4 items-start max-h-screen">
          <div class="flex flex-col w-full items-center mb-3">
            <h3 class="text-lg font-bold my-3">
              {category.name || "Unnamed Category"}
            </h3>

            <button
              class="flex orange-button"
              on:click={() => addTodoWrapper(category.ID)}>ADD</button
            >
          </div>

          <div
            class="flex flex-col w-full overflow-y-scroll overflow-x-hiddenh-full mb-3 h-full"
          >
            {#each todos as todo}
              {#if todo.category_id === category.ID}
                <div
                  draggable={true}
                  on:dragstart={(event) => onDragStart(event, todo.id)}
                >
                  <Todo
                    id={todo.id}
                    name={todo.name}
                    categoryId={category.ID}
                  />
                </div>
              {/if}
            {/each}
          </div>
        </div>
      </div>
    {/each}
  </div>
</main>

<style lang="postcss" global>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;

  .orange-button {
    @apply bg-white text-orange-600 border-2 rounded-md font-bold text-center justify-center border-orange-600 w-full py-2 mx-5 px-5;
  }
</style>
