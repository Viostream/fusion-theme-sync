# Fusion Theme Sync

[![Test](https://github.com/Viostream/fusion-theme-sync/actions/workflows/test.yml/badge.svg)](https://github.com/Viostream/fusion-theme-sync/actions/workflows/test.yml)
[![Build](https://github.com/Viostream/fusion-theme-sync/actions/workflows/build.yml/badge.svg)](https://github.com/Viostream/usion-theme-sync/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/viostream/fusion-theme-sync/.svg)](https://pkg.go.dev/github.com/viostream/fusion-theme-sync/)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Viostream_fusion-theme-sync&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Viostream_fusion-theme-sync)


Fusion's themes are a flat object that contain heaps of ever-changing
templates. Managing them via a traditional version control system is onerous.

This package can be used to sync down the themes from one fusion server (e.g.
your dev server), and then sync them up to others.
