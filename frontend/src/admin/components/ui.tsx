import React from "react";
import { platformColor, platformLabel } from "../format";

export function Panel({ title, actions, children, className = "" }: {
    title?: string;
    actions?: React.ReactNode;
    children: React.ReactNode;
    className?: string;
}) {
    return (
        <section className={`rounded-xl border border-slate-200 bg-white ${className}`}>
            {(title || actions) && (
                <header className="flex flex-wrap items-center justify-between gap-3 border-b border-slate-100 px-5 py-4">
                    {title && <h2 className="text-sm font-semibold text-slate-900">{title}</h2>}
                    {actions}
                </header>
            )}
            <div className="p-5">{children}</div>
        </section>
    );
}

export function StatCard({ label, value, hint }: { label: string; value?: string; hint?: string }) {
    return (
        <div className="rounded-xl border border-slate-200 bg-white px-5 py-4">
            <p className="text-sm text-slate-500">{label}</p>
            {value === undefined
                ? <div className="mt-2 h-8 w-20 animate-pulse rounded bg-slate-100" />
                : <p className="mt-1 text-3xl font-semibold tabular-nums text-slate-900">{value}</p>}
            {hint && <p className="mt-1 text-xs text-slate-400">{hint}</p>}
        </div>
    );
}

export function ErrorState({ error, onRetry }: { error: Error; onRetry: () => void }) {
    return (
        <div className="flex flex-col items-center gap-3 py-8 text-center">
            <p className="text-sm text-slate-600">Could not load data: {error.message}</p>
            <button onClick={onRetry} className="rounded-lg border border-slate-300 px-3 py-1.5 text-sm font-medium text-slate-700 hover:bg-slate-50">
                Try again
            </button>
        </div>
    );
}

export function EmptyState({ children }: { children: React.ReactNode }) {
    return <p className="py-8 text-center text-sm text-slate-500">{children}</p>;
}

export function PlatformBadge({ os }: { os: string }) {
    return (
        <span className="inline-flex items-center gap-1.5 text-sm text-slate-700">
            <span className="h-2 w-2 rounded-full" style={{ backgroundColor: platformColor(os) }} />
            {platformLabel(os)}
        </span>
    );
}

export function SegmentedControl<T extends string>({ value, options, onChange }: {
    value: T;
    options: { value: T; label: string }[];
    onChange: (value: T) => void;
}) {
    return (
        <div className="inline-flex rounded-lg bg-slate-100 p-0.5" role="tablist">
            {options.map((option) => (
                <button
                    key={option.value}
                    role="tab"
                    aria-selected={value === option.value}
                    onClick={() => onChange(option.value)}
                    className={`rounded-md px-3 py-1 text-sm font-medium transition-colors ${value === option.value
                        ? "bg-white text-slate-900 shadow-sm"
                        : "text-slate-500 hover:text-slate-700"}`}
                >
                    {option.label}
                </button>
            ))}
        </div>
    );
}
