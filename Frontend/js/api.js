/**
 * Swagger（api-document.yaml）に沿ったエンドポイントを叩く薄いクライアント。
 * バックエンド実装前はネットワークエラーになり、app.js でデモ用データにフォールバックします。
 */
(function () {
  const cfg = window.APP_CONFIG || { apiBaseUrl: "/api" };
  const base = cfg.apiBaseUrl.replace(/\/$/, "");

  async function parseJsonResponse(res) {
    const text = await res.text();
    if (!text) return null;
    try {
      return JSON.parse(text);
    } catch {
      return null;
    }
  }

  function apiError(res, body) {
    const err = new Error(body?.message || res.statusText || "API エラー");
    err.status = res.status;
    err.body = body;
    return err;
  }

  window.TodoApi = {
    baseUrl: base,

    async listTodos() {
      const res = await fetch(`${base}/todos`, { headers: { Accept: "application/json" } });
      const body = await parseJsonResponse(res);
      if (!res.ok) throw apiError(res, body);
      return body;
    },

    async createTodo(payload) {
      const res = await fetch(`${base}/todos`, {
        method: "POST",
        headers: { "Content-Type": "application/json", Accept: "application/json" },
        body: JSON.stringify(payload),
      });
      const body = await parseJsonResponse(res);
      if (!res.ok) throw apiError(res, body);
      return body;
    },

    async updateTodo(id, payload) {
      const res = await fetch(`${base}/todos/${encodeURIComponent(id)}`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json", Accept: "application/json" },
        body: JSON.stringify(payload),
      });
      const body = await parseJsonResponse(res);
      if (!res.ok) throw apiError(res, body);
      return body;
    },

    async deleteTodo(id) {
      const res = await fetch(`${base}/todos/${encodeURIComponent(id)}`, {
        method: "DELETE",
        headers: { Accept: "application/json" },
      });
      if (res.status === 204) return;
      const body = await parseJsonResponse(res);
      if (!res.ok) throw apiError(res, body);
    },
  };
})();
