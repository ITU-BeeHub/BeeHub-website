import { useState } from "react";
import { Link } from "react-router-dom";
import { endpoints, useApi } from "./api";
import DownloadsChart from "./components/DownloadsChart";
import DownloadsTable from "./components/DownloadsTable";
import { EmptyState, ErrorState, Panel, PlatformBadge, SegmentedControl, StatCard } from "./components/ui";
import { formatNumber, platformColor } from "./format";
import type { DownloadPage, SeriesPoint, Summary } from "./types";

type Range = "daily" | "monthly";

const RANGES: { value: Range; label: string; path: string }[] = [
    { value: "daily", label: "Last 30 days", path: endpoints.daily(30) },
    { value: "monthly", label: "Last 12 months", path: endpoints.monthly(12) },
];

export default function DashboardPage() {
    const summary = useApi<Summary>(endpoints.summary());
    const [range, setRange] = useState<Range>("daily");
    const series = useApi<SeriesPoint[]>(RANGES.find((r) => r.value === range)!.path);
    const recent = useApi<DownloadPage>(endpoints.downloads({ page: 1, pageSize: 8 }));

    const refresh = () => {
        summary.reload();
        series.reload();
        recent.reload();
    };

    const stat = (value?: number) => (value === undefined ? undefined : formatNumber(value));

    return (
        <div className="space-y-6">
            <div className="flex items-center justify-between">
                <h1 className="text-2xl font-semibold text-slate-900">Overview</h1>
                <button
                    onClick={refresh}
                    className="rounded-lg border border-slate-300 bg-white px-3 py-1.5 text-sm font-medium text-slate-700 hover:bg-slate-50"
                >
                    Refresh
                </button>
            </div>

            {summary.error ? (
                <Panel><ErrorState error={summary.error} onRetry={summary.reload} /></Panel>
            ) : (
                <div className="grid grid-cols-2 gap-4 lg:grid-cols-4">
                    <StatCard label="Total downloads" value={stat(summary.data?.total)}
                        hint={summary.data && `${formatNumber(summary.data.uniqueIps)} unique IP addresses`} />
                    <StatCard label="Today" value={stat(summary.data?.today)} />
                    <StatCard label="Last 7 days" value={stat(summary.data?.last7Days)} />
                    <StatCard label="Last 30 days" value={stat(summary.data?.last30Days)} />
                </div>
            )}

            <div className="grid gap-6 lg:grid-cols-3">
                <Panel
                    title="Downloads over time"
                    className="lg:col-span-2"
                    actions={<SegmentedControl value={range} options={RANGES} onChange={setRange} />}
                >
                    {series.error
                        ? <ErrorState error={series.error} onRetry={series.reload} />
                        : series.data
                            ? <DownloadsChart series={series.data} />
                            : <div className="h-72 animate-pulse rounded-lg bg-slate-50" />}
                </Panel>

                <Panel title="By platform">
                    {summary.data && <PlatformSplit summary={summary.data} />}
                </Panel>
            </div>

            <Panel
                title="Recent downloads"
                actions={<Link to="/admin/downloads" className="text-sm font-medium text-[#0372CE] hover:underline">View all</Link>}
            >
                {recent.error
                    ? <ErrorState error={recent.error} onRetry={recent.reload} />
                    : recent.data?.items.length === 0
                        ? <EmptyState>No downloads yet.</EmptyState>
                        : recent.data && <DownloadsTable downloads={recent.data.items} />}
            </Panel>
        </div>
    );
}

function PlatformSplit({ summary }: { summary: Summary }) {
    if (summary.total === 0) return <EmptyState>No downloads yet.</EmptyState>;

    return (
        <ul className="space-y-4">
            {summary.byPlatform.map(({ os, count }) => {
                const share = Math.round((count / summary.total) * 100);
                return (
                    <li key={os}>
                        <div className="mb-1.5 flex items-baseline justify-between">
                            <PlatformBadge os={os} />
                            <span className="text-sm tabular-nums text-slate-500">
                                <span className="font-medium text-slate-900">{formatNumber(count)}</span> · {share}%
                            </span>
                        </div>
                        <div className="h-2 overflow-hidden rounded-full bg-slate-100">
                            <div className="h-full rounded-full" style={{ width: `${share}%`, backgroundColor: platformColor(os) }} />
                        </div>
                    </li>
                );
            })}
        </ul>
    );
}
