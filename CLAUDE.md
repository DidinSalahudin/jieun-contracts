# jieun-contracts

## Knowledge base — graphify

Repo ini punya knowledge graph di `graphify-out/` (`graph.json`, `GRAPH_REPORT.md`,
`graph.html`) — per 2026-08-23 **sebagian**: AST lengkap + cuma 1 dari 3 chunk ekstraksi
semantik yang sempet kelar (1109 node) sebelum kena monthly spend limit.

- Sebelum jawab pertanyaan arsitektur/lintas file (termasuk relasi schema JSON → generated
  client Go/TS), cek `graphify-out/graph.json` ada dulu, lalu pakai
  `graphify query "<pertanyaan>"` daripada baca file satu-satu dari nol.
- Jalankan `/graphify . --update` begitu spend limit naik biar 2 chunk semantik yang bolong
  nyusul, dan setelah perubahan kode/schema signifikan supaya graph gak basi.
