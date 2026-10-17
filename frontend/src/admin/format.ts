// Statistics are bucketed by day in this time zone on the server
// (STATS_TIMEZONE), so timestamps are shown in it too.
const TIME_ZONE = "Europe/Istanbul";

export const PLATFORMS = [
    { id: "windows", label: "Windows", color: "#0372CE" },
    { id: "mac", label: "macOS", color: "#D97706" },
] as const;

export function platformLabel(os: string) {
    return PLATFORMS.find((platform) => platform.id === os)?.label ?? os;
}

export function platformColor(os: string) {
    return PLATFORMS.find((platform) => platform.id === os)?.color ?? "#64748b";
}

const numberFormat = new Intl.NumberFormat("en-US");
export function formatNumber(value: number) {
    return numberFormat.format(value);
}

const dateTimeFormat = new Intl.DateTimeFormat("en-GB", {
    dateStyle: "medium",
    timeStyle: "short",
    timeZone: TIME_ZONE,
});
export function formatDateTime(iso: string) {
    return dateTimeFormat.format(new Date(iso));
}

/** Formats a series period ("2026-10-17" or "2026-10") as a short axis label. */
export function formatPeriod(period: string) {
    const [year, month, day] = period.split("-").map(Number);
    const date = new Date(Date.UTC(year, month - 1, day ?? 1));
    return date.toLocaleDateString("en-GB", day
        ? { day: "numeric", month: "short", timeZone: "UTC" }
        : { month: "short", year: "2-digit", timeZone: "UTC" });
}

/** Turns an ISO 3166-1 alpha-2 code into its flag emoji. */
export function countryFlag(code: string) {
    if (!/^[A-Z]{2}$/.test(code)) return "";
    return String.fromCodePoint(...[...code].map((char) => 0x1f1e6 + char.charCodeAt(0) - 65));
}
