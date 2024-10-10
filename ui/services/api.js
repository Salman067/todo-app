
export async function fetchTodos(currentPage, perPage, searchTask, searchStatus) {
  try {
    const params = new URLSearchParams({
      page: currentPage.toString(),
      per_page: perPage.toString(),
    });

    if (searchTask.trim()) {
      params.append("task", searchTask.trim());
    }

    if (searchStatus) {
      params.append("status", searchStatus);
    }

    const url = `/api/v1/todos?${params.toString()}`;
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`Failed to get todo list. statusCode: ${response.status}`);
    }

    const data = await response.json();
    return data; 
  } catch (error) {
    console.error("Error fetching todos:", error);
    throw error; 
  }
}

// Function to add a new todo
export async function addTodo(newTask) {
  if (newTask.trim() === "") return;

  try {
    const response = await fetch("/api/v1/todos", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        task: newTask,
        status: "created",
      }),
    });

    if (!response.ok) {
      throw new Error(`Failed to create todo. statusCode: ${response.status}`);
    }

    return response.json();
  } catch (error) {
    console.error("Error creating todo:", error);
    throw error;
  }
}


// Edit an existing todo
export async function editTodo(editTaskId, editTaskContent) {
  if (!editTaskContent.trim()) {
    throw new Error("Please enter a valid task.");
  }

  try {
    const response = await fetch(`/api/v1/todos/${editTaskId}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        task: editTaskContent,
      }),
    });

    if (!response.ok) {
      throw new Error(`Failed to edit todo. statusCode: ${response.status}`);
    }

    return response.json();
  } catch (error) {
    console.error("Error editing task:", error.message);
    throw error;
  }
}



// Update the status of a todo
export async function updateTodoStatus(todo) {
  const statuses = ["created", "in-progress", "done"];
  const currentIndex = statuses.indexOf(todo.Status);
  const nextIndex = (currentIndex + 1) % statuses.length;
  const nextStatus = statuses[nextIndex];

  try {
    const response = await fetch(`/api/v1/todos/${todo.ID}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        Status: nextStatus,
      }),
    });

    if (!response.ok) {
      throw new Error(`Failed to update todo status. statusCode: ${response.status}`);
    }

    return response.json();
  } catch (error) {
    console.error("Error updating todo status:", error.message);
    throw error;
  }
}

// delete todo
export async function deleteTodo(id) {
  try {
    const response = await fetch(`/api/v1/todos/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error(`Failed to delete todo. statusCode: ${response.status}`);
    }

    return response.json();
  } catch (error) {
    console.error("Error deleting todo:", error.message);
    throw error;
  }
}

// view todo
export async function viewTodoDetails(id) {
  try {
    const response = await fetch(`/api/v1/todos/${id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error(`Failed to fetch todo details. statusCode: ${response.status}`);
    }

    const data = await response.json();
    return data.data;
  } catch (error) {
    console.error("Error fetching todo details:", error.message);
    throw error;
  }
}