import { useCallback, useEffect, useState } from "react";
import { clearSession, getSession } from "./session";
import type { LoginResponse } from "./types";

export class ApiError extends Error {
    constructor(public readonly status: number, message: string) {
        super(message);
    }
}

async function request<T>(path: string, init: RequestInit = {}): Promise<T> {
    const session = getSession();
    const headers = new Headers(init.headers);
    if (init.body) headers.set("Content-Type", "application/json");
    if (session) headers.set("Authorization", `Bearer ${session.token}`);

    const response = await fetch(`/api${path}`, { ...init, headers });
    const body = await response.json().catch(() => null);

    if (!response.ok) {
        if (response.status === 401 && session) clearSession();
        throw new ApiError(response.status, body?.error ?? `Request failed (${response.status})`);
    }
    return body as T;
}

export function login(username: string, password: string) {
    return request<LoginResponse>("/admin/login", {
        method: "POST",
        body: JSON.stringify({ username, password }),
    });
}

export const endpoints = {
    summary: () => "/admin/stats/summary",
    daily: (days: number) => `/admin/stats/daily?days=${days}`,
    monthly: (months: number) => `/admin/stats/monthly?months=${months}`,
    downloads: (params: { page: number; pageSize: number; q?: string; os?: string }) => {
        const query = new URLSearchParams({ page: String(params.page), pageSize: String(params.pageSize) });
        if (params.q) query.set("q", params.q);
        if (params.os) query.set("os", params.os);
        return `/admin/downloads?${query}`;
    },
};

type RequestState<T> = {
    data?: T;
    error?: Error;
    loading: boolean;
};

/** Fetches `path` from the API and refetches when it changes or on reload(). */
export function useApi<T>(path: string) {
    const [state, setState] = useState<RequestState<T>>({ loading: true });
    const [attempt, setAttempt] = useState(0);

    useEffect(() => {
        const controller = new AbortController();
        setState((previous) => ({ data: previous.data, loading: true }));

        request<T>(path, { signal: controller.signal }).then(
            (data) => setState({ data, loading: false }),
            (error: Error) => {
                if (!controller.signal.aborted) setState((previous) => ({ ...previous, error, loading: false }));
            },
        );
        return () => controller.abort();
    }, [path, attempt]);

    const reload = useCallback(() => setAttempt((n) => n + 1), []);
    return { ...state, reload };
}
