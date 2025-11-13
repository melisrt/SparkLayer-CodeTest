<script lang="ts">
  import type { TodoItem } from "./types";

  interface Props {
    id: number;
    title: string;
    description: string;
    priority: "urgent" | "medium" | "low";
    completed: boolean;
    onToggle: () => void;
    onDelete: () => void;
    onEdit: (updated: TodoItem) => void;
  }

  const { id, title, description, priority, completed, onToggle, onDelete, onEdit }: Props = $props();

  let isEditing = $state(false);
  let editTitle = $state(title);
  let editDescription = $state(description);
  let editPriority = $state<"urgent" | "medium" | "low">(priority);

  //Toggles edit mode or saves changes
  //If already editing: validates and saves the changes
  //If not editing: enters edit mode and initializes edit fields
  function handleEdit() {
    if (isEditing) {
      //Validate before saving
      if (!editTitle.trim() || !editDescription.trim()) {
        alert("Title and description are required");
        return;
      }
      onEdit({
        id,
        title: editTitle.trim(),
        description: editDescription.trim(),
        priority: editPriority,
        completed,
      });
      isEditing = false;
    } else {
      isEditing = true;
      editTitle = title;
      editDescription = description;
      editPriority = priority;
    }
  }

  function cancelEdit() {
    isEditing = false;
    editTitle = title;
    editDescription = description;
    editPriority = priority;
  }
</script>

<div class="todo" class:completed>
  <input 
    type="checkbox" 
    checked={completed} 
    onchange={onToggle}
    class="todo-checkbox"
    aria-label={completed ? "Mark todo as incomplete" : "Mark todo as complete"}
  />
  
  <div class="todo-details">
    {#if isEditing}
      <input 
        bind:value={editTitle} 
        class="edit-input" 
        placeholder="Title" 
        aria-label="Edit todo title"
      />
      <input 
        bind:value={editDescription} 
        class="edit-input" 
        placeholder="Description" 
        aria-label="Edit todo description"
      />
      <select 
        bind:value={editPriority} 
        class="edit-select"
        aria-label="Select todo priority"
      >
        <option value="urgent">Urgent</option>
        <option value="medium">Medium</option>
        <option value="low">Low</option>
      </select>
      <div class="edit-buttons">
        <button onclick={handleEdit} class="save-btn" aria-label="Save todo changes">Save</button>
        <button onclick={cancelEdit} class="cancel-btn" aria-label="Cancel editing todo">Cancel</button>
      </div>
    {:else}
      <span class="priority-badge" class:urgent={priority === "urgent"} class:medium={priority === "medium"} class:low={priority === "low"}>
        {priority}
      </span>
      <p class="todo-title">{title}</p>
      <p class="todo-description">{description}</p>
    {/if}
  </div>

  <!-- Action buttons (Edit/Delete) - only shown when not in edit mode -->
  {#if !isEditing}
    <div class="todo-actions">
      <button onclick={handleEdit} class="edit-btn" aria-label="Edit todo">Edit</button>
      <button onclick={onDelete} class="delete-btn" aria-label="Delete todo">Delete</button>
    </div>
  {/if}
</div>

<style>
  .todo {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 15px;

    padding: 0px 20px;

    background-color: #303540;

    border-style: solid;
    border-color: #464e61;
    border-width: 1px;
    border-radius: 15px;

    margin: 10px 0px;
    transition: all 0.3s ease;
  }

  /* Light mode styling for todo container */
  :global(.light-mode) .todo {
    background-color: white;
    border-color: #ddd;
  }

  /* Reduced opacity for completed todos */
  .todo.completed {
    opacity: 0.6;
  }

  .todo-checkbox {
    width: 20px;
    height: 20px;
    cursor: pointer;
  }

  .todo-details {
    display: flex;
    flex-direction: column;
    align-items: start;
    flex: 1;

    padding: 15px 0px;
  }

  .todo-title {
    font-size: 28px;
    font-weight: bold;
    margin: 0px;
    margin-bottom: 5px;
    text-decoration: none;
  }

  .todo.completed .todo-title {
    text-decoration: line-through;
  }

  .todo-description {
    display: flex;
    margin: 0px;
    font-size: 20px;
    color: #ddd;
    text-align: left;
  }

  :global(.light-mode) .todo-description {
    color: #666;
  }

  .todo.completed .todo-description {
    text-decoration: line-through;
  }

  /* Priority badge styling - color coded by priority level */
  .priority-badge {
    display: inline-block;
    padding: 4px 12px;
    border-radius: 12px;
    font-size: 14px;
    font-weight: bold;
    margin-bottom: 8px;
    text-transform: capitalize;
  }

  .priority-badge.urgent {
    background-color: #d53030;  /* Red for urgent */
    color: white;
  }

  .priority-badge.medium {
    background-color: #d69719;  /* Orange for medium */
    color: white;
  }

  .priority-badge.low {
    background-color: #21af21;  /* Green for low */
    color: white;
  }

  .todo-actions {
    display: flex;
    gap: 10px;
  }

  .edit-btn, .delete-btn {
    padding: 8px 16px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 14px;
    font-weight: bold;
  }

  .edit-btn {
    background-color: #4a90e2;
    color: white;
  }

  .edit-btn:hover {
    background-color: #357abd;
  }

  .delete-btn {
    background-color: #e24a4a;
    color: white;
  }

  .delete-btn:hover {
    background-color: #bd3535;
  }

  .edit-input, .edit-select {
    width: 100%;
    padding: 8px;
    margin-bottom: 8px;
    border: 1px solid #464e61;
    border-radius: 5px;
    background-color: #282c34;
    color: white;
    font-size: 16px;
  }

  :global(.light-mode) .edit-input,
  :global(.light-mode) .edit-select {
    border: 1px solid #ccc;
    background-color: white;
    color: #333;
  }

  .edit-buttons {
    display: flex;
    gap: 10px;
    margin-top: 8px;
  }

  .save-btn, .cancel-btn {
    padding: 8px 16px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    font-size: 14px;
    font-weight: bold;
  }

  .save-btn {
    background-color: #4a90e2;
    color: white;
  }

  .cancel-btn {
    background-color: #666;
    color: white;
  }
</style>
