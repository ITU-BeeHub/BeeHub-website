import { Bar, BarChart, CartesianGrid, Legend, ResponsiveContainer, Tooltip, XAxis, YAxis } from "recharts";
import { PLATFORMS, formatPeriod } from "../format";
import type { SeriesPoint } from "../types";

/** Stacked bar chart of downloads per period, split by platform. */
export default function DownloadsChart({ series }: { series: SeriesPoint[] }) {
    const rows = series.map((point) => ({ period: point.period, ...point.byPlatform }));

    return (
        <div className="h-72">
            <ResponsiveContainer width="100%" height="100%">
                <BarChart data={rows} margin={{ top: 4, right: 4, bottom: 0, left: -16 }}>
                    <CartesianGrid vertical={false} stroke="#e2e8f0" />
                    <XAxis
                        dataKey="period"
                        tickFormatter={formatPeriod}
                        tick={{ fontSize: 12, fill: "#64748b" }}
                        tickLine={false}
                        axisLine={false}
                        minTickGap={16}
                    />
                    <YAxis allowDecimals={false} tick={{ fontSize: 12, fill: "#64748b" }} tickLine={false} axisLine={false} />
                    <Tooltip
                        labelFormatter={(period: string) => formatPeriod(period)}
                        cursor={{ fill: "#f1f5f9" }}
                        contentStyle={{ borderRadius: 8, borderColor: "#e2e8f0", fontSize: 13 }}
                    />
                    <Legend iconType="circle" iconSize={8} wrapperStyle={{ fontSize: 13 }} />
                    {PLATFORMS.map((platform, index) => (
                        <Bar
                            key={platform.id}
                            dataKey={platform.id}
                            name={platform.label}
                            stackId="downloads"
                            fill={platform.color}
                            radius={index === PLATFORMS.length - 1 ? [3, 3, 0, 0] : 0}
                            maxBarSize={32}
                            isAnimationActive={false}
                        />
                    ))}
                </BarChart>
            </ResponsiveContainer>
        </div>
    );
}
