// Response types of the admin API (backend/internal/admin, backend/internal/auth).

export type LoginResponse = {
    token: string;
    expiresAt: string;
};

export type PlatformCount = {
    os: string;
    count: number;
};

export type Summary = {
    total: number;
    today: number;
    last7Days: number;
    last30Days: number;
    uniqueIps: number;
    byPlatform: PlatformCount[];
};

export type SeriesPoint = {
    period: string;
    total: number;
    byPlatform: Record<string, number>;
};

export type Download = {
    id: number;
    os: string;
    ipAddress: string;
    country: string;
    createdAt: string;
};

export type DownloadPage = {
    items: Download[];
    total: number;
    page: number;
    pageSize: number;
};
