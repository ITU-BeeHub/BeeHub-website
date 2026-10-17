import { useSyncExternalStore } from "react";

const STORAGE_KEY = "beehub.admin.session";

export type Session = {
    token: string;
    expiresAt: string;
};

const listeners = new Set<() => void>();

function emitChange() {
    listeners.forEach((listener) => listener());
}

function subscribe(listener: () => void) {
    listeners.add(listener);
    return () => listeners.delete(listener);
}

function readRaw(): string | null {
    return sessionStorage.getItem(STORAGE_KEY);
}

function parse(raw: string | null): Session | null {
    if (!raw) return null;
    try {
        const session = JSON.parse(raw) as Session;
        return new Date(session.expiresAt).getTime() > Date.now() ? session : null;
    } catch {
        return null;
    }
}

export function getSession(): Session | null {
    return parse(readRaw());
}

export function saveSession(session: Session) {
    sessionStorage.setItem(STORAGE_KEY, JSON.stringify(session));
    emitChange();
}

export function clearSession() {
    sessionStorage.removeItem(STORAGE_KEY);
    emitChange();
}

/** Returns the current admin session and re-renders when it changes. */
export function useSession(): Session | null {
    return parse(useSyncExternalStore(subscribe, readRaw));
}
