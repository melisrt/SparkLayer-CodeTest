<script lang="ts">
  import Todo from "./lib/Todo.svelte";
  import type { TodoItem } from "./lib/types";

  let todos: TodoItem[] = $state([]);
  let title = $state("");
  let description = $state("");
  let priority = $state<"urgent" | "medium" | "low" | "">("");
  
  //Load theme from localStorage or default to dark
  let isDarkMode = $state((() => {
    if (typeof window !== "undefined") {
      const saved = localStorage.getItem("theme");
      return saved ? saved === "dark" : true;
    }
    return true;
  })());
  
  //Save theme to localStorage when it changes
  $effect(() => {
    if (typeof window !== "undefined") {
      localStorage.setItem("theme", isDarkMode ? "dark" : "light");
    }
  });
  
  function showError(message: string) {
    alert(message);
  }

  async function fetchTodos() {
    try {
      const response = await fetch("http://localhost:8080/");
      if (response.status !== 200) {
        showError("Failed to load todos. Please try again.");
        return;
      }

      todos = await response.json();
    } catch (e) {
      showError("Could not connect to server. Please ensure the backend is running.");
    }
  }

  async function handleSubmit() {
    // Validate input
    const trimmedTitle = title.trim();
    const trimmedDescription = description.trim();
    
    if (!trimmedTitle || !trimmedDescription) {
      showError("Title and description are required");
      return;
    }
    
    const selectedPriority = priority || "medium";
    const newTodo = { title: trimmedTitle, description: trimmedDescription, priority: selectedPriority };
    try {
      const response = await fetch("http://localhost:8080/", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newTodo),
      });

      if (response.ok) {
        title = "";
        description = "";
        priority = "";
        await fetchTodos();
      } else {
        const errorText = await response.text();
        showError(`Failed to add todo: ${errorText}`);
      }
    } catch (e) {
      showError("Error adding todo. Please try again.");
    }
  }

  // Sort todos: urgent > medium > low, then completed at end
  function sortTodos(todos: TodoItem[]): TodoItem[] {
    const priorityOrder = { urgent: 0, medium: 1, low: 2 };
    return [...todos].sort((a, b) => {
      // Completed todos go to the end
      if (a.completed !== b.completed) {
        return a.completed ? 1 : -1;
      }
      // Sort by priority
      return priorityOrder[a.priority] - priorityOrder[b.priority];
    });
  }

  let sortedTodos = $derived(sortTodos(todos));

  async function updateTodo(id: number, updated: TodoItem) {
    try {
      const response = await fetch(`http://localhost:8080/?id=${id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(updated),
      });

      if (response.ok) {
        await fetchTodos();
      } else {
        const errorText = await response.text();
        showError(`Failed to update todo: ${errorText}`);
      }
    } catch (e) {
      showError("Error updating todo. Please try again.");
    }
  }

  async function deleteTodo(id: number) {
    try {
      const response = await fetch(`http://localhost:8080/?id=${id}`, {
        method: "DELETE",
      });

      if (response.ok) {
        await fetchTodos();
      } else {
        const errorText = await response.text();
        showError(`Failed to delete todo: ${errorText}`);
      }
    } catch (e) {
      showError("Error deleting todo. Please try again.");
    }
  }

  // Initially fetch todos on page load
  $effect(() => {
    fetchTodos();
  });
</script>

<main class="app" class:light-mode={!isDarkMode}>
  <button class="theme-toggle" onclick={() => isDarkMode = !isDarkMode}>
    {#if isDarkMode}
      ☀️
    {:else}
      🌙
    {/if}
  </button>
  <header class="app-header">
    <h1>TODO</h1>
  </header>

  <h2 class="todo-list-form-header">Add a Todo</h2>
  <form class="todo-list-form" onsubmit={(e) => { e.preventDefault(); handleSubmit(); }}>
    <input placeholder="Title" name="title" bind:value={title} />
    <input placeholder="Description" name="description" bind:value={description} />
    <select bind:value={priority}>
      <option value="" disabled>Select Priority</option>
      <option value="urgent">Urgent</option>
      <option value="medium">Medium</option>
      <option value="low">Low</option>
    </select>
    <button type="submit">Add Todo</button>
  </form>

  {#if sortedTodos.length > 0}
    <h2 class="current-todos-header">Current ToDos</h2>
  {/if}
  <div class="todo-list">
    {#each sortedTodos as todo (todo.id)}
      <Todo 
        id={todo.id}
        title={todo.title} 
        description={todo.description}
        priority={todo.priority}
        completed={todo.completed}
        onToggle={async () => {
          const todoId = todo.id;
          const currentTodo = todos.find(t => t.id === todoId);
          if (currentTodo) {
            const updated = { ...currentTodo, completed: !currentTodo.completed };
            await updateTodo(todoId, updated);
          }
        }}
        onDelete={async () => {
          await deleteTodo(todo.id);
        }}
        onEdit={async (updated) => {
          await updateTodo(todo.id, updated);
        }}
      />
    {/each}
  </div>
</main>

<style>
  .app {
    color: white;
    background-color: #282c34;
    text-align: center;
    font-size: 24px;
    min-height: 100vh;
    padding: 20px;
    position: relative;
    transition: background-color 0.3s ease, color 0.3s ease;
  }

  .app.light-mode {
    background-color: #f5f5f5;
    color: #333;
  }

  .theme-toggle {
    position: fixed;
    top: 20px;
    right: 20px;
    background: none;
    border: 2px solid currentColor;
    border-radius: 50%;
    width: 50px;
    height: 50px;
    font-size: 24px;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
    transition: all 0.3s ease;
  }

  .theme-toggle:hover {
    transform: scale(1.1);
    background-color: rgba(255, 255, 255, 0.1);
  }

  .light-mode .theme-toggle {
    border-color: #333;
  }

  .light-mode .theme-toggle:hover {
    background-color: rgba(0, 0, 0, 0.1);
  }

  .app-header {
    font-size: calc(10px + 4vmin);
    margin-top: 50px;
  }

  .current-todos-header {
    margin-top: 50px;
    margin-bottom: 20px;
    font-size: calc(10px + 3vmin);
  }

  .todo-list {
    margin: 0px 100px 0px 100px;
  }

  .todo-list-form-header {
    margin-top: 50px;
    margin-bottom: 20px;
  }

  .todo-list-form {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 15px;
    margin: 20px auto;
    max-width: 500px;
  }

  .todo-list-form input,
  .todo-list-form select {
    width: 100%;
    box-sizing: border-box;
    padding: 12px;
    border: 1px solid #464e61;
    border-radius: 8px;
    background-color: #303540;
    color: white;
    font-size: 16px;
    transition: all 0.3s ease;
  }

  .light-mode .todo-list-form input,
  .light-mode .todo-list-form select {
    border: 1px solid #ccc;
    background-color: white;
    color: #333;
  }

  .todo-list-form input::placeholder {
    color: #888;
  }

  .light-mode .todo-list-form input::placeholder {
    color: #999;
  }

  .todo-list-form button {
    padding: 12px 24px;
    background-color: #4a90e2;
    color: white;
    border: none;
    border-radius: 8px;
    font-size: 16px;
    font-weight: bold;
    cursor: pointer;
  }

  .todo-list-form button:hover {
    background-color: #357abd;
  }
</style>
