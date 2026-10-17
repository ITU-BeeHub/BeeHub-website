# BeeHub Website

Source of [beehubapp.com](https://beehubapp.com): the download page for the
BeeHub desktop app, its documentation, and an admin panel with download
statistics.

```
Cloudflare ── cloudflared tunnel ── nginx :80 ─┬─ /        static files  /var/www/beehubapp.com
                                               └─ /api/    beehub-web    127.0.0.1:9000
                                                               │
                                                               └─ PostgreSQL (Docker)  127.0.0.1:5429
```

| Path | Contents |
| --- | --- |
| `backend/` | Go API (gin, GORM). `cmd/website` is the server, `cmd/admin-user` manages admin accounts. |
| `frontend/` | React + Vite + Tailwind site. The admin panel lives in `src/admin` and is loaded lazily. |
| `deploy/` | Production configuration: Compose file for PostgreSQL, systemd unit, nginx site, environment template, SQL maintenance scripts. |

## API

| Endpoint | Purpose |
| --- | --- |
| `GET /api/version` | Current release, polled by the desktop app's update check. Keep its response shape stable. |
| `GET /api/download?os=windows\|mac` | Serves the installer and counts the download. |
| `GET /api/health` | Returns 200 when the API can reach the database. |
| `POST /api/admin/login` | Exchanges admin credentials for a token (rate limited per IP). |
| `GET /api/admin/stats/{summary,daily,monthly}`, `GET /api/admin/downloads` | Admin reports; require `Authorization: Bearer <token>`. |

Swagger UI is served at `/swagger/index.html` when `SWAGGER_ENABLED=true`.
After changing handler annotations, regenerate it with
`cd backend && go run github.com/swaggo/swag/cmd/swag@v1.16.3 init -g cmd/website/main.go -o docs --parseInternal`.

### How downloads are counted

One click can produce several requests: browsers and download managers
probe with `HEAD`, resume with `Range` requests and split large files into
parallel ranges. Only a `GET` without a `Range` header, or with one starting
at byte 0, is counted, and repeats from the same IP address and platform
within `DOWNLOAD_DEDUP_WINDOW` (10 minutes) are counted once. The check and
insert run under a PostgreSQL advisory lock, so concurrent requests cannot
both be counted. The country comes from Cloudflare's `CF-IPCountry` header.

## Development

Requirements: Go 1.24+, Node.js 20+, Docker.

```sh
# Backend: create backend/.env from deploy/env.example, pointing at a local
# database, then:
cd backend && go run ./cmd/website

# Frontend: proxies /api to 127.0.0.1:9000 (override with VITE_API_PROXY).
cd frontend && npm ci && npm run dev

make test               # go vet, unit tests, eslint
make test-integration   # all backend tests against a throwaway PostgreSQL
```

## Production

Everything runs on one host as the deploy user, except the steps marked
with `sudo`.

| What | Where |
| --- | --- |
| Configuration and secrets | `/etc/beehub/env` (mode 640, from `deploy/env.example`) |
| API binaries | `/opt/beehub/bin/` |
| Installers | `/opt/beehub/installers/installer_<version>.{exe,dmg}` |
| Frontend | `/var/www/beehubapp.com/` |
| Database | container `beehub-db`, data in Docker volume `beehub_pgdata` |
| Logs | `journalctl -u beehub-backend` |
| Backups | `~/backups/beehub/` |

### First-time setup

```sh
sudo install -d -m 755 /etc/beehub
sudo install -m 640 -o root -g "$USER" deploy/env.example /etc/beehub/env   # then fill in the secrets
sudo install -m 644 deploy/systemd/beehub-backend.service /etc/systemd/system/
sudo install -m 644 deploy/nginx/beehubapp.com.conf /etc/nginx/sites-available/beehubapp.com
sudo ln -sf /etc/nginx/sites-available/beehubapp.com /etc/nginx/sites-enabled/beehubapp.com
sudo nginx -t && sudo systemctl reload nginx

docker volume create beehub_pgdata
make db-up
make deploy
make admin-user USERNAME=<name>
sudo systemctl enable beehub-backend
```

`make deploy-backend` restarts the service with `sudo`. To run it without a
password prompt, allow only that command:

```sh
echo "$USER ALL=(root) NOPASSWD: /usr/bin/systemctl restart beehub-backend.service" \
  | sudo tee /etc/sudoers.d/beehub && sudo chmod 440 /etc/sudoers.d/beehub && sudo visudo -c
```

### Releasing a new desktop version

1. Copy `installer_<version>.exe` and `installer_<version>.dmg` to
   `/opt/beehub/installers/` and make them world-readable (`chmod 644`).
2. Set `APP_VERSION=<version>` in `/etc/beehub/env`.
3. `sudo systemctl restart beehub-backend`

### Routine operations

```sh
make deploy                       # build and deploy API and frontend
make logs                         # follow API logs
make db-backup                    # dump the database to ~/backups/beehub
make db-psql                      # interactive psql
make admin-user USERNAME=<name>   # create an admin or reset its password
```

Restore a backup into the running container with
`docker exec -i beehub-db sh -c 'pg_restore --clean --if-exists -U "$POSTGRES_USER" -d "$POSTGRES_DB"' < <file>.dump`.
