// src/pages/UsersPage.tsx
import { useEffect, useState } from "react";

type User = {
  id: number;
  name: string;
  email: string;
  status: string;
  last_modified: number; // unix seconds
};

type UsersResponse = {
  data: User[];
  limit: number;
  skip: number;
  total: number;
};

export default function UsersPage() {
  const [users, setUsers] = useState<User[]>([]);
  const [limit, setLimit] = useState(5);
  const [skip, setSkip] = useState(0);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const ctrl = new AbortController();
    (async () => {
      try {
        setLoading(true);
        setError(null);
        const url = `http://localhost:8085/users?limit=${limit}&skip=${skip}`;
        const res = await fetch(url, { signal: ctrl.signal });
        if (!res.ok) throw new Error(`HTTP ${res.status}`);
        const json: UsersResponse = await res.json();
        setUsers(json.data ?? []);
        setTotal(json.total ?? 0);
      } catch (e: any) {
        if (e.name !== "AbortError") setError(e.message ?? "Failed to load");
      } finally {
        setLoading(false);
      }
    })();
    return () => ctrl.abort();
  }, [limit, skip]);

  const from = users.length ? skip + 1 : 0;
  const to = Math.min(skip + limit, total);

  // Build rows with a plain loop (no .map)
  const rows: JSX.Element[] = [];
  for (const u of users) {
    rows.push(
      <tr key={u.id} className="border-b">
        <td className="px-3 py-2">{u.id}</td>
        <td className="px-3 py-2">{u.name}</td>
        <td className="px-3 py-2">{u.email}</td>
        <td className="px-3 py-2">{u.status}</td>
        <td className="px-3 py-2">
          {new Date(u.last_modified * 1000).toLocaleString()}
        </td>
      </tr>
    );
  }

  return (
    <section className="mx-auto max-w-5xl">
      <h1 className="text-2xl font-bold mb-4">Users</h1>

      {error && (
        <div className="mb-3 rounded border border-red-300 bg-red-50 px-3 py-2 text-sm text-red-700">
          {error}
        </div>
      )}

      <div className="mb-3 text-sm text-gray-600">
        {loading ? "Loading..." : `Showing ${from}-${to} of ${total}`}
      </div>

      <div className="overflow-x-auto rounded border">
        <table className="min-w-full text-sm">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-3 py-2 text-left">ID</th>
              <th className="px-3 py-2 text-left">Name</th>
              <th className="px-3 py-2 text-left">Email</th>
              <th className="px-3 py-2 text-left">Status</th>
              <th className="px-3 py-2 text-left">Last Modified</th>
            </tr>
          </thead>
          <tbody>{rows}</tbody>
        </table>
      </div>

      <div className="mt-4 flex items-center gap-2">
        <button
          className="px-3 py-2 rounded border hover:bg-gray-100 disabled:opacity-50"
          onClick={() => setSkip(Math.max(0, skip - limit))}
          disabled={skip === 0 || loading}
        >
          Prev
        </button>
        <button
          className="px-3 py-2 rounded border hover:bg-gray-100 disabled:opacity-50"
          onClick={() => setSkip(skip + limit)}
          disabled={skip + limit >= total || loading}
        >
          Next
        </button>

        <label className="ml-4 text-sm">
          Limit:
          <select
            className="ml-2 border rounded px-2 py-1"
            value={limit}
            onChange={(e) => {
              setSkip(0);
              setLimit(Number(e.target.value));
            }}
          >
            <option value={5}>5</option>
            <option value={10}>10</option>
            <option value={20}>20</option>
          </select>
        </label>
      </div>
    </section>
  );
}
