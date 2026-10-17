import { useEffect, useState } from "react";
import { endpoints, useApi } from "./api";
import DownloadsTable from "./components/DownloadsTable";
import { EmptyState, ErrorState, Panel } from "./components/ui";
import { PLATFORMS, formatNumber } from "./format";
import type { DownloadPage } from "./types";

const PAGE_SIZE = 25;
const SEARCH_DELAY_MS = 300;

export default function DownloadsPage() {
    const [searchInput, setSearchInput] = useState("");
    const [search, setSearch] = useState("");
    const [os, setOs] = useState("");
    const [page, setPage] = useState(1);

    // Query the API only once the user pauses typing.
    useEffect(() => {
        const timer = setTimeout(() => {
            setSearch(searchInput.trim());
            setPage(1);
        }, SEARCH_DELAY_MS);
        return () => clearTimeout(timer);
    }, [searchInput]);

    const { data, error, loading, reload } = useApi<DownloadPage>(
        endpoints.downloads({ page, pageSize: PAGE_SIZE, q: search, os }),
    );

    const pageCount = data ? Math.max(1, Math.ceil(data.total / PAGE_SIZE)) : 1;
    const firstItem = data && data.total > 0 ? (page - 1) * PAGE_SIZE + 1 : 0;
    const lastItem = data ? (page - 1) * PAGE_SIZE + data.items.length : 0;

    return (
        <div className="space-y-6">
            <h1 className="text-2xl font-semibold text-slate-900">Download log</h1>

            <Panel
                actions={
                    <div className="flex w-full flex-wrap gap-3">
                        <input
                            type="search"
                            value={searchInput}
                            onChange={(e) => setSearchInput(e.target.value)}
                            placeholder="Search IP address or country"
                            className="min-w-0 flex-1 rounded-lg border border-slate-300 px-3 py-1.5 text-sm focus:border-[#0372CE] focus:outline-none focus:ring-2 focus:ring-[#0372CE]/20"
                        />
                        <select
                            value={os}
                            onChange={(e) => { setOs(e.target.value); setPage(1); }}
                            className="rounded-lg border border-slate-300 bg-white px-3 py-1.5 text-sm text-slate-700"
                        >
                            <option value="">All platforms</option>
                            {PLATFORMS.map((platform) => (
                                <option key={platform.id} value={platform.id}>{platform.label}</option>
                            ))}
                        </select>
                    </div>
                }
            >
                {error ? (
                    <ErrorState error={error} onRetry={reload} />
                ) : !data ? (
                    <div className="h-96 animate-pulse rounded-lg bg-slate-50" />
                ) : data.items.length === 0 ? (
                    <EmptyState>No downloads match your filters.</EmptyState>
                ) : (
                    <div className={loading ? "opacity-60 transition-opacity" : ""}>
                        <DownloadsTable downloads={data.items} />
                    </div>
                )}

                {data && data.total > 0 && (
                    <footer className="mt-4 flex items-center justify-between text-sm text-slate-500">
                        <span className="tabular-nums">
                            {formatNumber(firstItem)}–{formatNumber(lastItem)} of {formatNumber(data.total)}
                        </span>
                        <div className="flex gap-2">
                            <PageButton disabled={page <= 1} onClick={() => setPage(page - 1)}>Previous</PageButton>
                            <PageButton disabled={page >= pageCount} onClick={() => setPage(page + 1)}>Next</PageButton>
                        </div>
                    </footer>
                )}
            </Panel>
        </div>
    );
}

function PageButton({ disabled, onClick, children }: { disabled: boolean; onClick: () => void; children: string }) {
    return (
        <button
            disabled={disabled}
            onClick={onClick}
            className="rounded-lg border border-slate-300 bg-white px-3 py-1.5 font-medium text-slate-700 hover:bg-slate-50 disabled:cursor-not-allowed disabled:opacity-40"
        >
            {children}
        </button>
    );
}
