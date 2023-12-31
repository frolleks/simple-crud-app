<script lang="ts">
  import { onMount } from "svelte";

  let todos: any[] = [];
  let newTodoTitle = "";

  async function getTodos() {
    try {
      const res = await fetch("http://localhost:3000/todos");
      if (!res.ok) {
        throw new Error(`Error: ${res.statusText}`);
      }
      let fetchedTodos = await res.json();

      // Sort the todos by the CreatedAt field, newest first
      fetchedTodos.sort(
        (a: any, b: any) =>
          new Date(b.CreatedAt).getTime() - new Date(a.CreatedAt).getTime()
      );

      todos = fetchedTodos;
    } catch (error) {
      console.error("Failed to fetch todos:", error);
    }
  }

  async function addTodo() {
    try {
      const response = await fetch("http://localhost:3000/todos", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ title: newTodoTitle, completed: false }),
      });

      if (!response.ok) {
        throw new Error(`Error: ${response.statusText}`);
      }

      // Clear the input field after adding
      newTodoTitle = "";

      // Optionally, fetch the updated list of todos
      getTodos();
    } catch (error) {
      console.error("Failed to add todo:", error);
    }
  }

  async function toggleTodoStatus(todo: any) {
    try {
      const response = await fetch(`http://localhost:3000/todos/${todo.ID}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ completed: todo.completed }), // Only sending the completed status
      });

      if (!response.ok) {
        throw new Error(`Error: ${response.statusText}`);
      }

      // Correctly toggle the completed status in the local todos array
      todos = todos.map((t) =>
        t.ID === todo.ID ? { ...t, completed: t.completed } : t
      );
    } catch (error) {
      console.error("Failed to toggle todo status:", error);
    }
  }

  async function deleteTodo(todoId: any) {
    try {
      const response = await fetch(`http://localhost:3000/todos/${todoId}`, {
        method: "DELETE",
      });

      if (!response.ok) {
        throw new Error(`Error: ${response.statusText}`);
      }

      // Remove the todo from the local array
      todos = todos.filter((t) => t.id !== todoId);

      getTodos();
    } catch (error) {
      console.error("Failed to delete todo:", error);
    }
  }

  onMount(() => {
    getTodos();
  });
</script>

<form on:submit|preventDefault={addTodo}>
  <input type="text" bind:value={newTodoTitle} placeholder="Add new todo" />
  <button type="submit" disabled={newTodoTitle.trim() === ""}>Add</button>
</form>

<ul>
  {#each todos as todo}
    <li>
      <input
        type="checkbox"
        bind:checked={todo.completed}
        on:change={() => toggleTodoStatus(todo)}
      />
      {todo.title}
      <button on:click={() => deleteTodo(todo.ID)}>Delete</button>
    </li>
  {/each}
</ul>
