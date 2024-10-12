export async function fetchTodos(
  currentPage,
  perPage,
  keyWord
) {
  try {
    const params = new URLSearchParams({
      page: currentPage.toString(),
      per_page: perPage.toString(),
    });

    if (keyWord.trim()) {
      params.append("key_word", keyWord.trim());
    }

    const url = `/api/v1/todos?${params.toString()}`;
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(
        `Failed to get todo list. statusCode: ${response.status}`
      );
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching todos:", error);
    throw error;
  }
}

export async function fetchTodosForStatus(currentPage, perPage, status) {
  try {
    const params = new URLSearchParams({
      page: currentPage.toString(),
      per_page: perPage.toString(),
    });

    if (status.trim()) {
      params.append("key_word", status.trim());
    }
    const url = `/api/v1/todos/status?${params.toString()}`;
    const response = await fetch(url);
  
    if (!response.ok) {
      throw new Error(
        `Failed to get todo list. statusCode: ${response.status}`
      );
    }

    const data = await response.json();
    console.log(data);
    return data;
  } catch (error) {
    console.error("Error fetching todos:", error);
    throw error;
  }
}

// Function to add a new todo
export async function addTodo(currentTask) {
  try {
    const response = await fetch("/api/v1/todos", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        task_title: currentTask.title,
        description: currentTask.description,
        priority_task: currentTask.priority.toLowerCase(),
        status: currentTask.status.toLowerCase(),
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
export async function editTodo(updatedTask) {
  try {
    const response = await fetch(`/api/v1/todos/${updatedTask.ID}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        task_title: updatedTask.Task,
        description: updatedTask.Description,
        priority_task: updatedTask.PriorityTask.toLowerCase(),
        status: updatedTask.Status.toLowerCase(),
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
      throw new Error(
        `Failed to update todo status. statusCode: ${response.status}`
      );
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
      throw new Error(
        `Failed to fetch todo details. statusCode: ${response.status}`
      );
    }

    const data = await response.json();
    return data.data;
  } catch (error) {
    console.error("Error fetching todo details:", error.message);
    throw error;
  }
}

export async function fetchTasksWithoutPagination() {
  try {
    const url = `/api/v1/todos/withoutPagination`;
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(
        `Failed to get todo list. statusCode: ${response.status}`
      );
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching todos:", error);
    throw error;
  }
}
