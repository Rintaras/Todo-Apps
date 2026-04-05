(function () {
  const form = document.getElementById("todo-form");
  const titleInput = document.getElementById("todo-title");
  const listEl = document.getElementById("todo-list");
  const statusEl = document.getElementById("status-message");
  const template = document.getElementById("todo-item-template");
  const filterButtons = document.querySelectorAll(".filter-btn");

  let todos = [];
  let filter = "all";
  let useDemo = false;

  const DEMO_TODOS = [
    { id: "demo-1", title: "API を実装する", completed: false },
    { id: "demo-2", title: "Swagger UI で契約を確認する", completed: true },
  ];

  function setStatus(text, isError) {
    statusEl.textContent = text || "";
    statusEl.classList.toggle("is-error", Boolean(isError));
  }

  function normalizeListResponse(data) {
    if (Array.isArray(data)) return data;
    if (data && Array.isArray(data.todos)) return data.todos;
    if (data && Array.isArray(data.items)) return data.items;
    return [];
  }

  function mapTodo(raw) {
    return {
      id: String(raw.id ?? raw.ID ?? ""),
      title: String(raw.title ?? raw.Title ?? ""),
      completed: Boolean(raw.completed ?? raw.done ?? raw.Completed),
    };
  }

  async function loadTodos() {
    setStatus("読み込み中…");
    try {
      const data = await window.TodoApi.listTodos();
      todos = normalizeListResponse(data).map(mapTodo);
      useDemo = false;
      setStatus("");
    } catch (e) {
      useDemo = true;
      todos = DEMO_TODOS.map((t) => ({ ...t }));
      setStatus(
        "API に接続できませんでした。デモ表示中です（実装後は自動で API データに切り替わります）。",
        false
      );
    }
    render();
  }

  function visibleTodos() {
    if (filter === "active") return todos.filter((t) => !t.completed);
    if (filter === "completed") return todos.filter((t) => t.completed);
    return todos;
  }

  function render() {
    listEl.innerHTML = "";
    const items = visibleTodos();
    if (items.length === 0) {
      const empty = document.createElement("p");
      empty.className = "empty-state";
      empty.textContent =
        filter === "all"
          ? "タスクはまだありません。上のフォームから追加してください。"
          : filter === "active"
            ? "未完了のタスクはありません。"
            : "完了したタスクはありません。";
      listEl.appendChild(empty);
      return;
    }

    for (const todo of items) {
      const node = template.content.firstElementChild.cloneNode(true);
      node.dataset.id = todo.id;
      node.classList.toggle("is-completed", todo.completed);
      const cb = node.querySelector(".todo-item__check");
      cb.checked = todo.completed;
      node.querySelector(".todo-item__title").textContent = todo.title;
      listEl.appendChild(node);
    }
  }

  listEl.addEventListener("change", async (ev) {
    const cb = ev.target;
    if (!cb.classList.contains("todo-item__check")) return;
    const li = cb.closest(".todo-item");
    const id = li?.dataset.id;
    if (!id) return;
    const todo = todos.find((t) => t.id === id);
    if (!todo) return;
    const next = cb.checked;

    if (useDemo) {
      todo.completed = next;
      li.classList.toggle("is-completed", next);
      return;
    }

    cb.disabled = true;
    try {
      await window.TodoApi.updateTodo(id, { completed: next });
      todo.completed = next;
      li.classList.toggle("is-completed", next);
    } catch (err) {
      cb.checked = !next;
      setStatus(err.message || "更新に失敗しました", true);
    } finally {
      cb.disabled = false;
    }
  });

  listEl.addEventListener("click", async (ev) => {
    const btn = ev.target.closest(".todo-item__delete");
    if (!btn) return;
    const li = btn.closest(".todo-item");
    const id = li?.dataset.id;
    if (!id) return;

    if (useDemo) {
      todos = todos.filter((t) => t.id !== id);
      render();
      return;
    }

    btn.disabled = true;
    try {
      await window.TodoApi.deleteTodo(id);
      todos = todos.filter((t) => t.id !== id);
      render();
      setStatus("");
    } catch (err) {
      setStatus(err.message || "削除に失敗しました", true);
    } finally {
      btn.disabled = false;
    }
  });

  form.addEventListener("submit", async (ev) => {
    ev.preventDefault();
    const title = titleInput.value.trim();
    if (!title) return;

    if (useDemo) {
      const id = "demo-" + Date.now();
      todos.unshift({ id, title, completed: false });
      titleInput.value = "";
      render();
      return;
    }

    const submitBtn = form.querySelector('[type="submit"]');
    submitBtn.disabled = true;
    try {
      const created = await window.TodoApi.createTodo({ title });
      const t = mapTodo(created);
      if (t.id) todos.unshift(t);
      else await loadTodos();
      titleInput.value = "";
      setStatus("");
      render();
    } catch (err) {
      setStatus(err.message || "作成に失敗しました", true);
    } finally {
      submitBtn.disabled = false;
    }
  });

  filterButtons.forEach((b) => {
    b.addEventListener("click", () => {
      filterButtons.forEach((x) => x.classList.remove("is-active"));
      b.classList.add("is-active");
      filter = b.dataset.filter || "all";
      render();
    });
  });

  loadTodos();
})();
