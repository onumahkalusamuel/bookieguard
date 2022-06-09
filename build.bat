@echo off
title Bookie Guard
echo --------------------------
echo make sure you have installed the following: 
echo --------------------------
echo go: https://go.dev/doc/install
echo wixtoolset: https://wixtoolset.org/releases/
echo go-msi: https://github.com/mh-cbon/go-msi/
echo --------------------------
go build
go-msi make --msi BookieGuard.msi --version 1.0.0