# jieun-contracts

Sumber kebenaran kontrak antar-repo. Perubahan kontrak **selalu** lewat PR
di sini terlebih dulu, tidak pernah langsung di service. Lihat
`../jieun-app/ARCHITECTURE.md` §2.2 dan `../jieun-app/INTEGRATION.md` §9.

## Isi

```
lint/golangci.yml     ← konfigurasi lint bersama untuk semua repo Go (ADA)
openapi/               ← spesifikasi OpenAPI per service (T-001a)
jobs/                  ← JSON Schema payload task (T-001a)
events/                ← JSON Schema event (T-001a)
errors/catalog.yaml    ← katalog kode error (T-001a)
```

`openapi/`, `jobs/`, `events/`, `errors/` dibangun di T-001a — repo ini baru
memuat konfigurasi lint untuk sekarang.

## Menjalankan

```bash
make setup
make lint    # golangci-lint config verify
```

## Dipakai oleh

Setiap repo Go (`jieun-platform`, `jieun-analyzer`, `jieun-clipper`) mengacu
`lint/golangci.yml` lewat path relatif di `Makefile` masing-masing. Merge ke
`main` nantinya menghasilkan klien Go dan tipe TypeScript ber-versi (T-001a).
