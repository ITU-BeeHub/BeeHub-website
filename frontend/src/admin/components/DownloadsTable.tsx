import { countryFlag, formatDateTime } from "../format";
import type { Download } from "../types";
import { PlatformBadge } from "./ui";

export default function DownloadsTable({ downloads }: { downloads: Download[] }) {
    return (
        <div className="-mx-5 overflow-x-auto">
            <table className="min-w-full text-left text-sm">
                <thead>
                    <tr className="border-b border-slate-100 text-xs uppercase tracking-wide text-slate-500">
                        <th className="px-5 py-2 font-medium">Time</th>
                        <th className="px-5 py-2 font-medium">Platform</th>
                        <th className="px-5 py-2 font-medium">IP address</th>
                        <th className="px-5 py-2 font-medium">Country</th>
                    </tr>
                </thead>
                <tbody className="divide-y divide-slate-100">
                    {downloads.map((download) => (
                        <tr key={download.id} className="hover:bg-slate-50">
                            <td className="whitespace-nowrap px-5 py-2.5 tabular-nums text-slate-700">{formatDateTime(download.createdAt)}</td>
                            <td className="whitespace-nowrap px-5 py-2.5"><PlatformBadge os={download.os} /></td>
                            <td className="whitespace-nowrap px-5 py-2.5 font-mono text-[13px] text-slate-700">{download.ipAddress}</td>
                            <td className="whitespace-nowrap px-5 py-2.5 text-slate-700">
                                {download.country
                                    ? <>{countryFlag(download.country)} {download.country}</>
                                    : <span className="text-slate-400">—</span>}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}
