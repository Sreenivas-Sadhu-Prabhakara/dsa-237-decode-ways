# Decode Ways

> NeetCode 150 · 1-D DP · Blind 75 — climbing stairs over a digit string: extend by a 1-digit or a valid 2-digit letter — except every rule bends around zero. Go, stdlib only, tested.

## Problem

Digits encode letters (1→A … 26→Z); count the decodings of a digit string. "226" has three (2·2·6, 22·6, 2·26); "06" has none — zero can only exist as the tail of 10 or 20.

## Approach

Structurally climbing stairs: ways(i) gets a contribution from ways(i−1) when s[i] alone is a letter, and from ways(i−2) when s[i−1..i] forms 10–26. All the actual difficulty is zero: it contributes nothing as a single digit, kills everything as a leading digit, and strands the string in cases like "100" — handled by the early-exit when a position accrues zero ways (continuing would just propagate zeros, but exiting makes the reasoning visible). Two rolling variables as ever. Enumerating the zero cases *as tests first* is the sane way to write this one.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
